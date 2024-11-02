# Inpatient Order System
## Description
TODO

## Running the project
### using Go
```bash
go run .
```
### using Air (for live reloading)
```bash
air
```

## Building the project
### for Linux
```bash
go build -o inpatient-order-system
```
### for Windows
```bash
GOOS=windows GOARCH=amd64 go build -o inpatient-order-system.exe
```

## Tools
### Air
- live reloading, saved changes in the code will apply after running the project 
- docs: https://github.com/air-verse/air?tab=readme-ov-file#installation
```bash
go install github.com/air-verse/air@latest
```