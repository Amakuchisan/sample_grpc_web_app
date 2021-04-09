package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := insertPictures(); err != nil {
		log.Fatal(err)
	}
}

func insertPictures() error {
	dir := "./files/"
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	var pictures []string
	for _, file := range files {
		data, err := os.ReadFile(dir + file.Name())
		if err != nil {
			return err
		}

		b64 := base64.StdEncoding.EncodeToString(data)
		pictures = append(pictures, b64)
	}

	query := fmt.Sprintf("INSERT INTO picture(image) VALUES (\"%s\");", strings.Join(pictures, "\"), (\""))

	file, err := os.Create("../init/2_insert.sql")
	if err != nil {
		return err
	}
	defer file.Close()

	output := query
	file.Write(([]byte)(output))

	return nil
}
