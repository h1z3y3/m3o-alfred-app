package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
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

	ip := net.ParseIP(args[1])
	if nil == ip {
		workflow.NewError(i18n.I("Please input valid IP address"), "").Display()
		return
	}

	micro := m3o.NewMicro(token)
	resp, err := micro.Post("/ip/Lookup", m3o.Data{
		"ip": ip.String(),
	}).String()

	if err != nil {
		workflow.NewError(i18n.I("Request error."), err.Error()).Display()
		return
	}

	data := gjson.Parse(resp)

	title := data.Get("city").String()

	if country := data.Get("country").String(); country != "" {
		if title != "" {
			title += ", "
		}
		title += country
	}

	if continent := data.Get("continent").String(); continent != "" {
		if title != "" {
			title += ", "
		}
		title += continent
	}

	subtitle := data.Get("timezone").String()

	if l := data.Get("latitude").Float(); l != 0 {
		if subtitle != "" {
			subtitle += "; "
		}

		subtitle += fmt.Sprintf("(%.3f,%.3f)",
			l, data.Get("longitude").Float())
	}

	if asn := data.Get("asn").Int(); asn != 0 {
		if subtitle != "" {
			subtitle += "; "
		}
		subtitle += fmt.Sprintf("ASN: %d", asn)
	}

	items := workflow.Items{
		workflow.Item{
			Title:    title,
			Subtitle: subtitle,
			Valid:    true,
			Arg:      jsonPrettyPrint(resp),
			Icon: workflow.Icon{
				Path: "./assets/ip2geo.png",
			},
			IconType: workflow.IconTypeLocal,
		},
	}

	items.Display()
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
