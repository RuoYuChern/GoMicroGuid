package main

import (
	"fmt"
	"math"
)

var gname string = "Ruo"

// 声明常量变量
const (
	ST_OK        int = 0
	ST_FAILED    int = -1
	ST_NOT_SOUND int = 400
)

func say(world string) {
	// say 内的局部变量
	var age int = 10
	const prefix string = "Mr."
	fmt.Printf("1>> Hello %s %s age:%d\n", prefix, world, age)
	if true {
		// 更小作用域的变量
		var age int = 11
		fmt.Printf("2>> Hello %s %s age:%d\n", prefix, world, age)
	}
	fmt.Printf("3 >> Hello %s%s age:%d\n", prefix, world, age)
}

func pintInt() {
	fmt.Printf("int8: [%d, %d)\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("uint8: [%d, %d)\n", 0, math.MaxUint8)
	fmt.Printf("int16: [%d, %d)\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("uint16: [%d, %d)\n", 0, math.MaxUint16)
	fmt.Printf("int32: [%d, %d)\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("uint32: [%d, %d)\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int64: [%d, %d)\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("uint64: [%d, %d)\n", 0, uint64(math.MaxUint64))
}

func pintFloat() {
	var f32 float32 = 10.123456
	var f64 float64 = 1000.123456
	fmt.Printf("float: [%f, %f)\n", f32, f64)
	fmt.Printf("float: [%.4f, %.4f)\n", f32, f64)
	fmt.Printf("float: [%.f, %.f)\n", math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat64)
}

func printBool() {
	var isHigh bool = false
	var isLow bool = !isHigh
	var isEqual bool = (3 == 5)
	fmt.Printf("isHigh:%t, isLow:%t, isEqual:%t\n", isHigh, isLow, isEqual)
}

func printChar() {
	var a byte = 'A'
	var b rune = '学'
	var c byte = '\''
	fmt.Printf("a: '%c'\tb:\"%c\"\tc:%c\n", a, b, c)
	fmt.Printf("next line")
}

func printStr() {
	var s string = "Hello Go"
	var m string = `Hello Go,
	 func(){
		fmt.Printf("%s",s)
	 }`
	fmt.Printf("s:%s\n", s)
	fmt.Printf("m:%s\n", m)
}

func printComplex() {
	var c64 complex64 = complex(10, 20)
	var c128 complex128 = complex(100, 200)
	fmt.Printf("c64:%v, real: %f, imag:%f\n", c64, real(c64), imag(c64))
	fmt.Printf("c64:%v, real: %f, imag:%f\n", c128, real(c128), imag(c128))
}

func typeChange() {
	var i8 int8 = 10
	var i32 int32 = int32(i8)
	var f32 float32 = float32(i32)
	var str string = string(i32)
	var f64 float64 = 10.345
	var i64 int64 = int64(f64)
	var a int8 = 65
	var as string = string(rune(a))
	fmt.Printf("i8:%d, i32:%d, f32:%f, str:%s, f64:%f, i64:%d\n", i8, i32, f32, str, f64, i64)
	fmt.Printf("a: %s\n", as)
}

func printSlice() {
	//声明但未初始化
	var ar1 [4]int
	fmt.Printf("ar1 len: %d, cap:%d\n", len(ar1), cap(ar1))
	for i := 0; i < 4; i++ {
		ar1[i] = (i + 1)
	}
	// 声明并初始化
	var ar2 = [4]int{1, 2, 3, 4}
	// 声明并初始化
	var ar3 = [...]int{1, 2, 3, 4}
	for i := 0; i < cap(ar1); i++ {
		fmt.Printf("ar1[%d] = %d, ar2[%d] = %d, ar3[%d] = %d\n", i, ar1[i], i,
			ar2[i], i, ar3[i])
	}

	// 通过截取数组来获取
	var sl1 []int = ar1[:]
	fmt.Printf("sl1 len: %d, cap:%d\n", len(sl1), cap(sl1))
	for _, v := range sl1 {
		fmt.Printf("v:%d\t", v)
	}
	fmt.Printf("\n")
	// 声明未初始化
	var sl2 []int
	fmt.Printf("sl2: value:%p len: %d, cap:%d\n", sl2, len(sl2), cap(sl2))
	// 赋值
	for i := 0; i < 4; i++ {
		sl2 = append(sl2, i+1)
	}
	for _, v := range sl2 {
		fmt.Printf("v:%d\t", v)
	}
	fmt.Printf("\n")
	fmt.Printf("sl2: value:%p len: %d, cap:%d\n", sl2, len(sl2), cap(sl2))
	for i := 0; i < len(sl2); i++ {
		fmt.Printf("%d\t", sl2[i])
	}
	fmt.Printf("\n")
	// 采用make 函数初始化
	var sl3 []int = make([]int, 3)
	fmt.Printf("sl3: value:%p len: %d, cap:%d\n", sl3, len(sl3), cap(sl3))
	// 采用copy 函数拷贝
	l := copy(sl3, sl2)
	fmt.Printf("copy len:%d\n", l)
	for i := 0; i < len(sl3); i++ {
		fmt.Printf("%d\t", sl3[i])
	}
	fmt.Printf("\n")
}

func pintMap() {
	// 声明初始化map, 未初始化是nil 指针，不能操作
	var m1 map[string]int = make(map[string]int)
	fmt.Printf("m1 len:%d", len(m1))
	// 添加元素
	m1["A"] = 10
	fmt.Printf("m1 len:%d\n", len(m1))
	// 访问元素: 返回(value, bool)
	v, ok := m1["B"]
	if ok {
		fmt.Printf("value:%d\n", v)
	}

	m1["B"] = 11
	// 遍历
	for k, v := range m1 {
		fmt.Printf("m1 key:%s, value:%d\n", k, v)
	}

	// 声明并初始化 map
	var m2 = map[string]int{
		"A": 10,
		"B": 11,
	}

	// 删除元素
	delete(m2, "B")
	for k, v := range m2 {
		fmt.Printf("m2 key:%s, value:%d\n", k, v)
	}
}

type Person struct {
	name string
	age  int
	sex  string
}

func printPerson() {
	var p1 Person
	fmt.Printf("p1: %v\n", p1)
	p1.age = 10
	p1.name = "Ry"
	p1.sex = "M"
	fmt.Printf("p1: %v\n", p1)
	var p2 = Person{
		name: "RR",
		age:  10,
		sex:  "M",
	}
	fmt.Printf("p2: %v\n", p2)
}

func printGen() {
	f1 := 10.45
	var p2 = Person{
		name: "RR",
		age:  10,
		sex:  "M",
	}
	fmt.Printf("f1: %f, %v, %+v, %#v, %T\n", f1, f1, f1, f1, f1)
	fmt.Printf("p21: %v, %+v, %#v, %T\n", p2, p2, p2, p2)
}

func main() {
	var age int
	var money int = 11
	age = 10
	fmt.Printf("gname=%s,age=%d, money=%d\n", gname, age, money)
	say("Go")
	fmt.Printf("UInt64:[%v,%v]", 0, uint64(math.MaxUint64))
	pintInt()
	pintFloat()
	printBool()
	printChar()
	printStr()
	printComplex()
	typeChange()
	printSlice()
	pintMap()
	printPerson()
	printGen()
}
