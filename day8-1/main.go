package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type Stack []int64

func (s *Stack) Pop() int64 {
	x := (*s)[0]
	*s = (*s)[1:]
	return x
}

type Node struct {
	Metadata []int64
	Children []*Node
}

func (n *Node) Sum() int64 {
	var sum int64

	for _, v := range n.Metadata {
		sum += v
	}

	for _, child := range n.Children {
		sum += child.Sum()
	}

	return sum
}

func main() {
	stack := readInput(os.Stdin)
	root := parseTree(&stack)

	log.Println("metadata sum:", root.Sum())
}

func readInput(r io.Reader) Stack {
	var input []int64

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatalln("failed to parse input:", err)
			return nil
		}
		input = append(input, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return Stack(input)
}

func parseTree(stack *Stack) *Node {
	childCount := stack.Pop()
	metaCount := stack.Pop()

	children := make([]*Node, childCount)
	for i := range children {
		children[i] = parseTree(stack)
	}

	metadata := make([]int64, metaCount)
	for i := range metadata {
		metadata[i] = parseMetadata(stack)
	}

	return &Node{
		Metadata: metadata,
		Children: children,
	}
}

func parseMetadata(stack *Stack) int64 {
	return stack.Pop()
}
