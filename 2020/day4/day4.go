package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type passport struct {
	BirthYear, IssueYear, ExpirationYear, Height, HairColor, EyeColor, PassportId, CountryId string
}

func main() {
	passports := readPassports(os.Stdin)
	validCount := 0
	validCountPart2 := 0
	for p, pp := range passports {
		if pp.isValid() {
			validCount += 1
		}
		if pp.isPart2Valid() {
			validCountPart2 += 1
		} else {
			log.Printf("%d: part2 invalid: %+v\n", p, pp)
		}
	}
	fmt.Printf("validCount = %d\n", validCount)
	fmt.Printf("validCountPart2 = %d\n", validCountPart2)
}

func (pp *passport) isPart2Valid() bool {
	return len(pp.BirthYear) != 0 && isBirthYearValid(pp.BirthYear) &&
		len(pp.IssueYear) != 0 && isIssueYearValid(pp.IssueYear) &&
		len(pp.ExpirationYear) != 0 && isExpirationYearValid(pp.ExpirationYear) &&
		len(pp.Height) != 0 && isHeightValid(pp.Height) &&
		len(pp.HairColor) != 0 && isHairColorValid(pp.HairColor) &&
		len(pp.EyeColor) != 0 && isEyeColorValid(pp.EyeColor) &&
		len(pp.PassportId) != 0 && isPassportIdValid(pp.PassportId)
}

var yearRegex = regexp.MustCompile("^\\d{4}$")

func isYearInRange(name, year string, min, max int) bool {
	if !yearRegex.MatchString(year) {
		log.Printf("%s year did not match regex: %s", name, year)
		return false
	}
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if min <= y && y <= max {
		return true
	} else {
		log.Printf("%s year not in range: %s; it should be between %d and %d inclusive", name, year, min, max)
		return false
	}
}
func isBirthYearValid(year string) bool {
	return isYearInRange("birth", year, 1920, 2002)
}

func isIssueYearValid(year string) bool {
	return isYearInRange("issue", year, 2010, 2020)
}

func isExpirationYearValid(year string) bool {
	return isYearInRange("expiration", year, 2020, 2030)
}

var heightRegex = regexp.MustCompile("^(\\d+)(cm|in)$")

func isHeightValid(height string) bool {
	parts := heightRegex.FindStringSubmatch(height)
	if len(parts) != 3 {
		log.Printf("height: Expected 3 parts: %+v", parts)
		return false
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("height: Invalid number string: %s", parts[1])
		return false
	}
	if parts[2] == "cm" {
		if 150 <= y && y <= 193 {
			return true
		} else {
			log.Printf("height: cm term out of range %d", y)
			return false
		}
	} else if parts[2] == "in" {
		if 59 <= y && y <= 76 {
			return true
		} else {
			log.Printf("height: in term out of range %d", y)
			return false
		}

	} else {
		log.Printf("height: invalid type: %s", parts[2])
		return false
	}
}

var hairColorRegex = regexp.MustCompile("^#[0-9a-f]{6}$")

func isHairColorValid(hairColor string) bool {
	if hairColorRegex.MatchString(hairColor) {
		return true
	} else {
		log.Printf("hairColor: did not match regex: %s", hairColor)
		return false
	}
}

var eyeColors = []string{"amb", "blu", "brn", "grn", "gry", "hzl", "oth"}
var eyeColorsLen = len(eyeColors)

func isEyeColorValid(eyeColor string) bool {
	p := sort.Search(eyeColorsLen, func(idx int) bool {
		return eyeColors[idx] >= eyeColor
	})
	if p < eyeColorsLen && eyeColors[p] == eyeColor {
		return true
	} else {
		log.Printf("eyeColor: was not one of: %v : %s, p = %d", eyeColors, eyeColor, p)
		return false
	}
}

var passportIdRegex = regexp.MustCompile("^[0-9]{9}$")

func isPassportIdValid(passportId string) bool {
	if passportIdRegex.MatchString(passportId) {
		return true
	} else {
		log.Printf("passportId: does not match regex: %s", passportId)
		return false
	}

}

func (pp *passport) isValid() bool {
	return len(pp.BirthYear) != 0 &&
		len(pp.IssueYear) != 0 &&
		len(pp.ExpirationYear) != 0 &&
		len(pp.Height) != 0 &&
		len(pp.HairColor) != 0 &&
		len(pp.EyeColor) != 0 &&
		len(pp.PassportId) != 0
}

func readPassport(scanner *bufio.Scanner) *passport {
	var pp *passport = nil
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			return pp
		}
		if pp == nil {
			pp = new(passport)
		}
		pairs := strings.Split(line, " ")
		for _, pairStr := range pairs {
			pair := strings.Split(pairStr, ":")
			switch pair[0] {
			case "byr":
				pp.BirthYear = pair[1]
			case "iyr":
				pp.IssueYear = pair[1]
			case "eyr":
				pp.ExpirationYear = pair[1]
			case "hgt":
				pp.Height = pair[1]
			case "hcl":
				pp.HairColor = pair[1]
			case "ecl":
				pp.EyeColor = pair[1]
			case "pid":
				pp.PassportId = pair[1]
			case "cid":
				pp.CountryId = pair[1]
			default:
				panic(fmt.Errorf("unknown field type: %s with value %s", pair[0], pair[1]))
			}
		}
	}
	return pp
}

func readPassports(r io.Reader) []*passport {
	scanner := bufio.NewScanner(r)
	passports := make([]*passport, 0, 1000)
	for {
		passport := readPassport(scanner)
		if passport == nil {
			break
		}
		passports = append(passports, passport)
	}
	return passports
}
