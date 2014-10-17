package main

import(
	"fmt"
)

const (
	INFINIT int = -1
)

/**
 * 배열로 구성된 유향-가중치 그래프 표현
 *
 *	0 => 'A'
 *  ...  25 => 'Z' 
 *
 */
type ArrayGraph struct {
	board [][]int
	d_board [][] int
	d_pre_board[][] int
}

func createArrayGraph( node_cnt int ) *ArrayGraph {

	var ret = &ArrayGraph{ };
	
	// Link state
    ret.board = make([][]int, node_cnt)
    for i := range ret.board {
        ret.board[i] = make([]int, node_cnt)
        for j:= range ret.board[i] {
        	ret.board[i][j] = INFINIT
        }
    }
	
	// All nodes shortest path 
	ret.d_board = make( [][]int, node_cnt )
	for i := range ret.d_board {
        ret.d_board[i] = make( []int, node_cnt )

        for j:= range ret.d_board[i] {
        	ret.d_board[i][j] = INFINIT
        }
    }

    // All nodes shortest path 
	ret.d_pre_board = make( [][]int, node_cnt )
	for i := range ret.d_pre_board {
        ret.d_pre_board[i] = make( []int, node_cnt )

        for j:= range ret.d_pre_board[i] {
        	ret.d_pre_board[i][j] = INFINIT
        }
    }

    //fmt.Println( ret.d_board );

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

	fmt.Println("Link ", from , " -> ", to, "(" , weight , ")" )
	ag.board[ stoi(from) ][ stoi(to) ] = weight;

}


func contains(s []int, e int) bool {
    for a := range s { 
    	if a == e { 
    		return true 
    	} 
    }
    return false
}

/** _dijkstra 에서 쓰이는 모두 방문했는가  함수 **/
func isAllVisit( q []int ) bool{
	for i := range q {
		if 0 == q[i] {
			return false
		}
	}
	return true
}

func ( ag *ArrayGraph ) _dijkstra( cur int ){

	var cur_d_board = ag.d_board[cur]
	var cur_d_p_board = ag.d_pre_board[cur]

	// current node is weight 0
	cur_d_board[cur] = 0;

	// job q & s 
	var Q = make( []int, len( cur_d_board ) ) // 0 은 아직 가지 안은것, 1이면 이미 간것
	//var S = make( []int, len( cur_d_board ) ) // 0 은 아직 가지 안은것, 1이면 이미 간것

	for !isAllVisit( Q ) {

		var min_node = -1;
		var min_val = INFINIT;

		// Find Shortest D node
		for i := range Q {
			if  Q[i] == 0 &&	// Q 에서 사용되지 않았고
				( ( min_val == INFINIT && cur_d_board[ i ] != INFINIT ) || ( cur_d_board[ i ] != INFINIT && cur_d_board[ i ] < min_val ) ){ // 현재의 min 보다 작은 것 
				
				min_node = i
				min_val = cur_d_board[i];

				//fmt.Println(itos(cur), " Cur min node ", itos( min_node ) );

			}
		}

		if min_node == -1 {
			//fmt.Println("Do Not found another link")
			break;
		}

		// Q에서 제거 하고
		Q[ min_node ] = 1;

		// min_node 에서 연결 되는 경로 들 중에서 가장 작은 값으로 d를 갱신한다.
		for i := range ag.board[ min_node ] {
			if 	ag.board[ min_node ][i] != INFINIT &&
				( cur_d_board[ i ] == INFINIT || cur_d_board[ i ] > cur_d_board[ min_node ] + ag.board[ min_node ][i] ) {

				cur_d_board[ i ] = cur_d_board[ min_node ] + ag.board[ min_node ][i]
				cur_d_p_board[ i ] = min_node

				//fmt.Println("Change Node ", itos( cur ), "->", itos(i), " : ", itos(min_node), " (" , cur_d_board[i] , ")" );

			}

		}

	}

}

func ( ag *ArrayGraph ) dijkstra( ){
	for cur_node := range ag.board {
		ag._dijkstra( cur_node )
	}
}

func ( ag *ArrayGraph ) shortist( from string, to string ) string {

	var from_i = stoi(from)
	var to_i = stoi(to)

	var total_weight = ag.d_board[ from_i ][ to_i ]

	if INFINIT == total_weight {
		return "Can Not Find Path : " + from + " -> " + to
	}

	var ret = to
	var cur_to_i = to_i;
	for from_i != ag.d_pre_board[from_i][cur_to_i] {

		//fmt.Println( itos(cur_to_i), ag.d_pre_board[from_i][cur_to_i] )

		ret = itos( ag.d_pre_board[from_i][cur_to_i] ) + " -> " + ret
		cur_to_i = ag.d_pre_board[from_i][cur_to_i]
	}

	ret = fmt.Sprintf("%s -> %s (%d)", from, ret, total_weight )

	return ret

}

func ( ag *ArrayGraph ) dump() {

	for i := range ag.board {
		fmt.Print( itos(i), " : " );
		var link_cnt = 0
		for j := range ag.board[i] {
			if 0 < ag.board[i][j] {
				fmt.Print( itos(j), "(", ag.board[i][j] ,")", "\t" );
				link_cnt ++
			}
		}

		fmt.Println("LinkCnt(",link_cnt,")");
	}
	
}

func main(){

	var ag = createArrayGraph( 6 )

	//fmt.Println("test ", indexToString(25), indexToString(26), indexToString(30), indexToString(702));
	//fmt.Println("test ", stringToIndex("Z"), stringToIndex("AA"), stringToIndex("AE"), stringToIndex("AAA"));

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

	ag.dijkstra()

	fmt.Println("-----------------------------------------")
	fmt.Println("shortest : ", ag.shortist( "A", "F" ) )
	fmt.Println("shortest : ", ag.shortist( "F", "A" ) )
	fmt.Println("shortest : ", ag.shortist( "F", "B" ) )	/// 신기한건 F에서 D로 가는 경로가 없는데 탐색을 수행 한다.. 이건 로직에 문제가 있는듯
	fmt.Println("shortest : ", ag.shortist( "A", "B" ) )
	fmt.Println("-----------------------------------------")
	ag.dump()

}