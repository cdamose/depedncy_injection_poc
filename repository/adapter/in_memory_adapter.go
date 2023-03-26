package adapter

import (
	"github/cdamose/depedncy_injection_poc/config"
	"github/cdamose/depedncy_injection_poc/dao"

	"github.com/sirupsen/logrus"
)

type InMemoryRepository struct {
	logger logrus.Entry
	config config.Config
}

func NewInMemoryAdapter(logger logrus.Entry, config config.Config) *InMemoryRepository {
	return &InMemoryRepository{
		logger: logger,
		config: config}
}

func (m InMemoryRepository) GetSampleData() (*dao.SampleData, error) {
	//mock data obj here
	data_obj := &dao.SampleData{
		ID:   "<dummy_id>",
		Name: "<sample_name>",
	}
	return data_obj, nil

}
