package quadtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type MapUnit struct {
	Uid int64
	Bounds
}

func Test_add_test(t *testing.T) {
	qt := Quadtree[MapUnit]{
		Bounds:     Bounds{0, 0, 307200, 307200},
		MaxObjects: 128,
		MaxLevels:  10,
		Level:      1,
		Objects:    make([]Entity[MapUnit], 128),
		Nodes:      nil,
		Total:      0,
	}
	startTime := time.Now().UnixMilli()
	for i := 0; i < 300000; i++ {
		x := rand.Int31n(307200)
		y := rand.Int31n(307200)
		//fmt.Printf("   pppppp x = %d y =%d \n", x, y)
		mapUnit := &MapUnit{
			Uid:    0,
			Bounds: Bounds{float64(x), float64(y), 256, 256},
		}
		entity := Entity[MapUnit]{
			obj:    mapUnit,
			Bounds: mapUnit.Bounds,
		}
		qt.Insert(entity)
	}
	endTime := time.Now().UnixMilli()
	fmt.Printf("--------  cost time =%d \n", endTime-startTime)

	sStart := time.Now().UnixMilli()
	count := 0
	for i := 0; i < 10000; i++ {
		retrieve := qt.Retrieve(Bounds{
			X:      600,
			Y:      700,
			Width:  2000,
			Height: 2000,
		})
		count = len(retrieve)
	}
	sEnd := time.Now().UnixMilli()

	fmt.Printf("-----------cost  = %d   count = %d\n", sEnd-sStart, count)
}

type Person struct {
	Name string
	Age  int
}

func Test_Reference(t *testing.T) {
	// 创建一个Person结构体实例
	p := Person{Name: "张三", Age: 18}
	// 创建一个指向Person结构体实例的指针
	p1 := &p

	// 输出Person结构体实例的信息
	fmt.Println(p.Name, p.Age)

	// 通过指针修改Person结构体实例的信息
	p1.Name = "李四"
	p1.Age = 20

	fmt.Println(p.Name, p.Age)

	p2 := p
	p2.Name = "王五"
	p2.Age = 22
	acceptValue(*p1)

	// 再次输出Person结构体实例的信息
	fmt.Println(p.Name, p.Age)

	acceptPoint(p1)

	fmt.Println(p.Name, p.Age)
}

func acceptValue(person Person) {
	person.Name = "hello"
}

func acceptPoint(person *Person) {
	person.Name = "hello1"
}
