package main

import (
	"fmt"
	"os"
	"testing"
)

//func TestAdd(t *testing.T){
//
//}

func TestMain(m *testing.M){
	fmt.Println("setup")
	res := m.Run()
	fmt.Println("tear-down")

	os.Exit(res)
}

func TestMultiple(t *testing.T){
var x, y, result = 2, 2, 4

realResult := Multiple(x,y)

if realResult != result{
	t.Errorf("expected result: %d != real result: %d", result, realResult)
}
}

func TestMultiple2(t *testing.T){
	//setup

	t.Cleanup(func(){
	fmt.Println("Teardown on cleanup")
	})

	// insert mock data for db
	t.Run("groupA", func(t *testing.T){
		t.Run("simple", func(t *testing.T){
		t.Parallel()
		t.Log("simple")
		var x, y, result = 2, 2, 4

		realResult := Multiple(x,y)

		if realResult != result{
			t.Errorf("expected result: %d != real result: %d", result, realResult)
		}
	})
		t.Run("medium", func(t *testing.T){
		t.Parallel()
		t.Log("medium")
		var x, y, result = 222, 222, 49284

		realResult := Multiple(x,y)

		if realResult != result{
			t.Errorf("expected result: %d != real result: %d", result, realResult)
		}
	})
		t.Run("negative", func(t *testing.T){
		t.Parallel()
		t.Log("negative")
		var x, y, result = -2, 4, -8

		realResult := Multiple(x,y)

		if realResult != result{
			t.Errorf("expected result: %d != real result: %d", result, realResult)
		}
		t.Run("1", func(t *testing.T) {
			r := Multiple(1,1)
			if r != 1{
				t.Errorf("failed")
			}
		})
	})
	})
	//TearDown
	// delete mock data for db
}

func TestAdd(t *testing.T){
	// insert mock data for db

	t.Run("simple", func(t *testing.T){

		var x, y, result = 2, 2, 4

		realResult := Multiple(x,y)

		if realResult != result{
			t.Errorf("expected result: %d != real result: %d", result, realResult)
		}
	})

	t.Run("medium", func(t *testing.T){

		var x, y, result = 2, 2, 4

		realResult := Add(x,y)

		if realResult != result{
			t.Errorf("expected result: %d != real result: %d", result, realResult)
		}
	})

}