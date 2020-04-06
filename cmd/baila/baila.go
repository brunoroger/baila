package main

import (
	"baila/api"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func printInstructionsForCommand() {
	fmt.Println("Baila CLI")
	fmt.Println("")
	fmt.Println("Basic Commands:")
	fmt.Println("	get: to consult for example the temperature")
}

func printInstructionsForCommandGet() {
	fmt.Println("Baila CLI")
	fmt.Println("")
	fmt.Println("Basic Commands for: baila get")
	fmt.Println("	temperature: to consult temperature")
}

func main() {
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)

	if len(os.Args) < 2 {
		printInstructionsForCommand()
		return
	}

	switch os.Args[1] {
	case "get":
		getCommand.Parse(os.Args[1:])
	default:
		printInstructionsForCommand()
		return
	}

	if getCommand.Parsed() {
		if len(os.Args) < 3 {
			printInstructionsForCommandGet()
			return
		}

		temperatureCommand := flag.NewFlagSet("temperature", flag.ExitOnError)

		switch os.Args[2] {
		case "temperature":
			temperatureCommand.Parse(os.Args[2:])
		default:
			printInstructionsForCommandGet()
			return
		}

		if temperatureCommand.Parsed() {
			fmt.Println("Enter city: ")

			in := bufio.NewReader(os.Stdin)

			city, errCity := in.ReadString('\n')
			if errCity != nil {
				fmt.Println(errCity)
				return
			}

			city = strings.ReplaceAll(city, "\n", "")

			fmt.Println("Enter unit (1 - celsius, 2 - fahrenheit): ")

			unit, errUnit := in.ReadString('\n')
			if errUnit != nil {
				fmt.Println(errUnit)
				return
			}

			unit = strings.ReplaceAll(unit, "\n", "")

			service := api.NewTemperatureServicer()

			response, err := service.Temperature(city, unit)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Temperature:", response.Main.Temp)
		}
	}
}
