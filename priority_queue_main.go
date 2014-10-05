package main

import(
	"fmt"
)

type PQElement struct { 
	priority int
	data int
	left *PQElement
	right *PQElement
}

type PQ struct {
	top *PQElement
}

func createPQ() *PQ {
	return &PQ{ &PQElement{0,0, nil, nil} }
}

/**
 *
 * @param priority 1 ~ 100 사이의 값 
 **/
func ( pq *PQ ) putq( priority int, data int ) bool {

	if priority < 1 || priority > 100{
		return false;
	}

	var p_pq = pq.top
	var c_pq = pq.top.left

	// 우선순위 위치 찾기
	//    c_pq nil 만들기
	for nil != c_pq {
		if priority == c_pq.priority {
			p_pq = c_pq
			c_pq = c_pq.right	// 오른쪽 자식 탐색 ( 같은 레벨 )
			break;
		} else if priority > c_pq.priority {
			p_pq = c_pq
			c_pq = c_pq.left	// 왼쪽 자식 탐색 ( 하위 레벨 )
		} else if priority < c_pq.priority {
			c_pq = nil			// travel end
		}
	}

	// 조건에 따른 element 만들기
	var n_pq = &PQElement{ priority, data, nil, nil }
	if priority == p_pq.priority {
		p_pq.right = n_pq
	} else if priority > p_pq.priority {
		n_pq.left = p_pq.left
		p_pq.left = n_pq
	} else if priority < p_pq.priority {
		// 여기에 들어온다는건 travel이 잘못 되었다는것 ~~!
		fmt.Println("putq travel error : wrong priority ", priority )
		return false;
	}

	return true;
}

func ( pq PQ ) pop() ( int, bool ) {
	if pq.isEmpty() {
		return -1, false
	}

	var c_pq = pq.top.left;

	if nil != c_pq.right {
		
		pq.top.left = c_pq.right
		c_pq.right.left = c_pq.left

	} else {
		pq.top.left = c_pq.left
	}

	return c_pq.data, true
}

func ( pq PQ ) isEmpty() bool {
	return pq.top.left == nil
}

func ( pq PQ ) dump() {
	var c_pq = pq.top.left
	var next_pq *PQElement = nil

	if nil != c_pq && nil != c_pq.left {
		next_pq = c_pq.left
	}

	fmt.Println("pq dump")
	if nil == c_pq {
		fmt.Println("empty queue");
	} else {
		fmt.Print("[", c_pq.priority , "] : " );
		for nil != c_pq {

			fmt.Print( c_pq.data, " " )

			if nil != c_pq.right {
				c_pq = c_pq.right
			} else {
				c_pq = next_pq
				if nil != next_pq {
					next_pq = next_pq.left
				}

				if nil != c_pq {
					fmt.Println("");
					fmt.Print("[", c_pq.priority , "] : " );
				}
			}
		}
		fmt.Println("");
	}
}

func main(){
	var pq = createPQ();

	pq.dump();

	pq.putq( 1, 100 );
	pq.putq( 10, 1 );
	pq.putq( 5, 2 );
	pq.putq( 5, 4 );
	pq.putq( 6, 5 );

	pq.dump();

	var data int
	var result bool

	for i := 0 ; i < 6 ; i++ {
		data, result = pq.pop()
		fmt.Println("pop : ", data, result )
	}

	pq.dump();
}