package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/utils/logger"
)

// Config stores the configuration of this program.
var Config *config = &config{}
var configLogger *logger.Logger = logger.NewLogger(
	"config",
	nil,
)
