package services

//DeleteProductSupplier is bussiness logic for delete product supplier by it's ID
func (s *Service) DeleteProductSupplier(ID uint) error {
	err := s.Repository.DeleteProductSupplier(ID)
	if err != nil {
		return err
	}
	return nil
}
