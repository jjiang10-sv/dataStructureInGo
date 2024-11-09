package list

import (
	"fmt"
	"testing"
)

func drivePath(n uint) uint {
	//var res uint = 0
	if n == 2 {
		return 1 + (n - 2)
	}
	return drivePath(n-1) + (n - 2)
}
func Test_Go(t *testing.T) {
	res := drivePath(3)
	fmt.Println(res)
	res = drivePath(4)
	fmt.Println(res)
	res = drivePath(5)
	fmt.Println(res)
	res = drivePath(6)
	fmt.Println(res)
	res = drivePath(7)
	fmt.Println(res)

}

func Test_Slice(t *testing.T) {
	// arr1 := [...]int{2,3,1,5,4,9,5}
	// slice := arr1[1:4]
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(slice), cap(slice))
	// fmt.Printf("the slice is %d \n", slice)
	// fmt.Printf("the pointer of slice %p \n", &slice)
	// fmt.Printf("the pointer of arr1 %p \n", &arr1[0])

	// sli1 := make([]int, 2,4)
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(sli1), cap(sli1))
	// sli2 := make([]int, 2)
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(sli2), cap(sli2))
	// sli3 := []int{1,3,4,2,7,5,9}
	// sli4 := []int{3:3}
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(sli3), cap(sli3))
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(sli4), cap(sli4))

	// var sliNil  []int
	// fmt.Printf("the length of slice %d, the capacity of slice %d , type %T\n", len(sliNil), cap(sliNil), sliNil)
	// fmt.Println(sliNil)

	// sliEmpty := make([]int,0)
	// sliEmpty1 := []int{}
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(sliEmpty), cap(sliEmpty))
	// fmt.Printf("the length of slice %d, the capacity of slice %d \n", len(sliEmpty1), cap(sliEmpty1))

	// sli3 := []int{1, 3, 4, 2, 7, 5, 9}
	// sli3 = append(sli3, 6,7,8)
	// sli10 := []int{1,4,5}
	// sli3 = append(sli3, sli10...)
	// for k,v := range sli3 {
	// 	fmt.Printf("keys as %d; value as %d \n", k,v)
	// }
	// fmt.Printf(" sli3 is %d ; \n ", sli3)
	// sli11 := sli3[:4]
	// sli12 := copy(sli3[4:], sli11)
	// fmt.Printf(" sli3 is %d ; sli12 is %d \n ", sli3, sli12)
	// fmt.Printf(" sli3[:4] is %d ; sli3[4:] is %d \n ", sli3[:4], sli3[4:])

	s15 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	s16 := s15[:4]
	newSlice := copy(s15[:4], s16)
	fmt.Printf("newSlice = %d，s15 = %v", newSlice, s15)

}

func Test_Map(t *testing.T) {
	map1 := make(map[string]int)
	fmt.Println("map1 = ", map1)
	fmt.Printf("map1的长%d, 类型：%T,地址：%p \n", len(map1), map1, map1)
	map2 := map[string]int{}
	fmt.Println("map2 = ", map2)
	fmt.Printf("map2的长度: %d, 类型: %T, 地址: %p \n", len(map2), map2, map2)
	var map3 map[string]int
	fmt.Println("map3 = ", map3)
	fmt.Printf("map3的长度:%d, 类型: %T，地址: %p\n", len(map3), map3, map3)
	if map1 == nil {
		fmt.Println("man1 # nil")
	} else if map2 == nil {
		fmt.Println("map2 # nil")
	} else if map3 == nil {
		fmt.Println("map3 # nil")
	} else {
		fmt.Println("都不是nil")
	}
	map2["One"] = 1
	map2["Two"] = 2
	map2["Three"] = 3
	map2["Four"] = 4
	map2["Five"] = 5
	fmt.Println("map2 = ", map2)
	fmt.Printf("map2的长度:%d, 类型：%T,地址：%p\n", len(map2), map2, map2)

	if value, ok := map2["Two"]; !ok {
		fmt.Print("map2中不存在键值为Two的元素")
	} else {
		fmt.Printf("value = %d, ok = %T\n", value, value)
	}

	fmt.Println("=====分割袋=")
	for key, value := range map2 {
		fmt.Printf("#-##5: key = %s, value = %d\n", key, value)
	}
	for key, value := range map2 {
		fmt.Printf("第二次遍：key= %s,value= %d", key, value)
	}
	delete(map2, "Three")
	for key, value := range map2 {
		fmt.Printf("map2：key=%s,value=%d\n", key, value)
	}
	map2["One"] = 11
	fmt.Printf("map2的键值为One的元素: %d \n", map2["One"])
}

type User struct {
	Name   string
	Age    uint8
	IsMale bool
}

func (u *User) talk() string {
	u.Name = "john"
	u.IsMale = true
	s := fmt.Sprintf(" the name is %v , age is %d , isMake %v", u.Name, u.Age, u.IsMale)
	return s
}

func Test_Struct(t *testing.T) {
	// type User struct {
	// 	Name   string
	// 	Age    uint8
	// 	IsMale bool
	// }

	type Admin struct {
		person User
		level  uint8
	}

	//var Bob User
	var lily = User{
		Name:   "lily",
		Age:    32,
		IsMale: false,
	}
	fmt.Println(lily.talk(), lily.Name)

	// var _ = User{"lucy", 43, false}
	// adminLily := Admin{
	// 	person: lily,
	// 	level:  3,
	// }
	// fmt.Printf("the val is %v; the type is %T, the address is %p \n", Bob, Bob, &Bob)
	// fmt.Printf("the val is %v; the type is %T, the address is %p \n", lily, lily, &lily)
	// fmt.Printf("the val is %v; the type is %T, the address is %p \n", adminLily, adminLily, &adminLily)

}

func Test_Err(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err is %v ", err)
		}
	}()

	defer func() {
		panic("err 2")
	}()

	panic("err 1")
	fmt.Println("exited")
}
