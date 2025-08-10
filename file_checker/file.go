package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const (
	RecentFileTime = 5 * time.Minute
	unit = 1024
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
	// env variable for folder path which needs to be checked
	//os.Setenv("folder_path", "/home/kartik/Videos")
	paths := os.Getenv("folder_path")

	// Get the current time
	currentTime := time.Now()

	// Get a list of files in the specified directory
	files, err := os.ReadDir(paths)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through each file and check if it was modified in the last 5 minutes
	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			log.Printf("Error retrieving info for file %s: %v", file.Name(), err)
			continue
		}
		mod_time := fi.ModTime()
		formated_time := mod_time.Format("2006 Jan 02 15:04:05")

		// If the file was modified in the last 5 minutes, print relevant details
		if currentTime.Sub(mod_time) <= RecentFileTime {
			fmt.Printf("[%s] File_Name:'%s', size: %s, Modified_time: [%s]\n",
				currentTime.Format("2006 Jan 02 15:04:05"),
				file.Name(), 
				HumanReadableSize(fi.Size()),
				formated_time)
		}
	}
}