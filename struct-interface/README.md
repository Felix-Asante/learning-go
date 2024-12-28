#### INTERFACES

In Go, an interface is a type that defines a set of method signatures.

- Interface can also be used as types.
- Interfaces are used to define a common behavior for a group of related types.

**The any Interface type**
Go also support a special interface type that allows a variable or function to accept `any` types of values.
**` interface {}`**

```go
var variable interface{}
variable = 10
variable = "hello"
variable = true
```

Go also has a concept that is called:
**The type assertions**
A type assertion provides access to an interface value's underlying concrete value.

`t := i.(T)`
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the `underlying value` and a `boolean` value that reports whether the assertion succeeded.

`t, ok := i.(T)`
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

**Type Switch (variable.(type))**
Type switches are a safer way to inspect and act on the dynamic type of an interface value.

```go
switch variable.(type) {
case int:
    fmt.Println("int")
case string:
    fmt.Println("string")
case bool:
    fmt.Println("bool")
default:
    fmt.Println("unknown")
}
```
