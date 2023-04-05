package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSelfRestartApp(t *testing.T) {
	go main()

	for  {
		select {
		case restarted := <-restartNotifier:
			assert.Equal(t, restarted, "App restarted")
			return
		case <-time.After(10 * time.Second):
			t.Error("App didn't restart within the expected time")
		}
	}
}
