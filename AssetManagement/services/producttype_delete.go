package services

//DeleteProductType is bussiness logic for delete product type by it's ID
func (s *Service) DeleteProductType(ID uint) error {
	err := s.Repository.DeleteProductType(ID)
	if err != nil {
		return err
	}
	return nil
}
