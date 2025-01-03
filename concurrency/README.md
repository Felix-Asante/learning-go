#### CONCURRENCY

Concurrency is the ability of a program to run multiple tasks at the same time.

In Go, concurrency is achieved through **goroutines** and **channels**

**Goroutines**
A goroutine is a light-weight thread of execution that can be scheduled to run concurrently with other goroutines.

- They are lighweight thread managed by the `Go runtime` and not by the `OS`.
- Goroutines run in the smae address spaces,so access to shared memory must be synchronized.

**Channel**
A channel is a communication mechanism that allows goroutines to send and receive data.
