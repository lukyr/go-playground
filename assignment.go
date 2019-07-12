/*
As others have explained already, := is for both declaration, assignment, and also for redeclaration; and it guesses (infer) the variable's type automatically.

It's a short-hand form of:
*/

var foo int
foo = 32

// OR:
var foo int = 32

// OR:
var foo = 32


/*
Some rules

You can't use := out of funcs. It's because, out of any func, a statement should start with a keyword.
example:
*/

illegal := 42

func foo() {
  legal := 42
}

/*
You can't use them twice (in the same scope):
Because, := introduces "a new variable", hence using it twice does not redeclare a second variable, so it's illegal.
*/

legal := 42
legal := 42 // <-- error

/*
	However, you can use them twice in "multi-variable" declarations, if one of the variables is new:
*/

foo, bar  := someFunc()
foo, jazz := someFunc()  // <-- jazz is new
baz, foo  := someFunc()  // <-- baz is new

/*
You can use the short declaration to declare a variable in a newer scope even that variable is already declared with the same name before:
*/

var foo int = 34

func some() {
  // because foo here is scoped to some func
  foo := 42  // <-- legal
  foo = 314  // <-- legal
  /*
  Here, foo := 42 is legal, because, it declares foo in some() func's scope. foo = 314 is legal, because, it just assigns a new value to foo.
  */
}

/*
You can use them for multi-variable declarations and assignments:
*/
foo, bar   := 42, 314
jazz, bazz := 22, 7

/*
You can declare the same name in short statement blocks like: if, for, switch:
*/
foo := 42
if foo := someFunc(); foo == 314 {
  // foo is scoped to 314 here
  // ...
}
// foo is still 42 here
//Because, foo in if foo := ..., only belongs to that if clause and it's in a different scope


/*
So, as a general rule: If you want to easily declare a variable you can use :=, or, if you want to overwrite an existing value, you can use =.
*/




