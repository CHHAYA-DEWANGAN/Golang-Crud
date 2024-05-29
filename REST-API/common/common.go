package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

func ToGetAlphaString(num int) string {
	var result string
	for num > 0 {
		remainder := (num - 1) % 26
		result = string('A'+remainder) + result
		num = (num - 1) / 26
	}
	return result
}

func GetWorkingDirectory() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}
	fmt.Println("Current working directory:", currentDir)

	// Get the directory of the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return ""
	}
	exeDir := filepath.Dir(exePath)
	fmt.Println("Executable directory:", exeDir)
	return currentDir
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result += string(charset[n.Int64()])
	}
	return result
}
