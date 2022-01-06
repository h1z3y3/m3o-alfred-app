package workflow

type Error Item

func NewError(title, subtitle string) Error {
	return Error{
		Title:    title,
		Subtitle: subtitle,
		Icon: Icon{
			Path: "./assets/error.png",
		},
	}
}

func (e Error) Display() {
	Items{Item(e)}.Display()
}
