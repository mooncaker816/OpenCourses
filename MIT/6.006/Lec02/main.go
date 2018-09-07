package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
)

func main() {
	f1Name, f2Name := "a.txt", "b.txt"
	m1, err := getWordMap(f1Name)
	if err != nil {
		panic("can not build map for file1")
	}
	m2, err := getWordMap(f2Name)
	if err != nil {
		panic("can not build map for file2")
	}
	fmt.Printf("The distance between the documents is: %0.6f (radians)\n", vectorAngle(m1, m2))
}

func getWordMap(path string) (map[string]float64, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(f))
	scanner.Split(scanAlphaNumToken)
	m := make(map[string]float64)
	for scanner.Scan() {
		m[scanner.Text()]++
	}
	return m, nil
}

func scanAlphaNumToken(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// skip non-alphanumeric char
	start := 0
	for ; start < len(data); start++ {
		if isAlphaNumeric(data[start]) {
			break
		}
	}
	// scan until first non-alphanumeric char
	for i := start; i < len(data); i++ {
		if !isAlphaNumeric(data[i]) {
			return i + 1, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}

func isAlphaNumeric(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9'
}

func innerProduct(m1, m2 map[string]float64) float64 {
	sum := 0.0
	if len(m1) <= len(m2) {
		for k, v1 := range m1 {
			if v2, ok := m2[k]; ok {
				sum += v1 * v2
			}
		}
		return sum
	}
	for k, v2 := range m2 {
		if v1, ok := m1[k]; ok {
			sum += v1 * v2
		}
	}
	return sum
}

func vectorAngle(m1, m2 map[string]float64) float64 {
	numerator := innerProduct(m1, m2)
	denominator := math.Sqrt(innerProduct(m1, m1) * innerProduct(m2, m2))
	return math.Acos(numerator / denominator)
}
