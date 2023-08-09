package product

import (
	"errors"
	"strings"
)

type Product struct {
	id          int
	name        string
	description string
	category    string
	obs         string
	enable      bool
}

func (p *Product) GetId() int {
	return p.id
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) GetCategory() string {
	return p.category
}

func (p *Product) GetObs() string {
	return p.obs
}

func (p *Product) IsEnable() bool {
	return p.enable
}

func (p *Product) SetID(id int) {
	p.id = id
}

func (p *Product) SetName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("name cant be empty")
	}
	p.name = name
	return nil
}
func (p *Product) SetDescription(description string) {
	p.description = description
}

func (p *Product) SetCategory(category string) error {
	category = strings.TrimSpace(category)

	if category == "" {
		return errors.New("category cant be empty")
	}
	p.category = category
	return nil
}

func (p *Product) SetObs(obs string) {
	p.obs = obs
}

func (p *Product) SetEnable(isEnable bool) {
	p.enable = isEnable
}

func NewProduct(name, description, category, obs string) (*Product, error) {

	name = strings.TrimSpace(name)
	category = strings.TrimSpace(category)

	if name == "" {
		return nil, errors.New("name cant be empty")
	}
	if category == "" {
		return nil, errors.New("category cant be empty")
	}

	product := Product{
		name:        name,
		description: description,
		category:    category,
		obs:         obs,
		enable:      false,
	}

	return &product, nil
}
