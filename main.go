package main

import (
	"github.com/thanhpk/randstr"

	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	STRING_ID       = "STRING"
	POSITIVE_INT_ID = "NUMBER"
	ISO_TIME_ID     = "DATE"
	ISO_TIME_T_ID   = "DATET"
	HEXADECIMAL_ID  = "HEXA"
	COLOR_ID        = "COLOR"
)

var HELP_TEXT = fmt.Sprint(
	`This program allows you to quickly create a CSV file called "generated.csv" with random values

Use program's argument to define numbers of rows and types that will be automagically generated;
for example: 5 NUMBER DATE will generate 5 rows each with a random number and a date separated by ", ".

The number of rows is optional, when omitted the program will randomize it (between 1 and 30)

The list of types currently avaiable is the following (Keep in mind that are case insensitive):`, "\n",
	STRING_ID,       "\t - a random-lenght chars sequence of alfanumeric values\n",
	POSITIVE_INT_ID, "\t - a random number between 1 and 100\n",
	ISO_TIME_ID,     "\t - the today's date complete in ISO fomrmat (ex. 2006-01-02 15:04:05)\n",
	ISO_TIME_T_ID,   "\t - like DATE but date and time are separated by a 'T' instead of ' '\n",
	HEXADECIMAL_ID,  "\t - a random hexadecimal value of 128 bit\n",
	COLOR_ID,        "\t - a random hexadecimal value of 6-chars preceded by '#'")

func main() {
	// length controls over program's args
	if argSize := len(os.Args); argSize < 2 {
		// Check if user did not insert any arguments when launching the program
		fmt.Println("Too few values to generate anything\n Use -help to check how to use me")
		return
	} else if argSize == 2 {
		// Check if user used help command
		if match, _ := regexp.MatchString(`(?i)^-{0,2}help$`, os.Args[1]); match {
			log.Fatal(HELP_TEXT)
		}
	}

	// Set the random seeds to the current nano seconds
	rand.Seed(time.Now().UnixNano())

	var (
		types      []string
		file       *os.File
		nRows, err = strconv.Atoi(os.Args[1])
	)
	if err != nil { // If user did not input a number of rows then randomize it
		nRows = randInt(1, 30)
		types = os.Args[1:]
	} else {
		types = os.Args[2:]
	}

	// Create the file
	file, err = os.Create("generated.csv")
	defer file.Close()
	if err != nil {
		log.Fatal("Error: Unable to create or truncate file 'generated.csv'")
	}

	// Fill the file
	for i := 0; i < nRows; i++ {
		file.WriteString(genRow(types, ", ") + "\n")
	}
}

// genRow generates a row of random values for each type passed
func genRow(types []string, joiner string) string {
	var row []string
	for _, selectedType := range types {
		var value string
		switch strings.ToUpper(selectedType) {
		case STRING_ID:
			value = randstr.String(randInt(5, 10))
		case POSITIVE_INT_ID:
			value = fmt.Sprint(randInt(1, 100))
		case ISO_TIME_ID:
			value = time.Now().Format("2006-01-02 15:04:05")
		case ISO_TIME_T_ID:
			value = time.Now().Format("2006-01-02T15:04:05")
		case HEXADECIMAL_ID:
			value = randstr.Hex(16)
		case COLOR_ID:
			value = "#" + randstr.Hex(16)[:6]

		default:
			arr := strings.Split(selectedType, ":")
			ind := randInt(0, uint(len(arr)))
			row = append(row, arr[ind])
		}
		row = append(row, value)
	}
	return strings.Join(row, joiner)
}

// randInt generates a random positive integer
func randInt(min, max uint) int {
	minInt := int(min)
	return rand.Intn(int(max)-minInt) + minInt
}
