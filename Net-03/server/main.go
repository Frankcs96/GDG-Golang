package main

func main() {
	cs := NewChatServer("tcp", ":8081")

	cs.Start()

}
