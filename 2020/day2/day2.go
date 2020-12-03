package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type rule struct {
	Character string
	Min, Max  int
}

type password struct {
	Line     string
	Rule     *rule
	Password string
}

var lineFormat = regexp.MustCompile("(\\d+)-(\\d+)\\s+([^:]+):\\s+(.*)")

func readPasswords(reader io.Reader) ([]*password, error) {
	scanner := bufio.NewScanner(reader)
	passwords := make([]*password, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		matches := lineFormat.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("error parsing line: `%s`: no match found")
		}
		if len(matches) != 5 {
			return nil, fmt.Errorf("error parsing line: `%s`: only %d submatches found", line, len(matches)-1)
		}
		min, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing line `%s`: min not an integer")
		}
		max, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing line `%s`: max not an integer")
		}
		check := fmt.Sprintf("%d-%d %s: %s", min, max, matches[3], matches[4])
		if check != line {
			log.Fatalf("line does not match reconstruction on line %d", len(passwords))
		}
		passwords = append(passwords, &password{line, &rule{matches[3], min, max}, matches[4]})
	}
	return passwords, nil
}

func checkPassword(pw *password) bool {
	count := 0
	for _, r := range pw.Password {
		if string(r) == pw.Rule.Character {
			count += 1
		}
	}
	ok := pw.Rule.Min <= count && count <= pw.Rule.Max
	return ok
}

func checkPasswordPart2(pw *password) bool {
	count := 0
	for p, r := range pw.Password {
		if (p + 1 == pw.Rule.Min || p + 1 == pw.Rule.Max) && string(r) == pw.Rule.Character {
			count += 1
		}
	}
	return count == 1
}

func main() {
	passwords, err := readPasswords(os.Stdin)
	if err != nil {
		log.Fatal(err.Error())
	}
	validCount := 0
	validCountPart2 := 0
	for _, pw := range passwords {
		if checkPassword(pw) {
			validCount += 1
		}
		if checkPasswordPart2(pw) {
			validCountPart2 += 1
		}
	}
	log.Printf("validCount = %d", validCount)
	log.Printf("validCountPart2 = %d", validCountPart2)
}
