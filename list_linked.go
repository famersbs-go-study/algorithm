package main

import(
	"fmt"
)

type Node struct { 
	data int
	prev *Node
	next *Node
}

func createNode( ) *Node {
	return &Node{ 0, nil, nil }
}

func (h *Node) insertAfter( data int ) *Node {
	
	var new_node = Node{ data, h, h.next }
	if nil != h.next {
		h.next.prev = &new_node;
	}
	h.next = &new_node;

	return &new_node;
}

func (h Node) findByValue( data int ) *Node {
	var curr *Node = h.next
	for nil != curr {

		if data == curr.data {
			return curr
		} 

		curr = curr.next
	}

	return nil
}

func (h Node) Next() *Node{
	return h.next
}

func (h *Node) delete( ) bool {

	if nil == h.prev {
		fmt.Println("do not delete head!!")
		return false;
	}

	h.prev.next = h.next
	if nil != h.next {
		h.next.prev = h.prev
	}

	return true

}

func (h Node) dump(){
	var curr *Node = h.next
	for nil != curr {
		fmt.Print( curr.data , " -> ");
		curr = curr.next
	}
	fmt.Println(" end");
}

func main(){
	fmt.Println("Start linked list");
	var list = createNode();

	fmt.Println("Insert 10 to 50 ");
	list.insertAfter( 10 ).
		 insertAfter( 20 ).
		 insertAfter( 30 ).
		 insertAfter( 40 ).
		 insertAfter( 50 );
	list.dump();

	fmt.Println("Find 30");
	var data_30 = list.findByValue(30);
	fmt.Println( data_30 );

	fmt.Println("Delete 30");
	data_30.delete();
	list.dump();


}