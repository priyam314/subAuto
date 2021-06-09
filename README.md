# subAuto

## Description

this project is written in golang which generates a folder inside which there are sub-folders each named after predefined programming languages  name, inside each sub-folder is file named after main folder with that particular programming language extension.

## build
`$ go build src/subAuto.go`

## Install
`$ go install src/subAuto.go`

## Usage
`$ subAuto hello_world`
```
hello_world/
├── c
│   └── hello_world.c
├── c++
│   └── hello_world.cpp
├── dart
│   └── hello_world.dart
├── dlang
│   └── hello_world.d
├── golang
│   └── hello_world.go
├── java
│   └── hello_world.java
├── javascript
│   └── hello_world.js
├── julia
│   └── hello_world.jl
├── python
│   └── hello_world.py
├── R
│   └── hello_world.r
├── rust
│   └── hello_world.rs
└── scala
    └── hello_world.scala
```
