package main

import (
	//"fmt"
	"fmt"
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

	domain_obj, _ := InitializeDomain(*logger)

	logger.WithFields(log.Fields{
		"confi_obj": domain_obj,
	}).Info("After config initialized")

	data, _ := domain_obj.GetUserDetails()
	fmt.Println(data)

}
