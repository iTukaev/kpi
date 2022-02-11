package sheets

type Code string

type Employee struct {
	MoID       string
	FirstName  string
	LastName   string
	MiddleName string
	Salary     []int
	FuncCode   Code
}

type SalaryFunc struct {
	X  []int
	Fx []int
}

func NewSalaryFunc(len int) *SalaryFunc {
	return &SalaryFunc{
		X:  make([]int, len),
		Fx: make([]int, len),
	}
}

type SalaryFuncs struct {
	Funcs map[Code]*SalaryFunc
}

func NewSalaryFuncs() *SalaryFuncs {
	return &SalaryFuncs{
		Funcs: make(map[Code]*SalaryFunc),
	}
}
