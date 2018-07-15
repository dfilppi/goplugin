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
