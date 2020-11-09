package main
import "fmt"

type Person struct{
	Name string
	age int
	sal float64
}

func NewPerson(name string) *Person{
	return &Person{
		Name:name,
	}
}

func (p *Person) SetAge(age int){
	if age>0 && age<150{
		p.age=age
	}else{
		fmt.Println("年龄范围不正确..")
	}
}


func (p *Person) GetAge() int{
	return p.age
}

func (p *Person) SetSal(sal float64){
	if sal>=3000 && sal<30000{
		p.sal=sal
	}else{
		fmt.Println("薪水范围不正确!")
	}
}


func (p *Person) GetSal() float64{
	return p.sal
}

func main(){
	p := NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name,"age=",p.GetAge(),"salary=",p.GetSal())
}
