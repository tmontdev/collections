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

## Working with Maps

> collections package provide the [IMap](https://godocs.io/github.com/tmontdev/collections/maps/#IMap) interface with helper methods to make your work with maps easier. The default implementation of [IMap](https://godocs.io/github.com/tmontdev/collections#Map) is the [Map](https://godocs.io/github.com/tmontdev/collectionsmaps/#Map) which is based on a built-in map type.

Let's make an example of Telephone Spelling Alphabet with Maps, and see how it works!

```go
import (
    "github.com/tmontdev/collections/maps"
)
// NATO Alphabet standard words is commonly used for accurately radio and telephone communication
// standardRadioAlphabet returns a maps.Map[rune, string]
func standardRadioAlphabet() maps.Map[rune, string] {
	// Note: returning a built-in map is accepted, as the maps.Map is a map type
	return map[rune]string{
		'A': "Alpha", 'B': "Bravo", 'C': "Charlie", 'D': "Delta",
		'E': "Echo", 'F': "Foxtrot", 'K': "Kilo", 'L': "Lima", 'M': "Mike",
		'N': "November", 'O': "Oscar", 'P': "Papa", 'Q': "Quebec", 'R': "Romeu",
		'S': "Sierra", 'T': "Tango", 'U': "Uniform", 'V': "Victor",
		'W': "Whiskey", 'X': "X-Ray", 'Y': "Yankee", 'Z': "Zulu",
	}
}

func main() {
	// built-in map handled here as maps.Map now have super-powers
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
func standardRadioNumbers() maps.Map[rune, string] {
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

func preferredRadioSpelling() maps.Map[rune, string] {
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
	// Note: Complement() method sets key/value pairs from the given map, into itself... on key conflict, keeps original value
	// Note: SetFrom() method sets key/value pairs from the given map, into itself... on key conflict, sets the new value from given map
	// Note: Clone() method create and returns a new map, identical to the original.
	// Note: In order to keep letterSpelling map unaltered, let's merge numberSpelling in its Clone
	standardSpelling := letterSpelling.Clone().Complement(numberSpelling)

	// now we are going to replace our preferred spelling words in the alphabet.
	// Note: In order to replace our preferred spelling words, let's use SetFrom() method instead of Complement()
	// Note: In order to keep standardSpelling map unaltered, let's merge preferredRadioSpelling() result in its Clone
	customSpelling := standardSpelling.Clone().SetFrom(preferredRadioSpelling())

	println(customSpelling.Get('X'))   // "X-men"
	println(standardSpelling.Get('X')) // "X-Ray"
	println(customSpelling.Get('9'))   // "Niner"
	println(standardSpelling.Get('9')) // "Nine"
	println(customSpelling.Get('?'))   // "What"
	println(standardSpelling.Get('?')) // empty string, as it was not mapped
}
```

Map interface have many other methods to make your work with maps easier, without giving up performance. [To know more about Maps, please refer to Map Godoc](https://godocs.io/github.com/tmontdev/collections/maps#IMap)

## Working with Lists

> We're gonna use the code of our previous section with maps, so keep it :D

> collections package provide the [IList](https://godocs.io/github.com/tmontdev/collections/lists#IList) interface with helper methods to make your work with arrays and slices easier. The default implementation of [IList](https://godocs.io/github.com/tmontdev/collections/lists#IList) is the [List](https://godocs.io/github.com/tmontdev/collections/lists/#List) which is based on a built-in slice pointer. (see [why do we prefer slice pointers](https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b#:~:text=The%20pointer%20to%20a%20slice,those%20who%20call%20the%20function.))
> We also provide the [SafeList](https://godocs.io/github.com/tmontdev/collections/lists#SafeList) implementation which is [thread-safe](https://en.wikipedia.org/wiki/Thread_safety).

Now that our spelling map is complete, make use of it! Let’s declare a list of words to spell, and see how [IList](https://godocs.io/github.com/tmontdev/collections/lists#IList) works.

```go
import (
    "github.com/tmontdev/collections/maps"
    "github.com/tmontdev/collections/lists"
)

// our official alphabet spelling
func customSpelling() maps.IMap[rune, string] {
	// uses old code here
	return standardRadioAlphabet().Merge(standardRadioNumbers(), false).Merge(preferredRadioSpelling(), true)
}

// returns a list of runes, which is the letters of the given word
func lettersFrom(word string) lists.IList[rune] {
	// creates a list from the slice of runes
	return lists.NewListFrom[rune]([]rune(word))
}

// returns a list of words, which is the spelling of the given word
func spell(word string) lists.IList[string] {
	spelling := lists.NewList[string]()
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
	words := lists.NewList[string]("foo", "bar", "berserk")
	// we may easily add, remove, and filter words
	// lets and some harder words to spell
	words.Push("bumfuzzle", "cattywampus", "Kakorrhaphiophobia")
	// I don't event know what the heck these words mean

	//See Reduce() docs for more information
	spelledWords := words.Reduce(func(acc any, word string, idx int) any {
		return acc.(maps.Map[string, lists.IList[string]]).Set(word, spell(word))
	}, maps.Map[string, lists.IList[string]]{})
	fmt.Printf("%v", spelledWords)
}
```

List interface have many other methods to easily filter, sort, and transform your slices, without giving up performance. [To know more about Lists, please refer to List Godoc](https://godocs.io/github.com/tmontdev/collections/lists#IList)

# Usage [![Go Documentation](https://godocs.io/github.com/tmontdev/collections?status.svg)](https://godocs.io/github.com/tmontdev/collections)

**To get more api usage instructions, see our [godoc](https://godocs.io/github.com/tmontdev/collections)**

# Issues and Discussions

Please, feel free to open an issues and discuss about this project on [github](https://github.com/tmontdev/collections)
