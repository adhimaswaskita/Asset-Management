package services

//DeleteManufacture is business logic for delete Manufacture
func (s *Service) DeleteManufacture(ID uint) error {
	err := s.Repository.DeleteManufacture(ID)

	if err != nil {
		return err
	}
	return nil
}
