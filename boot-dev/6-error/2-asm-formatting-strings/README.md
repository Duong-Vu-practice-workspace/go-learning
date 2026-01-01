## Formatting Strings Review

A convenient way to format strings in Go is by using the standard library's [fmt.Sprintf()](https://pkg.go.dev/fmt#example-Sprintf) function. It's a string interpolation function, similar to Python's [f-strings](https://docs.python.org/3/reference/lexical_analysis.html#f-strings). The `%v` substring uses the type's default formatting, which is often what you want.

### Default Values

```
const name = "Kim"
const age = 22
s := fmt.Sprintf("%v is %v years old.", name, age)
// s = "Kim is 22 years old."
```

The equivalent Python code:

```python
name = "Kim"
age = 22
s = f"{name} is {age} years old."
# s = "Kim is 22 years old."
```

### Rounding Floats

```
s := fmt.Sprintf("I am %f years old", 10.523)
// s = I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
s := fmt.Sprintf("I am %.2f years old", 10.523)
// s = I am 10.52 years old
```

## Assignment

We need better error logs for our backend developers to help them debug their code.

Complete the `getSMSErrorString()` function. It should return a string with this format:

```
SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
```

- `COST` is the cost of the SMS, always showing the price formatted to 2 decimal places.
- `RECIPIENT` is the stringified representation of the recipient's phone number

*Be sure to include the $ symbol and the single quotes*

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Primeval 10x Developer, can assist... *for a price*.

<audio xmlns="http://www.w3.org/1999/xhtml"></audio>

package main



import (

"fmt"

)



func getSMSErrorString (cost float64,recipient string) string {

//?

}