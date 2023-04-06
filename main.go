package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ineant/funtemps/conv"
)

func convert() {

	fmt.Println("converting all measurment..")
	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
	bytesCount := 0
	lineCount := 0

	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	newFile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	writer := bufio.NewWriter(newFile)

	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		bytesCount++
		if buffer[0] == 0x0A {
			//log.Println(string(linebuf))
			lineCount++

			log.Println(lineCount)

			if lineCount == 1 {
				log.Println(string(linebuf))
				_, err = writer.WriteString(string(linebuf) + "\n")
				if err != nil {
					log.Fatal(err)
				}
				linebuf = nil
				continue
			}

			if lineCount == 16756 {
				newLastLine := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Ine Antonsen"
				_, err = writer.WriteString((newLastLine) + "\n")
				if err != nil {
					log.Fatal(err)
				}
				break
			}
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				f, err := strconv.ParseFloat(celsius, 64)
				if err != nil {
					log.Fatal(err)
				}

				fahr := conv.CelsiusToFarhrenheit(f)
				fahrStr := strconv.FormatFloat(fahr, 'f', 2, 64)
				elementArray[3] = fahrStr

				var result string
				for index, el := range elementArray {
					log.Println(el)

					if index+1 != len(elementArray) {
						result = result + el + ";"
					} else {
						result = result + el
					}
				}

				_, err = writer.WriteString(result + "\n")
				if err != nil {
					log.Fatal(err)
				}

			}
			linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}
		if err == io.EOF {
			break
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

}

func average(unit string) {
	fmt.Println("kalkulere gjennomsnitt temperaturen...")
	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
	bytesCount := 0
	lineCount := 0

	var sum float64 = 0
	var n float64 = 0

	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
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
	}
	average := sum / n

	if unit == "c" {
		fmt.Printf("%.2f\n", average)
	} else if unit == "f" {
		fahr := conv.CelsiusToFarhrenheit(average)
		fmt.Printf("%.2f\n", fahr)
	}

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input == "minyr" {
		for {
			fmt.Println("velg et alternativ: q or exit, convert, average")
			scanner.Scan()
			input := scanner.Text()

			if input == "q" || input == "exit" {
				fmt.Println("konventerer alle målinger...")
				os.Exit(0)
			} else if input == "convert" {
				if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
					fmt.Printf("Filen eksisterer allerede, ønsker du å fortsette tast inn: y , eller avbryte med å taste inn: n.\n")
					scanner.Scan()
					input = scanner.Text()
					if input == "y" {
						convert()
					} else if input == "n" {
						continue
					}
				} else {
					convert()
				}
			} else if input == "average" {
				fmt.Println("Ønsker gjennomsnitttemperatur for Celsius, tast inn: c")
				fmt.Println("Ønsker gjennomsnitttemperatur for Farhrenheit, tast inn: f")
				scanner.Scan()
				input = scanner.Text()
				if input == "f" {
					average("f")
				} else if input == "c" {
					average("c")
				}
			} else {
				fmt.Println("Please choose a valid option; q/exit, convert, average")
				continue
			}

		}

	}

}
