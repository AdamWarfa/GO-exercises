package exercises

import (
	"fmt"
	"sync"
	"time"
)

type blud struct {
	message string
}

func RoutineTest() {
	var waitGroup sync.WaitGroup

	c := make(chan string)
	d := make(chan blud)
	e := make(chan int)

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go stampTime(i, &waitGroup, c, d, e)
	}

	waitGroup.Wait()
	close(c)
	close(d)
	fmt.Println("done med main flow")

}

func stampTime(id int, waitGroup *sync.WaitGroup, c chan string, d chan blud, e chan int) {
	defer waitGroup.Done()
	switch id {
	case 3:
		message := "Vi hørte noget"
		c <- message
		d <- blud{
			message: "big bludclart",
		}

	case 6:
		select {
		case receivedString := <-c:
			fmt.Println(receivedString)
		case receivedFromBlud := <-d:
			fmt.Println(receivedFromBlud.message)
		// case receivedInt := <-e:
		// 	fmt.Println(receivedInt)
		default:
			fmt.Println("No messages received")
		}
		return
	}

	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println(fmt.Sprintf("timeren med id'et %v er nu done diego", id))

}
