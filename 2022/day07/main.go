package day07

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tom5760/advent-of-code/2022/inpututils"
)

func Parse(name string) (Directory, error) {
	var (
		root Directory
		cur  *Directory
	)

	err := inpututils.Scan(name, func(scanner *bufio.Scanner) error {
		for scanner.Scan() {
			line := scanner.Bytes()
			if len(line) == 0 {
				return fmt.Errorf("empty line?")
			}

			fields := strings.Fields(string(line))
			if len(fields) == 0 {
				return fmt.Errorf("empty command?")
			}

			if fields[0] == "$" {
				if len(fields) < 2 {
					return fmt.Errorf("prompt with no command?")
				}

				cmd := fields[1]

				switch cmd {
				case "cd":
					if len(fields) != 3 {
						return fmt.Errorf("cd requires path")
					}

					arg := fields[2]

					if cur == nil {
						// Assume we are at the root
						cur = &root
						cur.name = arg
						continue
					}

					if arg == ".." {
						cur = cur.parent
						continue
					}

					newDir := &Directory{
						parent: cur,
						name:   arg,
					}

					cur.entries = append(cur.entries, newDir)
					cur = newDir

				case "ls":
					continue

				default:
					return fmt.Errorf("invalid command '%v'", cmd)
				}

				continue
			}

			if len(fields) != 2 {
				return fmt.Errorf("invalid file")
			}

			if fields[0] == "dir" {
				// We'll capture directories when we cd into them.
				continue
			}

			size, err := strconv.Atoi(fields[0])
			if err != nil {
				return fmt.Errorf("failed to parse size: %w", err)
			}

			// Append file to current directory
			cur.entries = append(cur.entries, &File{
				name: fields[1],
				size: size,
			})
		}

		return nil
	})

	fmt.Print(root.String())

	return root, err
}

func Part1(root Directory) int {
	const limit = 100000

	var total int

	root.Walk(func(entry Entry) {
		// skip non-directories
		if _, ok := entry.(*Directory); !ok {
			return
		}

		size := entry.Size()

		if size <= limit {
			total += size
		}
	})

	return total
}

func Part2(root Directory) int {
	const (
		total    = 70000000
		required = 30000000
	)

	available := total - root.Size()
	need := required - available

	var candidates []int

	root.Walk(func(entry Entry) {
		// skip non-directories
		if _, ok := entry.(*Directory); !ok {
			return
		}

		size := entry.Size()

		if size >= need {
			candidates = append(candidates, size)
		}
	})

	var min int = math.MaxInt
	for _, candidate := range candidates {
		if candidate < min {
			min = candidate
		}
	}

	return min
}

type (
	Directory struct {
		parent  *Directory
		name    string
		entries []Entry
	}

	File struct {
		name string
		size int
	}

	Entry interface {
		Name() string
		Size() int
	}
)

func (d *Directory) Name() string { return d.name }
func (d *Directory) Size() int {
	var total int

	for _, e := range d.entries {
		total += e.Size()
	}

	return total
}

func (d *Directory) Walk(fn func(e Entry)) {
	fn(d)

	for _, e := range d.entries {
		if subdir, ok := e.(*Directory); ok {
			subdir.Walk(fn)
		} else {
			fn(e)
		}
	}
}

func (d *Directory) String() string {
	var sb strings.Builder

	d.stringRecurse(&sb, 0)

	return sb.String()
}

func (d *Directory) stringRecurse(sb *strings.Builder, level int) {
	for i := 0; i < level; i++ {
		sb.WriteByte('\t')
	}

	fmt.Fprintf(sb, "- %v (dir, size=%v)\n", d.name, d.Size())

	level++

	for _, entry := range d.entries {
		switch e := entry.(type) {
		case *File:
			for i := 0; i < level; i++ {
				sb.WriteByte('\t')
			}
			sb.WriteString(e.String())
		case *Directory:
			e.stringRecurse(sb, level)
		}
	}
}

func (f *File) Name() string { return f.name }
func (f *File) Size() int    { return f.size }
func (f *File) String() string {
	return fmt.Sprintf("- %v (file, size=%v)\n", f.name, f.size)
}
