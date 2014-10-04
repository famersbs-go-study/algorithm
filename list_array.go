package main

import(
	"fmt"
)


/**
*
* array based list
*
*	init
*	insert
*	delete
*	clear
*	find
*
*	point value => int
*
**/
type List_BA struct { 
	idx []int
	data []int
	used []bool
	head int
}
func (h *List_BA) init( size int ) bool { 
	h.idx = make([]int, size);
	h.data = make([]int, size);
	h.used = make([]bool, size);
	h.head = -1;

	for i := 0; i < len(h.idx); i++ {
    	h.idx[i] = -1;
    	h.used[i] = false;
    }

	return true; 
}
func (h List_BA) findEmptyIDX( ) int {
	for i := 0 ; i < len(h.used) ; i++ {
		if h.used[i] == false {
			return i;
		}
	}
	return -1;
}

func (h List_BA) findBeforeIDX( index int ) int {
	if h.head == index {
		return -1;
	}

	for i := 0 ; i < len( h.idx ) ; i++ {
		if h.idx[i] == index && h.used[i] {
			return i
		}
	}
	return -1;
}

func (h List_BA) findByValue( data int ) int {
	for i := 0 ; i < len( h.idx ) ; i++ {
		if h.used[i] && h.data[i] == data {
			return i
		}
	}
	return -1;
}

/**
 * @ return  삽입에 성공하면 삽입된 index 리턴, 실패 하면 -1 리턴
 **/
func (h *List_BA) insertAfter( params ...int ) int {
	var index int = -1;
	var value int;

	if len( params ) <= 0 {
		fmt.Println( "insertAfter : parameter is empty." );
		return -1;
	}
	value = params[0];
	if len( params ) >= 2 {
		index = params[1];
	}

	// 오류 조건 탐색 - 잘못된 인덱스 지정
	if index != -1 && h.head == -1 {
		// 리스트가 비었는데, 인덱스가 -1이 아닐때 오류
		fmt.Println( "insertAfter : This list is empty.[", index, "]" );
		return -1;
	} else if index >= len( h.idx ) {
		// 리스트가 비었는데, 인덱스가 -1이 아닐때 오류
		fmt.Println( "insertAfter : Out of bound index.[", index, "]" );
		return -1;
	} else if index != -1 && !h.used[index] {
		// 끊어진 링크인지 확인 (?)
		fmt.Println( "insertAfter : Index link is broken.[", index, "]" );
		return -1;
	}

	// 삽입 

	// 넣을 공간 확인 하기
	// 지정된 인덱스 확인 하기
		// 지정된 인덱스가 없다면, head를 지정된 인덱스로 지정 하고, head 를 next로 설정
		// 지정된 인덱스가 있다면, 지정된 인덱스의 뒤에 삽입
		//	 만약 지정된 인덱스가 사용중이 아니라면 ( 즉 깨진 인덱스라면 ) 오류 발생
	var empty_idx = h.findEmptyIDX();
	if -1 == empty_idx {
		fmt.Println("insertAfter : This list is full ");
		return -1;
	}

	h.data[ empty_idx ] = value;
	h.used[ empty_idx ] = true;

	if index == -1 {
		h.idx[ empty_idx ] = h.head;
		h.head = empty_idx;
	}else {
		h.idx[ empty_idx ] = h.idx[ index ];
		h.idx[ index ] = empty_idx;
	}

	return empty_idx;
}

func (h *List_BA) deleteByIdx( index int ) bool {

	// index 유효성 확인
		// range
	if index < 0 || index >= len( h.idx ) {
		fmt.Println("delete : Out of Range ", index );
		return false;
	}
		// used flag
	if !h.used[index] {
		fmt.Println("delete : Index is not used ", index );
		return false;
	}

	// delete 절차 처리
	var before_idx = -1;
	if index != h.head {
		before_idx = h.findBeforeIDX( index );
	}

	h.used[ index ] = false;
	if before_idx != -1 {
		h.idx[ before_idx ] = h.idx[ index ];
	} else {
		h.head = h.idx[ index ];
	}

	return true;
}

func (h *List_BA) deleteByData( data int ) bool {
	var idx = h.findByValue( data );
	return h.deleteByIdx( idx );
}

func (h List_BA) dump( ) {

	if -1 == h.head {
		fmt.Println("is Empty");
		return;
	}

	var idx = h.head
	var cnt = 1;

	for idx != -1 {
		fmt.Println( "[", cnt, "/", idx , "] : ", h.data[ idx ] );
		idx = h.idx[idx];
		cnt += 1 ;
	}
}

func main(){

	var l = List_BA{};
	var idx_20 = -1;
	var idx_30 = -1;

	fmt.Println("Init : ", l.init(5) );
	l.dump();

	fmt.Println("insert 10 to 40");
	l.insertAfter( 10 );
	idx_20 = l.insertAfter( 20 );
	idx_30 = l.insertAfter( 30 );
	l.insertAfter( 40 );
	l.dump();

	fmt.Println(" insert 25 after 20");
	l.insertAfter( 25, idx_20 );
	l.dump();
	
	fmt.Println("Over insert ");
	l.insertAfter( 50 );

	fmt.Println("Delete 30 ( by idx )");
	l.deleteByIdx( idx_30 );
	l.dump();

	fmt.Println("Delete 40 ( by data )")
	l.deleteByData( 40 );
	l.dump();

	fmt.Println("find 25 : ", l.findByValue(25) );
	fmt.Println("find 100 : ", l.findByValue(100) );

}