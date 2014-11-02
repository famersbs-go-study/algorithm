package main

import(
	"fmt"
	"math"
)

const (
	BOARD_LEN int = 3
	BOARD_SIZE int = BOARD_LEN * BOARD_LEN
	COST_MANHATAN int = 1
	COST_MISPLACEDTILES int = 0
)

/** puzzle State **/
type puzzleState struct {
	board []int
	parent *puzzleState
	state string
	cost_h int
	cost_g int
	cost_f int
	space_idx int
} 

////////////////////////////////////////////////////////////
// 생성
func ( p *puzzleState ) init( board []int, parent *puzzleState) {

	p.board = board
	if nil != parent {
		p.parent = parent
	}
	p.createState()
	p.calcCost( COST_MISPLACEDTILES )
	//p.calcCost( COST_MANHATAN )
}
func ( p *puzzleState ) createState() {
	p.state = fmt.Sprintf( "%d %d %d\n%d %d %d\n%d %d %d\n",
							p.board[0], p.board[1], p.board[2],
							p.board[3], p.board[4], p.board[5],
							p.board[6], p.board[7], p.board[8] );
}

func ( p *puzzleState ) toString() string {
	return fmt.Sprintf( "========\ncost : %d\n%s", p.cost_h, p.state );
}

func createPuzzleState( board []int, parent *puzzleState ) *puzzleState {
	var ret = puzzleState{}

	ret.init( board, parent )

	return &ret

}

////////////////////////////////////////////////////////////
// 코스트 계산
func (p *puzzleState ) calcCost( t int ){

	p.cost_h = 0
	if COST_MISPLACEDTILES == t {

        for i := 0; i < BOARD_SIZE; i++ {
            
            var value = p.board[i] - 1

            // Space tile's value is -1
            if value == -2 {
                value = BOARD_SIZE - 1;
                p.space_idx = i;
            }

            if value != i {
                p.cost_h ++;
            }
        }

	}else{

		for i := 0 ; i < BOARD_SIZE ; i ++ {

			var value = p.board[i] - 1
			if value == -2 {
				value = 8
			}

			var c_x = i % BOARD_LEN;
			var c_y = i / BOARD_LEN;
			var i_x = value % BOARD_LEN;
			var i_y = value / BOARD_LEN;

			var calc_result = (int) ( math.Abs( (float64) (c_y - i_y) ) + math.Abs( (float64) (c_x - i_x) ) )
			p.cost_h += calc_result

			//fmt.Println( "Cost ", i, value ,c_x, i_x, c_y, i_y, calc_result  );

		}

		//fmt.Println( "Cost ", p.cost_h  );

	}

	if p.parent != nil {
		p.cost_g = p.parent.cost_g + 1
	}else {
		p.cost_g = 0
	}

	p.cost_f = p.cost_h + p.cost_g

}

////////////////////////////////////////////////////////////
// 다음 단계 가져오기
func ( p *puzzleState ) getNextStates() []*puzzleState {
	var nextStates = make( []*puzzleState, 4 );	// 일단 4개 다 만들고 N 개 버린다 ~!

	var translateIdxToXY = func ( idx int ) ( x int, y int ){
		if 0 == idx {
			return 0,0
		}
		return idx % BOARD_LEN, idx / BOARD_LEN
	}

	var createNextState = func ( curr *puzzleState, x int, y int ) *puzzleState {
		var nextidx = x + y * BOARD_LEN

		var nextboard = make( []int, BOARD_SIZE )
		copy( nextboard, curr.board )

		// swap
		var ttt = nextboard[ curr.space_idx ]
		nextboard[ curr.space_idx ] = nextboard[ nextidx ]
		nextboard[ nextidx ] = ttt

		// create board
		return createPuzzleState( nextboard, curr )

	}

	var x, y = translateIdxToXY( p.space_idx )
	var cur_idx = 0

	//현재 상태 에서 4방 확인
	if y > 0 {
		// up
		nextStates[ cur_idx ] = createNextState( p, x, y - 1 )
		cur_idx ++
	}
	if y < BOARD_LEN - 1 {
		// down
		nextStates[ cur_idx ] = createNextState( p, x, y + 1 )
		cur_idx ++
	}
	if x > 0 {
		// left
		nextStates[ cur_idx ] = createNextState( p, x - 1, y )
		cur_idx ++
	}
	if x < BOARD_LEN - 1 {
		// right
		nextStates[ cur_idx ] = createNextState( p, x + 1, y )
		cur_idx ++
	}

	// 정리 후 리턴 
	return nextStates[ : cur_idx ]

}

// 경로 추적
func ( p *puzzleState ) printresult() int {
	var step_cnt = 1

	if p.parent != nil {
		step_cnt += p.parent.printresult()
	}

	fmt.Println("----------------------------")
	fmt.Println("Step ", ( step_cnt ) ) 
	fmt.Println( p.toString() )

	return step_cnt

}

