package main

import (
	"fmt"
	"errors"
	"uhop/pancake"
	"os"
	"bufio"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if ok := scanner.Scan(); !ok {
		fmt.Printf("error while reading number of input lines\n")
		return
	}

	rawInput := scanner.Text()
	lines, err := strconv.Atoi(rawInput)
	if err != nil {
		fmt.Printf("didn't understand \"%s\" as the number of input lines\n", rawInput)
		return
	}

	// reject out-of-range values for the number of input lines
	if lines < 1 || lines > 100 {
		fmt.Printf("invalid number of input lines specified: %d\n", lines)
		return
	}

	caseCount := 1
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

		// stop if we've read 'lines' number of inputs; the
		// specification doesn't say what to do in this case,
		// but it doesn't seem to make sense to have the file
		// specify a limit and for us to then ignore it
		if caseCount > lines {
			break
		}
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

