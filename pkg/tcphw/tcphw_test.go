package main_test

import (
	"testing"

	"github.com/scottshotgg/proximity/pkg/tcphw/client"
	"github.com/scottshotgg/proximity/pkg/tcphw/server"
)

func Test_HW(t *testing.T) {
	go server.Start()
	client.Start("localhost:9090")
}
