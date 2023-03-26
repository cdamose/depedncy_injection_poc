package domain

import (
	"github/cdamose/depedncy_injection_poc/config"
	"github/cdamose/depedncy_injection_poc/dto"
	"github/cdamose/depedncy_injection_poc/repository"

	"github.com/sirupsen/logrus"
)

type UserDomain struct {
	logger     logrus.Entry
	config     config.Config
	repository repository.Repository
}

func NewUserDomain(logger logrus.Entry, config config.Config, repository repository.Repository) UserDomain {
	return UserDomain{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}

func (ud *UserDomain) GetUserDetails() (*dto.User, error) {
	data_obj, err := ud.repository.GetSampleData()
	dto_obj := &dto.User{}
	if err == nil {
		dto_obj.ID = data_obj.ID
		dto_obj.Name = data_obj.Name
	} else {
		return nil, err
	}
	return dto_obj, nil
}
