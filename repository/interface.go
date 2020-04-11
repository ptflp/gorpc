package repository

import "github.com/ptflp/gorpc"

type Repository interface {
	Create(item gorpc.Page) gorpc.Page
	Read(idx int) (gorpc.Page, error)
	Update(title string, edit gorpc.Page) (gorpc.Page, error)
	Delete(item gorpc.Page) (gorpc.Page, error)
}
