package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	LogLevel string
	Etcd     []string
	Consul   string
}

var C Config

func init() {
	ReadConfig()
}

// ReadConfig reads config from env
func ReadConfig() {
	C = Config{
		LogLevel: GetEnvOrDefault("LOG_LEVEL", "warn"),
		Etcd:     GetEnvSlice("ETCD"),
		Consul:   GetEnvOrDefault("CONSUL", ""),
	}
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value, has := os.LookupEnv(key); has {
		return value
	} else {
		return defaultValue
	}
}

func GetEnvOrDefaultInt(key string, defaultValue int) int {
	if value, has := os.LookupEnv(key); has {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Println("failed to parse env ", key, "use default:", defaultValue)
			return defaultValue
		}
		return num
	} else {
		return defaultValue
	}
}

func GetEnvSlice(key string) []string {
	if value, has := os.LookupEnv(key); has {
		values := strings.Split(value, ",")
		if len(values) == 1 && values[0] == "" {
			return nil
		}
		return values
	} else {
		return []string{}
	}
}

func GetEnvSliceOrDefault(key string, defaultValue string) []string {
	if value, has := os.LookupEnv(key); has {
		values := strings.Split(value, ",")
		if len(values) == 1 && values[0] == "" {
			return nil
		}
		return values
	} else {
		return []string{defaultValue}
	}
}
