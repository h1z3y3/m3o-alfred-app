package i18n

import (
	"os"
	"strings"
)

type translation map[string]string

var i18nList map[string]translation

func init() {
	i18nList = map[string]translation{
		"zh_CN": zhCh,
	}
}

func I(message string) string {
	country := strings.TrimSpace(os.Getenv("i18n"))

	if _, ok := i18nList[country]; !ok {
		return message
	}

	if _, ok := i18nList[country][message]; !ok {
		return message
	}

	return i18nList[country][message]
}
