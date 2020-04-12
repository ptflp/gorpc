package repository

import "github.com/ptflp/gorpc"

type Repository interface {
	Create(p gorpc.Page) gorpc.Page
	Read(p gorpc.Page) ([]gorpc.Page, error)
	Update(p gorpc.Page) (gorpc.Page, error)
	Delete(p gorpc.Page) (gorpc.Page, error)
}
