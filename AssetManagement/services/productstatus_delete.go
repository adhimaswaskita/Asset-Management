package services

//DeleteProductStatus is bussiness logic for delete product Status by it's ID
func (s *Service) DeleteProductStatus(ID uint) error {
	err := s.Repository.DeleteProductStatus(ID)
	if err != nil {
		return err
	}
	return nil
}
