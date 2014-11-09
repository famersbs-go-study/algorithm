package main

import(
	"fmt"
	"math"
)


/**
 * 
 */
type TreeNode struct{

	left 		*TreeNode
	right 		*TreeNode	
	data 		int

}

func createTree( root int ) *TreeNode {
	return &TreeNode{ nil, nil, root }
}

func ( t *TreeNode ) push( value int ) {

	if t.data > value {
		if t.left != nil {
			t.left.push( value )
		}else {
			t.left = &TreeNode{ nil, nil, value }
		}
	}else {
		if t.right != nil {
			t.right.push( value )
		}else {
			t.right = &TreeNode{ nil, nil, value }
		}
	}

}

func ( t *TreeNode ) travel(){
	fmt.Print( t.data," " );
	if t.left != nil {
		t.left.travel();
	}
	if t.right != nil {
		t.right.travel();
	}
}


func (t *TreeNode ) getWidth_recur( c_width int ) ( int, int ){

	var l_l, l_r, r_l, r_r int;

	if t.left != nil {
		l_l, l_r =  t.left.getWidth_recur( c_width - 1 );
	}else {
		l_l = c_width;
		l_r = c_width;
	}
	if t.right != nil {
		r_l, r_r = t.right.getWidth_recur(  c_width + 1 );
	}else{
		r_l = c_width;
		l_r = c_width;
	}

	// 일반 적인 상식으로는 왼쪽 자식의 가장 왼쪽에 있는 값이 제일 작지만, 만약 오른쪽 자식들의 최고 왼쪽 자식이 더 왼쪽으로 치우쳐졌다면 (더 작다면) l_l 을 r_l로 교체 한다. 
	// 오른쪽도 마찬가지
	if l_l > r_l {
		l_l = r_l;
	}

	if r_r < l_r {
		r_r = l_r
	}

	return l_l, r_r;
}

func (t *TreeNode ) getWidth() int{

	var left, right = t.getWidth_recur( 0 );

	return ( ( (int)(math.Abs( (float64)(left) )) ) + right + 1 );

}

func main(){

	fmt.Println("Start Tree Width check");

	var t = createTree( 10 );

	t.push( 100 )
	t.push( 9 )
	t.push( 101 )
	t.push( 8 )
	t.push( 102 )
	t.push( 7 )
	t.push( 6 )

	fmt.Println("Width", t.getWidth() );

	var t2 = createTree( 10 );
	t2.push(9)
	t2.push(8)
	t2.push(7)
	t2.push(6)
	t2.push(5)
	t2.push(4)
	t2.push(3)
	t2.push(2)
	t2.push(1)
	fmt.Println("Width", t2.getWidth() );


}