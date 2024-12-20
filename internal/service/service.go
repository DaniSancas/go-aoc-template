package service

import (
	"aoc-runner/internal/day"
	"fmt"
	"log"
	"os"
	"reflect"
)

// Invoke a function by reflection with an array of args
func invoke(funcValue reflect.Value, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return funcValue.Call(inputs)
}

func readFile(number int) string {
	path := fmt.Sprintf("assets/input/%02d.txt", number)
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)
}

// Entry point for the program:
//
// For all days (1-25) and exercise variants (a & b)
// - Try to get an existing function
// - Try to read it's associated text file
// - Execute them printing the result!
func ExecuteAllDays() {
	// Reference `day.D` for access all day functions attached to it
	var d day.D

	for dayNumber := 1; dayNumber <= 25; dayNumber++ {
		for _, dayVariant := range []string{"a", "b"} {
			funcName := fmt.Sprintf("Day%02d%s", dayNumber, dayVariant)

			funcValue := reflect.ValueOf(&d).MethodByName(funcName)
			if funcValue.IsValid() {
				fileContent := readFile(dayNumber)
				retValue := invoke(funcValue, fileContent)
				fmt.Println(fmt.Sprintf("%s: %v", funcName, retValue[0]))
			}
		}
	}
}
