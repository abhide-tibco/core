package engine

import (
	"os"
	"strconv"
	"strings"
)

const (
	EnvEnableSchemaValidation = "FLOGO_SCHEMA_VALIDATION"
	EnvKeyEnvName             = "FLOGO_ENV"
	EnvAppPropertyResolvers   = "FLOGO_APP_PROP_RESOLVERS"
	EnvKeyRunnerType          = "FLOGO_RUNNER_TYPE"
	DefaultRunnerType         = ValueRunnerTypePooled
	ValueRunnerTypePooled     = "POOLED"
	ValueRunnerTypeDirect     = "DIRECT"
	EnvKeyRunnerWorkers       = "FLOGO_RUNNER_WORKERS"
	DefaultRunnerWorkers      = 5
)

func IsSchemaValidationEnabled() bool {
	schemaValidationEnv := os.Getenv(EnvEnableSchemaValidation)
	return strings.EqualFold(schemaValidationEnv, "true")
}

// GetEnvName returns the name of the environment e.g. dev, test, prod
func GetEnvName() string {
	return os.Getenv(EnvKeyEnvName)
}

// GetRunnerType returns the runner type
func GetRunnerType() string {
	runnerTypeEnv := os.Getenv(EnvKeyRunnerType)
	if len(runnerTypeEnv) > 0 {
		return runnerTypeEnv
	}
	return DefaultRunnerType
}

// GetRunnerWorkers returns the number of workers to use
func GetRunnerWorkers() int {
	numWorkers := DefaultRunnerWorkers
	workersEnv := os.Getenv(EnvKeyRunnerWorkers)
	if len(workersEnv) > 0 {
		i, err := strconv.Atoi(workersEnv)
		if err == nil {
			numWorkers = i
		}
	}
	return numWorkers
}
