# logtest
Return import data about structs for logging


This small example shows data extracted from a struct for logging purposes.  The idea is that the names of the fields in the struct as well as their types are import for an index look up query of logs.  Then the actual contents of the struct is okay represented as a JSON blob.  So if you happen to use Elastic Search you can index the names, and types, and leave json representation as unindex field.


It requires the fields be Exported from Struct (starts with capital letter).  If you don't export the fields, but you still want to have items indexed and a JSON representation, you can implement the following interface:

```
//InnerLogger interface is used if struct wants to provide it's own way of returns names, types, and json string
type InnerLogger interface {
	InnerLogInfo() (names []string, types []string, jsonStruct string)
}
```

It is also possible to decerate the struct so that you admit fields:

```
`structs:"-"`
`structs:",omitnested"`
`structs:",omitempty"`
```
