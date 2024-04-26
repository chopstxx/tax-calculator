package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func (fm FileManager) ReadLines()([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("File opening failed")
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("Reading the file content failed")
	}
	file.Close()
	return lines, nil
}
func (fm FileManager) WriteResult(path string, data any)error{
	file, err := os.Create(fm.OutputFilePath)
	if err!=nil{
		return errors.New("Failed to create file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert to JSON")
	}
	file.Close()
	return nil
}
func New (inputPath, outputPath string) FileManager{
	return FileManager{
		inputPath, outputPath,
	}
}