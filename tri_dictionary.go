package main

import(
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

////////////////////////////////////////////////////////////
// Tri source
const (
	TRI_TABLE_SIZE = 27
)

type DicTriNode struct {

	
	descript []string
	table [TRI_TABLE_SIZE]*DicTriNode

}

func (n *DicTriNode) addDescript( desc string ) {
	n.descript = append( n.descript, desc )
}
func (n *DicTriNode) isNotWord() bool {
	return ( len( n.descript ) == 0 )
}

func (n *DicTriNode ) getDescript () string {
	
	var ret string = ""

	if n.isNotWord() {
		return "";
	}

	for i := 0 ; i < len( n.descript ) ; i++ {

		ret = fmt.Sprintf( "%s\t%d : %s\n", ret, ( i + 1 ), n.descript[i] );

	}

	return ret;
}

func (n *DicTriNode) dump( current_word string ){
	if !n.isNotWord() {
		fmt.Println( current_word, "\n", n.getDescript() );
	}

	for i := 0 ; i < TRI_TABLE_SIZE ; i++ {
		if nil != n.table[i] {
			n.table[i].dump( current_word + string( 'a' + i ) )
		}
	}
}

func ( node *DicTriNode ) statics( t int, n int, b int ) ( int, int, int ){

	var r_t, r_n, r_b int
	var c_b int = 0			// 현재 노드의 비어 있는 링크의 수

	r_t = t + TRI_TABLE_SIZE
	r_n = n
	r_b = b
	if node.isNotWord() {
		r_n = r_n + 1
	}

	for i := 0 ; i < TRI_TABLE_SIZE ; i++ {
		if nil == node.table[i] {
			c_b = c_b + 1;
		} else {
			r_t, r_n, r_b = node.table[i].statics( r_t, r_n, r_b )
		}
	}

	r_b = ( r_b + c_b ) / 2

	return r_t, r_n, r_b

}

type Dictionary struct {

	root DicTriNode

}

func createDictionary() *Dictionary {
	return &Dictionary{};
}

func charactorToInt( charactor byte ) byte {

	if charactor >= 'a' && charactor <= 'z' {
		return charactor - 'a'
	}else if charactor >= 'A' && charactor <= 'Z' {
		return charactor - 'A'
	}else{
		// undefine charictor is 26
		return 26
	}

}

func ( d *Dictionary ) add( word string, descript string ) bool {

	var length = len( word )
	var cur_node *DicTriNode = &d.root

	// find node
	for i := 0 ; i < length ; i++ {

		var cur_idx = charactorToInt( word[i] )

		if nil == cur_node.table[ cur_idx ] {
			// create a new node 
			cur_node.table[ cur_idx ] = &DicTriNode{};
		}

		cur_node = cur_node.table[ cur_idx ];
	}

	cur_node.addDescript( descript )

	return true;
}

func ( d *Dictionary ) find( word string ) ( int, string ) {

	var length = len( word )
	var cur_node *DicTriNode = &d.root
	var last_validate_node_depth int = 0
	var last_validate_node *DicTriNode = nil

	// find node
	for i := 0 ; i < length ; i++ {

		var cur_idx = charactorToInt( word[i] )


		if nil == cur_node.table[ cur_idx ] {
			break;
		}
		cur_node = cur_node.table[ cur_idx ];

		if nil != cur_node && !cur_node.isNotWord() {
			last_validate_node_depth = i
			last_validate_node = cur_node
		}

	}

	if nil != last_validate_node {

		var percentage int =  ( (last_validate_node_depth) * 100 ) / ( ( length - 1) )
		return percentage, string( word[:last_validate_node_depth + 1 ] ) + "\n" + cur_node.getDescript();

	}else{
		return 0, "Not defined word : " + word 		
	}

}

func ( d *Dictionary ) find_checkTime( word string ) (int, string ){

	var start_timestamp int64 = makeTimestamp();
	var p, r = d.find( word )
	var end_timestamp int64 = makeTimestamp();

	fmt.Println( "find : ", ( end_timestamp - start_timestamp ),"us" );

	return p, r 

}

func( d *Dictionary ) dump() {

	d.root.dump("");

}

func( d *Dictionary ) statics() ( t int, n int, b int ){

	var rt, rn, rb = d.root.statics( 0,0,-1 )

	fmt.Println( "total link cnt : ", rt )
	fmt.Println( "total none node link cnt : ", rn )
	fmt.Println( "total empty node link cnt : ", rb )

	return rt,rn,rb

}
// Tri source
////////////////////////////////////////////////////////////



////////////////////////////////////////////////////////////
// force Dictionary

var words []string
var descripts []string

func forceFind( word string ) string{
	var start_timestamp int64 = makeTimestamp();

	var length = len(words)
	var finded int = -1;
	for i := 0 ; i < length ; i ++ {
		if words[i] == word {
			finded = i
			break;
		}
	}

	var end_timestamp int64 = makeTimestamp();

	fmt.Println( "find : " , ( end_timestamp - start_timestamp ),"us" );

	if finded != -1 {
		return descripts[finded];
	}else{
		return " Not Find..."
	}

}

// force Dictionary
////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////
// External Dictionary load

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )
  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}

