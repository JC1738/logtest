package main

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
	"github.com/fatih/structs"
)

//InnerLogger interface is used if struct wants to provide it's own way of returns names, fields, and json string
type InnerLogger interface {
	InnerLogInfo() (names []string, fields []string, jsonStruct string)
}

//DefaultStructInfoFunc is the default implementation for structs to use to return names, fields, and jsonStruct
//It checks if the interface passed in implements InnerLogger and will use that instead
func DefaultStructInfo(o interface{}) (names []string, fields []string, jsonStruct string) {

	if innerLog, ok := o.(InnerLogger); ok {
		names, fields, jsonStruct = innerLog.InnerLogInfo()
		return
	}


	s := structs.New(o)

	names = s.Names()
	fields = []string{}
	

	//Open question, do we recursively call DefaultStructInfo on subsequent fields instuct
	for _, f := range s.Fields() {
		fmt.Printf("field name: %+v\n", f.Name())
		fields = append(fields, f.Name())
		if f.IsExported() {

			fmt.Printf("value   : %+v\n", f.Value())
			fmt.Printf("is zero : %+v\n", f.IsZero())
		}
	}

	jsonStruct = ""
	mapParent := structs.Map(o)
	if mapB, err := json.Marshal(mapParent); err == nil {
		jsonStruct = string(mapB)
	} else {
		fmt.Println("Failure to marshal map", err.Error())
	}
	fmt.Printf("struct: %s, fields: %s, value: %s\n", s.Name(), s.Names(), jsonStruct)

	return
}

//Note you can use `structs:"-"` to decerate a field to let the logger know that you don't a field logged

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
	Age  int
}

type aunt struct {
	name Name
}

func (a aunt) InnerLogInfo() (names []string, fields []string, jsonStruct string) {
	names = []string{"aunt"}
	fields = []string{"name"}
	jsonObj := gabs.New()

	jsonObj.Set(a.name.FullName, "aunt", "name")
	jsonStruct = jsonObj.String()

	return
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
				Age: 33,
			},
			uncle2: Uncle{
				Name: Name{
					FullName: "Mark Wu",
					first:    "Mark",
					last:     "Wu",
				},
				Age: 43,
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
				Age: 33,
			},
			uncle2: Uncle{
				Name: Name{
					FullName: "Mark Wu",
					first:    "Mark",
					last:     "Wu",
				},
				Age: 43,
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
			Age: 33,
		},
	}
	return myParent

}
func printVals(names []string, fields []string, jsonString string) {

	for _, n := range names {
		fmt.Println("name = ", n)
	}
	for _, field := range fields {
		fmt.Println("field = ", field)
	}
	fmt.Println("json = ", jsonString)

}
func main() {
	fmt.Println("vim-go")

	if err := fmt.Errorf("foo"); err != nil {
		fmt.Println("here")
	}

	myParent := populateStruct()
	names, fields, jsonString := DefaultStructInfo(myParent.Uncle1)
	printVals(names, fields, jsonString)


	names, fields, jsonString = DefaultStructInfo(myParent.Aunt1)
	printVals(names, fields, jsonString)

}
