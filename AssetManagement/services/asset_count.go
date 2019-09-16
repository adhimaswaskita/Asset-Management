package services

//CountAsset is bussiness logic for get one Asset from database
func (s *Service) CountAsset() (int, error) {
	total, err := s.Repository.CountAsset()
	if err != nil {
		return -1, err
	}

	return total, nil
}
