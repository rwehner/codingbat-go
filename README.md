# Codingbat for Golang

I found Nick Parlante's [codingbat.com](http://codingbat.com/) programming exercise site useful and enjoyable when I was learning Python. I wanted something similar to practice ["little code"](http://codingbat.com/about.html) as I learn [Golang](http://golang.org). Since there is no official support for Go at codingbat.com, I created this repo as a port of the Codingbat Python exercises to Go. 

I'm very grateful to Nick for having created Codingbat. **But this is not an official Codingbat project**. I've done my best to port the exercises and unit tests accurately. However, it is a port and I'm just learning Go so there could be errors and non-idiomatic code. These problems were almost certainly introduced by me, not CodingBat. If you find better ways to do things, feel free to submit a PR. 

## How to use

The repo is organized with each sub-directory representing the problems in a specific CodingBat category. Each directory has a `<catagory>.go` file with stub functions you need to finish coding to pass the unit tests in the `<category>_test.go` file in the same directory.

1. [Install Go](https://golang.org/doc/install) in your local environment in whatever way you want.
1. Clone the repo to your local environment: `git clone https://github.com/rwehner/codingbat-go.git`
1. Choose a directory and start coding solutions in its `<category>.go` file.
1. Validate your solutions by running `go test -v` in the sub-directory.
