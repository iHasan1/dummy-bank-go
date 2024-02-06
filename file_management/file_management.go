package file_management

import (
	"os"
	"errors"
	"strconv"
	"fmt"
)

func getFloatFromFile(fileName string) (float64, error){
	data, err := os.ReadFile(fileName) // using underscore simply means telling go u dont want to use that returned value right now

	if err != nil {
		return 1000.0, errors.New("failed to find file")
	}

	fileText := string(data)
	dataStr, err := strconv.ParseFloat(fileText, 64)

	if err != nil {
		return 1000.0, errors.New("failed to parse stored value")
	}

	return dataStr, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644) //0644 is a file permissions notation which means owner can read write and others can only read

}