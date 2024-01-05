package main

import (
	"github.com/Fajar3108/go-boilerplate/cmd/api/server"
	"github.com/Fajar3108/go-boilerplate/internal/config"
	"github.com/Fajar3108/go-boilerplate/internal/constants"
	"github.com/Fajar3108/go-boilerplate/pkg/logger"
	"github.com/sirupsen/logrus"
	"runtime"
)

func init() {
	if err := config.IntializeAppConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}

	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	numCPU := runtime.NumCPU()
	logger.InfoF("The project is running on %d CPU(s)", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig}, numCPU)

	if numCPU > 2 {
		runtime.GOMAXPROCS(numCPU / 2)
	}

	err := server.NewApp()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	}
}
