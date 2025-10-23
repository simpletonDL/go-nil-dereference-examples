# Spot the Nil Dereference
A collection of subtle Go bugs involving nil dereference from the talk
“Spot the Nil Dereference: How to Avoid a Billion-Dollar Mistake.”

## About
This repository contains real-world examples and minimal reproductions of tricky nil-related issues in Go —
cases where the language’s semantics make it easy to write code that looks correct but panics at runtime.
Each example is designed to illustrate a specific kind of mistake or misconception.

## Topics Covered
* Var shadowing 
* Interface vs pointer confusion
* Interprocedural nil dereference
* Nil basics

## Talk
For the full context and explanation, see the talk:
“Spot the Nil Dereference: How to Avoid a Billion-Dollar Mistake.”
(link coming soon)

## Requirements
Go 1.21+;
Run examples directly:
```shell
go run 0_default_init/main.go
```
