package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Jeffail/gabs"
	"github.com/fatih/structs"
)

/*
*
* The idea of this file is to show an example of Logging if you want to have a log line that consists of two indexed
* types and as JSON blob.
*
* names is an index field that contains all the names of the structs under examination
* types is an index field that contains all of the type names in the structs under examination
* jsonStruct is non index field that contains the blob of the struct
*
 */

//InnerLogger interface is used if struct wants to provide it's own way of returns names, types, and json string
type InnerLogger interface {
	InnerLogInfo() (names []string, types []string, jsonStruct string)
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

//DefaultStructInfoFunc is the default implementation for structs to use to return names, types, and jsonStruct
//It checks if the interface passed in implements InnerLogger and will use that instead
func DefaultStructInfo(o interface{}) (names []string, types []string, jsonStruct string) {

	if innerLog, ok := o.(InnerLogger); ok {
		names, types, jsonStruct = innerLog.InnerLogInfo()
		return
	}

	s := structs.New(o)

	names = s.Names()
	names = append(names, s.Name())
	types = []string{}

	types = append(types, getType(o))

	for _, f := range s.Fields() {
		if f.IsExported() {
			types = append(types, getType(f.Value()))
			if f.Kind() == reflect.Struct {
				innerNames, innerTypes, _ := DefaultStructInfo(f.Value())
				names = append(names, innerNames...)
				types = append(types, innerTypes...)
			}
			//fmt.Printf("value   : %+v\n", f.Value())
			//fmt.Printf("is zero : %+v\n", f.IsZero())
			//fmt.Printf("is kind : %s\n", f.Kind().String())
			//fmt.Printf("is type : %s\n", getType(f.Value()))
		}
	}

	jsonStruct = ""
	mapParent := structs.Map(o)
	if mapB, err := json.Marshal(mapParent); err == nil {
		jsonStruct = string(mapB)
	} else {
		fmt.Println("Failure to marshal map", err.Error())
	}

	return
}

//Note you can use `structs:"-"` to decerate a field to let the logger know that you don't a field logged
//You can also use the following:
//
/// A value with the option of "omitnested" stops iterating further if the type
// is a struct. Example:
//
//   // Fields is not processed further by this package.
//   Field time.Time     `structs:",omitnested"`
//   Field *http.Request `structs:",omitnested"`
//
// A tag value with the option of "omitempty" ignores that particular field and
// is not added to the values if the field value is empty. Example:
//
//   // Field is skipped if empty
//   Field string `structs:",omitempty"`
//
//
//
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

type Car struct {
	Make  string
	Model string
}

//Uncle struct
type Uncle struct {
	Name     Name
	Age      int
	UncleCar Car
}

type aunt struct {
	name Name
}

//InnerLogInfo used to replace default returing of names, types, and jsonStruct
//This is especially useful if fields are not exported and want to still log
func (a aunt) InnerLogInfo() (names []string, types []string, jsonStruct string) {
	names = []string{"aunt"}
	types = []string{"Name"}
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
				UncleCar: Car{
					Make:  "Buick",
					Model: "Regal",
				},
			},
			uncle2: Uncle{
				Name: Name{
					FullName: "Mark Wu",
					first:    "Mark",
					last:     "Wu",
				},
				Age: 43,
				UncleCar: Car{
					Make:  "Lexus",
					Model: "350",
				},
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
				UncleCar: Car{
					Make:  "Buick",
					Model: "Regal",
				},
			},
			uncle2: Uncle{
				Name: Name{
					FullName: "Mark Wu",
					first:    "Mark",
					last:     "Wu",
				},
				Age: 43,
				UncleCar: Car{
					Make:  "Lexus",
					Model: "350",
				},
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
			UncleCar: Car{
				Make:  "Buick",
				Model: "Regal",
			},
		},
	}
	return myParent

}
func printVals(names []string, types []string, jsonString string) {

	for _, n := range names {
		fmt.Println("name = ", n)
	}
	for _, t := range types {
		fmt.Println("type = ", t)
	}
	fmt.Println("json = ", jsonString)
	fmt.Println("***********************************************************")

}
func main() {
	fmt.Println("vim-go")

	if err := fmt.Errorf("foo"); err != nil {
		fmt.Println("here")
	}

	myParent := populateStruct()

	names, types, jsonString := DefaultStructInfo(myParent)
	printVals(names, types, jsonString)

	names, types, jsonString = DefaultStructInfo(myParent.Uncle1)
	printVals(names, types, jsonString)

	names, types, jsonString = DefaultStructInfo(myParent.Aunt1)
	printVals(names, types, jsonString)

}
