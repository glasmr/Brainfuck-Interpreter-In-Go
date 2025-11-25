package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file>\n", os.Args[0])
	}
	var fileName string = os.Args[1]
	path, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	program := string(data)

	brainfuck(program)
}

func brainfuck(program string) {
	var memory [30000]byte
	var ip uint = 0
	var dp uint = 0
	var jumpStack []uint

	reader := bufio.NewReader(os.Stdin)

	for ip < uint(len(program)) {
		var lookAhead int = 0
		switch program[ip] {
		case '>':
			dp++
			ip++
			break
		case '<':
			dp--
			ip++
			break
		case '+':
			memory[dp]++
			ip++
			break
		case '-':
			memory[dp]--
			ip++
			break
		case '.':
			fmt.Printf("%c", memory[dp])
			ip++
			break
		case ',':
			charByte, err := reader.ReadByte()
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
			memory[dp] = charByte
			ip++
			break
		case '[':
			jumpStack = append(jumpStack, ip)
			if memory[dp] == 0 {
				var tempIP uint = ip
				for {
					tempIP++
					if program[tempIP] == '[' {
						lookAhead++
					} else if program[tempIP] == ']' {
						if lookAhead != 0 {
							lookAhead--
						} else {
							jumpStack = jumpStack[:len(jumpStack)-1]
							ip = tempIP + 1 //Jump to instruction right after ]
							break
						}
					}
				}
			} else {
				ip++
			}
			break
		case ']':
			if memory[dp] != 0 {
				ip = jumpStack[len(jumpStack)-1] + 1
			} else {
				jumpStack = jumpStack[:len(jumpStack)-1]
				ip++
			}
			break
		default:
			ip++
			break
		}
	}
}
