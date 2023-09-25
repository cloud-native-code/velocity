package main

import (
	"github.com/cloud-native-code/velocity/config"
	"github.com/cloud-native-code/velocity/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	log := logger.Build(&logger.ProductionLogger{})
	defer log.Sync()

	config := config.BuildConfig(log)

	log.Info("Configured port", zap.Int("port", config.Server.Port))
}
