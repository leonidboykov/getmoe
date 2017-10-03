package getmoe

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/leonidboykov/getmoe/utils"
)

// Post contains post data, represents intersection of *boorus post structs
type Post struct {
	ID       int         `json:"id"`
	FileURL  string      `json:"file_url"`
	FileSize int         `json:"file_size"`
	Width    int         `json:"width"`
	Height   int         `json:"height"`
	Author   string      `json:"author"`
	Source   string      `json:"source"`
	Rating   string      `json:"rating"`
	Md5      string      `json:"md5"`
	Tags     interface{} `json:"tags"`
	Score    int         `json:"score"`
	// TODO: add tags as array
	// Tags     []string `json:"tags"`
}

// GetTags ...
// func (p *Post) GetTags() []string {
// 	return strings.Split(p.Tags, " ")
// }

// Save post to dir
func (p Post) Save(saveDir string) error {
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
