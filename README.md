# My Gobuster

## Description

This Go program is a simplified version of Gobuster, focusing solely on brute-forcing URLs. It was developed to learn and practice the Go programming language, especially implementing goroutines and concurrent programming.

## Execution Environment
```
Guaranteed compatible systems: Linux, Windows, macOS  
Guaranteed compatible Go version: Go 1.23.2 or higher
```
## Usage

### 1. Compiling the Program

To compile the program, use the following command:
```
go build main.go
```

### 2. Running the Program

#### Running the compiled program:
```
./main
```
#### Running the program without compiling:
```
go run main.go
```

### 3. Displaying Help (-h)

To display the help message:
```
./main -h
```
Output:
```
Usage of mygb: -d string Path to dictionary file (required) -q Quiet mode, only displays HTTP 200 responses -t string Target URL to test (required) -w int Number of workers to run concurrently (default: 1)
```

### 4. Basic Usage

Basic usage allows sending requests from a given list to a target URL.
```
./main -t <Target URL> -d <Path to dictionary>
```

Example:
```
go run main.go -t https://example.com -d wordlist.txt
```

Example output:
```
(^) ---
Target: https://example.com/List: wordlist.txt
Workers: 1---
(^) Starting scan...
(^) /test 404
/admin 200/secret 403 (^)
/login 200
Scan done.
```

## Flags

### 1. `-d` string (Required)
- **Description**: Path to the dictionary file that contains a list of potential URL paths to test.
- **Usage**: 
```
-d <Path to dictionary>
```
- **Example**: 
```
-d /path/to/wordlist.txt
```

### 2. `-q` (Optional)
- **Description**: Quiet mode. If enabled, the program will only display HTTP 200 responses (successful requests).
- **Usage**: 
```
-q
```
- **Example**: 
```
-q
```

### 3. `-t` string (Required)
- **Description**: The target URL to enumerate (e.g., `https://example.com`).
- **Usage**: 
```
-t <Target URL>
```
- **Example**: 
```
-t https://example.com
```
### 4. `-w` int (Optional)
- **Description**: The number of concurrent workers (goroutines) to run. The default value is 1.
- **Usage**: 
```
-w <Number of workers>
```
- **Example**: 
```
-w 10
```
**Note:**  
A higher number of workers can speed up the scan, but may cause resource saturation or errors on the system.

## Development and Structure

The program is structured for readability and scalability:


.├── main.go # Main program (^)

The code is commented in English for better understanding and extensibility.

## Dependencies

The program has no external dependencies. Here are the used imports:

- "bufio"
- "flag"
- "fmt"
- "log"
- "net/http"
- "net/url"
- "os"
- "sync"

## Author

Project developed by **Erwann IROUCHE**, inspired by Gobuster as part of learning the Go programming language.
