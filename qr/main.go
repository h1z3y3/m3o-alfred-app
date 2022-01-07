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

	size, _ := strconv.Atoi(os.Getenv("size"))
	if size <= 0 {
		size = 100
	}

	micro := m3o.NewMicro(token)
	resp, err := micro.Post("/qr/Generate", m3o.Data{
		"text": args[1],
		"size": size,
	}).String()

	if err != nil {
		workflow.NewError(i18n.I("Request error."), err.Error()).Display()
		return
	}

	data := gjson.Parse(resp)

	qr := data.Get("qr").String()

	cache, err := workflow.NewCache(qr).Cache()
	if err != nil {
		workflow.NewError(i18n.I("Download QR code Error."), err.Error()).Display()
		return
	}

	items := workflow.Items{
		workflow.Item{
			Title:    "Press Enter to Copy",
			Subtitle: "",
			Valid:    true,
			Arg:      fmt.Sprintf("%s##%s", cache, qr),
			Icon: workflow.Icon{
				Path: cache,
			},
			IconType: workflow.IconTypeLocal,
		},
	}

	items.Display()
}
