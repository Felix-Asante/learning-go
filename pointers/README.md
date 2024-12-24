### POINTERS

Pointers are variables that store the memory address of another variable.

### ADVANTAGES OR WHY YOU WOULD USE POINTERS

1. Pointers are useful when you want to pass a variable by reference instead of by value.
2. Pointers are useful when you want to access the memory address of a variable.
3. Pointers are useful when you want to modify the value of a variable without changing its address.
4. Pointers are useful when you want to create a new variable that points to the same memory address as another variable.

In summary, pointers are useful when you want to **avoid unnecessary copies of data** and **directly mutate values**.

##### HOW POINTERS AVOID UNNECESSARY COPIES

- Generally when you create a variablel and pass it as an argument to a function, Go create a copy of that variable and pass it to the function to be used. Thereby making having the same value occupying multiple space in the memory until these variables are not in used anymore.
- The `Garbage Collector` in Go is responsible for cleaning up these unnecessary copies of data when they are no longer in use.
- By using pointers we can avoid the need to make a copy of the variable and pass it to the function.

**dereferencing pointers**: When you want to access the value of a variable that is stored in a pointer, you can use the `*` operator to dereference the pointer.

```go
package main

import "fmt"

func main() {

    var a int = 10
    var b *int = &a

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(*b)
}
```

- The `null-value` of a pointer is `nil`. **nil**: represents the absence of an (address) value. That is a pointer pointing at no address / no value in memory.

** It is important to know that `const` variables cannot be used as pointers.**
-In Go, `constants` cannot be used as pointers because they are `immutable` and `do not occupy memory at runtime`. Constants are resolved at `compile time`, and they are not addressable (i.e., they do not have a memory address).
