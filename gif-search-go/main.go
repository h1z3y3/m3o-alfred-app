package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/h1z3y3/m3o-alfred-workflow/i18n"
	"github.com/h1z3y3/m3o-alfred-workflow/m3o"
	"github.com/h1z3y3/m3o-alfred-workflow/workflow"
	"github.com/tidwall/gjson"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	token := strings.TrimSpace(os.Getenv("m3o-token"))
	if token == "" {
		workflow.NewError(i18n.I("Please set environment variable `m3o-token` first."), "").Display()
		return
	}

	limit, _ := strconv.Atoi(os.Getenv("limit"))
	if limit <= 0 {
		limit = 10
	}

	micro := m3o.NewMicro(token)
	resp, err := micro.Post("/gifs/Search", m3o.Data{
		"limit": limit,
		"query": os.Args[1],
	}).String()

	if err != nil {
		workflow.NewError(i18n.I("Request error."), err.Error()).Display()
		return
	}

	items := make(workflow.Items, 0)

	gjson.Get(resp, "data").ForEach(func(key, val gjson.Result) bool {

		downsized := val.Get("images.downsized")
		size := formatFileSize(downsized.Get("size").Int())
		height := downsized.Get("height").Int()
		width := downsized.Get("width").Int()

		item := workflow.Item{
			Title:        val.Get("title").Str,
			Subtitle:     fmt.Sprintf("%dx%d %s %s", width, height, size, val.Get("slug").Str),
			Valid:        true,
			Uid:          "",
			Arg:          fmt.Sprintf("%s##%s", val.Get("images.downsized.url").Str, val.Get("id")),
			Autocomplete: "",
			Type:         "",
			Icon: workflow.Icon{
				Path: val.Get("images.downsized_still.url").Str,
			},
			IconType: workflow.IconTypeUrl,
		}

		items = append(items, item)

		return true
	})

	if len(items) == 0 {
		workflow.NewError(i18n.I("Not found."), "").Display()
		return
	}

	items.Display()
}

// formatFileSize
func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.1fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.1fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.1fMB", float64(fileSize)/float64(1024*1024))
	} else {
		return fmt.Sprintf("%.1fGB", float64(fileSize)/float64(1024*1024*1024))
	}
}
