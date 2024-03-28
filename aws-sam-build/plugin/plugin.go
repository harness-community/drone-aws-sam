// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
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

	var cmd *exec.Cmd

	if args.BuildImage != "" {
		cmd = exec.Command("sam", "build", "--use-container", "--build-image", args.BuildImage, "--template-file", args.TemplateFilePath, args.BuildCommandOptions)
	} else {
		cmd = exec.Command("sam", "build", "--template-file", args.TemplateFilePath, args.BuildCommandOptions)
	}

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
