[![Go Documentation](https://godocs.io/github.com/tmontdev/collections?status.svg)](https://godocs.io/github.com/tmontdev/collections)
[![Go Report Card](https://goreportcard.com/badge/github.com/tmontdev/collections)](https://goreportcard.com/report/github.com/tmontdev/collections)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=tmontdev_iterable&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=tmontdev_iterable)
[![Sourcegraph](https://sourcegraph.com/github.com/tmontdev/collections/-/badge.svg)](https://sourcegraph.com/github.com/tmontdev/collections?badge)
![visitors](https://visitor-badge.laobi.icu/badge?page_id=tmontdev.collections)
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://github.com/tmontdev/collections/blob/main/LICENSE)

# Collections

collections provide interfaces for easy handling of data collections (such as maps and lists). It was inspired by modern languages such as Dart, Javascript and C#, which have a more convenient and simpler way of doing this.

# Installation

collections is a modern golang package, and requires [golang programing language](https://go.dev/doc/install) 1.20 or above

You may install collections by running:

```bash
go get github.com/tmontdev/collections
```

# Project Status

**experimental/alpha**.
It was created for personal use, and is still growing. We are not in v1 yet.
However there is intention to keep improving, so feel free to suggest, discuss and issue.

# Quick Start

Once installed, simply import it to your code with

```go
import "github.com/tmontdev/collections"
```

## Working with Maps

> collections package provide the [Map](https://godocs.io/github.com/tmontdev/collections#Map) interface with helper methods to make your work with maps easier. The default implementation of [Map](https://godocs.io/github.com/tmontdev/collections#Map) is the [HashMap](https://godocs.io/github.com/tmontdev/collections#HashMap) which is based on a built-in map type.

Let's make an example of Telephone Spelling Alphabet with Maps, and see how it works!

```go
// NATO Alphabet standard words is commonly used for accurately radio and telephone communication
// standardRadioAlphabet returns a collections.HashMap[rune, string]
func standardRadioAlphabet() collections.HashMap[rune, string] {
	// Note: returning a built-in map is accepted, as the collections.HashMap is a map type
	return map[rune]string{
		'A': "Alpha", 'B': "Bravo", 'C': "Charlie", 'D': "Delta",
		'E': "Echo", 'F': "Foxtrot", 'K': "Kilo", 'L': "Lima", 'M': "Mike",
		'N': "November", 'O': "Oscar", 'P': "Papa", 'Q': "Quebec", 'R': "Romeu",
		'S': "Sierra", 'T': "Tango", 'U': "Uniform", 'V': "Victor",
		'W': "Whiskey", 'X': "X-Ray", 'Y': "Yankee", 'Z': "Zulu",
	}
}

func main() {
	// built-in map handled here as collections.HashMap now have super-powers
	// use .Where() to get a new Map with all the key/value pairs which satisfies your predicate
	longSpells := standardRadioAlphabet().Where(func(letter rune, word string) bool {
		return len(word) > 4 // here we want all spelling with length greater than 4
	})
	fmt.Printf("%v", longSpells)
	// note that runes are referenced by numbers
	// outputs only values with length greater than 4
	// map[65(A):Alpha 66(B):Bravo 67(C):Charlie 68(D):Delta 70(F):Foxtrot 78(N):November 79(O):Oscar 81(Q):Quebec 82(R):Romeu 83(S):Sierra 84(T):Tango 85(U):Uniform 86(V):Victor 87(W):Whiskey 88(X):X-Ray 89(Y):Yankee]
}
```

Ok! Now we may add number spelling too. And maybe we prefer to use custom spelling in our phonetic alphabet…

So let's add some code to our file!

```go
func standardRadioNumbers() collections.HashMap[rune, string] {
	return map[rune]string{
		'1': "One",
		'2': "Two",
		'3': "Three",
		'4': "Four",
		'5': "Five",
		'6': "Six",
		'7': "Seven",
		'8': "Eight",
		'9': "Nine",
		'0': "Zero",
	}
}

func preferredRadioSpelling() collections.HashMap[rune, string] {
	return map[rune]string{
		'9': "Niner", // commonly spelled niner, cause nine and five are easily confused
		'X': "X-Men", // because you may work at Marvel comics now lol
		'A': "Apple",
		'B': "Bingo",
		'?': "What", // not spelled yet on the other maps... lol
	}
}

// now lets mix it up!!
func main() {
	letterSpelling := standardRadioAlphabet()
	numberSpelling := standardRadioNumbers()
	// let's get words and letters together in a single map
	// Note: Merge() method receives a second parameter, the "replace" here passed as false, which means that we don't want to replace values, in case of conflicted keys.
	// Note: Merge() method alters the reference map. If you don't want your original map to be altered, use Clone() before it.
	// Note: Clone() method create and returns a new map, identical to the original.
	// Note: In order to keep letterSpelling map unaltered, let's merge numberSpelling in its Clone
	standardSpelling := letterSpelling.Clone().Merge(numberSpelling, false)

	// now we are going to replace our preferred spelling words in the alphabet.
	// Note: In order to replace our preferred spelling words, let's pass "replace" parameter as true
	// Note: In order to keep standardSpelling map unaltered, let's merge preferredRadioSpelling() result in its Clone
	customSpelling := standardSpelling.Clone().Merge(preferredRadioSpelling(), true)

	println(customSpelling.Get('X'))   // "X-men"
	println(standardSpelling.Get('X')) // "X-Ray"
	println(customSpelling.Get('9'))   // "Niner"
	println(standardSpelling.Get('9')) // "Nine"
	println(customSpelling.Get('?'))   // "What"
	println(standardSpelling.Get('?')) // empty string, as it was not mapped
}
```

Map interface have many other methods to make your work with maps easier, without giving up performance. [To know more about Maps, please refer to Map Godoc](https://godocs.io/github.com/tmontdev/collections#Map)

## Working with Lists

> We're gonna use the code of our previous section with maps, so keep it :D

> collections package provide the [List](https://godocs.io/github.com/tmontdev/collections#List) interface with helper methods to make your work with arrays and slices easier. The default implementation of [List](https://godocs.io/github.com/tmontdev/collections#List) is the [SimpleList](https://godocs.io/github.com/tmontdev/collections#SimpleList) which is based on a built-in slice pointer. (see [why do we prefer slice pointers](https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b#:~:text=The%20pointer%20to%20a%20slice,those%20who%20call%20the%20function.))
> We also provide the [SafeList](https://godocs.io/github.com/tmontdev/collections#SafeList) implementation which is [thread-safe](https://en.wikipedia.org/wiki/Thread_safety).

Now that our spelling map is complete, make use of it! Let’s declare a list of words to spell, and see how [List](https://godocs.io/github.com/tmontdev/collections#List) works.

```go
// our official alphabet spelling
func customSpelling() collections.Map[rune, string] {
	// uses old code here
	return standardRadioAlphabet().Merge(standardRadioNumbers(), false).Merge(preferredRadioSpelling(), true)
}

// returns a list of runes, which is the letters of the given word
func lettersFrom(word string) collections.List[rune] {
	// creates a list from the slice of runes
	return collections.NewListFrom[rune]([]rune(word))
}

// returns a list of words, which is the spelling of the given word
func spell(word string) collections.List[string] {
	spelling := collections.NewList[string]()
	spellingReference := customSpelling()
	letters := lettersFrom(word)
	// letters is a List of runes (*[]rune).
	// Note: Elements() method return a built-in slice value, with all elements of the list
	// Note: See Map(), and Reduce() List methods too.
	for _, letter := range letters.Elements() {
		// for each letter, get the spelling reference, and add in the spelling word list
		spelling.Push(spellingReference.Get(unicode.ToUpper(letter)))
	}
	return spelling
}

func main() {
	words := collections.NewList[string]("foo", "bar", "berserk")
	// we may easily add, remove, and filter words
	// lets and some harder words to spell
	words.Push("bumfuzzle", "cattywampus", "Kakorrhaphiophobia")
	// I don't event know what the heck these words mean

	//See Reduce() docs for more information
	spelledWords := words.Reduce(func(acc any, word string, idx int) any {
		return acc.(collections.HashMap[string, collections.List[string]]).Set(word, spell(word))
	}, collections.HashMap[string, collections.List[string]]{})
	fmt.Printf("%v", spelledWords)
}
```

List interface have many other methods to easily filter, sort, and transform your slices, without giving up performance. [To know more about Lists, please refer to List Godoc](https://godocs.io/github.com/tmontdev/collections#Map)

# Usage [![Go Documentation](https://godocs.io/github.com/tmontdev/collections?status.svg)](https://godocs.io/github.com/tmontdev/collections)

**To get more api usage instructions, see our [godoc](https://godocs.io/github.com/tmontdev/collections)**

# Issues and Discussions

Please, feel free to open an issues and discuss about this project on [github](https://github.com/tmontdev/collections)
