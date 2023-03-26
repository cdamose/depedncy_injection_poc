package main

import (
	//"fmt"
	"github/cdamose/depedncy_injection_poc/common/logs"
	//"github/cdamose/depedncy_injection_poc/config"
	//"github/cdamose/depedncy_injection_poc/domain"
	//"github/cdamose/depedncy_injection_poc/repository/adapter"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	// logger parr initialization
	logs.Init()
	logger := logrus.NewEntry(logrus.StandardLogger())
	// config part initialization
	//cfg := config.InitConfig(logger)
	cfg, _ := InitializeEvent(*logger)
	
	 logger.WithFields(log.Fields{
	 	"confi_obj": cfg,
	 }).Info("After config initialized")

	// //Repository Initialization

	// repo := adapter.NewInMemoryAdapter(logger, cfg)

	// //service initilazation
	// domain := domain.NewUserDomain(logger, cfg, repo)

	// data, _ := domain.GetUserDetails()
	// fmt.Printf("%+v\n", data)

}
