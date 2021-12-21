package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func Next(v, sn int) int {
	return (v * sn) % 20201227
}

func Transform(sn, ls int) int {
	v := 1
	for ls > 0 {
		v = Next(v, sn)
		ls -= 1
	}
	return v
}

func FindLoopSize(target int) int {
	ls := 1
	v := 1
	for {
		v = Next(v, 7)
		if v == target {
			return ls
		}
		ls += 1
	}
}

func ReadPublicKeys(reader io.Reader) []int {
	pks := make([]int, 0, 2)
	for scanner := bufio.NewScanner(reader); scanner.Scan(); {
		line := scanner.Text()
		if len(line) != 0 {
			pk, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			pks = append(pks, pk)
		}
	}
	return pks
}

func main() {
	publicKeys := ReadPublicKeys(os.Stdin)
	cardPk, doorPk := publicKeys[0], publicKeys[1]
	cardLs := FindLoopSize(cardPk)
	doorLs := FindLoopSize(doorPk)
	encryptionKey1 := Transform(cardPk, doorLs)
	encryptionKey2 := Transform(doorPk, cardLs)
	if encryptionKey1 != encryptionKey2 {
		panic("encryption keys are not the same")
	}
	log.Printf("Encryption key = %d", encryptionKey1)
}
