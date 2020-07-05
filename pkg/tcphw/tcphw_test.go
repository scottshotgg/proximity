package main_test

import (
	"testing"
	"time"

	"github.com/scottshotgg/proximity/pkg/tcphw/client"
	"github.com/scottshotgg/proximity/pkg/tcphw/server"
)

func Test_HW(t *testing.T) {
	var n = server.New()

	go n.Start("localhost")

	time.Sleep(100 * time.Millisecond)

	go client.New().Start("localhost", 1, true)

	time.Sleep(100 * time.Millisecond)

	client.New().Start("localhost", 1, false)
}
