#!/bin/bash

# Display the directory structure
echo "Directory structure:"
tree .

# Define a function to display file content
function show_files {
    for file in "$1"/*
    do
        if [ -d "$file" ]; then
            show_files "$file"
        elif [ -f "$file" ]; then
            echo -e "\nContent of $file:"
            cat "$file"
        fi
    done
}

# Call the function with the current directory
show_files .
