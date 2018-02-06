# Telosys-cmd

## Introduction

This CLI is a simple implementation of the [Telosys CLI]("http://www.telosys.org/cli.html") in Go.
Only the basic functionalities are implemented. 

## Requirements

To work on this project you need to install [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)

```bash
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

## Usage

First build the project. To simplify we'll call the file tcmd.exe .

```bash
go build -o tcmd.exe
```

Then in your telosys project folder you can use the basic functions

```bash
tcmd init 
tcmd nm model1
tcmd ne ent1
tcmd b install
```
