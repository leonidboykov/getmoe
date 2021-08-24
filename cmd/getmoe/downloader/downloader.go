package downloader

import (

	// init all known providers.
	"fmt"
	"sync"

	"github.com/leonidboykov/getmoe"
	_ "github.com/leonidboykov/getmoe/provider/danbooru"
	_ "github.com/leonidboykov/getmoe/provider/gelbooru"
	_ "github.com/leonidboykov/getmoe/provider/moebooru"
	_ "github.com/leonidboykov/getmoe/provider/sankaku"
	_ "github.com/leonidboykov/getmoe/provider/sankaku/v2"
)

const cacheFile = "getmoe_cache.json"

type Downloader struct {
	Boards     map[string]*getmoe.Board
	boardNames []string

	// cache stores info about downloaded images with image hash as key.
	// Hashes are more or less consistent across boorus.
	cache sync.Map
}

func NewDownloader(config map[string]getmoe.BoardConfiguration) (*Downloader, error) {
	d := &Downloader{
		Boards: make(map[string]*getmoe.Board),
	}
	d.loadCache(cacheFile)

	for name, board := range config {
		b, err := getmoe.NewBoard(name, board)
		if err != nil {
			return nil, fmt.Errorf("unable to create a board '%s': %w", name, err)
		}
		d.Boards[name] = b
		d.boardNames = append(d.boardNames, name)
	}

	return d, nil
}

// func (d *Downloader) Execute(cmds []getmoe.DownloadConfiguration) error {
// 	wg := new(sync.WaitGroup)
// 	wg.Add(len(cmds))
// 	for _, cmd := range cmds {
// 		cmd := cmd
// 		go func() {
// 			d.execCommand(cmd)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	return nil
// }

// func (d *Downloader) execCommand(cmd getmoe.DownloadConfiguration) error {
// 	// Execute command on all boards if there are no boards specified.
// 	if len(cmd.Boards) == 0 {
// 		cmd.Boards = d.boardNames
// 	}

// 	wg := new(sync.WaitGroup)
// 	for _, name := range cmd.Boards {
// 		if board, ok := d.Boards[name]; ok {

// 		}
// 	}

// 	return nil
// }

// func (d *Downloader) work(b *getmoe.Board, cmd getmoe.DownloadConfiguration) error {
// 	posts, err := b.RequestAll(cmd.Tags)
// 	if err != nil {
// 		return err
// 	}
// 	return
// }

// func (d *Downloader) filterPosts(posts []getmoe.Post, filters []string) []getmoe.Post {
// 	sort.SearchStrings()
// }
