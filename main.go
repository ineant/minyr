package main

import (
        "os"
        "log"
        "io"
        "strings"
        "github.com/ineant/funtemps/conv"
        "strconv"
        "bufio"
	"fmt"
)       

func main() {
        src, err := os.Open("/home/ineant/project/minyr/kjevik-temp-celsius-20220318-20230318.csv")
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

	scanner := bufio.NewScanner(os.Stdin)

	for {
	fmt.Println("enter an option")
	scanner.Scan()
	input := scanner.Text()

	if input == "q" || input == "exit"{
		fmt.Println("Converting all measurments..")
		os.Exit(0)
	}else if input == "convert"{
		fmt.Println("converting all measurment..")
		os.Open("/home/ineant/project/minyr/kjevik-temp-fahr-20220318-20230318.csv")
	}else {
            fmt.Println("Please choose a valid option.")
            continue
        }

        var buffer []byte
        var linebuf []byte // nil
        buffer = make([]byte, 1)
        bytesCount := 0
        lineCount := 0
        
        for {
                _, err := src.Read(buffer)
                if err != nil && err != io.EOF {
                        log.Fatal(err)
                }
		bytesCount++
                if buffer[0] == 0x0A {
                   //log.Println(string(linebuf))
                        lineCount++

                        if  lineCount == 1 || lineCount == 16755 {
                        log.Println(string(linebuf))
                         _, err = writer.WriteString(string(linebuf) + "\n")
                        if err != nil {
                                log.Fatal(err)
                        }
                        linebuf = nil
                        continue
                        }
                   elementArray := strings.Split(string(linebuf), ";")
                   if len(elementArray) > 3 {
                         celsius := elementArray[3]
                        f, err := strconv.ParseFloat(celsius, 64)
                        if err != nil{
                                log.Fatal(err)
                        }

                        fahr := conv.CelsiusToFarhrenheit(f)
                        fahrStr := strconv.FormatFloat(fahr, 'f', 2, 64)
                        elementArray[3] = fahrStr
                        //
                        var result string
                        for index, el := range(elementArray) {
                                log.Println(el)
                        
                                if index +1 != len(elementArray) {
                                        result = result + el + ";"
                                } else {
                                        result = result + el
                                }
                        }

                        //
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

}
