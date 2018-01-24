package bench

import "reflect"

type Data struct {
	X int
}

func (d *Data) Inc() {
	d.X++
}

var d = new(Data)

func set(x int) {
	d.X = x
}

var dv = reflect.ValueOf(d).Elem()
var f = dv.FieldByName("X")

func rset(x int) {
	f.Set(reflect.ValueOf(x))
}

func call() {
	d.Inc()
}

var dp = reflect.ValueOf(d)
var m = dp.MethodByName("Inc")

func rcall() {
	m.Call(nil)
}

