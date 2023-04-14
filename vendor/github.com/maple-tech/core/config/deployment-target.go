package config

import "errors"

//DeploymentTarget is an ENUM-style specifier for which config option should be used when loading/starting
//the service. This value directly corrisponds to the config version used, by filename suffix.
//
//Constant variables are provided with the values set for easier use.
//Values available:
// 0 = Development: Local development options only
// 1 = Beta: Production testing and beta usage. Designed for real-world, but while still logging more.
// 2 = Production: Full production version. Limits logging to important information.
type DeploymentTarget byte

const (
	//DeploymentTargetDevelopment (0) will use 'config.dev.json' as the configuration source
	DeploymentTargetDevelopment DeploymentTarget = iota

	//DeploymentTargetBeta (1) will use 'config.beta.json' as the configuration source
	DeploymentTargetBeta

	//DeploymentTargetProduction (2) will use 'config.json' as the configuration source
	DeploymentTargetProduction
)

func (d DeploymentTarget) String() string {
	if d == DeploymentTargetProduction {
		return "production"
	} else if d == DeploymentTargetBeta {
		return "beta"
	}
	return "development"
}

//UnmarshalText implements the Text unmarshalling for converting string versions into their enum value
func (d *DeploymentTarget) UnmarshalText(src []byte) error {
	str := string(src)
	switch str {
	case "production":
		*d = DeploymentTargetProduction
	case "beta":
		*d = DeploymentTargetBeta
	case "development":
		*d = DeploymentTargetDevelopment
	default:
		return errors.New("failed to unmarshal DeploymentTarget, expected one of {development|beta|production}")
	}

	return nil
}

//GetSuffix returns the extension suffix that postfixes the file path, and pre-fixes the actual JSON extension.
func (d DeploymentTarget) GetSuffix() string {
	switch d {
	case DeploymentTargetDevelopment:
		return ".dev"
	case DeploymentTargetBeta:
		return ".beta"
	default:
		return ""
	}
}
