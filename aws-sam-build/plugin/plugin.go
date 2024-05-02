// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/shlex"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	BuildImage              string `envconfig:"PLUGIN_BUILD_IMAGE"`
	TemplateFilePath        string `envconfig:"PLUGIN_TEMPLATE_FILE_PATH"`
	BuildCommandOptions     string `envconfig:"PLUGIN_BUILD_COMMAND_OPTIONS"`
	PrivateRegistryURL      string `envconfig:"PLUGIN_PRIVATE_REGISTRY_URL"`
	PrivateRegistryUsername string `envconfig:"PLUGIN_PRIVATE_REGISTRY_USERNAME"`
	PrivateRegistryPassword string `envconfig:"PLUGIN_PRIVATE_REGISTRY_PASSWORD"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	if args.TemplateFilePath == "" {
		if _, err := os.Stat("./template.yaml"); err == nil {
			args.TemplateFilePath = "./template.yaml"
		} else if _, err := os.Stat("./template.yml"); err == nil {
			args.TemplateFilePath = "./template.yml"
		}
	}

	if args.PrivateRegistryURL != "" {
		if err := loginToPrivateRegistry(args); err != nil {
			return err
		}
	}

	// Create the initial command slice
	commandArgs := []string{"sam", "build"}

	// Include build image option if provided
	if args.BuildImage != "" {
		commandArgs = append(commandArgs, "--use-container", "--build-image", args.BuildImage)
	}

	// Include the template file path
	commandArgs = append(commandArgs, "--template-file", args.TemplateFilePath)

	// Add the build command options by parsing them correctly respecting quotes
	if args.BuildCommandOptions != "" {
		options, err := shlex.Split(args.BuildCommandOptions)
		if err != nil {
			return fmt.Errorf("error parsing build command options: %v", err)
		}
		commandArgs = append(commandArgs, options...)
	}

	cmd := exec.CommandContext(ctx, commandArgs[0], commandArgs[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %v\nOutput:\n%s", err, output)
	}

	fmt.Println(string(output))

	return nil
}

func loginToPrivateRegistry(args Args) error {
	if args.PrivateRegistryUsername == "" || args.PrivateRegistryPassword == "" {
		return fmt.Errorf("private registry credentials not provided")
	}
	cmd := exec.Command("docker", "login", args.PrivateRegistryURL, "--username", args.PrivateRegistryUsername, "--password", args.PrivateRegistryPassword)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command: %v\nOutput:\n%s", err, output)
	}

	fmt.Println(string(output))
	return nil

}
