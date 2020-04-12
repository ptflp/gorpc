package rpca

import (
	"github.com/ptflp/gorpc"
	"github.com/ptflp/gorpc/repository"
)

type API struct {
	PageRepository repository.PageRepository
}

func (a *API) GetAll(empty string, reply *[]gorpc.Page) error {
	_ = empty
	*reply = a.PageRepository.GetAll()
	return nil
}

func (a *API) GetByTitle(title string, reply *gorpc.Page) error {
	res, err := a.PageRepository.GetByTitle(title)

	if err == nil {
		reply = &res[0]
	}

	return nil
}

func (a *API) AddItem(item gorpc.Page, reply *gorpc.Page) error {
	*reply = a.PageRepository.Create(item)
	return nil
}

func (a *API) EditItem(item gorpc.Page, reply *gorpc.Page) error {
	res, err := a.PageRepository.Update(item)
	if err == nil {
		*reply = res
	}
	return nil
}

func (a *API) DeleteItem(item gorpc.Page, reply *gorpc.Page) error {
	*reply, _ = a.PageRepository.Delete(item)
	return nil
}

func NewAPI(pageRepository *repository.PageRepository) *API {
	return &API{PageRepository: *pageRepository}
}
