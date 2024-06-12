package main

func suma(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func UneficientFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return UneficientFibonacci(n-1) + UneficientFibonacci(n-2)
}

func Efficientfibonacci(n int) int {
    if n <= 1 {
        return n
    }

    fib := make([]int, n+1)
    fib[0], fib[1] = 0, 1

    for i := 2; i <= n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }

    return fib[n]
}