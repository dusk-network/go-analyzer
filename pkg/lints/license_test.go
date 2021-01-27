// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

package lints

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test doubles as a lint for the current package, but makes sure the
// license header check works in its entirety.
func TestLicenseLint(t *testing.T) {
	dirPath, err := os.Getwd()
	assert.NoError(t, err)

	assert.Empty(t, CheckLicenseHeaders(dirPath))
}

// Ensure that files with a correct header pass the check.
func TestCorrectHeader(t *testing.T) {
	assert := assert.New(t)

	// Let's quickly create a go file with a correct license header.
	fileName := "dummyfile.go"

	file, err := os.Create(fileName)
	assert.NoError(err)

	defer func() {
		_ = os.Remove(fileName)
	}()

	defer func() {
		_ = file.Close()
	}()

	// Write in some code...
	_, err = file.Write([]byte(`// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`))
	assert.NoError(err)

	// Now, let's run the license header lint on it.
	assert.NoError(checkFile(file.Name()))
}

// Ensure that files with an incorrect header result in an error being returned.
func TestIncorrectHeader(t *testing.T) {
	assert := assert.New(t)

	// Let's now make a couple go files with incorrect license headers.
	fileName := "dummyfile.go"

	file, err := os.Create(fileName)
	assert.NoError(err)

	defer func() {
		_ = os.Remove(fileName)
	}()

	defer func() {
		_ = file.Close()
	}()

	// No spacing between package name and license header
	_, err = file.Write([]byte(`// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`))
	assert.NoError(err)

	assert.Error(checkFile(file.Name()))

	// Empty the file, and retry
	assert.NoError(file.Truncate(0))

	// No license header
	_, err = file.Write([]byte(`package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`))
	assert.NoError(err)

	assert.Error(checkFile(file.Name()))
	assert.NoError(file.Truncate(0))

	// Incorrect spacing at the start
	_, err = file.Write([]byte(`
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT License was not distributed with this
// file, you can obtain one at https://opensource.org/licenses/MIT.
//
// Copyright (c) DUSK NETWORK. All rights reserved.

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`))
	assert.NoError(err)

	assert.Error(checkFile(file.Name()))
}
