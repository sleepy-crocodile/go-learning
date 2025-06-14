# Rececing input from terminal

For receiving the input from the terminal we will use `bufio` package which provide the Reader. Let's create the Reader using Stdin.

```go
reader := bufio.NewReader(os.stdin)
```

`reader` has the function called `ReadString(delimeter)`, let's read the line (delimeter='\n')
```go
line, err := reader.ReadString('\n')
```
This return the line read from the terminal or error if occured.

Note: This not just read the lines from the terminal, it can read the lines from the file as well.

example: 
```shell
cat sample-io.txt | go run main.go
```

this will take each line from the sample-io.txt as input. Make sure that last line is using \n as delimeter otherwise error will be thrown.

### String to Number
`strconv` is the package which can be used to convert the string into number but they have be the numbers in string format. Any other string can not be converted into the number.

Note: In Windows, program can throw the error for delimeter so every time we read the string make sure that line is trimmed.

```shell
cat stringToNumber.txt | go run main.go
```