package test

import (
	"fmt"
	"testing"
)

// 基类 Animal
type Animal struct {
	Name string
}

// Animal 的方法
func (a *Animal) Speak() {
	a.CallA()
}

// Animal 的方法
func (a *Animal) CallA() {
	fmt.Println(" call  animl ")
}

// 子类 Dog
type Dog struct {
	Animal // 嵌入 Animal，实现组合
}

func (d *Dog) CallA() {
	fmt.Println(" call dog callA")
}

func TestCall(t *testing.T) {
	// 创建 Dog 实例
	dog := &Dog{
		Animal: Animal{Name: "Buddy"},
	}
	dog.Speak()
}
