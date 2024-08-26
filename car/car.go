package car

import "errors"

type Car struct {
	Name  string
	Price float32
}

func (c *Car) SetName(name string) (newName string) {
	if name != "" {
		c.Name = name
	}

	return c.Name
}

func New(name string, price float32) (car *Car, err error) {
	if name == "" {
		return nil, errors.New("missing name")
	}

	return &Car{
		Name:  name,
		Price: price,
	}, nil
}
