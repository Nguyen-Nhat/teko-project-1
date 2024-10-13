package repository

type IStudentRepository interface {
}

type studentRepository struct{}

func NewStudentRepository() IStudentRepository {
	return &studentRepository{}
}
