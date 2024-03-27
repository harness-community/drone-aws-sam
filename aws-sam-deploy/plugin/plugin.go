// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"encoding/json"
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
	AWSSessionToken      string `envconfig:"PLUGIN_AWS_SESSION_TOKEN"`
	AWSStsExternalID     string `envconfig:"PLUGIN_AWS_STS_EXTERNAL_ID"`
	AWSRoleARN           string `envconfig:"PLUGIN_AWS_ROLE_ARN"`
	AWSRegion            string `envconfig:"PLUGIN_AWS_REGION"`
	TemplateFilePath     string `envconfig:"PLUGIN_TEMPLATE_FILE_PATH"`
	StackName            string `envconfig:"PLUGIN_STACK_NAME"`
	S3Bucket             string `envconfig:"PLUGIN_S3_BUCKET"`
	DeployCommandOptions string `envconfig:"PLUGIN_DEPLOY_COMMAND_OPTIONS"`
}

type AwsAssumeRoleOutput struct {
	Credentials struct {
		AccessKeyID     string `json:"AccessKeyId"`
		SecretAccessKey string `json:"SecretAccessKey"`
		SessionToken    string `json:"SessionToken"`
		Expiration      string `json:"Expiration"`
	} `json:"Credentials"`
	AssumedRoleUser struct {
		AssumedRoleID string `json:"AssumedRoleId"`
		Arn           string `json:"Arn"`
	} `json:"AssumedRoleUser"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	if err := verifyArgs(args); err != nil {
		return err
	}

	var cmd *exec.Cmd

	if args.AWSAccessKey != "" && args.AWSSecretKey != "" {
		if args.AWSSessionToken == "" && args.AWSRoleARN != "" {
			os.Setenv("AWS_ACCESS_KEY_ID", args.AWSAccessKey)
			os.Setenv("AWS_SECRET_ACCESS_KEY", args.AWSSecretKey)

			stsArgs := []string{"aws", "sts", "assume-role", "--role-arn", args.AWSRoleARN, "--role-session-name", "sam-deploy"}
			if args.AWSStsExternalID != "" {
				stsArgs = append(stsArgs, "--external-id", args.AWSStsExternalID)
			}
			stsCmd := exec.Command(stsArgs[0], stsArgs[1:]...)
			output, err := stsCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error executing sts assume-role command: %v\nOutput:\n%s", err, output)
			}
			var assumeRoleOutput map[string]interface{}
			if err := json.Unmarshal(output, &assumeRoleOutput); err != nil {
				return fmt.Errorf("error parsing sts assume-role output: %v\nOutput:\n%s", err, output)
			}

			credentials, ok := assumeRoleOutput["Credentials"].(map[string]interface{})
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: Credentials field not found or invalid format")
			}

			accessKeyID, ok := credentials["AccessKeyId"].(string)
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: AccessKeyId field not found or invalid format")
			}

			secretAccessKey, ok := credentials["SecretAccessKey"].(string)
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: SecretAccessKey field not found or invalid format")
			}

			sessionToken, ok := credentials["SessionToken"].(string)
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: SessionToken field not found or invalid format")
			}

			os.Setenv("AWS_ACCESS_KEY_ID", accessKeyID)
			os.Setenv("AWS_SECRET_ACCESS_KEY", secretAccessKey)
			os.Setenv("AWS_SESSION_TOKEN", sessionToken)

			cmd = exec.Command("sam", "deploy", "--region", args.AWSRegion, "--template-file", args.TemplateFilePath, "--stack-name", args.StackName, "--s3-bucket", args.S3Bucket, args.DeployCommandOptions, "--no-confirm-changeset")
		} else if args.AWSSessionToken != "" && args.AWSRoleARN != "" {

			os.Setenv("AWS_ACCESS_KEY_ID", args.AWSAccessKey)
			os.Setenv("AWS_SECRET_ACCESS_KEY", args.AWSSecretKey)
			os.Setenv("AWS_SESSION_TOKEN", args.AWSSessionToken)

			stsArgs := []string{"aws", "sts", "assume-role", "--role-arn", args.AWSRoleARN, "--role-session-name", "sam-deploy"}
			if args.AWSStsExternalID != "" {
				stsArgs = append(stsArgs, "--external-id", args.AWSStsExternalID)
			}
			stsCmd := exec.Command(stsArgs[0], stsArgs[1:]...)
			output, err := stsCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error executing sts assume-role command: %v\nOutput:\n%s", err, output)
			}
			var assumeRoleOutput map[string]interface{}
			if err := json.Unmarshal(output, &assumeRoleOutput); err != nil {
				return fmt.Errorf("error parsing sts assume-role output: %v\nOutput:\n%s", err, output)
			}

			credentials, ok := assumeRoleOutput["Credentials"].(map[string]interface{})
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: Credentials field not found or invalid format")
			}

			accessKeyID, ok := credentials["AccessKeyId"].(string)
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: AccessKeyId field not found or invalid format")
			}

			secretAccessKey, ok := credentials["SecretAccessKey"].(string)
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: SecretAccessKey field not found or invalid format")
			}

			sessionToken, ok := credentials["SessionToken"].(string)
			if !ok {
				return fmt.Errorf("error parsing sts assume-role output: SessionToken field not found or invalid format")
			}

			os.Unsetenv("AWS_ACCESS_KEY_ID")
			os.Unsetenv("AWS_SECRET_ACCESS_KEY")
			os.Unsetenv("AWS_SESSION_TOKEN")

			os.Setenv("AWS_ACCESS_KEY_ID", accessKeyID)
			os.Setenv("AWS_SECRET_ACCESS_KEY", secretAccessKey)
			os.Setenv("AWS_SESSION_TOKEN", sessionToken)

			cmd = exec.Command("sam", "deploy", "--region", args.AWSRegion, "--template-file", args.TemplateFilePath, "--stack-name", args.StackName, "--s3-bucket", args.S3Bucket, args.DeployCommandOptions, "--no-confirm-changeset")
		} else if args.AWSSessionToken != "" && args.AWSRoleARN == "" {
			os.Setenv("AWS_ACCESS_KEY_ID", args.AWSAccessKey)
			os.Setenv("AWS_SECRET_ACCESS_KEY", args.AWSSecretKey)
			os.Setenv("AWS_SESSION_TOKEN", args.AWSSessionToken)

			cmd = exec.Command("sam", "deploy", "--region", args.AWSRegion, "--template-file", args.TemplateFilePath, "--stack-name", args.StackName, "--s3-bucket", args.S3Bucket, args.DeployCommandOptions, "--no-confirm-changeset")
		} else {
			os.Setenv("AWS_ACCESS_KEY_ID", args.AWSAccessKey)
			os.Setenv("AWS_SECRET_ACCESS_KEY", args.AWSSecretKey)
			cmd = exec.Command("sam", "deploy", "--region", args.AWSRegion, "--template-file", args.TemplateFilePath, "--stack-name", args.StackName, "--s3-bucket", args.S3Bucket, args.DeployCommandOptions, "--no-confirm-changeset")
		}
	} else if args.AWSRoleARN != "" {
		return fmt.Errorf("error: AWS Access Key and Secret Key are required when using a Role ARN")
	} else if args.AWSAccessKey == "" && args.AWSSecretKey == "" && args.AWSRoleARN == "" {
		return fmt.Errorf("error: AWS credentials are required")
	}

	if args.AWSRoleARN != "" {
		stsCmd := exec.Command("aws", "sts", "assume-role-with-web-identity", "--role-arn", args.AWSRoleARN, "--role-session-name", "sam-deploy", "--web-identity-token", os.Getenv("AWS_WEB_IDENTITY_TOKEN_FILE"))
		output, err := stsCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error executing sts assume-role command: %v\nOutput:\n%s", err, output)
		}

		var assumeRoleOutput map[string]interface{}
		if err := json.Unmarshal(output, &assumeRoleOutput); err != nil {
			return fmt.Errorf("error parsing sts assume-role output: %v\nOutput:\n%s", err, output)
		}

		credentials, ok := assumeRoleOutput["Credentials"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("error parsing sts assume-role output: Credentials field not found or invalid format")
		}

		accessKeyID, ok := credentials["AccessKeyId"].(string)
		if !ok {
			return fmt.Errorf("error parsing sts assume-role output: AccessKeyId field not found or invalid format")
		}

		secretAccessKey, ok := credentials["SecretAccessKey"].(string)
		if !ok {
			return fmt.Errorf("error parsing sts assume-role output: SecretAccessKey field not found or invalid format")
		}

		sessionToken, ok := credentials["SessionToken"].(string)
		if !ok {
			return fmt.Errorf("error parsing sts assume-role output: SessionToken field not found or invalid format")
		}

		os.Unsetenv("AWS_ACCESS_KEY_ID")
		os.Unsetenv("AWS_SECRET_ACCESS_KEY")
		os.Unsetenv("AWS_SESSION_TOKEN")

		os.Setenv("AWS_ACCESS_KEY_ID", accessKeyID)
		os.Setenv("AWS_SECRET_ACCESS_KEY", secretAccessKey)
		os.Setenv("AWS_SESSION_TOKEN", sessionToken)

		cmd = exec.Command("sam", "deploy", "--region", args.AWSRegion, "--template-file", args.TemplateFilePath, "--stack-name", args.StackName, "--s3-bucket", args.S3Bucket, args.DeployCommandOptions, "--no-confirm-changeset")
	}

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("error executing command: %v\nOutput:\n%s", err, output)
	}

	fmt.Println(string(output))
	return nil
}
