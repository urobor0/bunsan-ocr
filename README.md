# Bunsan-OCR

This project is an api rest to transform ocr codes to real numbers. The input data is a text file, it produces a file with a number of entries that look like this with a number of entries that look like this:

```
    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|
```

Each entry consists of 4 lines and each line is 27 characters long. The first 3 lines of each entry contain an account number written using pipes and underscores, and the fourth line is blank.
Each account number must have 9 digits, all of which must be in the range 0-9.

To realize the requirements, an asynchronous processing solution is proposed. 

## Sequence diagram

![](resources/sequence-diagram.png)

## Creating ocr job flow petition.

![](resources/create-ocr-job-flow-petition.png)

## How to Install and Run the Project

You need the following tools

* Go 1.16 or later.

1. Download all dependencies with ```go mod download```

This version of the project uses in-memory volatile persistence for ease of use so just having go installed works.

At this point if everything is ok has been correct you can run the project tests. with the following command.

```shell
go test ./...
```

Now you can run the project with the command:

```shell
go run cmd/api/main.go
```

## Api rest endpoint

```http request
POST http://localhost:8080/api/item
Content-Type: multipart/form-data; boundary=WebAppBoundary

--WebAppBoundary
Content-Disposition: form-data; name="file"; filename="input.txt"
Content-Type: text/plain

```