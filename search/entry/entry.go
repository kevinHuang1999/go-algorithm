package main

func Slice() {
	s := make([]int, 10000, 10000)
	for idx, _ := range s {
		s[idx] = idx
	}
}

func main() {
	Slice()
}
