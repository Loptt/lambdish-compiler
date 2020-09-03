# Lambdish Compiler
Lambdish is a compiled functional programming language inspired on other functional languages such as Racket and Haskell. It's main purpose is to provide a simple experiece of functional programming, while keeping complexity at a minimal. Thus, it is meant to be an introduction to the functional world for programmers of other paradigms.

## Features
Below are some of the main characteristics of the language.
* Statically typed: Lambdish is a statically typed language and thus types are well defined and checked at compile time. 
* Strongly typed: Operation type checking at compile time, and inmutability allows for easier programming .
* Simple data types:
  * Numbe (num): Any numberic value wether integer or float, positive or negative.
  * Character (char): Any alphanumeric element.
  * Boolean (bool): Classic boolean value, true or false.
* More complex data types: 
  * Lists: Lambdish supports lists of any of the forementioned types, as well as list types. Thus, allowing for nested lists of any level.
  * Functions: As a functional programming language, functions can be used as data types. The specific type of a function is defined by it's parameters and return type.
* Lambda functions: Complete support for anonymous functions and higher order functions.
* Single output: In a Lambdish file, functions are declared first, and at the end a single function call can be made at the end of the file. The return value of that function is printed to standard output when the program is executed.

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

## Main Contributors:
* Carlos Estrada
* Erick González
