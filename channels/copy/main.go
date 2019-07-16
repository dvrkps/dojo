package main

func main() {
	ch := make(chan []byte)

	content := []byte{65, 66, 67}
	println("init", &content)

	go func() {
		println("send", &content)
		ch <- content
	}()

	r := <-ch
	println("receive", &r)
}
