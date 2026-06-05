package types

type Id int
type Book struct {
	Name          string `json:"name"`
	Content       string `json:"content"`
	DateOfRelease int    `json:"date_of_release"`
}

type Library map[Id]Book

func (l Library) FindBook(id int) *Book {
	x, y := l[Id(id)]
	if !y {
		return nil
	}

	return &x
}

func (l *Library) DeleteBook(id int) bool {
	_, y := (*l)[Id(id)]

	if !y {
		return false
	}

	delete((*l), Id(id))
	return true
}
