package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"encoding/json"
)

var (
	deployment DeploymentTarget
	loaded     bool
	options    *Options
)

// IsLoaded returns true if the configuration has been loaded from file
func IsLoaded() bool {
	return loaded
}

// Get returns the current Options pointer, or nil, from what was loaded
func Get() *Options {
	return options
}

// IsDevelopment returns true if the deployment target is set to development
func IsDevelopment() bool {
	return deployment == DeploymentTargetDevelopment
}

// SetDeploymentTarget sets the internal deployment target variable.
// Performs no loading of any files, just validates the value and sets the variable
func SetDeploymentTarget(target DeploymentTarget) error {
	if target < DeploymentTargetDevelopment || target > DeploymentTargetProduction {
		return errors.New("invalid deployment target provided to config.SetDeploymentTarget()")
	}

	deployment = target
	return nil
}

// LoadFile directly loads a path given and unmarshals the JSON into an Options object.
// Makes no assumptions on the deployment target, nor any path building.
func LoadFile(path string) (*Options, error) {
	loaded = false
	if len(path) == 0 {
		return nil, errors.New("provided file path for config.LoadFile() was empty")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load the configuration file %s; error = %s", path, err.Error())
	}

	options = &Options{}
	err = json.Unmarshal(data, options)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the configuration file %s; error = %s", path, err.Error())
	}

	loaded = true

	return options, nil
}

// Load attempts to load in the configuration options using a multitude of different processes.
// First, it will run the LoadFile() function to build the config file path and attempt to load it via JSON.
// Secondly, it will run the LoadEnv() function to modify the values based on environment variables.
// If either of these steps fail, no error is currently returned
func Load(target DeploymentTarget, pathPrefix string) (*Options, error) {
	if target < DeploymentTargetDevelopment || target > DeploymentTargetProduction {
		return nil, errors.New("invalid deployment target provided to config.SetDeploymentTarget()")
	}

	deployment = target

	fullPath := pathPrefix + deployment.GetSuffix() + ".json"

	opts, err := LoadFile(fullPath)
	if err != nil {
		//Failure to load the file will now be a warning since we have new ways to complete the process
		fmt.Printf("config.Load() failed at retrieving options from file; error = %s\n", err.Error())
	}

	if err = LoadEnv(opts); err != nil {
		fmt.Printf("config.Load() failed at retrieving options form environment; error = %s\n", err.Error())
	}

	//TODO: Check that the config is valid and use that as the error return

	return opts, nil
}
