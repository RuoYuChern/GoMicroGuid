package main

import "fmt"

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

func main() {
	var age int
	var money int = 11
	age = 10
	fmt.Printf("gname=%s,age=%d, money=%d\n", gname, age, money)
	say("Go")
}
