package main

import (
	"fmt"
	"math"
	"math/bits"
)

type Segment int

const (
	SegmentA Segment = iota
	SegmentB
	SegmentC
	SegmentD
	SegmentE
	SegmentF
	SegmentG
)

type (
	Signal uint8

	Signals struct {
		Samples [10]Signal // random digits 0-9
		Output  [4]Signal  // output value
	}
)

func ParseSignal(segments []byte) (Signal, error) {
	var d Signal
	if len(segments) > 7 {
		return d, fmt.Errorf("got more than 7 segments: %v", len(segments))
	}

	for _, segment := range segments {
		if segment < 'a' || segment > 'g' {
			return d, fmt.Errorf("unexpected segment '%v'", segment)
		}

		d |= 1 << (segment - 97)
	}

	return d, nil
}

func (s Signal) String() string {
	return fmt.Sprintf("%#.8b", s)
}

func (s *Signals) Decode() uint64 {
	display := make(map[Signal]uint8, 10)

	var candidates [7]byte

	for i := range candidates {
		candidates[i] = 0b01111111
	}

	// Find well-known samples
	for _, sample := range s.Samples {
		s := uint8(sample)
		switch bits.OnesCount8(s) {
		case 2: // digit 1
			display[sample] = 1

			candidates[SegmentC] &= s
			candidates[SegmentF] &= s

			candidates[SegmentA] &^= s
			candidates[SegmentB] &^= s
			candidates[SegmentD] &^= s
			candidates[SegmentE] &^= s
			candidates[SegmentG] &^= s

		case 4: // digit 4
			display[sample] = 4

			candidates[SegmentB] &= s
			candidates[SegmentC] &= s
			candidates[SegmentD] &= s
			candidates[SegmentF] &= s

			candidates[SegmentA] &^= s
			candidates[SegmentE] &^= s
			candidates[SegmentG] &^= s

		case 3: // digit 7
			display[sample] = 7

			candidates[SegmentA] &= s
			candidates[SegmentC] &= s
			candidates[SegmentF] &= s

			candidates[SegmentB] &^= s
			candidates[SegmentD] &^= s
			candidates[SegmentE] &^= s
			candidates[SegmentG] &^= s

		case 7: // digit 8
			display[sample] = 8
		}
	}

	// We should know what segment a is now.
	if bits.OnesCount8(uint8(candidates[SegmentA])) != 1 {
		panic("did not find segment a")
	}

	// There should be 3 samples with 5 segments, corresponding to digits 2 3 5.
	// Each of them should have the 'g' segment set.
	for _, sample := range s.Samples {
		s := uint8(sample)
		if bits.OnesCount8(s) != 5 {
			continue
		}

		candidates[SegmentG] &= s
	}

	// We should know what segment g is now.
	if bits.OnesCount8(uint8(candidates[SegmentG])) != 1 {
		panic("did not find segment g")
	}

	for i := range candidates {
		if i == int(SegmentG) {
			continue
		}

		candidates[i] &^= candidates[SegmentG]
	}

	// We should know what segment e is now.
	if bits.OnesCount8(uint8(candidates[SegmentE])) != 1 {
		panic("did not find segment e")
	}

	// We can now find digit 2, it's the only 5 segment sample with segment 'e'.
	for _, sample := range s.Samples {
		s := uint8(sample)
		if bits.OnesCount8(s) == 5 && (s&candidates[SegmentE] == candidates[SegmentE]) {
			display[sample] = 2

			// Digit 2 doesn't use segment f, only c.
			candidates[SegmentC] &= s
			candidates[SegmentF] &^= s

			// Digit 2 doesn't use segment b, only d.
			candidates[SegmentD] &= s
			candidates[SegmentB] &^= s
		}
	}

	// Should now have all candidates figured out.
	for _, c := range candidates {
		if bits.OnesCount8(uint8(c)) != 1 {
			panic("didn't figure out all candidates")
		}
	}

	// Figure out the rest of the digits: 0 3 5 6 9
	zero := candidates[SegmentA] | candidates[SegmentB] | candidates[SegmentC] |
		candidates[SegmentE] | candidates[SegmentF] | candidates[SegmentG]

	three := candidates[SegmentA] | candidates[SegmentC] | candidates[SegmentD] |
		candidates[SegmentF] | candidates[SegmentG]

	five := candidates[SegmentA] | candidates[SegmentB] | candidates[SegmentD] |
		candidates[SegmentF] | candidates[SegmentG]

	six := candidates[SegmentA] | candidates[SegmentB] | candidates[SegmentD] |
		candidates[SegmentE] | candidates[SegmentF] | candidates[SegmentG]

	nine := candidates[SegmentA] | candidates[SegmentB] | candidates[SegmentC] |
		candidates[SegmentD] | candidates[SegmentF] | candidates[SegmentG]

	for _, sample := range s.Samples {
		s := uint8(sample)
		switch {
		case zero == s:
			display[sample] = 0

		case three == s:
			display[sample] = 3

		case five == s:
			display[sample] = 5

		case six == s:
			display[sample] = 6

		case nine == s:
			display[sample] = 9
		}
	}

	// 	log.Printf("NINE: %.7b", nine)
	// 	log.Println("CANDIDATES:")
	// 	for i, c := range candidates {
	// 		log.Printf("%c: %.7b", byte(97+i), c)
	// 	}
	//
	// 	log.Println("DIGITS:")
	// 	for sample, digit := range display {
	// 		log.Printf("%.7b: %v", sample, digit)
	// 	}

	// Print output
	var output uint64

	for i, out := range s.Output {
		output += uint64(display[out]) * uint64(math.Pow10(len(s.Output)-(i+1)))
	}

	return output
}
