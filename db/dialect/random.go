package dialect

import "github.com/astaxie/beego/orm"

/*
	Dialect: RandomInt
*/
func RandomBuiltinFunc() string {
	d := orm.NewOrm().Driver()
	if d.Type() == orm.DR_Oracle {
		return "dbms_random.value(1,100000)"
	} else if d.Type() == orm.DR_MySQL {
		return "RAND()"
	} else {
		return "RANDOM()"
	}
}
