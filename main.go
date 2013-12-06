package main

import (
	"errors"
	"fmt"
	"math"
    "io"
    "os"
	"strconv"
)

var stack = Stack{}

type Stack []float64

func (s *Stack) Pop() (float64, error) {
	result, err := s.Peek()
	if err != nil {
		return 0, err
	}
	*s = (*s)[:len(*s)-1]
	return result, nil
}

func (s *Stack) PopTwo() (float64, float64, error) {
	first, err := s.Pop()
	if err != nil {
		return 0, 0, err
	}
	second, err := s.Pop()
	if err != nil {
		return 0, 0, err
	}
	return first, second, nil
}

func (s *Stack) Push(input float64) {
	*s = append(*s, input)
}

func (s *Stack) Peek() (float64, error) {
	if len(*s) == 0 {
		return 0, errors.New("Cannot fetch item, stack is empty!")
	}
	return (*s)[len(*s)-1], nil
}

func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
        if err == io.EOF {
            os.Exit(0)
        } else if err != nil {
			panic(err)
		}
		number, err := strconv.ParseFloat(input, 64)
		if err == nil {
			stack.Push(number)
			continue
		}
		numError, ok := err.(*strconv.NumError)
		if !ok {
			fmt.Println(err)
			continue
		} else if numError.Err == strconv.ErrRange || numError.Err != strconv.ErrSyntax {
			fmt.Println(err)
			continue
		}
		switch input {
		case "+":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(first + second)
		case "-":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(second - first)
		case "*":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(first * second)
		case "/":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(second / first)
		case "%":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(math.Mod(second, first))
		case "~":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(math.Floor(second / first))
			stack.Push(math.Mod(second, first))
		case "^":
			first, second, err := stack.PopTwo()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(math.Pow(second, first))
		case "v":
			first, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
				continue
			}
			stack.Push(math.Sqrt(first))
		case "p":
			num, err := stack.Peek()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(num)
		case "c":
			stack = Stack{}
		case "pa":
			if len(stack) == 0 {
				fmt.Println("Stack is empty!")
				continue
			}
			for _, value := range stack {
				fmt.Printf("%v ", value)
			}
			fmt.Print("\n")
		default:
			fmt.Println("UNKNOWN COMMAND!")
			continue
		}
	}
}
