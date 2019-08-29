package services

//DeleteOrganization is business logic for delete Organization
func (s *Service) DeleteOrganization(ID uint) error {
	err := s.Repository.DeleteOrganization(ID)

	if err != nil {
		return err
	}
	return nil
}
