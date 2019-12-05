package main

import (
	"errors"
	"math"
)

const (
	pwLength          = 6
	maxAdjacentDigits = 2
)

// Errors returned from password validation.
var (
	ErrInvalidLength      = errors.New("invalid length")
	ErrMonotonicDigits    = errors.New("not monotonically increasing digits")
	ErrAdjacentDigits     = errors.New("no equivalent adjacent digits")
	ErrManyAdjacentDigits = errors.New("too many equivalent adjacent digits")
)

// ValidatePasswordV1 checks a password against the rules for part 1.
func ValidatePasswordV1(pw int) error {
	digits := getDigits(pw)

	// It is a six-digit number.
	if len(digits) != pwLength {
		return ErrInvalidLength
	}

	// Two adjacent digits are the same (like 22 in 122345).
	if !checkAdjacentDigits(digits) {
		return ErrAdjacentDigits
	}

	// Going from left to right, the digits never decrease; they only ever
	// increase or stay the same (like 111123 or 135679).
	if !checkMonotonicDigits(digits) {
		return ErrMonotonicDigits
	}

	return nil
}

// ValidatePasswordV2 checks a password against the stronger rules for part 2.
func ValidatePasswordV2(pw int) error {
	if err := ValidatePasswordV1(pw); err != nil {
		return err
	}

	digits := getDigits(pw)

	// The two adjacent matching digits are not part of a larger group of
	// matching digits.
	if !checkManyAdjacentDigits(digits) {
		return ErrManyAdjacentDigits
	}

	return nil
}

func getDigits(n int) []int {
	length := int(math.Log10(float64(n))) + 1
	digits := make([]int, 0, length)

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// reverse slice
	// From: https://github.com/golang/go/wiki/SliceTricks#reversing
	for i := len(digits)/2 - 1; i >= 0; i-- {
		opp := len(digits) - 1 - i
		digits[i], digits[opp] = digits[opp], digits[i]
	}

	return digits
}

// Going from left to right, the digits never decrease; they only ever increase
// or stay the same (like 111123 or 135679).
func checkMonotonicDigits(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			return false
		}
	}

	return true
}

// Two adjacent digits are the same (like 22 in 122345).
func checkAdjacentDigits(digits []int) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			return true
		}
	}

	return false
}

// The two adjacent matching digits are not part of a larger group of matching
// digits.
//
// NOTE: This means that there must be at least one pair of matching digits.
// It DOES NOT mean that every run of matching digits must be pairs.
func checkManyAdjacentDigits(digits []int) bool {
	count := 1

	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			count++
		} else {
			if count == 2 {
				return true
			}
			count = 1
		}
	}

	if count == 2 {
		return true
	}

	return false
}
