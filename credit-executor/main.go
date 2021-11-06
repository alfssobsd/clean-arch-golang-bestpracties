package main

import (
	"clean-arch-golang-best-practices/credit-executor/utils/appconfig"
	"clean-arch-golang-best-practices/credit-library/loggerhelper"
	"clean-arch-golang-best-practices/credit-shared-module/utils/heavyprocessor"
)

func main() {
	logger := loggerhelper.MakeNewZapProductionLogger()
	appConfig := appconfig.NewAppConfigurationFromEnvFile(".env")
	heavyProcessor, _ := heavyprocessor.NewHeavyProcessor(logger.Sugar(), 10)

	go func() {
		MakeBackgroundTaskServer(logger.Sugar(), appConfig, heavyProcessor)
	}()
	MakeHttpServer(logger, appConfig, heavyProcessor)
}
