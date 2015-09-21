package template

import (
	"container/list"
	"reflect"

	"github.com/ikeikeikeike/shuffler"
)

func List(l *list.List) chan interface{} {
	e := l.Front()
	c := make(chan interface{})
	go func() {
		for e != nil {
			c <- e.Value
			e = e.Next()
		}
		close(c)
	}()
	return c
}

func Shuffle(in interface{}) (list []interface{}) {
	elems := shuffler.Shuffler(in).(reflect.Value)
	for i := 0; i < elems.Len(); i++ {
		list = append(list, elems.Index(i).Interface())
	}
	return list
}

func Reverse(in interface{}) (list []interface{}) {
	elems := reverser(in).(reflect.Value)
	for i := 0; i < elems.Len(); i++ {
		list = append(list, elems.Index(i).Interface())
	}
	return list
}

func Slice(in interface{}, from, to int) (list []interface{}) {
	elems := slicer(in, from, to).(reflect.Value)
	for i := 0; i < elems.Len(); i++ {
		list = append(list, elems.Index(i).Interface())
	}
	return list
}

func slicer(src interface{}, from, to int) interface{} {
	s := reflect.ValueOf(src)
	t := reflect.TypeOf(src)

	dest := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, 0)
	for i := s.Len() - 1; i >= 0; i-- {
		if from <= i && i <= to {
			dest = reflect.Append(dest, s.Index(i))
		}
	}
	return dest
}

func reverser(src interface{}) interface{} {
	s := reflect.ValueOf(src)
	t := reflect.TypeOf(src)

	dest := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, 0)
	for i := s.Len() - 1; i >= 0; i-- {
		dest = reflect.Append(dest, s.Index(i))
	}
	return dest
}
