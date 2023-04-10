package main

func updateProgress(fileId string, progress int) {
	progressMutex.Lock()
    defer progressMutex.Unlock()
    progressMap[fileId] = progress

    if updateChan, ok := progressChannels[fileId]; ok {
        updateChan <- progress
    }
}