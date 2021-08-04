package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)


func TestContext( t *testing.T){
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T){
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")
	
	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))
}

func CreateCounter(ctx context.Context) chan int{
	destination := make(chan int)

	go func ()  {
		defer close(destination)
		counter := 1
		for{
			select{
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)//simulasi slow
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T){
	fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)	
	defer cancel()//mengirim signal cancel ke context
	defer fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())
	
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	
	time.Sleep(2 * time.Second)

}

func TestContextWithTimeout(t *testing.T){
	fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)	
	defer cancel()//mengirim signal cancel ke context
	defer fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())
	
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n) 
	}
	
	time.Sleep(2 * time.Second)

}


func TestContextWithDeadline(t *testing.T){
	fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))	
	defer cancel()//mengirim signal cancel ke context
	defer fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())
	
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutiner : ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("counter", n) 
	}
	
	time.Sleep(2 * time.Second)

}