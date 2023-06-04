# GO

## GO variables
~~~
var card string = "Ace of Spades"
~~~
We are going to break through this line:
* var is word for variable -> it means we are gonna create a new variable
* card is name of the variable
* string label at the end tells Go compiler that the only strings are going to be assigned to this variable

Go is static type programming language, here are the basic Go types
<table>
    <tr>
        <th>Type</th>
        <th>Example</span></th>
    </tr>
    </tr>
        <th>bool</th>
        <th>true / false</th>
    </tr>
    </tr>
        <th>string</th>
        <th>"Hi!" / "Hows it going?"</th>
    </tr>
        </tr>
        <th>int</th>
        <th>0 / -10000 / 99999</th>
    </tr>
        </tr>
        <th>float64</th>
        <th>10.0000001 / 0.0000009 / -100.003</th>
    </tr>
</table>

Alternative way of making same variable form up (initialziation statement)
~~~
card := "Ace of Spades"
~~~
- We use this : colon sign at very first usage of variable -> when it is created/defined
- If we what to reassing the value for already create variable we can do it like this:
~~~
card = "Five of Diamonds"
~~~

Function with return types must have it mentioned after parenthesesis like this:
~~~
func helloString() string {
    return "hello"
}
~~~

## Go structures for collections

* They can be devided in two groups:
    1. Array - Fixed length list of things
    2. Slice - An array that can grow or shrink

Slices and arrays both must be defined as array of data type, so it must have identical data type for items (elements)

To declare a variable which is slice of strings:
~~~
cards := []string{}
~~~

<i> Go is not Object Oriented programming languange <br/>
So there is no idea about classes !!!
But there is OO approach within golang </i>

## Go OOP approach

- To extend basic (primitive type) we can create our own custom extension (type)
```
type deck []string
```
This line basically means: create new type deck which is equal/or has all properties and possiblities like the slice of string
