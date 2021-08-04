package BelajarGolangGoroutine

import(
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T){
	group := &sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{}{
			return "New"
		},
	}

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func(){
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("Selesai")

}