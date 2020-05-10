package godash

import (
	"reflect"
	"testing"
)

func identificator(search interface{}) func(interface{}) bool {
	return func(value interface{}) bool {
		return reflect.DeepEqual(search, value)
	}
}

func TestFindIndex(t *testing.T) {
	slice := make([]int, 8, 20)
	pred := identificator(0)
	test1, err1 := FindIndex(slice, pred, 0)
	if test1 != 0 || err1 != nil {
		t.Errorf("FindIndex test 1 expected 1 and no error, got %d", test1)
	}

	pred2 := identificator(1)
	test2, err2 := FindIndex(slice, pred2, 0)
	if test2 != -1 || err2 != nil {
		t.Errorf("FindIndex test 2 expected -1 and no error, got %d", test2)
	}

	slice[0] = 1
	test3, err3 := FindIndex(slice, pred2, 2)
	if test3 != -1 || err3 != nil {
		t.Errorf("FindIndex test 3 expected -1 and no error, got %d", test3)
	}

	slice[7] = 1
	test4, err4 := FindIndex(slice, pred2, 0)
	if test4 != 0 || err4 != nil {
		t.Errorf("FindIndex test 4 expected 0 and no error, got %d", test4)
	}

	test5, err5 := FindIndex(slice, pred2, 2)
	if test5 != 7 || err5 != nil {
		t.Errorf("FindIndex test 5 expected 7 and no error, got %d", test5)
	}

	type box struct {
		x int
		y int
	}

	pred3 := identificator(box{1, 2})
	boxes := make([]box, 4, 8)
	test6, err6 := FindIndex(boxes, pred3, 0)
	if test6 != -1 || err6 != nil {
		t.Errorf("FindIndex test 6 expected -1 and no error, got %d", test6)
	}

	boxes[1] = box{1, 2}
	test7, err7 := FindIndex(boxes, pred3, 0)
	if test7 != 1 || err7 != nil {
		t.Errorf("FindIndex test 7 expected 1 and no error, got %d", test7)
	}

	b := box{1, 2}
	test8, err8 := FindIndex(b, pred3, 0)
	if test8 != -1 || err8.Error() != badParameterMsg {
		t.Errorf("FindIndex test 7 expected -1 and error, got %d", test8)
	}
}

func TestFind(t *testing.T) {
	slice := make([]int, 8, 20)
	pred := identificator(0)
	test1, err1 := Find(slice, pred, 0)
	if test1 != 0 || err1 != nil {
		t.Errorf("Find test 1 expected 0 and no error, got %d", test1)
	}

	pred2 := identificator(1)
	test2, err2 := Find(slice, pred2, 0)
	var empty interface{}
	if test2 != empty || err2 != nil {
		t.Errorf("Find test 2 expected empty interface and no error, got %d", test2)
	}

	slice[0] = 1
	test3, err3 := Find(slice, pred2, 2)
	if test3 != empty || err3 != nil {
		t.Errorf("Find test 3 expected empty interface and no error, got %d", test3)
	}

	slice[7] = 1
	test4, err4 := Find(slice, pred2, 0)
	if test4 != 1 || err4 != nil {
		t.Errorf("Find test 4 expected 1 and no error, got %d", test4)
	}

	test5, err5 := Find(slice, pred2, 2)
	if test5 != 1 || err5 != nil {
		t.Errorf("Find test 5 expected 1 and no error, got %d", test5)
	}

	type box struct {
		x int
		y int
	}

	pred3 := identificator(box{1, 2})
	boxes := make([]box, 4, 8)
	test6, err6 := Find(boxes, pred3, 0)
	if test6 != empty || err6 != nil {
		t.Errorf("Find test 6 expected empty interface and no error, got %d", test6)
	}

	boxes[1] = box{1, 2}
	test7, err7 := Find(boxes, pred3, 0)
	if test7 != boxes[1] || err7 != nil {
		t.Errorf("Find test 7 expected box{1,2} and no error, got %d", test7)
	}

	b := box{1, 2}
	test8, err8 := Find(b, pred3, 0)
	if test8 != empty || err8.Error() != badParameterMsg {
		t.Errorf("Find test 7 expected empty interface and error, got %d", test8)
	}

}
