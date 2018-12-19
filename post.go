package getmoe

import (
	"io"
	"net/http"
	"os"
	"path"
	"time"

	utils "github.com/leonidboykov/getmoe/internal/helper"
)

// Post contains post data, represents intersection of *boorus post structs
type Post struct {
	ID        int       `json:"id"`
	FileURL   string    `json:"file_url"`
	FileSize  int       `json:"file_size"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []string  `json:"tags"`
	Author    string    `json:"author"`
	Source    string    `json:"source"`
	Rating    string    `json:"rating"`
	Hash      string    `json:"hash"`
	Score     int       `json:"score"`
}

// HasTag returns true if post has specified tag
func (p *Post) HasTag(tag string) bool {
	for i := range p.Tags {
		if p.Tags[i] == tag {
			return true
		}
	}
	return false
}

// Save post to dir
func (p *Post) Save(saveDir string) error {
	// Getting the actual URL
	// TODO: support JPG sources forcing
	fileName, err := utils.FileURLUnescape(p.FileURL)
	if err != nil {
		return err
	}

	file, err := os.Create(path.Join(saveDir, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(p.FileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	io.Copy(file, resp.Body)

	return nil
}
