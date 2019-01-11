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

import "fmt"

// Example plugin function
func func1(ctx_port string, args []string) {
	fmt.Println("FUNC1 CALLED")
	res, _ := call_context(ctx_port, "node", "properties", "prop1")
	fmt.Printf("GOT VALUE of PROP1 = %s\n", res)
	call_context(ctx_port, "logger", "info", "LOGGING SOMETHING FROM PLUGIN")
}

// Example plugin function
func func2(ctx_port string, args []string) {
	fmt.Println("FUNC2 CALLED")
}
