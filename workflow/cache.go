package workflow

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type cache struct {
	url *url.URL
}

func NewCache(link string) *cache {
	u, _ := url.Parse(link)
	return &cache{
		url: u,
	}
}

// Cache downloads the cache to alfred cache directory and return the location
func (i *cache) Cache() (string, error) {
	if i.url.String() == "" {
		return "", nil
	}

	filename := fmt.Sprintf("%x", md5.Sum([]byte(i.url.String())))
	ext := filepath.Ext(i.url.Path)
	path := fmt.Sprintf("%s/%s%s", AlfredWorkflowCacheDir, filename, ext)

	_, err := os.Stat(path)
	if err == nil {
		return path, nil
	}

	resp, err := http.Get(i.url.String())
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("download status code is not 200")
	}

	// create an empty file
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// write the bytes to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}
