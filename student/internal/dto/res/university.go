package res

import "student/internal/model"

type UniversityDto struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	EstablishmentYear int    `json:"establishment_year"`
}

func (u *UniversityDto) FromModel(university model.University) {
	u.Id = university.ID
	u.Name = university.Name
	u.EstablishmentYear = university.EstablishmentYear
}
