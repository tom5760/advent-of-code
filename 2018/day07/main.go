package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

const (
	// These are the values for the example.
	// WorkerCount  = 2
	// StepBaseTime = 0

	WorkerCount  = 5
	StepBaseTime = 60
)

type Constraint struct {
	Pre, Post string
}

type Node struct {
	Step string
	Prev []*Node
	Next []*Node
}

type Worker struct {
	Task     *Node
	DoneTime uint64
}

type Scheduler struct {
	clock     uint64
	numTasks  int
	taskOrder []string

	eligible map[*Node]bool
	complete map[*Node]bool
	workers  []*Worker
}

func (s *Scheduler) Schedule() {
	for !s.isFinished() {
		s.Tick()
	}
}

func (s *Scheduler) Tick() {
	// Find any workers that have completed their tasks.
	for _, worker := range s.workers {
		if worker.Task != nil && s.clock >= worker.DoneTime {
			s.done(worker.Task)
			worker.Task = nil
		}
	}

	// Assign tasks to idle workers
	for {
		worker := s.idleWorker()
		if worker == nil {
			break
		}

		worker.Task = s.next()
		if worker.Task == nil {
			break
		}

		worker.DoneTime = s.clock + taskTime(worker.Task.Step)
	}

	// If there is more work to do, advance the clock.
	s.clock = s.nextClock()
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

	delete(s.eligible, nodes[0])

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

func (s *Scheduler) idleWorker() *Worker {
	for _, worker := range s.workers {
		if worker.Task == nil {
			return worker
		}
	}
	return nil
}

func (s *Scheduler) done(task *Node) {
	s.complete[task] = true
	s.taskOrder = append(s.taskOrder, task.Step)

	for _, next := range task.Next {
		s.update(next)
	}
}

func (s *Scheduler) isFinished() bool {
	// If there are no eligible tasks, and nobody is working, we are
	if len(s.eligible) == 0 {
		for _, worker := range s.workers {
			if worker.Task != nil {
				return false
			}
		}
		return true
	}

	return false
}

func (s *Scheduler) nextClock() uint64 {
	if s.isFinished() {
		return s.clock
	}

	var minTime uint64 = math.MaxUint64

	for _, worker := range s.workers {
		if worker.Task != nil && worker.DoneTime < minTime {
			minTime = worker.DoneTime
		}
	}

	return minTime
}

func main() {
	constraints := readInput(os.Stdin)

	log.Println("(part 1) step order:", singleOrder(constraints))
	log.Println("(part 2) total time:", paralellOrder(constraints))
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

func singleOrder(constraints []Constraint) string {
	scheduler := Scheduler{
		numTasks: len(constraints),
		eligible: buildGraph(constraints),
		complete: make(map[*Node]bool),
		// Only add one worker for single-threaded work.
		workers: []*Worker{{}},
	}

	scheduler.Schedule()

	var sb strings.Builder
	for _, step := range scheduler.taskOrder {
		sb.WriteString(step)
	}

	return sb.String()
}

func paralellOrder(constraints []Constraint) uint64 {
	scheduler := Scheduler{
		numTasks: len(constraints),
		eligible: buildGraph(constraints),
		complete: make(map[*Node]bool),
	}

	for i := 0; i < WorkerCount; i++ {
		scheduler.workers = append(scheduler.workers, new(Worker))
	}

	scheduler.Schedule()

	return scheduler.clock
}

func taskTime(step string) uint64 {
	// ASCII 'A' == 65
	return StepBaseTime + (uint64(step[0]) - 64)
}
