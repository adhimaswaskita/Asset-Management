package services

//DeleteOrganizationSite is business logic for delete OrganizationSite
func (s *Service) DeleteOrganizationSite(ID uint) error {
	err := s.Repository.DeleteOrganizationSite(ID)

	if err != nil {
		return err
	}
	return nil
}
