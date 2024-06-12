package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	/* 	total := suma(5, 5)

	   	if total != 10 {
	   		t.Errorf("Sum was incorrect, got %d, expected %d", total, 10)
	   	}

	*/

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := suma(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was wrong, expected %d, got %d", item.n, total)
		}
	}

}

func TestGetMax(t *testing.T) {
	/*
		if GetMax(12, 15) != 15 {
			t.Errorf("GetMax was incorrect, got %d, expected %d", GetMax(12, 15), 15)
		}
	*/

	table := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, item := range table {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("GetMax was wrong")
		}

	}

}

func TestBadFibonacci(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range table {
		fib := UneficientFibonacci(item.a)

		if fib != item.n {
			t.Errorf("Fibonacci was wrong, expected %d, got %d", item.n, fib)
		}
	}
}

func TestGoodFibonacci(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range table {
		fib := Efficientfibonacci(item.a)

		if fib != item.n {
			t.Errorf("Fibonacci was wrong, expected %d, got %d", item.n, fib)
		}
	}
}
