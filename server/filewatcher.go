package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func startFileWatcher(fileName string, rootNode *s3.Node) (*fsnotify.Watcher, error) {
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
					rootNode.LoadData(fileName)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileName)
	if err != nil {
		return nil, err
	}

	return watcher, nil
}
