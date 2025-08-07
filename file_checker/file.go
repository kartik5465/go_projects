package main

import (
	"fmt"
	"log"
	"os"
	"math"
	"strconv"
	"time"
)

func HumanReadableSize(bytes int64) string {
	if bytes == 0 {
		return "0 B"
	}
	const unit = 1024 // Use 1000 for SI units (e.g., KB = 1000 bytes)
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := int(math.Floor(math.Log(float64(bytes)) / math.Log(unit)))
	size := float64(bytes) / math.Pow(unit, float64(i))
	return strconv.FormatFloat(size, 'f', 2, 64) + " " + sizes[i]
}

func main() {
	// Get the current time
	currentTime := time.Now()

	// Get a list of files in the specified directory
	files, err := os.ReadDir("/home/kartik/Videos/")
	if err != nil {
		log.Fatal(err)
	}

	// Loop through each file and check if it was modified in the last 5 minutes
	for _, file := range files {
		fi, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}

		// If the file was modified in the last 5 minutes, print relevant details
		if currentTime.Sub(fi.ModTime()) <= 5*time.Minute {
			fmt.Printf("[%s] NEWLY ADDED: The file '%s' is %s long\n",
				currentTime.Format("2006 Jan 02 15:04:05"),
				file.Name(), 
				HumanReadableSize(fi.Size()))
		}
	}
}