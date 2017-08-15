//a few basic examples
package main

import "fmt"
import "errors"
	
func add(a int,b int) (int, int) {
	return a+b,a-b
}

func addPlus(a, b, c int) int {
	return a+b+c
}

func multi(args ...int) int {
	// total := 0
	var total int = 0
	for _,v := range args {
		total = total + v 

	}
	return total
}

func closure() func() int {
	i := 0;
	return func() int {
		i += 1
		return i
	}
}

func pointer(i *int) {
	*i = 0
}
	
type comb interface {
	cov() string
}

type myError struct {
	arg int
	reason string
}
func (m *myError) Error() string {
	return fmt.Sprintf("%d - %s ",m.arg,m.reason)
}

func f1(i int) (int, error) {
	if i == 42 {
		return -1,errors.New("cant deal with 42")
	}
	return i,nil
}

func f2(arg int) (int,error) {
	if arg == 42 {
		return -1,&myError{arg, "cant deal with"}
	}
	return arg,nil
}

type mystr struct{
	num int
	name string
	last string
}



func (m *mystr) cov() string {
	return m.name + " " + m.last
}

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=",1+1)
	fmt.Println(true)

	/*
	var s string = "just"
	fmt.Printf("%s\n",s)

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	i = -1
	if i > 0 {
		fmt.Println("i 大于 0")
	} else if i > -1 {
		fmt.Println("zz")
	} else {
		i = 1
		fmt.Println("dd")
	}
	*/
	

	whatAmI := func (i interface {}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("is boolen")
		case int:
			fmt.Println("is int")
		default:
			fmt.Printf("is %T\n",t)
		}

	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("dd")

	var mys mystr = mystr{num: 1,name: "hello"}
	whatAmI(mys)

	var a [5]int
	whatAmI(a)
	a[4] = 23
	fmt.Println(a)
	b := []int{1,2,3,4}
	whatAmI(b)
	// var td [][]int

	/*
	var c = make([]string,3)
	whatAmI(c)
	fmt.Println(c)
	fmt.Println(len(c))

	c[0] = "ss"
	fmt.Println(c)
	fmt.Println(len(c))

	d := append(c,"cc","dd","dd")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(d[2:5])
	*/

	/*
	var mapdata = make(map[string]int)
	mapdata["att1"] = 10
	mapdata["att2"] = 20
	delete(mapdata,"att2")
	fmt.Println(mapdata)
	v,prs := mapdata["att2"]
	fmt.Println(v,prs)
	n := map[string]int{"num":21,"age":44}
	fmt.Println(n)

	for v := range n {
		fmt.Println(n[v])
	}
	*/

	/*
	for i,c := range "9Aa" {
		fmt.Println(i,c)
	}

	ddd, err := add(1,2)
	fmt.Println(ddd,err)

	fmt.Println(addPlus(1,2,3))

	fmt.Println(multi(4,5,6))

	clo := closure()
	fmt.Println(clo())
	fmt.Println(clo())

	pr := 4
	pointer(&pr)
	fmt.Println(pr)
	fmt.Println(&pr)

	fmt.Println(&mys)
	*/
	mdd := mystr{name:"jack",last:"uncle"}
	fmt.Println(mdd.cov())

	sli := []int{2,42}

	for _,v := range sli {
		if _,err := f2(v); err != nil {
			fmt.Println(err)
		}
	}





}