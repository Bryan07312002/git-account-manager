package main

import (
	"encoding/json"
	"os"
)

func openJson(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return file
}

func saveJson(encoded []byte, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write(encoded)
}

func decodeJson[T any](file *os.File, accounts *[]T) {
	err := json.NewDecoder(file).Decode(&accounts)
	if err != nil {
		panic(err)
	}
}

func encodeToJson[T any](accounts []T) []byte {
	jsonData, err := json.Marshal(accounts)
	if err != nil {
		panic(err)
	}

	return jsonData
}
