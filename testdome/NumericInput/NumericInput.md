
https://www.testdome.com/t

En cours de réflexion :)

User interface contains NumericInput control ,which accepts only digits.
Extend  NumericInput structure so that:
-it implement UserInput Interface,
- Add(rune) should add only decimal digits to the input. Other runes
should be ignored,
- Getvalues() should return the current input.

For examples the following code should return 10:
var input UserInput = &NumericInput{}
    input.Add('1')
    input.Add('a')
    input.Add('0')
    fmt.Println(input.GetValue())

