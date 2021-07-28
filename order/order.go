package order

import "errors"

type order struct {
	id string
}

func New(id string) (*order, error) {
	if id == "" {
		return &order{}, errors.New("ID must not be empty")
	}

	return &order{id: id}, nil
}

func (o order) GetId() string {
	return o.id
}

func (o *order) SetId(id string) error {
	if id == "" {
		return errors.New("not a valid id")
	}

	o.id = id
	return nil
}
