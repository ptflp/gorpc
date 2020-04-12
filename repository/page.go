package repository

import (
	"errors"
	"github.com/ptflp/gorpc"
	"unicode/utf8"
)

const NotFound = "not found"

type PageRepository struct {
	storage []gorpc.Page
}

func NewPageRepository() Repository {
	return &PageRepository{}
}

func (d *PageRepository) GetAll() []gorpc.Page {
	return d.storage
}

func (d *PageRepository) Read(p gorpc.Page) ([]gorpc.Page, error) {
	var res []gorpc.Page
	for i := range d.storage {
		if p.ID == d.storage[i].ID {
			p.ID = i
			return []gorpc.Page{p}, nil
		}
		if p.Title == d.storage[i].Title && p.Body == d.storage[i].Body {
			p.ID = i
			res = append(res, p)
			continue
		}
		if p.Title == d.storage[i].Title && utf8.RuneCountInString(p.Body) == 0 {
			p.ID = i
			res = append(res, p)
			continue
		}
		if p.Body == d.storage[i].Body && utf8.RuneCountInString(p.Title) == 0 {
			p.ID = i
			res = append(res, p)
			continue
		}
	}

	if len(res) > 0 {
		return res, nil
	}

	return res, errors.New(NotFound)
}

func (d *PageRepository) Create(p gorpc.Page) gorpc.Page {
	p.ID = len(d.storage)
	d.storage = append(d.storage, p)

	return p
}

func (d *PageRepository) Update(p gorpc.Page) (gorpc.Page, error) {
	for i := range d.storage {
		if d.storage[i].ID == p.ID {
			d.storage[i] = p
			return p, nil
		}
	}

	return gorpc.Page{}, errors.New(NotFound)
}

func (d *PageRepository) Delete(item gorpc.Page) (gorpc.Page, error) {
	for i := range d.storage {
		if d.storage[i].Title == item.Title && d.storage[i].Body == item.Body {
			d.storage = append(d.storage[:i], d.storage[i+1:]...)

			return item, nil
		}
	}

	return gorpc.Page{}, errors.New(NotFound)
}

func (d *PageRepository) GetByTitle(title string) ([]gorpc.Page, error) {
	res, err := d.Read(gorpc.Page{Title: title})

	if err == nil {
		return res, nil
	}

	return res, errors.New(NotFound)
}
