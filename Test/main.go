package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	Solution("011100")
}

func Solution(S string) int {
    //convert binary string into int
    var counter int64 =0
	num := big.NewInt(0)
    //convert binary string into int
    if num, err:= strconv.ParseInt(S,2,64); err != nil{
        fmt.Println(err)
	}else{
        _,counter = CountProcess(num,-1)
    }

	fmt.Println(counter)
    return int(counter )
    

}
func CountProcess(number int64,counter int64)(int64,int64){
	fmt.Println("Masuk")
	fmt.Printf("Counter:%d\n",counter);
    counter += 1
    if number <= 0{
        return 0,counter
    }

    if number % 2 == 0{
        number /=2
    }else{
        number -= 1
    }	

    return CountProcess(number,counter)

    
}
