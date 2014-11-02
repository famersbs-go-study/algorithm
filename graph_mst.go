package main

import(
	"fmt"
)

/**
 *
 	신장 트리 ( spanning tree )
		: 모든 정점을 포함 하고, 정점간 서로 연결 되면서 사이클이 존재 하지 않는 그래프

	최소 신장 트리 ( minimal spanning tree )
		: 모든 간선의 비용의 합이 최소인 신장 트리


	Kruscal 방법
		1. 간선중 비용이 적은 순으로 소트
		2. 가장 적은 비용이 드는 간선부터 차례대로 추가
		3. 추가중 사이클이 존재하는 간선은 제외
		4. 모든 정점이 연결될 때 까지 2-3 과정을 반복

	탐욕적 방법 ( Prim's algorithm )

		1. 임의의 정점을 선택
		2. 이 정점에서 다른 정점으로 갈 수 있는 최소 비용의 정점을 선택 ( 이 정점은 제외 )
		3. 이 정점에서 다른 정점으로 가는 비용과 기존의 비용과 비교후 더 작은 비용이 있으면 갱신
		4. 2-3 번 과정을 n(정점의 수) - 1 번 반복 

 */	


const (
	INFINIT int = -1
)

/**
 * 배열로 구성된 무향-가중치 그래프 표현
 *
 *	0 => 'A'
 *  ...  25 => 'Z' 
 *
 */
type ArrayGraph struct {
	board [][]int
}

func createArrayGraph( node_cnt int ) *ArrayGraph {

	var ret = &ArrayGraph{ };
	
	// Link state
	// example 3x3 graph
	// ba 
	// ca cb 
    ret.board = make([][]int, node_cnt)
    for i := range ret.board {
        ret.board[i] = make([]int, node_cnt - ( node_cnt - i ) )
        for j:= range ret.board[i] {
        	ret.board[i][j] = INFINIT
        }
    }

    return ret;
}

func itos( index int ) string {
	var ret = ""
	for index >= 0 {
		ret = string( ( index % 26 ) + 'A' ) + ret
		index = ( index / 26 ) - 1
	}
	return ret;
}
func stoi( index string ) int {
	var ret = 0
	for i := 0 ; i < len(index) ; i ++ {
		ret = ret * 26
		ret = ret + int( index[i] - 'A' ) + 1;
	}
	return ret - 1
}
func ( ag *ArrayGraph ) link( from string, to string, weight int ) {

	var from_i = stoi( from )
	var to_i = stoi( to )
	if( from_i == to_i ){
		return
	} 
	if( from_i < to_i ){
		var tmp = from_i
		from_i = to_i
		to_i = tmp
	}
	ag.board[ from_i ][ to_i ] = weight;

}



func main(){

	var ag = createArrayGraph(6)
		fmt.Println("Array Graph " )

	fmt.Println("-----------------------------------------")
	ag.link("A","B", 18 )
	ag.link("A","C", 30 )
	ag.link("A","D", 5 )
	ag.link("D","C", 10 )
	ag.link("D","E", 5 )
	ag.link("B","F", 15 )
	ag.link("C","F", 5 )
	ag.link("E","F", 5 )
	ag.link("F","B", 1 )

	fmt.Println( ag )

}

