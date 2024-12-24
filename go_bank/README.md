### GO BANK APP

This is a simple bank app that allows you to create accounts and make transactions between them.

**The purpose of this project is to learn about control structures, loops, error handling and basics of working with files in Go, and to practice using functions, variables, and data types.**

**`App doesnot crash in go when an error occurs instead it return an error`**

- Errors in Go are treated as **values**, not **exceptions.**
- A function that can fail typically returns two values

```go
- The results
- An error value
```

- The most common error type in Go is the `error` interface. An error can be thrown using the `errors.New()` function from the `errors` package

- The fmt package formats an error value by calling its `Error()` string method.

**Why nil?**

`Representation of "No Error"`: In Go, the `nil` value for the error type explicitly signifies the absence of an error.
`Simplicity`: It aligns with Go's philosophy of simplicity. By checking

```go
if err == nil
```

you can clearly see whether an error occurred.
`Consistency`: nil is used in Go to represent zero values for pointers, interfaces, maps, slices, and channels. Using nil for errors keeps the language consistent.

**`Unlike exceptions, errors are not thrown and caught; they are explicitly returned and checked.`**
