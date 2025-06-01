# Variable Declaration
## Option 1
Strictly declartion
```gotemplate
var <variableName> <type>
<variableName> = <value>
```

## Option 2
Loose declaration
```gotemplate
var <variableName> = <value>
```
In this way goLang compiler will automatically detect that which type of data type will be right for the assignment.

## Option 3 (Preferred)
```gotemplate
<variableName> := <value>
```

Note: `:=` means that we are declaring and defining the value together. `=` means assign the value it could be declaration level or after that.

## Types
### Primitves:
- string
- number
    - int   -> contains same size for 32/64-bit(based on our system.)
        - int8      -> 8-bit signed integer
        - int16     -> 16-bit signed integer
        - int32     -> 32-bit signed integer
        - int64     -> 64-bit signed integer
    - uint  -> contains same size for 32/64-bit
        - uint8     -> 8-bit unsigned integer
        - uint16    -> 16-bit unsigned integer
        - uint32    -> 32-bit unsigned integer
        - uint64    -> 64-bit unsigned integer
    - rune  ->  synonyms of int32, also represent the unicode code points or characters
    - byte  ->  synonym of unit8
    - uintpr    ->  It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
    - float
        - float32   -> 32-bit floating number
        - float64   -> 64-bit floating number
    - complex
        - complex64     -> 32bit real and imaginary
        - complex128    -> 32bit real and imaginary
- char
- bool
### Composite
Collection of primitives/composite data types are called as composite data types.

## Non-initialized/Default value
Default value differse for each data type.
|Data Type|Default Value|
|---------|-------------|
|String |""|
|Number|0|
|bool|false|

Example snippet:
```gotemplate
var message string
var isTrue bool
var number int

fmt.Println(message)    //  ""
fmt.Println(isTrue)     //  false
fmt.Println(number)     //  0
```
## Exceeding limits of data types:
When the value assigned to a data type exceeds its limit, the value wraps around due to overflow.
```gotemplate
var smallInt int8 = 100
fmt.Println(smallInt + 28) // -128

var smallInt int8 = 100
fmt.Println(smallInt + 29)  // -127

var smallInt int8 = -127
fmt.Println(smallInt - 1)  // -128

var smallInt int8 = -127
fmt.Println(smallInt - 2)  // 127
```