package types

import "fmt"

type i int32
type s string
type vector []int32
type foo func(v vector) string

func (v vector) push_back(i int32) {

	v = append(v, i)
}

func (v vector) size() int32 {
	return int32(len(v))
}

func (v vector) del() {

	return
}

func (v vector) pop() int32 {

	return v[len(v)-1]
}

func (v vector) String() string {
	return string(v)
}

func Slice() {
	var sl []int
	//fmt.Println(sl)

	sl = append(sl)
	sl = make([]int, 0, 0)

	fmt.Println(sl)

}
func sCheck(mys s, m string) {

	/*	muf := func(v vector) string {
		return ""
	}*/

	/*sl := []int
	fmt.Println(sl)
	var df foo = muf

	var v vector = make([]byte,2)
	 v.push_back(2)
	 d := v.pop() //get last
	 v.push_back(d)

	 fmt.Println( v.size())

	 v.del(idx) //

	func (v vector) string{
		return v.String()
	}(v)

	v.String()

	v2 := make([]byte,3)

	 ddd := []byte(v)
	*/

	//fmt.Println(vector(v2))

	if mys == s(m) {

	}
}
