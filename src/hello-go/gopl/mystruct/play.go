package mystruct

import (
	"time"
	"hello-go/gopl"
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func init() {
	gopl.AddRunFunc(run)
}

func run() {

	// 写成员
	dilbert.Salary -= 5000
	// 成员取址
	position := &dilbert.Position
	// 用指针访问成员变量
	*position = "Senior" + *position

	// 指向结构体的指针
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proact team player)"
	// 等价
	// (*employeeOfTheMonth).Position += "(proact team player)"

	fmt.Println(dilbert)
	// 根据ID 返回一个员工
	fmt.Println(EmployeeByID(3))
	fmt.Println(EmployeeByID(4))
	var emp *Employee
	emp = EmployeeByID(2)
	if emp != nil {
		fmt.Println(emp)
		fmt.Printf("工资 %v \n ", emp.Salary)
		// 更新工资
		/*
		id := emp.ID
		EmployeeByID(id).Salary = 5
		*/
		emp.Salary = 5
		fmt.Println(emp)
	}

	emp = EmployeeByID(5)
	if emp == nil {
		fmt.Println("not found the emp with id 5")
	}
}
/**
*   如果将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，
*   因为在赋值语句的左边并不确定是一个变量（译注：调用函数返回的是值，并不是一个可取地址的变量）
* ——————————————————————————————————————————————————
*                        如果返回结构体类型 而非指针 时  函数结果只读 不可写！
 */
func EmployeeByID(id int) *Employee {
	db := []Employee{
		Employee{ManagerID:1, Name:"yiqing"},
		Employee{ManagerID:2, Name:"Bob"},
		Employee{ManagerID:3, Name:"lucy"},
		Employee{ManagerID:1, Name:"Henrry"},
	}
	for _, emp := range db {
		if (emp.ManagerID == id) {
			return &emp
		}
	}
	return nil
}