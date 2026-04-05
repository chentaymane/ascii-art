# ASCII Art Generator

A Go program that converts text into graphic representations using ASCII characters and supports writing the result into a file using a command-line flag.

---

## 📋 Description

This project implements an ASCII art generator that takes a string as input and outputs it in stylized ASCII art format.  
The program supports multiple banner styles and handles letters, numbers, spaces, special characters, and newlines.

An optional flag `--output=<fileName.txt>` allows saving the generated ASCII art into a file instead of printing it to standard output.

---

## ✨ Features

- Converts text into ASCII art
- Supports banner styles: standard, shadow, thinkertoy
- Handles printable ASCII characters (32–126)
- Supports newlines (`\n`) in input
- Writes output to a file using `--output=<fileName.txt>`
- Written in Go using only standard libraries

---

## 📥 Installation

1. Install Go (version 1.16 or higher recommended)
2. Clone the repository
3. Navigate to the project directory

git clone <repository-url>
cd ascii-art


---

## 🚀 Usage

### General Format

go run . [OPTION] [STRING] [BANNER]


If the option format is invalid, the program must print:

Usage: go run . [OPTION] [STRING] [BANNER]


---

### Supported OPTION

- `--output=<fileName.txt>`  
  Writes the ASCII art result into the specified file.

The option must follow **exactly** this format.  
Any other format is considered invalid.

---

## ✅ Examples

### Shadow banner example
```tx

student$ go run . --output=banner.txt "hello" standard
student$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
The program must still work without the output flag.

go run . "Hello"

go run . "Hello" thinkertoy


---
student$ go run . --output=banner.txt 'Hello There!' shadow
student$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
student$

## 🧩 Running Without OPTION

The program must still work without the output flag.

go run . "Hello"

go run . "Hello" thinkertoy


---

## ❌ Invalid Usage Examples

The following commands must return the usage message:

go run . --output banner.txt "hello"
go run . --output:banner.txt "hello"
go run . --out=banner.txt "hello"


Output:

Usage: go run . [OPTION] [STRING] [BANNER]


---

## 🎭 Banner Styles

Banner files are located in the `banner/` directory:

- standard.txt
- shadow.txt
- thinkertoy.txt

If an invalid banner name is provided, the program defaults to `standard`.

---

## 📁 Project Structure

.
├── main.go
├── banner/
│ ├── standard.txt
│ ├── shadow.txt
│ └── thinkertoy.txt
├── go.mod
├── README.md
└── tests/


---

## 🧪 Testing

Recommended test coverage includes:

- Empty input
- Single character
- Multiple words
- Newlines
- Special characters
- Invalid option formats
- File output correctness
- Different banner styles

---

## 📌 Technical Requirements

- Written in Go
- Uses only standard Go packages
- Follows Go best practices
- Unit tests are recommended

---


