# Lambdish Compiler
Lambdish is a compiled functional programming language inspired on other functional languages such as Racket and Haskell. It's main purpose is to provide a simple experiece of functional programming, while keeping complexity at a minimal. Thus, it is meant to be an introduction to the functional world for programmers of other paradigms.

## Features
Below are some of the main characteristics of the language.
* Statically typed: Lambdish is a statically typed language and thus types are well defined and checked at compile time. 
* Strongly typed: Operation type checking at compile time, and inmutability allows for easier programming.
* Native operations: Lambdish provides a set of built in operations (functions) to provide arithmetic, relational, and logical operations, if-else operations, and other native operations.
* Simple data types:
  * Numbe (num): Any numberic value wether integer or float, positive or negative.
  * Character (char): Any alphanumeric element.
  * Boolean (bool): Classic boolean value, true or false.
* More complex data types: 
  * Lists: Lambdish supports lists of any of the forementioned types, as well as list types. Thus, allowing for nested lists of any level.
  * Functions: As a functional programming language, functions can be used as data types. The specific type of a function is defined by it's parameters and return type.
* Lambda functions: Complete support for anonymous functions and higher order functions.
* Single output: In a Lambdish file, functions are declared first, and at the end a single function call can be made at the end of the file. The return value of that function is printed to standard output when the program is executed.

### Operations
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

More examples can be found in the `examples` directory.

## Main Contributors:
* Carlos Estrada
* Erick González
