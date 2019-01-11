########
# Copyright (c) 2019 Cloudify Platform Ltd. All rights reserved
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
############


package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Send JSON request to context. Format { "args": [ "arg1", ..., "argn" ] }
func call_context(port string, args ...string) ([]byte, error) {
	var buffer bytes.Buffer

	buffer.WriteString("{ \"args\": [ ")
	for i, arg := range args {
		buffer.WriteString("\"")
		buffer.WriteString(arg)
		buffer.WriteString("\"")
		if i < len(args)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("] }")

	resp, err := http.Post("http://localhost:"+port, "application/json", bytes.NewReader(buffer.Bytes()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
