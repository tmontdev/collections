[![Go Documentation](https://godocs.io/github.com/tmontdev/iterable?status.svg)](https://godocs.io/github.com/tmontdev/iterable)
[![Go Report Card](https://goreportcard.com/badge/github.com/tmontdev/iterable)](https://goreportcard.com/report/github.com/tmontdev/iterable)

<!-- TOC -->

- [Iterable](#iterable)
  - [Project Status](#project-status)
  - [Install](#install)
  - [Usage ](#usage-)
  - [Release Notes](#release-notes)
    - [v0.0.1](#v001)
    - [v0.0.2](#v002)
    - [v0.0.3](#v003)
    - [v0.0.4](#v004)

<!-- TOC -->

# Iterable

**Iterable** package provides a simple interface to handle data collections in golang.

## Project Status

**Currently, in Alpha**

## Install

```shell
go get github.com/tmontdev/iterable
```

## Usage [![Go Documentation](https://godocs.io/github.com/tmontdev/iterable?status.svg)](https://godocs.io/github.com/tmontdev/iterable)
**To get get usage instructions, see our [godoc](https://godocs.io/github.com/tmontdev/iterable)**

## Release Notes

### v0.0.1

**Iterable** interface with the primary methods: **Length, IsEmpty, IsNotEmpty, At, ElementAt, Elements**.
**List** implementation

### v0.0.2
Added new methods: **Push, Clone, FirstElement, First, LastElement, Last, FirstIndexWhere, LastIndexWhere, IndexWhere**

### v0.0.3
Added new methods: **Where**

### v0.0.4
Added new methods: **Map**

### v0.1.0
Added new methods: **Reduce, Every, Some, None, Pop, Shift**
String representation