package mockdatabase

func NewMockConnection() (*OrderRepo, error) {
	inv := OrderRepo{
		List: []Order{},
	}

	return &inv, nil
}
