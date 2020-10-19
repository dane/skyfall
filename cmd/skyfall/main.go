package main

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/dane/skyfall/service/web"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	config, err := web.LoadConfig()
	if err != nil {
		logger.Fatal("failed to parse config", zap.Error(err))
	}

	w, err := web.New(
		config,
		web.SetGetSignInHandler(web.DefaultGetSignInHandler),
		web.SetPostSignInHandler(web.DefaultPostSignInHandler),
		web.SetRender(web.DefaultRender(config.TemplatePath)),
		web.SetLogger(logger),
	)

	if err != nil {
		logger.Fatal("failed to initialize web service", zap.Error(err))
	}

	logger.Info("starting web service", zap.String("address", config.HTTPAddr))
	http.ListenAndServe(config.HTTPAddr, w)
}
