package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var b58 = []byte(alphabet)
var b58Index = func() map[byte]int {
	idx := make(map[byte]int)
	for i, b := range b58 {
		idx[b] = i
	}
	return idx
}()

func EncodeBase58(input []byte) string {
	intVal := new(big.Int).SetBytes(input)
	var result []byte

	zero := big.NewInt(0)
	fiftyEight := big.NewInt(58)
	mod := new(big.Int)

	for intVal.Cmp(zero) > 0 {
		intVal.DivMod(intVal, fiftyEight, mod)
		result = append([]byte{b58[mod.Int64()]}, result...)
	}

	// Handle leading zeros
	for _, b := range input {
		if b != 0 {
			break
		}
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}

func DecodeBase58(input string) ([]byte, error) {
	intVal := big.NewInt(0)
	fiftyEight := big.NewInt(58)

	for i := 0; i < len(input); i++ {
		char := input[i]
		val, ok := b58Index[char]
		if !ok {
			return nil, fmt.Errorf("invalid base58 character: %c", char)
		}
		intVal.Mul(intVal, fiftyEight)
		intVal.Add(intVal, big.NewInt(int64(val)))
	}

	decoded := intVal.Bytes()

	// Add back leading zeros
	leadingOnes := 0
	for i := 0; i < len(input) && input[i] == '1'; i++ {
		leadingOnes++
	}

	return append(make([]byte, leadingOnes), decoded...), nil
}


func main() {
	if len(os.Args) != 2 || (os.Args[1] != "-e" && os.Args[1] != "-d") {
		fmt.Fprintln(os.Stderr, "Usage: base58 [-e|-d]")
		fmt.Fprintln(os.Stderr, "  -e  Encode raw input to Base58")
		fmt.Fprintln(os.Stderr, "  -d  Decode Base58 input to raw")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		switch os.Args[1] {
		case "-e":
			data := []byte(line)
			fmt.Println(EncodeBase58(data))

		case "-d":
			data, err := DecodeBase58(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid base58: %s\n", err)
				continue
			}
			fmt.Println(string(data)) // Print raw string, not hex
		}
	}
}

