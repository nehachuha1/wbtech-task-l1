package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	ID      int
	Message string
}

func StartWorker(workerID int, out chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range out {
		time.Sleep(3 * time.Second)
		fmt.Printf("Worker ID %v | Task ID %v | Message: %v\n",
			workerID, task.ID, task.Message)
	}
	fmt.Println("Exit from worker...")
}

func GenerateTasks(out chan Task, quit chan bool, taskCount int) {
	for i := 0; i < taskCount; i++ {
		newTask := Task{
			ID:      i + 1,
			Message: fmt.Sprintf("This is new task with ID %v", i+1),
		}
		select {
		case _, ok := <-quit:
			if !ok {
				return
			}
		default:
			out <- newTask
		}
	}
	close(out)
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	const workerCount = 2
	var wg sync.WaitGroup

	dataChannel := make(chan Task)
	quitChan := make(chan bool)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go StartWorker(i+1, dataChannel, &wg)
	}

	go func() {
		select {
		case <-sigchan:
			fmt.Println("Received signal from sigchan. Exit from program")
			quitChan <- true
			close(dataChannel)
		}
	}()

	go GenerateTasks(dataChannel, quitChan, 10)

	wg.Wait()
	fmt.Println("All tasks are completed!")
}
