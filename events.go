package main

import (
	"fmt"
)

func updateProgress(fileId string, progress int) {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	progressMap[fileId] = progress
}

func addClient(fileId string, client chan string) {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	progressMap[fileId] = 0
	go func() {
		for progress := range progressMap {
			client <- fmt.Sprintf("fileId: %s, progress: %s%%", fileId, progress)
		}
	}()
}

func removeClient(fileId string, client chan string) {
	progressMutex.Lock()
	defer progressMutex.Unlock()
	delete(progressMap, fileId)
	close(client)
}
