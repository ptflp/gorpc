package repository

import (
	"errors"
	"github.com/ptflp/gorpc"
)

const NotFound = "not found"

type PageRepository struct {
	storage []gorpc.Page
}

func NewPageRepository() Repository {
	return &PageRepository{}
}

func (d *PageRepository) GetByName(title string) (gorpc.Page, error) {
	for i := range d.storage {
		if d.storage[i].title == title {
			return d.storage[i], nil
		}
	}

	return gorpc.Page{}, errors.New(NotFound)
}

func (d *PageRepository) Read(idx int) (gorpc.Page, error) {
	if idx >= 0 || idx < len(d.storage) {
		return d.storage[idx], nil
	}

	return gorpc.Page{}, errors.New(NotFound)
}

func (d *PageRepository) Create(item gorpc.Page) gorpc.Page {
	d.storage = append(d.storage, item)

	return item
}

func (d *PageRepository) Update(title string, edit gorpc.Page) (gorpc.Page, error) {
	for i := range d.storage {
		if d.storage[i].title == title {
			d.storage[i] = edit
			return edit, nil
		}
	}

	return gorpc.Page{}, errors.New(NotFound)
}

func (d *PageRepository) Delete(item gorpc.Page) (gorpc.Page, error) {
	for i := range d.storage {
		if d.storage[i].title == item.title && d.storage[i].body == item.body {
			d.storage = append(d.storage[:i], d.storage[i+1:]...)

			return item, nil
		}
	}

	return gorpc.Page{}, errors.New(NotFound)
}
