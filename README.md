# quiz-game-in-golang

## General info
Simple quiz game created with Golang.
* The CSV file should default to quiz_input.csv & the default time limit should be 20 seconds, but the user should be able to customize the filename & default time limit via a flag.

## Technologies
Project is created with:
* go 1.13

## Setup
To run this project
```
go build && ./quiz-game-in-golang -csv="quiz_input.csv" -limit=20
```