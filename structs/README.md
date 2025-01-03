#### STRUCT

Structs are user-defined types that help us create a collection of data describing a single entity.

- Structs are typed collections of fields. They are useful for grouping related data together to form records.
- Structs are used to describe entities in your application, such as users, products, or orders.

```go
type Person struct {
    name string
    age  int
    city string
}
```

In this example, the `Person` struct has three fields: `Name`, `Age`, and `City`.
The keyword `type` is there to create a new type. It is called `type definition`.
The sample code is basically saying **create a type called `Person` based on a struct of name string, age int and city string**.

**Struct instances:**

```go
var p1 Person
p1 = Person{name:"John Doe", age:30, city:"New York"}
p2 := Person{name:"Jane Doe", age:25, city:"London"}
```

or

```go
var p1 Person
p1 = Person{"John Doe", 30, "New York"} // This is also called struct literal or composite literal
p2 := Person{"Jane Doe", 25, "London"}
```

- You can also created a struct instance with a null value

```go
var p1 Person
p1 = Person{}
```

NB: Structs can be passed as pointers to functions. And the values of the struct fields can be accessed using the `.` operator without `dereferencing`. The Go compiler automatically dereferences the pointer for you. But technically you it's a good idea to dereferencing the struct to access its fields.

#### METHODS

Methods are functions that are associated with a struct. Methods can be defined for either pointer or value receiver types.

```go
func (u User) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}
```

`func (u User)` in the method definition is called receiver
In other to be able to mutate struct data in a method, the `receiver` should `not accept the struct as value but instead a pointer`

** Constructor function **: a function that creates a new instance of a struct

```go
func newUser(firstName string, lastName string, birthdate time.Time) *User {
    return &User{
        firstName: firstName,
        lastName:  lastName,
        birthdate: birthdate,
    }
}
```

constructor function can provide extra benefit than just creating a struct instance. It can be used for validation and initialization of the struct.

```go
func newUser(firstName string, lastName string, birthdate time.Time) *User {
    if firstName == "" || lastName == ""  || birthdate == "" {
        return nil, error.New("invalid input")
    }
    return &User{
        firstName: firstName,
        lastName:  lastName,
        birthdate: birthdate,
    }
}
```

**STRUCT EMBEDDING**

Struct embedding is a technique that allows a struct to embed another struct as a field. This means that the fields of the embedded struct are also part of the struct that is embedding it.

```go
type User struct {
    Name string
    Age  int
    Address Address

}

type Address struct {
    Street string
    City   string
    Zip    string
    User User // just put User struct here without any name
}
```
