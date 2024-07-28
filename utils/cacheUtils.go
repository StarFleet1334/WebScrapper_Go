package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ClearCache() {
	cacheDir := "cache_dir"
	err := os.RemoveAll(cacheDir)
	if err != nil {
		fmt.Println("Error clearing cache:", err)
	} else {
		fmt.Println("Cache cleared successfully.")
	}
}

func GetCacheFilePath(cacheDir, linkName string) string {
	cacheFileName := strings.ReplaceAll(linkName, "/", "_")
	return filepath.Join(cacheDir, cacheFileName)
}

func ReadCache(filePath string) ([]byte, error) {
	// Debugging print
	fmt.Println("Reading cache from:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteCache(filePath string, data []byte) error {
	// Debugging print
	fmt.Println("Writing cache to:", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
