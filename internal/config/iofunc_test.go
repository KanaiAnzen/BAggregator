package config

import (
	"fmt"
	"testing"
)

func TestGetconfigfilepath(t *testing.T) {

	path, _ := GetConfigFilePath()

	fmt.Println(path)

}
