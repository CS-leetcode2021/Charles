package main

import (
	"fmt"
	"sort"
)

/**
 *  @ClassName:sort_package
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/5/25 下午7:15
 *  {@Link}:https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/
 */
// go语言的排序思路：C和C++有些差别。C默认是对数组进行排序，C++是对一个序列进行排序，GO则更加宽泛一些，待排序的可以是任何对象，多数情况是一个slice或者是包含slice的一个对象
// 排序接口主要有三个：待排序的元素个数n；第i个和第j个元素的比较函数cmp；第i个和第j个的元素交换swap；

// 基本排序类型：go本身提供了sort.Ints() 、 sort.Float64s() 和 sort.Strings() 函数， 默认都是从小到大排序。
// 1、升序排序

func main01() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}

	sort.Ints(intList)
	sort.Float64s(float8List)
	// 按照字典排序
	sort.Strings(stringList)

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

// 2、降序排序
// 默认的是升序排序，所以可以使用sort.Sort(obj)实现，就是需要对Type类型绑定三个方法，Len()求长度，Less(i,j)比较大小的法则，Swap(i,j)交换法则

func main02() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	float8List := []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "d", "f", "i", "z", "x", "w", "y"}
	//fmt.Printf("%T",intList)
	// sort.IntSlice(intList) 这是一个类型转换
	//a := sort.IntSlice(intList)
	//fmt.Printf("%T",a)
	//fmt.Println()
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

	fmt.Printf("%v\n%v\n%v\n", intList, float8List, stringList)
}

// 3、深入理解排序
// sort包中有一个sort.Interface接口，该接口有三个方法 Len() 、 Less(i,j) 和 Swap(i,j) 。
// 通用排序函数 sort.Sort 可以排序任何实现了 sort.Inferface 接口的对象(变量)。
// 下面使用自定义的Reverse结构体，而不是sort.Reverse函数，来实现逆向排序

type Reverse struct {
	sort.Interface
}

func (r Reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

/*
 *  @Description:   sort.SearchInts
 *  @Param:
 *  @Return:
 */
func main03() {
	ints := []int{5, 2, 6, 3, 1, 4}

	sort.Ints(ints)
	fmt.Println("after sort by ints:\t", ints)

	doubles := []float64{2.3, 3.2, 6.7, 10.9, 5.4, 1.8}

	sort.Float64s(doubles)
	fmt.Println("after sort by Float64s:\t", doubles) // [1.8 2.3 3.2 5.4 6.7 10.9]

	strings := []string{"hello", "good", "students", "morning", "people", "world"}
	sort.Strings(strings)
	fmt.Println("after sort by Strings:\t", strings)

	// ！！！在递增的顺序中搜索x，如果存在，返回x的索引，如果不存在，返回x应该插入的数据
	ipos := sort.SearchInts(ints, -1) // int 搜索
	fmt.Printf("pos of 5 is %d th\n", ipos)

	dpos := sort.SearchFloat64s(doubles, 20.1) // float64 搜索
	fmt.Printf("pos of 5.0 is %d th\n", dpos)

	// Float64sAreSorted判定是否是升序
	fmt.Printf("doubles is asc ? %v\n", sort.Float64sAreSorted(doubles))

	doubles = []float64{3.5, 4.2, 8.9, 100.98, 20.14, 79.32}
	sort.Sort(sort.Float64Slice(doubles))         // float64 排序方法 2
	fmt.Println("after sort by Sort:\t", doubles) // [3.5 4.2 8.9 20.14 79.32 100.98]

	(sort.Float64Slice(doubles)).Sort()           // float64 排序方法 3
	fmt.Println("after sort by Sort:\t", doubles) // [3.5 4.2 8.9 20.14 79.32 100.98]

	sort.Sort(Reverse{sort.Float64Slice(doubles)})         // float64 逆序排序
	fmt.Println("after sort by Reversed Sort:\t", doubles) // [100.98 79.32 20.14 8.9 4.2 3.5]
}

// 4、结构体类型的排序
// 结构体类型的排序是通过使用sort.Sort(slice)实现的，只要slice实现了sort.interface的三个方法就可以

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (a PersonSlice) Len() int {
	return len(a)
}

func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PersonSlice) Less(i, j int) bool {
	return a[j].Age < a[i].Age
}

func main04() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonSlice(people)) // 按照 Age 的逆序排序
	fmt.Println(people)

	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	fmt.Println(people)

}

// 5、封装成Wrapper
// 这种方法将 [] Person 和比较的准则 cmp 封装在了一起，形成了 PersonWrapper 函数，然后在其上绑定 Len 、 Less 和 Swap 方法。
// 实际上 sort.Sort(pw) 排序的是 pw 中的 people， 这就是前面说的， go 的排序未必就是针对的一个数组或是 slice， 而可以是一个对象中的数组或是 slice 。

type PersonWrapper struct {
	people []Person
	by     func(p, q *Person) bool
}

func (pw PersonWrapper) Len() int { // 重写 Len() 方法
	return len(pw.people)
}
func (pw PersonWrapper) Swap(i, j int) { // 重写 Swap() 方法
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
func (pw PersonWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return pw.by(&pw.people[i], &pw.people[j])
}

func main05() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age // Age 递减排序
	}})

	fmt.Println(people)
	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return p.Name < q.Name // Name 递增排序
	}})

	fmt.Println(people)

}

// 6、进一步封装
//
type SortBy func(p, q *Person) bool

// 封装成 SortPerson 方法
func SortPerson(people []Person, by SortBy) {
	sort.Sort(sort.Reverse(PersonWrapper{people, by}))
}

func main06() {
	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	fmt.Println(people)

	sort.Sort(PersonWrapper{people, func(p, q *Person) bool {
		return q.Age < p.Age // Age 递减排序
	}})

	fmt.Println(people)

	SortPerson(people, func(p, q *Person) bool {
		return p.Name < q.Name // Name 递增排序
	})

	fmt.Println(people)

	SortPerson(people, func(p, q *Person) bool {
		return p.Name > q.Name
	})
	fmt.Println(people)

}

// 7、另一种思路
//
type PersonPlus struct {
	Name   string
	Weight int
}

type PersonPlusSlice []PersonPlus

func (s PersonPlusSlice) Len() int {
	return len(s)
}

func (s PersonPlusSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByName struct {
	PersonPlusSlice
} // 将 PersonSlice 包装起来到 ByName 中

func (s ByName) Less(i, j int) bool {
	return s.PersonPlusSlice[i].Name < s.PersonPlusSlice[j].Name
} // 将 Less 绑定到 ByName 上

type ByWeight struct {
	PersonPlusSlice
} // 将 PersonSlice 包装起来到 ByWeight 中

func (s ByWeight) Less(i, j int) bool {
	return s.PersonPlusSlice[i].Weight < s.PersonPlusSlice[j].Weight
} // 将 Less 绑定到 ByWeight 上

func main07() {
	s := []PersonPlus{
		{"apple", 12},
		{"pear", 20},
		{"banana", 50},
		{"orange", 87},
		{"hello", 34},
		{"world", 43},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("People by weight:")
	printPeople(s)

	sort.Sort(ByName{s})
	fmt.Println("\nPeople by name:")
	printPeople(s)

}

func printPeople(s []PersonPlus) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}

// test: sort.Slice

func main() {
	ints := []int{2, 4, 56, 1, 6, 2, 6, 7, 3}
	fmt.Println(ints)

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]		// 升序
	})

	fmt.Println(ints)

	pos := sort.SearchInts(ints,56)
	fmt.Println(pos)
}
