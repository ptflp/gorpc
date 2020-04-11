package repository

import (
	"../gorpc"
	"errors"
)

const NotFound = "not found"

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

func (d *PageRepository) Read(idx int) (Page, error) {
	if idx >= 0 || idx < len(d.storage) {
		return d.storage[idx], nil
	}

	return Page{}, errors.New(NotFound)
}

func (d *PageRepository) Create(item Page) Page {
	d.storage = append(d.storage, item)

	return item
}

func (d *PageRepository) Update(title string, edit Page) (Page, error) {
	for i := range d.storage {
		if d.storage[i].title == title {
			d.storage[i] = edit
			return edit, nil
		}
	}

	return Page{}, errors.New(NotFound)
}

func (d *PageRepository) Delete(item Page) (Page, error) {
	for i := range d.storage {
		if d.storage[i].title == item.title && d.storage[i].body == item.body {
			d.storage = append(d.storage[:i], d.storage[i+1:]...)

			return item, nil
		}
	}

	return Page{}, errors.New(NotFound)
}
