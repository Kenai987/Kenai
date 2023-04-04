package main

import "fmt"

// Phone 接口
type Phone interface {
	call()
	down()
}

// NokiaPhone 实现类
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) down() {
	//TODO implement me
	panic("implement me")
}

// 实现接口的方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type MiPhone struct {
}

func (miPhone MiPhone) down() {
	fmt.Println("我先挂了")
}

// IPhone 实现类
type IPhone struct {
}

func (iPhone IPhone) down() {
	//TODO implement me
	panic("implement me")
}

// 实现接口的方法
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

//interface可以被任意的对象实现
//一个对象可以实现任意多个interface
//任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface
func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	miPhone := new(MiPhone)
	miPhone.down()
}
