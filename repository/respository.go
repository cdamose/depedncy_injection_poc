package repository

import "github/cdamose/depedncy_injection_poc/dao"

type Repository interface {
	GetSampleData() (*dao.SampleData, error)
}
