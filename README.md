# Lambdish Compiler
Lambdish is a compiled functional programming language inspired on other functional languages such as Racket and Haskell. It's main purpose is to provide a simple experiece of functional programming, while keeping complexity at a minimal. Thus, it is meant to be an introduction to the functional world for programmers of other paradigms.

## Usage
This section describes how to run your own lambdish programs.

### Prerequisites
You must have Go installed and configured in your system. For a guide on how to do this please follow this [guide](https://golang.org/doc/install).

### Configure your PATH and GOPATH
The GOPATH is the directory where Go installs binaries by default. Add this to your regular PATH variable to access Go installed binaries from anywhere. First check the value of your GOPATH by running the following command.
```sh
$ go env
```
Look for the line containing GOPATH, and the value to the right is what we are looking for. In linux systems, this path is usually something like $HOME/go/bin. Now, to add this value to your PATH, in linux systems you can modify the $HOME/.profile file and add a line at the end like the following.
```sh
$ export PATH=$PATH:$HOME/go/bin
```
For how to configure this in MacOS you can check this [tutorial](https://www.architectryan.com/2012/10/02/add-to-the-path-on-mac-os-x-mountain-lion/) and this [tutorial](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/) for Windows 10.

### Install the commands

Install the compiler command.
```sh
$ go get github.com/Loptt/lambdish-compiler/cmd/clamb
```

Install the execution command.
```sh
$ go get github.com/Loptt/lambdish-compiler/cmd/rlamb
```

`clamb` takes a lambdish source file (`lsh`) and transforms it into object code (`obj`).
`rlamb` then takes an object code, executes it and prints its output.

### Compile lambdish
Run the following command to compile a lambdish source file.
```sh
$ clamb test.lsh
```
This will generate test.obj upon successful compilation.

### Run lambdish
Run the following command to execute a lambdish obj file and see the output.
```
$ rlamb test.obj
```

## Features
Below are some of the main characteristics of the language.
* Single output: In a Lambdish file, functions are declared first, and at the end a single function call can be made at the end of the file. The return value of that function is printed to standard output when the program is executed.
* Statically typed: Lambdish is a statically typed language and thus types are well defined and checked at compile time. 
* Strongly typed: Operation type checking at compile time, and inmutability allows for easier programming.
* Native operations: Lambdish provides a set of built in operations (functions) to provide arithmetic, relational, and logical operations, if-else operations, and other native operations.
* Simple data types:
  * Number (num): Any numeric value wether integer or float, positive or negative.
  * Character (char): Any alphanumeric element.
  * Boolean (bool): Classic boolean value, true or false.
* More complex data types: 
  * Lists: Lambdish supports lists of any of the forementioned types, as well as list types. Thus, allowing for nested lists of any level.
  * Functions: As a functional programming language, functions can be used as data types. The specific type of a function is defined by it's parameters and return type.
* Lambda functions: Complete support for anonymous functions and higher order functions.

## Operations
* `+` Add
* `-` Subtract
* `*` Multiply
* `/` Divide
* `%` Modulus
* `<` Less than
* `>` Greaer than
* `equal` Equals
* `and` And
* `or` Or
* `!` Not
* `head` Get first element of list
* `tail` Get list minus first element
* `insert` Insert element into beginning of list
* `append` Joins two lists into one
* `empty` Checks whether a list is empty
* `if` Performs if-else operation

### Data Types
* `num` Number data type, represents any real number. (integer and float).
* `char` Character data type. Currently only supports alphanumeric characters and space.
* `bool` Boolean data type, can be either `true` or `false`.
* `[TYPE]` List data type, represents a list of the given type, `TYPE` can be another list, allowing for n nested lists.
* `(TYPE1 => TYPE2)` A function data type that takes TYPE1 parameters and returns a TYPE2. Note that TYPE can also be a function data type.
* `(TYPE1, TYPEN => TYPE2)` A function data type that takes TYPE1 and TYPEN as arguments, and return a TYPE2, showing how multiple parameter functions data types are represented.

### Lambdas
A lambda function is declared with the following syntax.
```
(# TYPEN x => TYPE2 (BODY))
```
Where TYPEN represents the types for the parameters of the functions (which can be more than 1, separated by commas) and TYPE2 represents the return value of the lambda function. For an example of a lambda function, check the filter example below.

For a more details in the Lambdish syntax, you can consult the grammar definition for the language in [grammar/lambdish.ebnf](https://github.com/Loptt/lambdish-compiler/blob/master/grammar/lambdish.ebnf), which describes the grammar in EBNF syntax, as well as the tokens.

## Examples
A program that sums two numbers.
```
func sum :: num x, num y => num (
    +(x, y)
)

sum(3, 4) // Outputs 7
```
Note that the syntax for using any operator is also as a function.

A program that sums a list of numbers.
```
func sum :: [num] l => num (
    if( empty(l) , 
        0,
        +(head(l), sum(tail(l)))
    )
)

sum([1, 2, 3]) // Outputs 6
```
Note that an if statement also works as a function, taking the condition as first argument, the true case as second, and the false case as third.

A program that counts the amount of even numbers in a list
```
func countEven :: [num] l => num (
    if(empty(l), 
        0,
        if( 
            equal(%(head(l),2) , 0),
            +(countEven(tail(l)), 1),
            countEven(tail(l))
        )
    )
)

countEven([4, 3, 4, 6, 5, 4])
```
A program that implements a filter higher order function
```
func filter :: [num] l, (num => bool) f => [num] (
    if (empty(l),
        [num],
        if (f(head(l)),
            insert(
                head(l),
                filter(tail(l), f)
            ),
            filter(tail(l), f)
        )
    )
)

filter([1,2,3,4,5,6], (# num x => bool (
    equal(x, 7)
)))
```

A program that implements a map higher order function
```
func map :: [num] l, (num => num) f => [num] (
    if (empty(l),
        [num],
        insert(
            f(head(l)),
            map(tail(l), f)
        )
    )
)

map([1,2,3,4], (# num x => num (/(x, 2))))
```

A more complete example of a merge sort implementation
```
func length :: [num] l => num (
    if (empty(l),
        0,
        +(1, length(tail(l)))
    )
)

func take :: num n, [num] l => [num] (
    if (empty(l),
        l,
        if (or(equal(n, 0), <(n, 0)),
            [num],
            insert(head(l), take(-(n, 1), tail(l)))
        )
    )
)

func drop :: num n, [num] l => [num] (
    if (or(equal(n, 0), <(n, 0)),
        l,
        drop(-(n, 1), tail(l))
    )
)


func mergeSortAux :: [num] xl, [num] yl => [num] (
    if (empty(yl),
        xl,
        if (empty(xl),
            yl,
            if (<(head(xl), head(yl)),
                insert(
                    head(xl),
                    mergeSortAux(tail(xl), yl)
                ),
                insert(
                    head(yl),
                    mergeSortAux(xl, tail(yl))
                )
            )
        )
    )
)

func mergeSort :: [num] l => [num] (
    if (empty(l),
        [num],
        if (equal(length(l), 1),
            [head(l)],
            mergeSortAux(
                mergeSort(
                    take(/(length(l), 2), l)
                ),
                mergeSort(
                    drop(/(length(l), 2), l)
                )
            )
        )
    )
)

mergeSort([5,3,2,3,1])
```

More examples can be found in the [examples](https://github.com/Loptt/lambdish-compiler/blob/master/examples) directory.

## Video Tutorial

Small video tutorial in spanish of a sample program can be found [here](https://drive.google.com/file/d/1Hr8IU5pF8xBGNyOdQBs8KrM68Zn82JsP/view?usp=sharing) (English version coming soon).

## Main Contributors:
* [Carlos Estrada](https://github.com/Loptt)
* [Erick González](https://github.com/vsapiens)
