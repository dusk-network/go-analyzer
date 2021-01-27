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
func CheckLicenseHeaders(dirPath string) []error {
	return checkDirectory(dirPath)
}

func checkDirectory(dirPath string) []error {
	//nolint:gosec
	dir, err := os.Open(dirPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = dir.Close()
	}()

	files, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	errs := make([]error, 0)

	for _, file := range files {
		if file.IsDir() {
			errs = append(errs, checkDirectory(dirPath+"/"+file.Name())...)
		}

		if strings.TrimSuffix(file.Name(), ".go") != file.Name() {
			if err := checkFile(dirPath + "/" + file.Name()); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}

func checkFile(fileName string) error {
	//nolint:gosec
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = file.Close()
	}()

	buf := make([]byte, len(correctHeader))
	if _, err := file.Read(buf); err != nil {
		return err
	}

	if strings.Compare(string(buf), correctHeader) != 0 {
		return fmt.Errorf("%v does not have the correct license header", fileName)
	}

	return nil
}
