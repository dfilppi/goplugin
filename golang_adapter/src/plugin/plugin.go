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
	"fmt"
	"os"
)

// The executable just proxies to underlying functions, which are specified in the pluging
// operation mapping or blueprint
//  arg[1] = port of context proxy ( provided by go.py )
//  arg[2] = the function to call
//  arg[3:] = the args for the function if any
func main() {

	// By convention, args[1] is a function name.  The equivalent to
	// the operation in the Python plugin

	ctx_port := os.Args[1]
	fname := os.Args[2]
	args := os.Args[3:]

	if fname == "func1" {
		func1(ctx_port, args)
	} else if fname == "func2" {
		func2(ctx_port, args)
	} else {
		fmt.Println("Function " + fname + " not found")
		os.Exit(1)
	}
}
