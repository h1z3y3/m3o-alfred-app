package workflow

const (
	IconTypeLocal = "local"
	IconTypeUrl   = "url"
)

type Icon struct {
	Path string `json:"path"`
}

type Item struct {
	Title        string `json:"title,omitempty"`
	Subtitle     string `json:"subtitle,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
	Uid          string `json:"uid,omitempty"`
	Arg          string `json:"arg,omitempty"`
	Autocomplete string `json:"autocomplete,omitempty"`
	Type         string `json:"type,omitempty"`
	Icon         Icon   `json:"icon,omitempty"`
	IconType     string `json:"-"`
}

func (i Item) SetTitle(t string) Item {
	i.Title = t
	return i
}

func (i Item) SetSubtitle(t string) Item {
	i.Subtitle = t
	return i
}

func (i Item) SetValid(v bool) Item {
	i.Valid = v
	return i
}

func (i Item) SetUid(u string) Item {
	i.Uid = u
	return i
}

func (i Item) SetArg(a string) Item {
	i.Arg = a
	return i
}

func (i Item) SetAutocomplete(a string) Item {
	i.Autocomplete = a
	return i
}

func (i Item) SetType(t string) Item {
	i.Type = t
	return i
}

func (i Item) SetIcon(c Icon) Item {
	i.Icon = c
	return i
}

func (i Item) SetIconType(c string) Item {
	i.IconType = c
	return i
}
