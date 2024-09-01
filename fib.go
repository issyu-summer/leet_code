package main

func main() {

}

func fib(n int) int {
	memo := make(map[int]int)
	return fibMem(n, memo)
}

func fibMem(n int, mem map[int]int) int {
	if n <= 1 {
		return n
	}
	if v, ok := mem[n]; ok {
		return v
	}
	mem[n] = fibMem(n-1, mem) + fibMem(n-2, mem)
	return mem[n]
}
