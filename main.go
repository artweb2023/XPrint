package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	SizeX   = 50
	SizeY   = 5
	MaxSize = SizeX * SizeY
)

func ReadChar(ch rune) bool {
	var ValidLetters = map[rune]bool{
		'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true,
		'G': true, 'H': true, 'I': true, 'J': true, 'K': true, 'L': true,
		'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true,
		'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true,
		'Y': true, 'Z': true,
	}
	return ValidLetters[ch]
}

func LowerToUpper(ch rune) rune {
	var chars []rune
	if unicode.IsLower(ch) {
		char := strings.ToUpper(string(ch))
		chars = []rune(char)
		return chars[0]
	} else {
		return ch
	}
}

func ReadDataInFile(ch string) []int {
	matrix := make(map[string][]int)
	file, err := os.Open("charset.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var numbers []int
		var letters []string
		lines := scanner.Text()
		line := strings.Split(lines, " ")
		for _, val := range line {
			num := ""
			for _, ch := range val {
				if unicode.IsDigit(ch) {
					num += string(ch)
				}
				if unicode.IsLetter(ch) {
					letters = append(letters, string(ch))
				}
			}
			if num != "" {
				number, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number)
			}
		}
		for i := range letters {
			for v := range numbers {
				matrix[letters[i]] = append(matrix[letters[i]], numbers[v])
			}
		}
	}
	return matrix[ch]
}

func GetLettersMatrix(ch rune) []int {
	switch ch {
	case 'A':
		return ReadDataInFile(string(ch))
	case 'B':
		return ReadDataInFile(string(ch))
	case 'C':
		return ReadDataInFile(string(ch))
	case 'D':
		return ReadDataInFile(string(ch))
	case 'E':
		return ReadDataInFile(string(ch))
	case 'F':
		return ReadDataInFile(string(ch))
	case 'G':
		return ReadDataInFile(string(ch))
	case 'H':
		return ReadDataInFile(string(ch))
	case 'I':
		return ReadDataInFile(string(ch))
	case 'J':
		return ReadDataInFile(string(ch))
	case 'K':
		return ReadDataInFile(string(ch))
	case 'L':
		return ReadDataInFile(string(ch))
	case 'M':
		return ReadDataInFile(string(ch))
	case 'N':
		return ReadDataInFile(string(ch))
	case 'O':
		return ReadDataInFile(string(ch))
	case 'P':
		return ReadDataInFile(string(ch))
	case 'Q':
		return ReadDataInFile(string(ch))
	case 'R':
		return ReadDataInFile(string(ch))
	case 'S':
		return ReadDataInFile(string(ch))
	case 'T':
		return ReadDataInFile(string(ch))
	case 'U':
		return ReadDataInFile(string(ch))
	case 'V':
		return ReadDataInFile(string(ch))
	case 'W':
		return ReadDataInFile(string(ch))
	case 'X':
		return ReadDataInFile(string(ch))
	case 'Y':
		return ReadDataInFile(string(ch))
	case 'Z':
		return ReadDataInFile(string(ch))
	default:
		return nil
	}
}

func GetMatrix(ch rune, position int) []int {
	var matrixX []int
	shift := 0
	n := 0
	if ReadChar(ch) {
		switch position {
		case 1:
			shift = 0
		case 2:
			shift = 5
		case 3:
			shift = 10
		case 4:
			shift = 15
		case 5:
			shift = 20
		case 6:
			shift = 25
		case 7:
			shift = 30
		case 8:
			shift = 35
		case 9:
			shift = 40
		case 10:
			shift = 45
		}
		matrix := GetLettersMatrix(ch)
		for val := range matrix {
			switch {
			case matrix[val] >= 1 && matrix[val] <= 5:
				n = matrix[val] + shift
			case matrix[val] >= 6 && matrix[val] <= 10:
				n = matrix[val] - 5 + SizeX + shift
			case matrix[val] >= 11 && matrix[val] <= 15:
				n = matrix[val] - 10 + SizeX*2 + shift
			case matrix[val] >= 16 && matrix[val] <= 20:
				n = matrix[val] - 15 + SizeX*3 + shift
			case matrix[val] >= 21 && matrix[val] <= 25:
				n = matrix[val] - 20 + SizeX*4 + shift
			}
			matrixX = append(matrixX, n)
		}
	} else {
		return []int{}
	}

	return matrixX
}

func PrintMatrix(matrix []int) {
	for i := 1; i <= MaxSize; i++ {
		if contains(matrix, i) {
			fmt.Print("\033[34m*")
		} else {
			fmt.Print(" ")
		}
		if i%SizeX == 0 {
			fmt.Println()
		} else if i%SizeY == 0 {
			fmt.Print(" ")
		}
	}
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func main() {
	var s string
	var bigMatrix []int
	fmt.Printf("Введите строку:")
	fmt.Scan(&s)
	position := 0
	for _, ch := range s {
		position++
		char := LowerToUpper(ch)
		matrix := GetMatrix(char, position)
		if len(matrix) > 0 {
			for i := range matrix {
				bigMatrix = append(bigMatrix, matrix[i])
			}
		} else {
			fmt.Printf("Для этого символа: %s печати нет\n", string(ch))
		}
	}
	if position < 10 {
		PrintMatrix(bigMatrix)
	} else {
		fmt.Println("Вы ввели больше 10 символов")
	}
}
