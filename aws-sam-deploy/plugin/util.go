// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "fmt"

func verifyArgs(args Args) error {
	if args.AWSRegion == "" {
		return fmt.Errorf("please specify AWS Region")
	}

	if args.S3Bucket == "" {
		return fmt.Errorf("please specify AWS S3 Bucket")
	}

	if args.StackName == "" {
		return fmt.Errorf("please specify stack name")
	}

	if args.TemplateFilePath == "" {
		return fmt.Errorf("plaese specify template file path")
	}
	return nil
}
