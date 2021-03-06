package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

// Config represents the composition of yml settings.
type Config struct {
	Database struct {
		Dialect   string `default:"sqlite3"`
		Host      string `default:"book.db"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"false"`
	}
	Redis struct {
		Enabled            bool `default:"false"`
		ConnectionPoolSize int  `yaml:"connection_pool_size" default:"10"`
		Host               string
		Port               string
	}
	Extension struct {
		MasterGenerator bool `yaml:"master_generator" default:"false"`
		CorsEnabled     bool `yaml:"cors_enabled" default:"false"`
		SecurityEnabled bool `yaml:"security_enabled" default:"false"`
	}
	Log struct {
		RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status}"`
	}
	Security struct {
		ExculdePath []string `yaml:"exclude_path"`
		UserPath    []string `yaml:"user_path"`
		AdminPath   []string `yaml:"admin_path"`
	}
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

var config *Config
var env *string

// Load reads the settings written to the yml file
func Load() {
	env = flag.String("env", "develop", "To switch configurations.")
	flag.Parse()
	config = &Config{}
	if err := configor.Load(config, "application."+*env+".yml"); err != nil {
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(2)
	}
}

// GetConfig returns the configuration data.
func GetConfig() *Config {
	return config
}

// SetConfig sets configuration data.
func SetConfig(conf *Config) {
	config = conf
}

// GetEnv returns the environment variable.
func GetEnv() *string {
	return env
}
