package services


import (
	// "errors"
	"speakbuddy/internal/models"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/internal/repository"
)

type DataAnakService interface {
	CreateOrUpdate(userID uint, req request.CreateDataAnakRequest) (*models.DataAnak, error)
	Get(userID uint) (*models.DataAnak, error)
}


type dataAnakService struct {
	repo repository.DataAnakRepository
}


func NewDataAnakService(repo repository.DataAnakRepository) DataAnakService {
	return &dataAnakService{repo}
}


func (s *dataAnakService) CreateOrUpdate(userID uint, req request.CreateDataAnakRequest) (*models.DataAnak, error) {
	data, err := s.repo.FindByUserID(userID)
	if err != nil {
		newData := models.DataAnak{
			UserID: userID,
			ChildName: req.ChildName,
			ChildAge: req.ChildAge,
			ChildSex: req.ChildSex,
		}
		err = s.repo.Create(&newData)
		return &newData, err
	}


	data.ChildName = req.ChildName
	data.ChildAge = req.ChildAge
	data.ChildSex = req.ChildSex


	err = s.repo.Update(data)
	return data, err
}


func (s *dataAnakService) Get(userID uint) (*models.DataAnak, error) {
	return s.repo.FindByUserID(userID)
}