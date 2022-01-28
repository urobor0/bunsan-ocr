package ocr

import (
	"bufio"
	"bunsan-ocr/kit/projectpath"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	linesOfEachEntry            = 4
	linesPerNumber              = 3
	zero              TextDigit = " _ | ||_|"
	one               TextDigit = "     |  |"
	two               TextDigit = " _  _||_ "
	three             TextDigit = " _  _| _|"
	four              TextDigit = "   |_|  |"
	five              TextDigit = " _ |_  _|"
	six               TextDigit = " _ |_ |_|"
	seven             TextDigit = " _   |  |"
	eight             TextDigit = " _ |_||_|"
	nine              TextDigit = " _ |_| _|"
)

// TextDigit represents the ocr digit in one line.
type TextDigit string

// ocrListOfNumericValue makes a text digit matches the numeric value.
var ocrListOfNumericValue = map[TextDigit]int{
	zero:  0,
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}

// String type converts the TextDigit into string.
func (o TextDigit) String() string {
	if n, ok := ocrListOfNumericValue[o]; ok {
		return strconv.Itoa(n)
	}
	return "?"
}

// Number type converts the TextDigit into int.
func (o TextDigit) Number() int {
	return ocrListOfNumericValue[o]
}

// AccountNumberTxt is the data structure that represent a AccountNumberTxt.
type AccountNumberTxt struct {
	d1, d2, d3, d4, d5, d6, d7, d8, d9 TextDigit
}

// NewOcrAccountNumberTxt creates a new AccountNumberFromTxt.
func NewOcrAccountNumberTxt(value []string) AccountNumberTxt {
	length := len(value)
	if length != linesPerNumber {
		return AccountNumberTxt{}
	}

	return AccountNumberTxt{
		d9: TextDigit(value[0][0:3] + value[1][0:3] + value[2][0:3]),
		d8: TextDigit(value[0][3:6] + value[1][3:6] + value[2][3:6]),
		d7: TextDigit(value[0][6:9] + value[1][6:9] + value[2][6:9]),
		d6: TextDigit(value[0][9:12] + value[1][9:12] + value[2][9:12]),
		d5: TextDigit(value[0][12:15] + value[1][12:15] + value[2][12:15]),
		d4: TextDigit(value[0][15:18] + value[1][15:18] + value[2][15:18]),
		d3: TextDigit(value[0][18:21] + value[1][18:21] + value[2][18:21]),
		d2: TextDigit(value[0][21:24] + value[1][21:24] + value[2][21:24]),
		d1: TextDigit(value[0][24:27] + value[1][24:27] + value[2][24:27]),
	}
}

// String type converts the AccountNumberTxt into string.
func (ocr AccountNumberTxt) String() string {
	return fmt.Sprintf("%s %s", ocr.Value(), ocr.Status())
}

// Status returns a status of AccountNumberTxt
func (ocr AccountNumberTxt) Status() string {
	switch {
	case strings.Contains(ocr.Value(), "?"):
		return "ILL"
	case !ocr.CheckSum():
		return "ERR"
	default:
		return "OK"
	}
}

// Value returns a value of AccountNumberTxt joining each of the digits
func (ocr AccountNumberTxt) Value() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s",
		ocr.d9.String(),
		ocr.d8.String(),
		ocr.d7.String(),
		ocr.d6.String(),
		ocr.d5.String(),
		ocr.d4.String(),
		ocr.d3.String(),
		ocr.d2.String(),
		ocr.d1.String(),
	)
}

// CheckSum returns if the module rule by 11 is passed successfully.
func (ocr AccountNumberTxt) CheckSum() bool {
	sum := 1*ocr.d1.Number() +
		2*ocr.d2.Number() +
		3*ocr.d3.Number() +
		4*ocr.d4.Number() +
		5*ocr.d5.Number() +
		6*ocr.d6.Number() +
		7*ocr.d7.Number() +
		8*ocr.d8.Number() +
		9*ocr.d9.Number()

	return sum%11 == 0
}

func TextConverter(filePath string, id JobID) error {
	fileOutputPath := fmt.Sprintf("%s/attachments/%s-output.txt", projectpath.RootDir(), id.String())
	fileOutput, err := os.Create(fileOutputPath)
	if err != nil {
		return err
	}

	defer fileOutput.Close()

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	var accountNumberSlice []string
	for scanner.Scan() {
		accountNumberSlice = append(accountNumberSlice, scanner.Text())

		if i%linesPerNumber == 0 {
			account := NewOcrAccountNumberTxt(accountNumberSlice).String()
			fmt.Println(account)
			_, err := fileOutput.WriteString(fmt.Sprintf("%s\n", account))
			if err != nil {
				return err
			}
		}

		if i%linesOfEachEntry == 0 {
			i = 1
			accountNumberSlice = []string{}
			continue
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
