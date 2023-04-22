# battery

A Go module for handy tools including an iterator implementation, a basic map-set implementation, and builtin slice and map utility function sets. This module is still under heavy development and should not be used in production. Contributions are welcome.

## Installation

go1.18 or higher version required.

```bash
go get github.com/nnnewb/battery
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/nnnewb/battery/iter"
)

// The following code demonstrates how to use the `iter` package from the `battery` module.
// First, import the `iter` package from `github.com/nnnewb/battery/iter`.
// Then, create a new iterator with a slice of integers using `iter.Lift([]int{1, 2, 3, 4, 5})`.
// Finally, iterate over the iterator using `for it.Next()`, and print each value using `fmt.Println(it.Value())`.
// The output should be:
// 1
// 2
// 3
// 4
// 5

func main() {
    it := iter.Lift([]int{1, 2, 3, 4, 5})

    for it.Next() {
        fmt.Println(it.Value())
    }
}
```

### battery/iter

The iterator pattern is a design pattern that provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation. It involves defining an interface for accessing the elements of a collection and implementing it in a way that allows the elements to be accessed sequentially. We should use the iterator pattern because it provides a standardized way to iterate over collections, making it easier to write reusable code that can work with different types of collections. Additionally, it can help to simplify code by abstracting away the details of iterating over a collection, allowing the developer to focus on the higher-level logic of their program.

The iterator pattern and monad functional programming are both design patterns that provide a way to work with collections of data in a standardized and reusable way. The iterator pattern provides a way to access the elements of a collection sequentially without exposing its underlying representation, while monads provide a way to chain together operations on data in a way that is composable and easy to reason about. While they are not directly related, they can be used together to create powerful and flexible code that is easy to maintain and extend. For example, an iterator can be used to iterate over a collection of data and apply a monadic operation to each element, allowing for complex transformations to be performed on the data in a way that is easy to understand and modify.

### battery/set

`map-set` is a basic implementation of a set data structure using a Go map. It provides a way to store a collection of unique values and perform set operations such as union, intersection, and difference. We should use `map-set` when we need to store a collection of unique values and perform set operations efficiently. The use of a map allows for constant time complexity for adding, removing, and checking for the existence of an element in the set.

### battery/slices

TODO

### battery/maps

## Contributions

Contributions are welcome! Please follow these guidelines when contributing to the project:

1. Fork the repository and create a new branch for your feature or bug fix.
2. Ensure that your code adheres to the project's coding style and conventions.
3. Write tests for your code and ensure that all tests pass before submitting a pull request.
4. Keep your pull requests small and focused on a single feature or bug fix.
5. Provide a clear and detailed description of your changes in the pull request.
6. Be responsive to feedback and be willing to make changes to your code if requested.
7. Finally, be respectful and professional in all interactions with other contributors and maintainers.

For more information on contributing to open source projects, please refer to the [GitHub Guide to Contributing to Open Source](https://opensource.guide/how-to-contribute/).

## License

MIT License

Copyright (c) 2023 weak_ptr

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.