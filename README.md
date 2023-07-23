# Thunder

## Description

Thunder is a POC project for better understanding of [cobra-cli](https://github.com/spf13/cobra-cli). It sends HEAD requests to a target host, allowing traffic generation for general purpose (such as small load tests, for example). It is not intended to be used for large scale tests.

## Installation

Just clone the project and build it
```
git clone https://github.com/pvskp/thunder.git
cd thunder
go build
```

This will generate a `thunder` binarie. Don't forget to put it on your PATH to use it anywhere.

## Usage

`thunder -h` shows the help message with the available options.
