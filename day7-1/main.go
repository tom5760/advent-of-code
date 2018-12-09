package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Constraint struct {
	Pre, Post string
}

type Node struct {
	Step string
	Prev []*Node
	Next []*Node
}

type Scheduler struct {
	eligible map[*Node]bool
	complete map[*Node]bool
}

func (s *Scheduler) Schedule() []*Node {
	order := make([]*Node, 0, len(s.eligible))

	for len(s.eligible) > 0 {
		node := s.next()

		order = append(order, node)
		s.complete[node] = true
		delete(s.eligible, node)

		for _, next := range node.Next {
			s.update(next)
		}
	}

	return order
}

func (s *Scheduler) next() *Node {
	if len(s.eligible) == 0 {
		return nil
	}

	nodes := make([]*Node, 0, len(s.eligible))
	for node := range s.eligible {
		nodes = append(nodes, node)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Step < nodes[j].Step
	})

	return nodes[0]
}

func (s *Scheduler) update(n *Node) {
	// If this step is complete, skip it.
	if s.complete[n] {
		return
	}
	// Unless all of this step's parents are complete, skip it.
	for _, prev := range n.Prev {
		if !s.complete[prev] {
			return
		}
	}
	s.eligible[n] = true
}

func main() {
	constraints := readInput(os.Stdin)
	eligible := buildGraph(constraints)

	scheduler := Scheduler{
		eligible: eligible,
		complete: make(map[*Node]bool),
	}

	var sb strings.Builder
	for _, node := range scheduler.Schedule() {
		sb.WriteString(node.Step)
	}

	log.Println("order is:", sb.String())
}

const inputFormat = "Step %s must be finished before step %s can begin."

func readInput(r io.Reader) []Constraint {
	var constraints []Constraint

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var c Constraint
		if _, err := fmt.Sscanf(scanner.Text(), inputFormat, &c.Pre, &c.Post); err != nil {
			log.Fatalln("failed to scan input line:", err)
			return nil
		}

		constraints = append(constraints, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("failed to read input:", err)
		return nil
	}

	return constraints
}

// buildGraph returns a set of unconstrained nodes.
func buildGraph(constraints []Constraint) map[*Node]bool {
	nodes := make(map[string]*Node)

	getNode := func(step string) *Node {
		node, ok := nodes[step]
		if !ok {
			node = &Node{Step: step}
			nodes[step] = node
		}
		return node
	}

	for _, c := range constraints {
		preNode := getNode(c.Pre)
		postNode := getNode(c.Post)

		postNode.Prev = append(postNode.Prev, preNode)
		preNode.Next = append(preNode.Next, postNode)
	}

	// Find unconstrained nodes
	startNodes := make(map[*Node]bool)
	for _, node := range nodes {
		if len(node.Prev) == 0 {
			startNodes[node] = true
		}
	}

	return startNodes
}
