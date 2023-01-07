package typgo

import (
	"github.com/kmrtftech/tg-framework/pkg/envkit"
)

type (
	// EnvLoader responsible to load env
	EnvLoader interface {
		EnvLoad() (map[string]string, error)
	}
	// DotEnv file
	DotEnv string
	// Environment variable
	Environment map[string]string
)

//
// DotEnv
//

var _ EnvLoader = (DotEnv)("")

// EnvLoad load environment from dotenv file
func (d DotEnv) EnvLoad() (map[string]string, error) {
	return envkit.ReadFile(string(d))
}

//
// Environments
//

var _ EnvLoader = (Environment)(nil)

// EnvLoad load environment from dotenv file
func (e Environment) EnvLoad() (map[string]string, error) {
	return e, nil
}
