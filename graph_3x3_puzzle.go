package main

import(
	"fmt"
)

func main(){

	var buffer [256]byte;

	buffer[0] = 0;
	buffer[1] = 1;

	var slice_buf = buffer[100:200];

	fmt.Println("buffer ", buffer );
	fmt.Println("buffer slice ", slice_buf );

}