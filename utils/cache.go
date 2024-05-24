package utils

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type CacheInfo map[string]time.Time

func ReadCacheInfo(filePath string) (CacheInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cacheInfo CacheInfo
	if err := json.Unmarshal(byteValue, &cacheInfo); err != nil {
		return nil, err
	}
	return cacheInfo, nil
}

func WriteCacheInfo(filePath string, cacheInfo CacheInfo) error {
	byteValue, err := json.MarshalIndent(cacheInfo, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, byteValue, 0644)
}

func CreateCacheFile(filePath string) error {
	defaultContent := CacheInfo{
		"createCacheFile": time.Now(),
	}
	return WriteCacheInfo(filePath, defaultContent)
}
