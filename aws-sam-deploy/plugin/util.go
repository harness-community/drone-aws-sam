// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "fmt"

func verifyArgs(args Args) error {
	if args.AWSAccessKey == "" {
		return fmt.Errorf("missing AWS access key")
	}
	if args.AWSSecretKey == "" {
		return fmt.Errorf("missing AWS secret key")
	}
	if args.AWSRegion == "" {
		return fmt.Errorf("missing AWS region")
	}
	if args.TemplateFilePath == "" {
		return fmt.Errorf("missing deploy template file path")
	}
	if args.StackName == "" {
		return fmt.Errorf("missing stack name")
	}
	if args.S3Bucket == "" {
		return fmt.Errorf("missing S3 bucket")
	}
	return nil
}
