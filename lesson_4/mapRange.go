package main

import "fmt"
//import "github.com/mitchellh/mapstructure"

type Point struct {
	X int
	Y int
}

func(p Point) method() {
	fmt.Println("call Point method")
}


func main(){

	//range

	//arr := []string{"a", "b", "c"}
	//
	//for i, l := range arr {
	//	fmt.Println("i: ", i, "  l: ",l)
	//}

	//initilation map

	//pointsMap := map[string] Point{}
	//otherMap := map[string]int{}
	otherPointMap := make(map[string]Point)

	otherPointMap["key"] = Point{1,2}
	fmt.Println(otherPointMap)
	fmt.Println(otherPointMap["key"])


	var pointsMap map[string]Point
	pointsMap["stringKey"] = Point{1,2}

	//Code and Decode map to struct

	//pointMap := map[string]int{
	//	"x": 1,
	//	"y": 2,
	//}

	//mapstructure.Decode(pointMap, )

}