/******************************* min pirote queue **/
type MinpriorityQueue struct {
	q []*puzzleState
	nextidx int
}

func crateMinpriorityQueue() MinpriorityQueue {
	var ret = MinpriorityQueue{}
	ret.q = make( []*puzzleState, 10000 ) // 일단 큐사이즈를 1만으로 고정 하자 ~!
	ret.nextidx = 0;
	return ret
}

func (q * MinpriorityQueue ) putq( s *puzzleState ) {
	q.q[ q.nextidx ] = s;
	q.nextidx ++
}

func (q * MinpriorityQueue ) popq( ) *puzzleState {

	if q.nextidx <= 0 {
		return nil
	}

	var min_idx = 0
	var min_f = q.q[ 0 ].cost_f

	

	// 제일 작은 녀석을 찾는다
	for i := 1 ; i < q.nextidx ; i ++ {
		if min_f > q.q[ i ].cost_f {
			min_idx = i
			min_f = q.q[ i ].cost_f
		}
	}

	// 마지막 과 바꾸고 리턴~~!
	var ret = q.q[ min_idx ]
	q.q[ min_idx ] = q.q[ q.nextidx - 1 ]
	q.nextidx --

	return ret

}
func (q * MinpriorityQueue) isEmpty() bool {
	return q.nextidx == 0
}

func (q * MinpriorityQueue) getStateAndIdx( state_code string ) ( *puzzleState, int ) {

	for i:= 0 ; i < q.nextidx ; i ++ {
		if q.q[ i ].state == state_code {
			return q.q[i], i
		}
	}

	return nil, -1

}

func( q* MinpriorityQueue) changeIdx( idx int, cc *puzzleState ) {

	q.q[ idx ] = cc

}

func startSolve( start []int ) *puzzleState{

	var openStates = make( map[string]bool )
	var closeStateQ = make( map[string]*puzzleState )
	var state_count = -1

	// Init queue & openstate
	var openStateQ = crateMinpriorityQueue()
	var start_state = createPuzzleState( start, nil )
	openStateQ.putq( start_state )
	openStates[ start_state.state ] = true

	var curr_state *puzzleState = nil

	// Main roop
	for !openStateQ.isEmpty() {

		curr_state = openStateQ.popq()
		delete( openStates, curr_state.state )

		state_count ++

		// check. is it complete
		if curr_state.cost_h == 0 {
			fmt.Println("Complete !!! " )
			return curr_state;
		}

		// get next States
		var nexts = curr_state.getNextStates()
		
		var openState *puzzleState = nil
		var closeState *puzzleState = nil
		var nextState *puzzleState = nil

		if len( nexts ) == 0 {
			continue;
		}

		for i := 0 ; i < len( nexts ) ; i ++ {
			openState = nil
			closeState = nil
			nextState = nexts[i]

			if openStates[ nextState.state ] {
				
				var open_idx = -1
				openState, open_idx = openStateQ.getStateAndIdx( nextState.state )

				if openState.cost_f > nextState.cost_f {
					// 교체 
					openStateQ.changeIdx( open_idx, nextState )
				}

			}else {

				closeState = closeStateQ[ nextState.state ]

				if nil != closeState {

                    if closeState.cost_f > nextState.cost_f {
                        closeStateQ[ nextState.state ] = nextState;
                    }

				}

			}

			if nil == openState && nil == closeState {
				openStateQ.putq( nextState )
				openStates[ nextState.state ] = true
			}

		}

		closeStateQ[curr_state.state] = curr_state;

	}

	if curr_state != nil {
		if curr_state.cost_h == 0 {
			return curr_state
		}
	}

	return nil

}

func main(){
	var start = []int{1,4,3,2,-1,6,7,8,5};
	//var goal = []int{1,2,3,4,5,6,7,8,-1};
	/*
	var state = createPuzzleState( start, nil )
	var state_g = createPuzzleState( goal, nil )

	var q = crateMinpriorityQueue();
	q.putq( state )
	q.putq( state_g )

	fmt.Println("state_g ", state_g.cost_h )

	fmt.Println("q test " )
	fmt.Println( q.getStateAndIdx( state.state ) )
	fmt.Println("---------------------------")

	fmt.Println( q.popq().toString() )
	fmt.Println( q.popq().toString() )

	fmt.Println( state.toString() )

	var nextg = state_g.getNextStates()

	for i := 0 ; i < len( nextg ) ; i ++ {
		fmt.Println("hh ", nextg[i].toString() )	
	}


	*/
	

	fmt.Println("---------------------------")
	var final = startSolve( start )

	if nil == final {
		fmt.Println("Find fail...")
	}else{
		var last_step_cnt = final.printresult();
		fmt.Println("Last Step cnt : ", last_step_cnt )
	}

}