package course

type Service interface {
	GetCourses(userID int, division string) ([]Course, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCourses(userID int, division string) ([]Course, error) {
	if userID != 0 {
		courses, err := s.repository.GetByUserID(userID)
		if err != nil {
			return courses, err
		}
		return courses, nil
	}

	if division != "" {
		courses, err := s.repository.GetByDivision(division)
		if err != nil {
			return courses, err
		}
		return courses, nil
	}

	courses, err := s.repository.GetAll()
	if err != nil {
		return courses, err
	}

	return courses, nil
}
