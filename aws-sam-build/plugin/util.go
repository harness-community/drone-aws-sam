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

	if args.TemplateFilePath == "" {
		return fmt.Errorf("missing template file path")
	}
	return nil
}
