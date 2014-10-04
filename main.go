// export GOPATH=/Users/famersbs/project/go_test/test
package main

import(
	"fmt"
	"time"
)


/** Example chan + range + close **/
func fibonacci( n int, c chan int ){
	x, y := 0, 1
	for i:=0 ; i < n ; i ++ {
		c <- x
		x, y = y, x+y
		time.Sleep( 100 * time.Millisecond )
	}
	close(c)
}
func sample_change_range_close(){
	c := make( chan int, 10 )
	go fibonacci( cap(c), c )
	for i := range c {
		fmt.Println( i )
	}
}

/** Example channel **/
func fibonacci_select_test( c, quit chan int ){
	x, y := 0, 1
	fmt.Println(" Start fibonacci... ");
	for{
		select{
		case c <- x:
			x, y = y, x+ y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func select_channel_test(){
	c := make( chan int )
	quit := make( chan int )

	//c <- 10

	go func(){
		fmt.Println(" Start func... ");
		for i := 0; i < 10 ; i ++ {
			fmt.Println( <- c )
		}
		quit <- 0
	}()
	time.Sleep( 1000 * time.Millisecond );
	fibonacci_select_test( c, quit )
}

func main(){
	//sample_change_range_close();
	select_channel_test();
}