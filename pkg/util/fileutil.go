// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// File is for implementing the operations on files.
type FileUtil interface {
	// Read reads the file named by filename and returns all the contents unitl EOF or an error.
	Read(string) ([]byte, error)
	// Write writes data to a file named by filename.
	Write(string, []byte, os.FileMode) error
}

// FileImpl is an implementation of File.
type FileUtilImpl struct {
}

// Read reads the file named by filename and returns all the contents until EOF or an error.
func (f FileUtilImpl) Read(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("file read error: %v", err)
	}
	return content, nil
}

// Write writes data to a file named by filename.
func (f FileUtilImpl) Write(filename string, content []byte, perm os.FileMode) error {
	if err := ioutil.WriteFile(filename, content, perm); err != nil {
		return fmt.Errorf("file write error: %v", err)
	}
	return nil
}