package config

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/logger"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/transport/http"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/transport/rpc"
	"github.com/theahmadzafar/resilient-order-processing-system/services/proto/inventry"
)

var config *Config
var once sync.Once

type Config struct {
	RPC      *rpc.Config      `validate:"required"`
	Server   *http.Config     `validate:"required"`
	Inventry *inventry.Config `validate:"required"`
	Logger   *logger.Config
}

func New() (*Config, error) {
	var initErr error

	once.Do(func() {
		config = &Config{}
		v := viper.New()
		v.AddConfigPath("./services/order-service")
		v.SetConfigName("config")

		if err := v.ReadInConfig(); err != nil {
			initErr = fmt.Errorf("error reading config: %w", err)

			return
		}

		if err := v.Unmarshal(&config); err != nil {
			initErr = fmt.Errorf("error unmarshalling config: %w", err)

			return
		}

		if err := parseSubConfig("rpc", &config.RPC, v); err != nil {
			initErr = err

			return
		}

		if v.Sub("logger") == nil {
			config.Logger = &logger.Config{
				LogLevel: "info",
			}
		} else {
			if initErr = parseSubConfig("logger", &config.Logger, v); initErr != nil {
				return
			}
		}

		if err := validator.New().Struct(config); err != nil {
			initErr = handleValidationError(err)

			return
		}

		if err := parseSubConfig("inventry", &config.Inventry, v); err != nil {
			initErr = err

			return
		}
	})

	return config, initErr
}

func parseSubConfig[T any](key string, parseTo *T, v *viper.Viper) error {
	subConfig := v.Sub(key)
	if subConfig == nil {
		return fmt.Errorf("can not read %s config: subconfig is nil", key)
	}

	if err := subConfig.Unmarshal(parseTo); err != nil {
		return fmt.Errorf("error unmarshalling %s config: %w", key, err)
	}

	return nil
}

func handleValidationError(err error) error {
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		var errStr strings.Builder
		for _, e := range validationErrs {
			errStr.WriteString(fmt.Sprintf("validation failed for field '%s': %s\n", e.Field(), e.ActualTag()))
		}

		return fmt.Errorf("config validation failed: %s", errStr.String())
	}

	return fmt.Errorf("error validating config: %w", err)
}
