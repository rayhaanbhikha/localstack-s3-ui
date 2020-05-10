package file

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// Watch ... watches file for any write changes.
func Watch(filePath string, handleOnWrite func()) (*fsnotify.Watcher, error) {
	// TODO: need to check if the file exists.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					handleOnWrite()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		return nil, err
	}

	return watcher, nil
}