func makeTimestamp() int64 {
    return time.Now().UnixNano() // / int64(time.Millisecond)
}

func loadDictionary( d *Dictionary ){

	fmt.Println("Start Load Dictionary ", int64(time.Millisecond));

	f, err := os.Open("./dictionary/EK51873.txt")

	if err != nil {
	    fmt.Println("error opening file= ",err)
	    os.Exit(1)
	}

	r := bufio.NewReader(f)

	s, e := Readln(r)

	for e == nil {
	    //fmt.Println(s)
	    s,e = Readln(r)
	    var split = strings.Index(s, "///" )

	    if split >= len( s ) -1 {
	    	fmt.Println("Error ", s );
	    	continue;
	    }

	    var word string = s[:split]
	    var descript string = s[split+3:]
	    //lines = append( lines, s );
	    words = append( words, strings.TrimSpace( word ) )
	    descripts = append( descripts, strings.TrimSpace( descript ) )
	}

	fmt.Println("Lines : ", len( words ) );
//	fmt.Println("Lines : ["+ words[0]+ "] [" + descripts[0] + "]" );
	var start_timestamp int64 = makeTimestamp()

	for i := 0 ; i < len( words ) ; i ++ {
		d.add( words[i], descripts[i] )
	}

	var end_timestamp int64 = makeTimestamp();

	fmt.Println("Dictionary load time ", ( end_timestamp - start_timestamp ), "us" );

}
// External Dictionary load
////////////////////////////////////////////////////////////

func checkfineTimeout( d *Dictionary, word string ){

	fmt.Println( "---------- tri " )
	fmt.Println( d.find_checkTime( word ) )
	fmt.Println( "---------- array " )
	fmt.Println( forceFind( word ) )


}

func main(){

	var d *Dictionary = createDictionary();

	d.add( "bee", "벌")
	d.add( "farmer", "농부")
	d.add( "sort", "정렬")
	d.add( "sort", "소금?" )
	d.add( "in", "~안에" )


	fmt.Println( d.find( "sort" ) );
	fmt.Println( d.find( "sortist" ) );
	fmt.Println( d.find( "bebrato" ) );
	fmt.Println( d.find( "into" ) );
	d.dump();


	var d_full *Dictionary = createDictionary();
	loadDictionary(d_full)

	checkfineTimeout( d_full, "disket" )
	checkfineTimeout( d_full, "apple" )
	checkfineTimeout( d_full, "history" )
	checkfineTimeout( d_full, "zoo" )

	fmt.Println("---Statics ---------")
	d_full.statics()


}