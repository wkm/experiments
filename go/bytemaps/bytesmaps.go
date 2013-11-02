package main

type key []byte

func main() {
	m := make(map[key]string)
	m[[]byte("hi")] = "there"
}
