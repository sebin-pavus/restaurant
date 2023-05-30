package mocks

import "fmt"

var Sample = []string{"125", "475", "879"}

type MockServiceSuccess struct{}

func (m MockServiceSuccess) GetTopMenuItems(logFile string) ([]string, error) {
	return Sample, nil
}

type MockServiceFailure struct{}

func (m MockServiceFailure) GetTopMenuItems(logFile string) ([]string, error) {
	var sample []string
	return sample, fmt.Errorf("Error getting menu items")
}
