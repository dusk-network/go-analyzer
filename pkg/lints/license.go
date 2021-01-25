// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

package lints

import (
	"fmt"
	"os"
	"strings"
)

const correctHeader = `// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

`

// CheckLicenseHeaders checks that all *.go files contained in the current
// directory have the proper license headers.
func CheckLicenseHeaders() []error {
	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return checkDirectory(dirPath)
}

func checkDirectory(dirPath string) []error {
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	errs := make([]error, 0)
	for _, file := range files {
		if file.IsDir() {
			errs = append(errs, checkDirectory(dirPath+"/"+file.Name())...)
		}

		if strings.Contains(file.Name(), ".go") {
			if err := checkFile(dirPath + "/" + file.Name()); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}

func checkFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, len(correctHeader))
	if _, err := file.Read(buf); err != nil {
		return err
	}

	if strings.Compare(string(buf), correctHeader) != 0 {
		return fmt.Errorf("%v does not have the correct license header", fileName)
	}

	return nil
}
