# alphanum
Library for sorting lists of strings with the [alphanum algorithm](http://davekoelle.com/alphanum.html) in the Go programming language.

## Background
The alphanum algorithm is designed to sort a list of strings in a natural order. Numbers in the list are ordered in the way that a human might think about them, rather than a pure alphabetical comparison. For example, a standard lexicographic sort might order a list like so:
```
File 1
File 10
File 11
File 12
File 2
File 3
File 4
File 5
```

While alphanum would return:
```
File 1
File 2
File 3
File 4
File 5
File 10
File 11
File 12
```

For more details, see the [alphanum webpage](http://davekoelle.com/alphanum.html),

## Usage
### sort.Interface
This package exports a `StringSlice` type, which implements `sort.Interface`. This means if you want to sort a `[]string`, you can do something like:
```go
data := []string{
	"20X Radonius Prime",
	"Allegia 50 Clasteron",
	"20X Radonius",
	"Xiph Xlater 3",
	"Xiph Xlater 1",
	"Xiph Xlater 58",
	"Xiph Xlater 12",
	"Xiph Xlater 10",
	"Xiph Xlater 2",
	"Callisto Morphamax 6000 SE",
	"Xiph Xlater 11",
}
sort.Sort(alphanum.StringSlice(data))
fmt.Printf("%#v\n", data) // => []string{"20X Radonius", "20X Radonius Prime", "Allegia 50 Clasteron", "Callisto Morphamax 6000 SE", "Xiph Xlater 1", "Xiph Xlater 2", "Xiph Xlater 3", "Xiph Xlater 10", "Xiph Xlater 11", "Xiph Xlater 12", "Xiph Xlater 58"}
```

### Less
If you're writing your own sort code, the package exports a `Less` function, which takes two strings and returns `true` if the first string should sorted before the second one.

For example:
```go
alphanum.Less("File 2", "File 10") // => true
alphanum.Less("File 17", "File 10") // => false
alphanum.Less("Allegia 50 Clasteron", "Allegia 500 Clasteron") // => true
alphanum.Less("Callisto Morphamax 6000 SE2", "Callisto Morphamax 6000 SE") // => false
```