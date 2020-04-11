package main

import (
	"errors"
	"fmt"
)

const NotFound = "not found"

type Page struct {
	title string
	body  string
}

type Repository interface {
	GetByName(title string) (Page, error)
	AddItem(item Page) Page
	EditItem(title string, edit Page) (Page, error)
	DeleteItem(item Page) (Page, error)
}

type PageRepository struct {
	storage []Page
}

func NewPageRepository() Repository {
	return &PageRepository{}
}

func (d *PageRepository) GetByName(title string) (Page, error) {
	for i := range d.storage {
		if d.storage[i].title == title {
			return d.storage[i], nil
		}
	}

	return Page{}, errors.New(NotFound)
}

func (d *PageRepository) AddItem(item Page) Page {
	d.storage = append(d.storage, item)

	return item
}

func (d *PageRepository) EditItem(title string, edit Page) (Page, error) {
	for i := range d.storage {
		if d.storage[i].title == title {
			d.storage[i] = edit
			return edit, nil
		}
	}

	return Page{}, errors.New(NotFound)
}

func (d *PageRepository) DeleteItem(item Page) (Page, error) {
	for i := range d.storage {
		if d.storage[i].title == item.title && d.storage[i].body == item.body {
			d.storage = append(d.storage[:i], d.storage[i+1:]...)

			return item, nil
		}
	}

	return Page{}, errors.New(NotFound)
}

func main() {
	pageRepository := NewPageRepository()
	fmt.Println(pageRepository)
}
