package services_test

type MockRepository struct {
	ErrMap       map[string]bool
	ErrStatement string
}
