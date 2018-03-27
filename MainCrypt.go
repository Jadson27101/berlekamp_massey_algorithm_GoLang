package main

import "fmt"
//var b [] int
//var t[] int
//var c[] int
var arr[] int
var s = [13] int{1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1}
var N, L, m int
func initConst(){
	N = 0
	L = 0
	m = -1
}
func createB(length int) []int{
	b:=make([]int, length)
	for i:=0; i<length; i++ {
		b[i] = 0
	}
	b[0] = 1
	return b
}
func createT(length int) []int{
	t:=make([]int, length)
	for i:=0; i<length; i++ {
		t[i] = 0
	}
	return t
}
func createC(length int) []int{
	c:=make([]int, length)
	for i:=0; i<length; i++ {
		c[i] = 0
	}
	c[0] = 1
	return c
}

func run(length int) []int  {
	var d int
	var b = createB(length);
	var c = createC(length);
	var t = createT(length);
	initConst();
	for N < len(s)  {
		d = 0
		for i:=0;i<=L; i++ {
			d+= s[N-i] * c[i]
		}
		d = d%2

		if d!=0 {
			copy(t,c)
			for i:=0; i<len(s)+m-1-N; i++ {
				c[N-m+i] = c[N-m+i] ^ b[i]
			}
			if L<=(N/2) {
				L = N+1-L
				m = N
				copy(b,t)
			}
		}
		N++
	}
	return c
}
func main()  {
	fmt.Println(run(13))
}