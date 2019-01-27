package funcache

import "testing"

func fib(in int) int {
	if in <= 1 {
		return 1
	}
	return fib(in-2) + fib(in-1)
}

func TestCacheFunWithLru(t *testing.T) {
	wrapped := func(in []interface{}) ([]interface{}, error){
		res := []interface{}{
			fib(in[0].(int)),
		}
		return res, nil
	}
	wrapper := CacheFunWithLru(32, wrapped, 3)
	if out, err:=wrapper([]interface{}{40});out[0].(int)!=fib(40) || err!=nil{
		t.Errorf("expect %d got %d with %v", fib(40), out[0].(int), err)
	}
}

func TestCacheFunWithLfu(t *testing.T) {
	wrapped := func(in []interface{}) ([]interface{}, error){
		res := []interface{}{
			fib(in[0].(int)),
		}
		return res, nil
	}
	wrapper := CacheFunWithLfu(32, wrapped, 3)
	if out, err:=wrapper([]interface{}{40});out[0].(int)!=fib(40) || err!=nil{
		t.Errorf("expect %d got %d with %v", fib(40), out[0].(int), err)
	}
}

func TestCacheFunWithArc(t *testing.T) {
	wrapped := func(in []interface{}) ([]interface{}, error){
		res := []interface{}{
			fib(in[0].(int)),
		}
		return res, nil
	}
	wrapper := CacheFunWithArc(32, wrapped, 3)
	if out, err:=wrapper([]interface{}{40});out[0].(int)!=fib(40) || err!=nil{
		t.Errorf("expect %d got %d with %v", fib(40), out[0].(int), err)
	}
}

func BenchmarkRawFib40(b *testing.B) {
	for i:=0;i<b.N;i++{
		fib(40)
	}
}

func BenchmarkCacheFunWithLruFib40(b *testing.B) {
	b.StopTimer()
	wrapped := func(in []interface{}) ([]interface{}, error){
		res := []interface{}{
			fib(in[0].(int)),
		}
		return res, nil
	}
	wrapperFib := CacheFunWithLru(64, wrapped, 3)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		if out, err:=wrapperFib([]interface{}{40});out[0].(int)!=165580141 || err!=nil{
			b.Errorf("expect %d got %d with %v", 165580141, out[0].(int), err)
		}
	}
}

func BenchmarkCacheFunWithLfuFib40(b *testing.B) {
	b.StopTimer()
	wrapped := func(in []interface{}) ([]interface{}, error){
		res := []interface{}{
			fib(in[0].(int)),
		}
		return res, nil
	}
	wrapperFib := CacheFunWithLfu(64, wrapped, 3)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		if out, err:=wrapperFib([]interface{}{40});out[0].(int)!=165580141 || err!=nil{
			b.Errorf("expect %d got %d with %v", 165580141, out[0].(int), err)
		}
	}
}

func BenchmarkCacheFunWithArcFib16(b *testing.B) {
	b.StopTimer()
	wrapped := func(in []interface{}) ([]interface{}, error){
		res := []interface{}{
			fib(in[0].(int)),
		}
		return res, nil
	}
	wrapperFib := CacheFunWithArc(64, wrapped, 3)
	b.StartTimer()
	for i:=0;i<b.N;i++{
		if out, err:=wrapperFib([]interface{}{40});out[0].(int)!=165580141 || err!=nil{
			b.Errorf("expect %d got %d with %v", 165580141, out[0].(int), err)
		}
	}
}