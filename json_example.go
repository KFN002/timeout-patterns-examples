package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		close(result)
		return
	}
	defer file.Close()

	select {
	case <-ctx.Done():
		fmt.Println("Reading JSON file canceled due to timeout or context cancellation")
		close(result)
		return
	default:
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			close(result)
			return
		}
		result <- data
	}
}
