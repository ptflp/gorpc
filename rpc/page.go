package rpc

import (
	"github.com/ptflp/gorpc"
	"github.com/ptflp/gorpc/repository"
)

type RPC struct {
	pageRepository repository.PageRepository
}

func (a *RPC) GetAll(empty string, reply *[]gorpc.Page) error {
	_ = empty
	*reply = a.pageRepository.GetAll()
	return nil
}

func (a *RPC) GetByTitle(title string, reply *gorpc.Page) error {
	res, err := a.pageRepository.GetByTitle(title)

	if err == nil {
		reply = &res[0]
	}

	return nil
}

func (a *RPC) AddItem(item gorpc.Page, reply *gorpc.Page) error {
	*reply = a.pageRepository.Create(item)
	return nil
}

func (a *RPC) EditItem(item gorpc.Page, reply *gorpc.Page) error {
	res, err := a.pageRepository.Update(item)
	if err == nil {
		*reply = res
	}
	return nil
}

func (a *RPC) DeleteItem(item gorpc.Page, reply *gorpc.Page) error {
	*reply, _ = a.pageRepository.Delete(item)
	return nil
}
