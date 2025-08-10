package main

import (
    "log"
	"math"
	"strconv"
	"os"
    "github.com/fsnotify/fsnotify"
)
const (
	unit = 1000
)

func HumanReadableSize(bytes int64) string {
	if bytes == 0 {
		return "0 B"
	}
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
    logUnit := math.Log(float64(unit))
    i := int(math.Floor(math.Log(float64(bytes)) / logUnit))
    size := float64(bytes) / math.Pow(unit, float64(i))
    return strconv.FormatFloat(size, 'f', 2, 64) + " " + sizes[i]
}

func main() {
    // Create new watcher.
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // Start listening for events.
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
				fileinfo, err := os.Stat(event.Name)
                    if err != nil {
                        log.Println("Error getting file info:", err)
                        continue
                    }
				//log.Printf("size:%s",HumanReadableSize(fileinfo.Size()))
                log.Printf("event: %s, File-size: %s", event, HumanReadableSize(fileinfo.Size()))
                if event.Has(fsnotify.Write) {
                    log.Println("modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add("/home/kartik/Videos")
    if err != nil {
        log.Fatal(err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}