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

type typeIdentifier = string

const (
	STRING_ID       typeIdentifier = "STRING"
	POSITIVE_INT_ID                = "NUMBER"
	ISO_TIME_ID                    = "DATE"
	ISO_TIME_T_ID                  = "DATET"
	HEXADECIMAL_ID                 = "HEXA"
	COLOR_ID                       = "COLOR"
)

const CUSTOM_SEPARATOR = ":"

var HELP_TEXT = fmt.Sprint(
	`This program allows you to quickly create a CSV file called "generated.csv" with random values

Use program's argument to define numbers of rows and types that will be automagically generated;
for example: 5 NUMBER DATE will generate 5 rows each with a random number and a date separated by ", ".

The number of rows is optional, when omitted the program will randomize it (between 1 and 30)

The list of types currently avaiable is the following (Keep in mind that are case insensitive):`, "\n",
	STRING_ID, "\t - a random-lenght sequence of random alfanumeric values\n",
	POSITIVE_INT_ID, "\t - a random number between 1 and 100\n",
	ISO_TIME_ID, "\t - the today's date complete in ISO fomrmat (ex. 2006-01-02 15:04:05)\n",
	ISO_TIME_T_ID, "\t - like DATE but date and time are separated by a 'T' instead of ' '\n",
	HEXADECIMAL_ID, "\t - a random hexadecimal value of 32 chars\n",
	COLOR_ID, "\t - a random hexadecimal value of 6 chars preceded by '#'\n", `
Use '`, CUSTOM_SEPARATOR, `' to separate two or more custom strings as a type to choose randomly from from them.
For example 'Pippo`, CUSTOM_SEPARATOR, "pluto", CUSTOM_SEPARATOR, `TOPOLINO' will make csvgen randomly pick one from 'Pippo', 'pluto' or 'TOPOLINO'.
This time case is sensitive`)

func main() {
	// length controls over program's args
	if argSize := len(os.Args); argSize < 2 {
		// Check if user did not insert any arguments when launching the program
		log.Fatal("Too few values to generate anything\n Use -help to check how to use me")
	} else if argSize == 2 {
		// Check if user used help command
		if match, _ := regexp.MatchString(`(?i)^-{0,2}help$`, os.Args[1]); match {
			fmt.Println(HELP_TEXT)
			return
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
	fmt.Println("Done!")
}

// genRow generates a row of random values for each type passed
func genRow(types []typeIdentifier, joiner string) string {
	var row = make([]string, len(types))
	for i, selectedType := range types {
		switch strings.ToUpper(selectedType) {
		case STRING_ID:
			row[i] = randstr.String(randInt(5, 10))
		case POSITIVE_INT_ID:
			row[i] = fmt.Sprint(randInt(1, 100))
		case ISO_TIME_ID:
			row[i] = time.Now().Format("2006-01-02 15:04:05")
		case ISO_TIME_T_ID:
			row[i] = time.Now().Format("2006-01-02T15:04:05")
		case HEXADECIMAL_ID:
			row[i] = randstr.Hex(16)
		case COLOR_ID:
			row[i] = "#" + randstr.Hex(3)
		default:
			row[i] = randValue(strings.Split(selectedType, CUSTOM_SEPARATOR))
		}
	}
	return strings.Join(row, joiner)
}

// randInt generates a random positive integer
func randInt(min, max uint) int {
	minInt := int(min)
	return rand.Intn(int(max)-minInt) + minInt
}

// randValue get one random value from a list of possible values
func randValue(values []string) string {
	return values[rand.Intn(len(values))]
}
