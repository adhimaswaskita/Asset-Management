package services

//DeleteOrganizationRegion is business logic for delete OrganizationRegion
func (s *Service) DeleteOrganizationRegion(ID uint) error {
	err := s.Repository.DeleteOrganizationRegion(ID)

	if err != nil {
		return err
	}
	return nil
}
