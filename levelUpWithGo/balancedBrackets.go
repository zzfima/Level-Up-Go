// this task is about: using stack

package main

import (
	"flag"
	"log"
	"strings"

	"github.com/golang-collections/collections/stack"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	var parentheses stack.Stack
	for _, c := range strings.Split(expr, "") {
		if c == "(" || c == "[" || c == "{" {
			parentheses.Push(c)
		} else if c == ")" {
			p := parentheses.Pop()
			if p != "(" {
				return false
			}
		} else if c == "]" {
			p := parentheses.Pop()
			if p != "[" {
				return false
			}
		} else if c == "}" {
			p := parentheses.Pop()
			if p != "{" {
				return false
			}
		}
	}

	return true
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

// 1 rename to main
// 2 run in terminal: go run .\balancedBrackets.go -expr "(1+3+[5*6]*{4/6+(5+7+[7/8])})"
func mainBalancedBrackets() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
