package logs

import "testing"

func TestInfo(t *testing.T) {
	Info("infoddd", Content{
		"id": 1, "name": "xxx",
	})

	Warning("warning拉", Content{
		"id": 2,
	})
}
