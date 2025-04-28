package mockdatabase

func NewMockConnection() (*OrderRepo, error) {
	inv := OrderRepo{}

	return &inv, nil
}
