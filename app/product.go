package app

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

func (p *Product) IsValid() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Product) Enable() error {
	//TODO implement me
	panic("implement me")
}

func (p *Product) Disable() error {
	//TODO implement me
	panic("implement me")
}

func (p *Product) GetID() string {
	//TODO implement me
	panic("implement me")
}

func (p *Product) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (p *Product) GetStatus() string {
	//TODO implement me
	panic("implement me")
}

func (p *Product) GetPrice() float64 {
	//TODO implement me
	panic("implement me")
}
