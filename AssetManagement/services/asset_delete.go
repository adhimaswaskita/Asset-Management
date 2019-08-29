package services

//DeleteAsset is bussiness logic for delete Asset  by it's ID
func (s *Service) DeleteAsset(ID uint) error {
	err := s.Repository.DeleteAsset(ID)
	if err != nil {
		return err
	}
	return nil
}
