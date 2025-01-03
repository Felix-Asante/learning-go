### A SIMPLE INVESTMENT CALCULATOR IN GO

This is a simple investment calculator in Go.

The program takes three arguments: the investment amount, the expected return rate, and the number of years. It then calculates the future value of the investment and prints it to the console.

#### PUPROSE OF THIS PROJECT

- To understand the basics of variables in go
- To understand data types and type conversion in Go
- To understand how to use some of Go standard library functions

#### GO VARIABLES

Variables in Go are defined using the `var` keyword followed by a variable name and a type. For example:

```go
var name string = "John Doe"
```

Go will create a variable called `name` of type `string` and assign it the value "John Doe".
or

```go
var age = 12
```

In this case, Go will create a variable called `age` and infer its type as `int` and assign it the value 12. `Int` because we didn't explicitly assign a type to the variable , so go will infer the type from the value.

**A short cut to declare a variable is to use the `:=` operator to assign a value to a variable without explicitly specifying the type\_**

```go
name := "John Doe"
```

It is also important to note that multiple variables can be declared on the same line using comma-separated list of variables and their values:

```go
name, age := "John Doe", 12
```

In this case, Go will create two variables called `name` and `age` and assign them the values "John Doe" and 12 respectively.

You can also tell Go to use a the same data type for multiple variables by using the `:=` operator:\*\*

```go
investmentAmount, expectedReturnRate Float64 := 1000.0, 12.0
```

#### GO TYPES AND NULL VALUES

##### Basic Types

Go comes with a couple of built-in basic types:

`int` => A number WITHOUT decimal places (e.g., -5, 10, 12 etc)

`float64` => A number WITH decimal places (e.g., -5.2, 10.123, 12.9 etc)

`string` => A text value (created via double quotes or backticks: "Hello World", `Hi everyone')

`bool` => true or false

<b>But there also are some noteworthy "niche" basic types which you'll typically not need that often but which you should still know about:</b>

`uint` => An unsigned integer which means a strictly non-negative integer (e.g., 0, 10, 255 etc)

`int32` => A 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890). This because `32bit` can represent `Math.Pow(2, 32)` unique values. They occupy `4 byte of memory space`

`rune` => An alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')

`uint32` => A 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295

`int64` => A 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807. This because `64bit` can represent `Math.Pow(2, 64)` unique values. They occupy `8 byte of memory space`

There also are more types like `int8` or `uint8` which work in the same way (=> integer with smaller number range)

##### Null Values

<em>All Go value types come with a so-called **`null value`** which is the value stored in a variable if no other value is explicitly set.</em>

For example, the following int variable would have a default value of 0 (because 0 is the null value of int, int32 etc):

var age int // age is 0
Here's a list of the null values for the different types:

`int => 0`

`float64 => 0.0`

`string => "" (i.e., an empty string)`

`bool => false`

##### **Summary of Differences between int and int32 and int64**

| Type    | Size                                                 | Range                                                                                                       | Platform-Dependent?          | Use Case                                                                          |
| ------- | ---------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ---------------------------- | --------------------------------------------------------------------------------- |
| `int`   | 32 bits on 32-bit systems, 64 bits on 64-bit systems | -2,147,483,648 to 2,147,483,647 (on 32-bit systems) or -9 quintillion to +9 quintillion (on 64-bit systems) | Yes, depends on architecture | General-purpose integer type, platform-dependent                                  |
| `int32` | 32 bits                                              | -2,147,483,648 to 2,147,483,647                                                                             | No                           | Fixed 32-bit integer, useful for smaller numbers or memory-sensitive applications |
| `int64` | 64 bits                                              | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807                                                     | No                           | Fixed 64-bit integer, useful for large numbers or high-precision calculations     |

##### fmt.scan() and it limitations

The fmt.Scan() function in Go is used to read input values from the standard input stream (stdin) and assign them to variables. However, it has a few limitations that you should be aware of:

**You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.**
