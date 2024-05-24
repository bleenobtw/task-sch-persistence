package main

import (
	"io"
	"log"
	"os"
	"task-persist/utils"
	"time"
)

const (
	taskName      = "Persistence Test"
	cacheInfoPath = "./cache.json"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("")
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	logger := log.New(multiWriter, "Info: ", log.LstdFlags)

	// Log first execution.
	logger.Println("Program has been executed..")

	if _, err := os.Stat(cacheInfoPath); os.IsNotExist(err) {
		logger.Printf("%s does not exist. Creating it..", cacheInfoPath)
		if err := utils.CreateCacheFile(cacheInfoPath); err != nil {
			logger.Printf("Failed to create %s: %v", cacheInfoPath, err)
		}
	}

	cacheInfo, err := utils.ReadCacheInfo(cacheInfoPath)
	if err != nil {
		logger.Fatalf("Failed to read cache file contents: %v", err)
	}

	if _, ok := cacheInfo["configureTaskScheduler"]; !ok {
		// Configure task scheduler task.
		taskExists, err := utils.CheckIfTaskExists(taskName)
		if err != nil {
			logger.Panicf("Error checking if task exists: %v", err)
			return
		}

		if !taskExists {
			// Create the task.

			cacheInfo["configureTaskScheduler"] = time.Now()
			logger.Println("Successfully created task.")
		}
	}

	if err := utils.WriteCacheInfo(cacheInfoPath, cacheInfo); err != nil {
		logger.Printf("There was an error writing to build info file: %v", err)
		return
	}
}
