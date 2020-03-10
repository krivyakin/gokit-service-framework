package main

import (
	"flag"
	kitzap "github.com/go-kit/kit/log/zap"
	"github.com/krivyakin/gokit-service-framework/pkg/config"
	httpserv "github.com/krivyakin/gokit-service-framework/pkg/http"
	"github.com/krivyakin/gokit-service-framework/pkg/log"
	"github.com/krivyakin/gokit-service-framework/pkg/service1"
	"github.com/krivyakin/gokit-service-framework/pkg/service1/implementation"
	"github.com/krivyakin/gokit-service-framework/pkg/service1/middleware"
	http_transport "github.com/krivyakin/gokit-service-framework/pkg/service1/transport/http"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func NewZapLogger() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	return zap.New(core)
}

func main() {
	var (
		configDir = flag.String("config_dir", "./cmd/service1/config", "directory with config files")
	)
	flag.Parse()

	logger := log.NewLogger(kitzap.NewZapSugarLogger(NewZapLogger(), zapcore.DebugLevel)).WithLocation("root")

	if err := config.InitConfig(*configDir); err != nil {
		logger.Errorm("can't load config file", "error", err)
		return
	}

	var service service1.Service
	{
		service = implementation.NewService(logger)
		service = middleware.NewLoggingMiddleware(logger)(service)
	}

	timeout := viper.GetDuration("http_server.timeout") * time.Millisecond
	router := httpserv.NewDefaultRouter(logger, timeout)
	http_transport.RegisterService(service, logger, router)

	{
		logger.Infom("service started")
		httpAddr := ":" + viper.GetString("http_server.port")
		srv := httpserv.NewServer(router, logger, httpAddr)
		logger.Infom("service stopped", "status", srv.Start())
	}
}
