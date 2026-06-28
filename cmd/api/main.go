package main

import (
	"log"

	"github.com/Wrehat/E-wallet-framework/internal/config"
	"github.com/Wrehat/E-wallet-framework/pkg/logger"
	"go.uber.org/zap"
)

func main() {

	cfg := config.SetupConfig()

	logger, err := logger.NewLogger(cfg.AppEnv)
	if err != nil {
		log.Fatalf("Gagal inisialisasi logger: %v", err)
	}

	defer logger.Sync()

	logger.Info("Aplikasi E-Wallet berhasil berjalan!",
		zap.String("APP_ENV", cfg.AppEnv),
		zap.String("APP_PORT", cfg.AppPort),
		zap.String("APP_GRPC_PORT", cfg.AppGrpcPort),
	)

}
