package services

type itemsService struct {
}

func ItemsService() *itemsService {
	return &itemsService{}
}

type itemServiceInterface interface {
	GetItem()
	SaveItem()
}

func (i *itemsService) GetItem() {
	panic("implement me")
}

func (i *itemsService) SaveItem() {
	panic("implement me")
}
