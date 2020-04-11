package repository

type Repository interface {
	Create(item Page) Page
	Read(idx int) (Page, error)
	Update(title string, edit Page) (Page, error)
	Delete(item Page) (Page, error)
}
