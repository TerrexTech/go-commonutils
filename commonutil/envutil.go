package commonutil

import (
	"errors"
	"os"
)

// ValidateEnv checks if the provided environment variables are set to some value.
// Otherwise the missing variable's name is returned along with an error.
// If multiple variables are missing, this will return error on the first variable
// which was not found.
func ValidateEnv(envVars ...string) (string, error) {
	for _, varname := range envVars {
		envVar := os.Getenv(varname)
		if envVar == "" {
			err := errors.New(
				"Error while bootstrapping Cassandra table: " +
					"Following env-var is required but was not found: " +
					varname,
			)
			return varname, err
		}
	}
	return "", nil
}
