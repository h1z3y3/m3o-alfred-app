package workflow

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Items []Item

func (is Items) Display() {
	wg := &sync.WaitGroup{}

	for k, val := range is {
		wg.Add(1)

		go func(item Item, k int) {
			defer wg.Done()

			// download the icon
			if item.IconType == IconTypeUrl {
				p, _ := NewCache(item.Icon.Path).Cache()
				is[k].Icon.Path = p
			}

		}(val, k)

	}

	wg.Wait()

	output := map[string]Items{
		"items": is,
	}

	bs, err := json.Marshal(output)

	if err != nil {
		NewError("结果解析失败", err.Error()).Display()
		return
	}

	fmt.Println(string(bs))
}
