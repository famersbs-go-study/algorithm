package main

import(
	"fmt"
)

type ExprType int

const (
	SYMBOL 	ExprType = 1
	NUMBER 	ExprType = 2
	BLANK  	ExprType = 3
	UNKNOWN	ExprType = 4
	ROOT 	ExprType = 5
)

/**
 * 
 */
type ExpTreeNode struct{

	left 		*ExpTreeNode
	right 		*ExpTreeNode	
	exprtype 	ExprType
	data 		string

}

/**
 * 
 */
type ExperTree struct { 
	entryp *ExpTreeNode
}

func (t *ExperTree) init( midexpr string ) bool {

	//t.entryp = ExpTreeNode { nil, nil, ROOT, "" }

	// make Tree
	var token_idx = 0

	var checkChar = func ( c byte ) ExprType {
		switch c {
			case '0','1','2','3','4','5','6','7','8','9': return NUMBER
			case '+','-','*','/': return SYMBOL
			case ' ': return BLANK
			default: return UNKNOWN
		}
	}

	var symbolToPirority = func( s string ) int {

		switch s{
		case "*", "/":
			return 5
		case "+","-":
			return 10
		}

		return 100
	}

	var nextToken = func() ( string, ExprType ) {
		var token string = ""	
		var mark_idx = -1

		for token_idx < len( midexpr ) {
			var cur_type ExprType = checkChar( midexpr[ token_idx ] )
			// get token
			switch cur_type {
			case NUMBER:
				if mark_idx == -1 {
					// not marked means - this token is start of number
					mark_idx = token_idx
				}
				token = token + string( midexpr[ token_idx ] );
			case SYMBOL, BLANK :
				if mark_idx != -1 {
					return token, NUMBER
				}else{
					token_idx += 1;
					return string( midexpr[ token_idx - 1 ] ), cur_type
				}
			case UNKNOWN:
				return "", UNKNOWN;
			}
			token_idx += 1;
		}

		if mark_idx == -1 {
			return "", UNKNOWN
		}else{
			return token, NUMBER
		}
	}

	for token_idx < len( midexpr ) {

		// get Token
		var token, current_type = nextToken();

		var tmp_node *ExpTreeNode = &ExpTreeNode{ nil, nil, current_type, token }

		if current_type == UNKNOWN {
			fmt.Println(" unknown token ")
			break
		}else if current_type == BLANK {
			continue
		}else if current_type == SYMBOL {

			var p_node *ExpTreeNode = nil
			var c_node *ExpTreeNode = t.entryp

			for true {

				if nil == c_node {
					fmt.Println("First Input is only number or symbol next symbol : ", token_idx )
					return false;
				}

				if c_node.exprtype == NUMBER {
					
					tmp_node.left = c_node

					if p_node == nil {
						// it's a root 
						t.entryp = tmp_node;
					}else{
						p_node.right = tmp_node
					}

					break;

				} else if c_node.exprtype == SYMBOL {

					var c_p = symbolToPirority( c_node.data )
					var t_p = symbolToPirority( tmp_node.data )

					if t_p < c_p {
						//tmp priority is hight then current
						p_node = c_node
						c_node = c_node.right

					}else {
						tmp_node.left = c_node

						if p_node == nil {
							t.entryp = tmp_node
						}else{
							p_node.right = tmp_node
						}
						break;
					}
				}
			}


		} else if current_type == NUMBER {

			if nil == t.entryp {
				t.entryp = tmp_node
			}else{

				//var p_node *ExpTreeNode = nil
				var c_node *ExpTreeNode = t.entryp

				for true {

					if c_node.exprtype == NUMBER {
						// It's an error
						fmt.Println("Number token after Number token : ", token_idx )
						return false

					}else if c_node.exprtype == SYMBOL {
						if c_node.right == nil {
							c_node.right = tmp_node
							break
						} else {
							//p_node = c_node
							c_node = c_node.right
						}
					}

				}
			}
		}
		//fmt.Println( token, " : " , current_type );
	}
	return true;
}

func preOrder ( node *ExpTreeNode ) string {
	if nil == node {
		return "";
	}
	return node.data + " " + preOrder( node.left ) + preOrder( node.right );
}

func (t *ExperTree) toPre( ) string {
	return preOrder( t.entryp );
}
func postOrder ( node *ExpTreeNode ) string {
	if nil == node {
		return "";
	}
	return postOrder( node.left ) + postOrder( node.right ) + " " + node.data ;
}

func (t *ExperTree) toPost( ) string {
	return postOrder( t.entryp );
}


func main(){

	fmt.Println("Expression Bin Tree")

	var ttt ExperTree

	fmt.Println("--- Normal")
	ttt.init( "10 + 9 * 10 + 2")
	fmt.Println("preOrder : ", ttt.toPre() );
	fmt.Println("postOrder : ", ttt.toPost() );

	fmt.Println("--- Error")
	ttt.init( "10 + 100 * 2 20")
	

}