package services

//DeleteProduct is bussiness logic for delete product  by it's ID
func (s *Service) DeleteProduct(ID uint) error {
	err := s.Repository.DeleteProduct(ID)
	if err != nil {
		return err
	}
	return nil
}
