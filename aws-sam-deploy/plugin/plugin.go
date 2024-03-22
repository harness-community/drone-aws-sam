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

	AWSAccessKey         string `envconfig:"PLUGIN_AWS_ACCESS_KEY"`
	AWSSecretKey         string `envconfig:"PLUGIN_AWS_SECRET_KEY"`
	AWSRegion            string `envconfig:"PLUGIN_AWS_REGION"`
	TemplateFilePath     string `envconfig:"PLUGIN_TEMPLATE_FILE_PATH"`
	StackName            string `envconfig:"PLUGIN_STACK_NAME"`
	S3Bucket             string `envconfig:"PLUGIN_S3_BUCKET"`
	DeployCommandOptions string `envconfig:"PLUGIN_DEPLOY_COMMAND_OPTIONS"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	if err := verifyArgs(args); err != nil {
		return err
	}

	os.Setenv("AWS_ACCESS_KEY_ID", args.AWSAccessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", args.AWSSecretKey)

	cmd := exec.Command("sam", "deploy", "--region", args.AWSRegion, "--template-file", args.TemplateFilePath, "--stack-name", args.StackName, "--s3-bucket", args.S3Bucket, args.DeployCommandOptions, "--no-confirm-changeset")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("error executing command: %v\nOutput:\n%s", err, output)
	}

	fmt.Println(string(output))
	return nil
}
