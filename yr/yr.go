package yr

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ineant/funtemps/conv"
)

func CelsiusToFarhrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFarhrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {
	elementsInLine := strings.Split(line, ";")
	var err error
	if len(elementsInLine) == 4 {
		elementsInLine[3], err = CelsiusToFarhrenheitString(elementsInLine[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(elementsInLine, ";"), nil
}

func AverageCelsius(unit string) (string, error) {
	var buffer []byte
	var linebuf []byte
	buffer = make([]byte, 1)
	bytesCount := 0
	lineCount := 0
	result := ""

	var sum float64 = 0
	var n float64 = 0

	src, err := os.Open("../kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		bytesCount++
		if buffer[0] == 0x0A {
			lineCount++

			if lineCount == 1 {
				linebuf = nil
				continue
			}
			if lineCount == 16756 {
				break
			}
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				f, err := strconv.ParseFloat(celsius, 64)
				if err != nil {
					log.Fatal(err)
				}
				sum += f
				n += 1
			}
			linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}
		if err == io.EOF {
			break
		}
		average := sum / n

		if unit == "c" {
			result = fmt.Sprintf("%.2f", average)
		} else if unit == "f" {
			fahr := conv.CelsiusToFarhrenheit(average)
			result = fmt.Sprintf("%.2f", fahr)
		}
	}
	return result, errors.New("linje har ikke forventet format")
}

func CountLinesInFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strconv.Itoa(lineCount), nil
}
