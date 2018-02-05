package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
)


var DefaultStructInfo = func(s interface{}) (name string, fields string[], )


//Name struct
type Name struct {
	FullName string
	first    string
	last     string
}

//Parent struct
type Parent struct {
	Name      Name
	Age       int
	FirstKid  Child
	secondKid Child
	Aunt1     aunt
	aunt2     aunt
	Uncle1    Uncle
}

//Child struct
type Child struct {
	name   Name
	age    int
	Aunt   aunt
	Uncle1 Uncle
	uncle2 Uncle
}

//Uncle struct
type Uncle struct {
	Name Name
	age  int
}

type aunt struct {
	name Name
}

func populateStruct() *Parent {

	myParent := &Parent{
		Name: Name{
			FullName: "Jim Castillo",
			first:    "Jim",
			last:     "Castillo",
		},
		Age: 43,
		FirstKid: Child{
			name: Name{
				FullName: "Kate Castillo",
				first:    "Kate",
				last:     "Castillo",
			},
			age: 6,
			Aunt: aunt{
				name: Name{
					FullName: "Kristin Castillo",
					first:    "Kristin",
					last:     "Castillo",
				},
			},
			Uncle1: Uncle{
				Name: Name{
					FullName: "Eric Castillo",
					first:    "Eric",
					last:     "Castillo",
				},
				age: 33,
			},
			uncle2: Uncle{
				Name: Name{
					FullName: "Mark Wu",
					first:    "Mark",
					last:     "Wu",
				},
				age: 43,
			},
		},
		secondKid: Child{
			name: Name{
				FullName: "Alex Castillo",
				first:    "Alex",
				last:     "Castillo",
			},
			age: 5,
			Aunt: aunt{
				name: Name{
					FullName: "Kristin Castillo",
					first:    "Kristin",
					last:     "Castillo",
				},
			},
			Uncle1: Uncle{
				Name: Name{
					FullName: "Eric Castillo",
					first:    "Eric",
					last:     "Castillo",
				},
				age: 33,
			},
			uncle2: Uncle{
				Name: Name{
					FullName: "Mark Wu",
					first:    "Mark",
					last:     "Wu",
				},
				age: 43,
			},
		},
		Aunt1: aunt{
			name: Name{
				FullName: "Kristin Castillo",
				first:    "Kristin",
				last:     "Castillo",
			},
		},
		aunt2: aunt{
			name: Name{
				FullName: "Yvonne Hao",
				first:    "Yvonne",
				last:     "Hao",
			},
		},
		Uncle1: Uncle{
			Name: Name{
				FullName: "Eric Castillo",
				first:    "Eric",
				last:     "Castillo",
			},
			age: 33,
		},
	}
	return myParent

}
func main() {
	fmt.Println("vim-go")

	if err := fmt.Errorf("foo"); err != nil {
		fmt.Println("here")
	}

	myParent := populateStruct()
	fmt.Println("%v", myParent)

	n := structs.Names(myParent)
	for _, name := range n {
		fmt.Printf("%s\n", name)
	}

	s := structs.New(myParent)
	for _, f := range s.Fields() {
		fmt.Printf("field name: %+v\n", f.Name())

		if f.IsExported() {
			fmt.Printf("value   : %+v\n", f.Value())
			fmt.Printf("is zero : %+v\n", f.IsZero())
		}
	}

	mapJSONString := ""
	mapParent := structs.Map(myParent)
	if mapB, err := json.Marshal(mapParent); err == nil {
		mapJSONString = string(mapB)
	}
	fmt.Printf("struct: %s, fields: %s, value: %s\n", s.Name(), s.Names(), mapJSONString)
}
