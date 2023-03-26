//go:build wireinject
// +build wireinject

package main

import (
	"github/cdamose/depedncy_injection_poc/config"
	"github/cdamose/depedncy_injection_poc/domain"
	"github/cdamose/depedncy_injection_poc/repository"
	"github/cdamose/depedncy_injection_poc/repository/adapter"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func InitializeDomain(logger logrus.Entry) (domain.UserDomain, error) {
	wire.Build(domain.NewUserDomain, adapter.NewInMemoryAdapter, config.InitConfig,
		wire.Bind(new(repository.Repository), new(*adapter.InMemoryRepository)))
	//wire.Build(config.InitConfig)
	return domain.UserDomain{}, nil
}
