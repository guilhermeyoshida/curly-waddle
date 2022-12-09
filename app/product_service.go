package app

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (p *ProductService) Get(id string) (ProductInterface, error) {
	product, err := p.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) Save(product ProductInterface) (ProductInterface, error) {
	//TODO implement me
	panic("implement me")
}
