#!/bin/bash

# Function to convert bytes to human-readable size
human_readable_size() {
  local bytes=$1
  if (( bytes == 0 )); then
    echo "0 B"
    return
  fi

  local units=("B" "KB" "MB" "GB" "TB" "PB" "EB")
  local unit=1024
  local i=0
  local size=$bytes

  while (( size >= unit )); do
    size=$((size / unit))
    ((i++))
  done

  echo "$size ${units[i]}"
}

# Set folder path from environment variable
folder_path="${folder_path:-/home/kartik/Videos}"

# Get the current time in seconds since epoch
current_time=$(date +%s)

# Loop through all files in the directory
for file in "$folder_path"/*; do
  if [ -f "$file" ]; then
    # Get the file's last modification time in seconds since epoch
    mod_time=$(stat -c %Y "$file")
    
    # Calculate the difference between current time and file's modification time
    time_diff=$((current_time - mod_time))
    
    # If the file was modified in the last 5 minutes (300 seconds)
    if (( time_diff <= 300 )); then
      # Get the file size
      file_size=$(stat -c %s "$file")
      
      # Convert file size to human-readable format
      readable_size=$(human_readable_size "$file_size")
      
      # Format modification time into readable format
      formatted_time=$(date -d @$mod_time "+%Y %b %d %H:%M:%S")

      # Print the details
      echo "[$(date "+%Y %b %d %H:%M:%S")] File_Name:'$(basename "$file")', size: $readable_size, Modified_time: $formatted_time"
    fi
  fi
done
