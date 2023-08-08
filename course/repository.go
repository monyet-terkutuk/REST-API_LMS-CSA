package courses

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Course, error)
	GetByDivision(division string) ([]Course, error)
	GetByUserID(userID int) ([]Course, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Course, error) {
	var courses []Course
	err := r.db.Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *repository) GetByDivision(division string) ([]Course, error) {
	var courses []Course
	err := r.db.Where("division = ?", division).Preload("CourseImages", "course_images.is_primary = 1").Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}

func (r *repository) GetByUserID(userID int) ([]Course, error) {
	var courses []Course
	err := r.db.Where("user_id = ?", userID).Preload("CourseImages", "course_images.is_primary = 1").Find(&courses).Error
	if err != nil {
		return courses, err
	}

	return courses, nil
}
