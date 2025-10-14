25-30

# Self-Introduction
1. Greeting
   * Arseniy + GoLand IDE developer + **static code analysis** => nil dereference

# Problem-introduction / Whay does it matter
## Version 1



Языки программирования => язык программирования go => nil pointers => nil pointer dereference



* A little bit history: One-Billion-Dollar
  * Nil value was invented in 1969
  * Они проросли в разные языки
  * Все программисты мира страдают от nil dereference exception
  * 2008 it was called => One-Billion-Dollar
  * [???] => Its a real issue today ... go-sepcific exapple

* Nil in Go (Мини-ликбез по nil значениям в go)
  * What nil can be, reference types (pointers, slices, maps, e.t.c.)
  * Safe/unsafe operation on nil values


## Version 2


* Why does it matter
   * Billion-dollar-mistake
     * Single concept ⇒ countless crashes, debugging, ...
     * https://www.reddit.com/r/golang/comments/18sncxt/go_nil_panic_and_the_billion_dollar_mistake/
   * It is not a history — real issue today


# Cool examples of nil dereference

## Default initialization
* Variables with no initialization
* Fields with no initialization

## Pointer to nil conversion
* Custom error => condition err != nil is always true

## Name shadowing
* result redeclaration, error redeclaration
* error redeclaration???

## Interprocedural bugs

* return nil as a result
* `return nil, nil`
* passing nil as a parameter

# Linters



# Probel statement / Why does it matter
* One-billion-dollar mistake
  * 










4. What the talk will cover
    * What nil can be [Mental Model]
    * Tricky examples of nil dereference
    * Tools to detect nil dereference

Introduction:
* никто не идеален, не все языки nil-safe ...
* nil-safe / not nil-safe languages
* (проверки на nil не всегда помогают)

(Что такое nil в целом и как можно обосраться в go в чатности)

* Nil-Safe / Nil-Unsafe
* Hook & what is it
    * `panic: runtime error: invalid memory address or nil pointer dereference`
    * nil dereference meaning ???


1. Зачем слушать? => навык


Какая реакция на выступление?

Дано: 














