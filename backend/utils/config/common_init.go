package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/utils/logger"
)

var Config *config
var configLogger *logger.Logger = logger.NewLogger(
	"config",
	nil,
)
