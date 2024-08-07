package main

func main() {
	forever := make(chan bool) //vazio

	go func() {
		for i := 0; i < 10; i++ {
			print(i)
		}
		forever <- true
	}()
	<-forever
}
