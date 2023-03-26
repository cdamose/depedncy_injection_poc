// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/sirupsen/logrus"
	"github/cdamose/depedncy_injection_poc/config"
	"github/cdamose/depedncy_injection_poc/domain"
	"github/cdamose/depedncy_injection_poc/repository/adapter"
)

// Injectors from wire.go:

func InitializeDomain(logger logrus.Entry) (domain.UserDomain, error) {
	configConfig := config.InitConfig(logger)
	inMemoryRepository := adapter.NewInMemoryAdapter(logger, configConfig)
	userDomain := domain.NewUserDomain(logger, configConfig, inMemoryRepository)
	return userDomain, nil
}
