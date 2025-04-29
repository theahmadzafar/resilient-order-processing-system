package mockdatabase

func NewMockConnection() (*Repo, error) {
	repo := Repo{
		List: []Invoice{},
	}

	return &repo, nil
}
