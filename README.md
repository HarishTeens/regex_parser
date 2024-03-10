## Dynamic regex parser 

A dynamic parser to conveniently parse any string based on a simpler regex.

### Example

lets say you need to parse a string like this: "Apple : 200 red; 50 green = 1000"
and you want to get a variable map like 
{
    "fruit": "Apple",
    "red": "200",
    "green": "50",
    "price": 1000
}

Now Instead of using split and reading elements in an array that would take multiple steps.
use the regex_parser.Parse() function

```
line := "Apple : 200 red; 50 green = 1000"
pattern := "%fruit : %red red; %green green = %price"

varMap, err := regex_parser.Parse(line, pattern)
```
`varMap` contains the desired map, simple as that :)