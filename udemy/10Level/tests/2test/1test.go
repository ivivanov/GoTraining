package main

type Message struct {
	str  string
	wait chan bool
}

func main() {
	msg1 := <-c; fmt.Println(msg1.str)
}
