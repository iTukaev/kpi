package sheets

import (
	"strconv"
	"strings"
)

var rNotation = [...]int{0, -1, 100, 150, 200}

func GetSalaryFuncs(e []*Employee) *SalaryFuncs {
	sFuncs := NewSalaryFuncs()
	for i, emp := range e {
		sf := calculate(emp.Salary)

		fc := fxToString(sf)
		if _, ok := sFuncs.Funcs[fc]; !ok {
			sFuncs.Funcs[fc] = sf
		}
		e[i].FuncCode = fc
	}
	return sFuncs
}

func calculate(salary []int) *SalaryFunc {
	sf := NewSalaryFunc(len(rNotation))

	for i := 0; i < len(rNotation); {
		if rNotation[i] == -1 {
			sf.X[i] = salary[0] * 100 / salary[1]
		} else {
			sf.X[i] = rNotation[i]
		}
		if sf.X[i] == 0 {
			sf.Fx[i] = salary[0] * 100 / salary[1]
		} else {
			sf.Fx[i] = salary[i-1] * 100 / salary[1]
		}
		i++
	}

	return sf
}

func fxToString(sf *SalaryFunc) Code {
	res := ""
	for i := 0; i < len(sf.Fx); i++ {
		res += strconv.Itoa(sf.Fx[i]) + "_"
	}
	res = strings.TrimSuffix(res, "_")
	return Code(res)
}
