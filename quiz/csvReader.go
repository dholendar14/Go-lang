package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problem.csv", "a csv fle in the format of 'question,answser'")
	csvLines := file(*csvFileName)
	val := parseLines(csvLines)
	count := 1
	for i, ques := range val {
		fmt.Printf("Problem #%d %s=\n", i+1, ques.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ques.a == ans {
			count++
		}
	}
	fmt.Printf("you have answered %v out of %v", count, len(csvLines))
}
func file(fileName string) [][]string {
	csvfile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	csvLines, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvLines
}

func parseLines(lines [][]string) []quiz {
	ret := make([]quiz, len(lines))
	for i, line := range lines {
		ret[i] = quiz{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}

	}
	return ret
}
