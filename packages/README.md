#### PACKAGES

**Packages are a way to group related code together in a single file**

- every package should have be in its own directory

**In other to use a an item from another package the package needs to import it. To export an item from a package the variable name of the item must start with a capital letter**

How to install GO packages
Go packages can be found in the [Go packages](https://pkg.go.dev) website.
And in order to install a package you need to run the following command
`go get <package name>`

You can also use the `go get` command to install and make available packages listed in the `go.mod` file.

In addition to the `go.mod` file, Go also uses the `go.sum` file to specify the checksums for the packages in the `go.mod` file.
