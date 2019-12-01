package main

import (
	"reflect"
	"runtime"
)

type Instruction struct {
	OP, A, B, C int
}

func addr(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] + r[i.B]
	return r
}

func addi(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] + i.B
	return r
}

func mulr(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] * r[i.B]
	return r
}

func muli(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] * i.B
	return r
}

func banr(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] & r[i.B]
	return r
}

func bani(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] & i.B
	return r
}

func borr(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] | r[i.B]
	return r
}

func bori(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A] | i.B
	return r
}

func setr(i Instruction, r Registers) Registers {
	r[i.C] = r[i.A]
	return r
}

func seti(i Instruction, r Registers) Registers {
	r[i.C] = i.A
	return r
}

func gtir(i Instruction, r Registers) Registers {
	if i.A > r[i.B] {
		r[i.C] = 1
	} else {
		r[i.C] = 0
	}
	return r
}

func gtri(i Instruction, r Registers) Registers {
	if r[i.A] > i.B {
		r[i.C] = 1
	} else {
		r[i.C] = 0
	}
	return r
}

func gtrr(i Instruction, r Registers) Registers {
	if r[i.A] > r[i.B] {
		r[i.C] = 1
	} else {
		r[i.C] = 0
	}
	return r
}

func eqir(i Instruction, r Registers) Registers {
	if i.A == r[i.B] {
		r[i.C] = 1
	} else {
		r[i.C] = 0
	}
	return r
}

func eqri(i Instruction, r Registers) Registers {
	if r[i.A] == i.B {
		r[i.C] = 1
	} else {
		r[i.C] = 0
	}
	return r
}

func eqrr(i Instruction, r Registers) Registers {
	if r[i.A] == r[i.B] {
		r[i.C] = 1
	} else {
		r[i.C] = 0
	}
	return r
}

type OpCode func(Instruction, Registers) Registers

func (o OpCode) String() string {
	return runtime.FuncForPC(reflect.ValueOf(o).Pointer()).Name()
}

var ISA = []OpCode{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti,
	gtir, gtri, gtrr, eqir, eqri, eqrr}

func Check(sample Sample) []OpCode {
	var opcodes []OpCode

	for _, opcode := range ISA {
		result := opcode(sample.Instruction, sample.Before)
		if result == sample.After {
			opcodes = append(opcodes, opcode)
		}
	}

	return opcodes
}

func ReduceISA(sample Sample) {
	for i, opcode := range ISA {
		result := opcode(sample.Instruction, sample.Before)
		if result == sample.After {
			ISA = append(ISA[:i], ISA[i+1:]...)
		}
	}
}
