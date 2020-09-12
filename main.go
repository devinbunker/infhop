package main

import (
	"fmt"
	"errors"
	"uhop/pancake"
	"os"
	"bufio"
)

func main() {
	caseCount := 1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		if in == "" {
			continue
		}

		inList, err := parseInput(in)
		if err != nil {
			fmt.Printf("failed to parse input: %s\n", err.Error())
			return
		}

		stack := pancake.Stack{}
		stack.Add(inList...)
		ops, ok := stack.Normalize()
		if !ok {
			fmt.Printf("failed to normalize pancake stack\n")
			return
		}
		fmt.Printf("Case #%d: %d\n", caseCount, ops)
		caseCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error while scanning input: %s\n", err.Error())
	}
}

func parseInput(in string) (out []bool, err error) {
	// preallocate enough space to fit the entire input
	out = make([]bool, 0, len(in))

	for _, val := range in {
		if val == '+' {
			out = append(out, true)
		} else if val == '-' {
			out = append(out, false)
		} else {
			return out, errors.New(fmt.Sprintf("invalid character: %s", val))
		}
	}

	return out, nil
}

