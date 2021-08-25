package downloader

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/leonidboykov/getmoe"

	// init all known providers
	_ "github.com/leonidboykov/getmoe/provider/danbooru"
	_ "github.com/leonidboykov/getmoe/provider/gelbooru"
	_ "github.com/leonidboykov/getmoe/provider/moebooru"
	_ "github.com/leonidboykov/getmoe/provider/sankaku"
	_ "github.com/leonidboykov/getmoe/provider/sankaku/v2"
)

const cacheFile = "getmoe_cache.json"

type Downloader struct {
	boards     map[string]*getmoe.Board
	boardNames []string

	// cache stores info about downloaded images with image hash as key.
	// Hashes are more or less consistent across boorus.
	cache sync.Map
}

func NewDownloader(config map[string]getmoe.BoardConfiguration) (*Downloader, error) {
	d := &Downloader{
		boards: make(map[string]*getmoe.Board),
	}
	d.loadCache(cacheFile)

	for name, board := range config {
		b, err := getmoe.NewBoard(name, board)
		if err != nil {
			return nil, fmt.Errorf("unable to create a board '%s': %w", name, err)
		}
		d.boards[name] = b
		d.boardNames = append(d.boardNames, name)
	}

	return d, nil
}

func (d *Downloader) Execute(cmd getmoe.DownloadConfiguration) error {
	sort.Strings(cmd.Filters)

	wg := new(sync.WaitGroup)
	wg.Add(len(cmd.Searches))
	for _, search := range cmd.Searches {
		search := search
		_ = search
		go func() {
			d.execCommand(search, cmd.Filters, cmd.SavePath)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func (d *Downloader) execCommand(cmd getmoe.SearchConfiguration, filters []string, savePath string) error {
	// Execute command on all boards if there are no boards specified.
	if len(cmd.Boards) == 0 {
		cmd.Boards = d.boardNames
	}

	wg := new(sync.WaitGroup)
	for _, name := range cmd.Boards {
		if board, ok := d.boards[name]; ok {
			board := board
			go func() {
				d.requestFromBoard(board, cmd.Tags, filters, savePath)
				wg.Done()
			}()
		}
	}
	wg.Done()
	return nil
}

type templateData struct {
	BoardName  string
	PostID     int
	PostAuthor string
	FilePath   string
	FileExt    string
	FileHash   string
}

func (d *Downloader) requestFromBoard(b *getmoe.Board, t getmoe.Tags, filters []string, savePath string) error {
	posts, err := b.RequestAll(t)
	if err != nil {
		return err
	}
	posts = d.filterPosts(posts, filters)

	tmpl, err := template.New("savepath").Parse(savePath)
	if err != nil {
		return err
	}
	for _, p := range posts {
		var fname bytes.Buffer
		filePath, err := basePathFromURL(p.FileURL)
		if err != nil {
			return err
		}
		ext := filepath.Ext(filePath)
		bname := strings.TrimSuffix(filePath, ext)
		data := templateData{
			BoardName:  b.Name,
			PostID:     p.ID,
			PostAuthor: p.Author,
			FileHash:   p.Hash,
			FilePath:   bname,
			FileExt:    ext,
		}
		if err := tmpl.Execute(&fname, data); err != nil {
			return err
		}
		fmt.Println("Saving to", fname.String())
		d.cache.Store(p.Hash, cacheValue{
			Board: b.Name,
			URL:   p.FileURL,
		})
	}
	return nil
}

func (d *Downloader) filterPosts(posts []getmoe.Post, filters []string) []getmoe.Post {
	var filteredPosts []getmoe.Post
	for _, post := range posts {
		if sliceContains(post.Tags, filters) {
			continue
		}
		if _, ok := d.cache.Load(post.Hash); ok {
			continue
		}
		filteredPosts = append(filteredPosts, post)
	}
	return filteredPosts
}

func sliceContains(a, b []string) bool {
	for i := range a {
		if sort.SearchStrings(b, a[i]) != len(b) {
			return true
		}
	}
	return false
}

func basePathFromURL(fileURL string) (string, error) {
	u, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}
	return filepath.Base(u.Path), nil
}
