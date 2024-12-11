#!/bin/bash

# Number of days
num_days=11  # You can change this number to the desired number of days

# Loop through each day and create the folders and files
for (( day=10; day<=num_days; day++ ))
do
    # Create the folder
    folder_name="day_$day"
    mkdir -p $folder_name

    # Create the files
    touch $folder_name/day_$day.go
    touch $folder_name/real_input.txt
    touch $folder_name/sample_input.txt
done

echo "Folders and files created successfully."