# Quotes

# What does this code do?

This code implements a back end that returns a JSON file containing a random inspirational quote and the respective author.

The quotes are stored in a sqllite database.

The following data is returned to the API consumer:

```
Quote
Author
```

Please note that the sample Quotes database file contains a list of quotes in Portuguese. You are welcome to adjust to your language as needed.


## Prerequisites

[Install Go](https://golang.org/doc/install)

## How does it work?

Build the code

```
go build
```

Run the program

```
./quotes
```

Basic test

```
curl localhost:8080/
```

Expected result

```
{"Quote":"Random quote.","Author":"Random Quote Author"}
```


## Author

[**Marcelo Pinheiro**](https://github.com/mpinheir)

## License

Copy and use as you wish.

## Thank you

* [Go Language creators](https://en.wikipedia.org/wiki/Go_(programming_language))
* To [Marco Paganini](https://github.com/marcopaganini) for his guidance

