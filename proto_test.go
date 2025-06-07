package main

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {

	rawMsg := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	cmd, err := parseCommand(rawMsg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cmd)

}
