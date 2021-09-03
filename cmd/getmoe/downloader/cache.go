package downloader

import (
	"encoding/json"
	"errors"
	"os"
)

// cacheValue stores information about downloaded image.
type cacheValue struct {
	Board string `json:"board,omitempty"`
	ID    int    `json:"id,omitempty"`
	URL   string `json:"url,omitempty"`
}

// loadCache reads cache from local file.
func (d *Downloader) loadCache(in string) error {
	var cacheMap map[string]cacheValue
	file, err := os.Open(in)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&cacheMap); err != nil {
		return err
	}
	for key, value := range cacheMap {
		d.cache.Store(key, value)
	}
	return nil
}

// saveCache writes cache to a local file.
func (d *Downloader) saveCache(out string) error {
	cacheMap := make(map[string]cacheValue)
	d.cache.Range(func(key, value interface{}) bool {
		cacheMap[key.(string)] = value.(cacheValue)
		return true
	})
	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(cacheMap); err != nil {
		return err
	}
	return nil
}
