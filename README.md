# ASCII Art Interpreter

A Go program that converts text into graphic representations using ASCII characters.

## 📋 Description

This project implements an ASCII art generator that takes a string as input and outputs it in stylized ASCII art format. The program supports multiple banner styles and handles various characters including letters, numbers, spaces, special characters, and newlines.

## ✨ Features

- 🎨 Converts text to ASCII art using different banner styles
- 🔤 Supports multiple characters: letters, numbers, spaces, and special characters
- ↵ Handles newlines (`\n`) in input
- 🏷️ Three built-in banner styles: `shadow`, `standard`, and `thinkertoy`
- 📏 Each character is 8 lines high with proper spacing
- 🚀 Written in Go with standard libraries only

## 📥 Installation

1. Ensure you have Go installed (version 1.16 or higher recommended)
2. Clone the repository or download the source files
3. Navigate to the project directory

```bash
git clone <repository-url>
cd ascii-art
```

## 🚀 Usage

### Basic Commands

```bash
# Run the program with text input
go run . "Hello"

# Pipe output with cat -e to see line endings
go run . "Hello" | cat -e

# Empty string
go run . ""

# Just a newline
go run . "\n"
```

### 📝 Examples

#### Simple Text
```bash
go run . "Hello" | cat -e
```
Output:
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

#### Text with Spaces
```bash
go run . "Hello There" | cat -e
```

#### Text with Numbers and Special Characters
```bash
go run . "1Hello 2There" | cat -e
```

#### Text with Brackets
```bash
go run . "{Hello There}" | cat -e
```

#### Multiple Lines
```bash
go run . "Hello\n\nThere" | cat -e
```

### 🎭 Added by us just for fun (if not available will return to standard) 🎭 (thinkertoy/standard/shadow)
```bash
go run . "Hello" "font_type" | cat -e
```

### 🎭 Banner Styles

The program includes three banner files in the `banner/` directory:

| Banner | Description | Example |
|--------|-------------|---------|
| `shadow.txt` | Shadow style with decorative elements | More artistic look |
| `standard.txt` | Standard ASCII art style | Clean and readable |
| `thinkertoy.txt` | Thinkertoy style | Unique, rounded characters |

To use a specific banner style, you'll need to modify the source code to load the desired banner file (implementation detail).

## 📊 Banner Format

- Each character has a fixed height of **8 lines**
- Characters are separated by newlines (`\n`)
- The banner files are pre-formatted and should not be modified

### Example Character Representations

**Space (' '):**
```
......
......
......
......
......
......
......
......
```

**Exclamation mark ('!'):**
```
._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....
```

**Double quote ('"'):**
```
._._..
(.|.).
.V.V..
......
......
......
......
......
```

## 📁 Project Structure

```
.
├── main.go                 # Main program entry point
├── banner/                 # Directory containing banner files
│   ├── shadow.txt         # Shadow style banner
│   ├── standard.txt       # Standard style banner
│   └── thinkertoy.txt     # Thinkertoy style banner
├── converter/              # Internal packages (if any)
│   └── converter            # ASCII art generation logic
├── go.mod                 # Go module file
├── README.md              # This file
└── tests/                 # Test files (recommended)
    └── main_test.go       # Unit tests
```

## 🔧 Technical Details

### Requirements
- **Go 1.22.3**
- Only **standard Go packages** are used (no external dependencies)

### Key Components
1. **File Reading**: Reads banner templates from files
2. **Text Parsing**: Processes input string character by character
3. **ASCII Assembly**: Builds the final ASCII art line by line
4. **Newline Handling**: Properly processes `\n` characters in input

###  Process Steps - ASCII Art Generator
---
1. **Command-Line Input Validation**
   - Check if at least one command-line argument is provided
   - If no arguments provided, exit program silently
   - Accept format: `program "text" [font]`

2. **Font Selection & Configuration**
   - Set default font to "standard"
   - If second argument provided:
     - Check if it matches "thinkertoy", "shadow", or "standard"
     - If valid, use specified font; otherwise, keep default
   - Construct font file path: `./banner/[font-name]`

3. **Text Preprocessing**
   - Replace all `\n` escape sequences with actual newline characters
   - Split input into individual lines using newline separator
   - Handle special case of empty first line

4. **Font File Loading**
   - Read font file from `banner/[font].txt`
   - Convert file content to string
   - Normalize line endings (replace `\r\n` with `\n`)
   - Split font file content into individual lines

5. **Character Validation**
   - Convert each input line to runes (Unicode characters)
   - Check each character's ASCII value
   - Accept only printable ASCII characters (32-126)
   - Return error if unsupported character found

6. **ASCII Art Character Mapping**
   - For each valid character:
     - Calculate index in font file: `(char - ' ') * 9 + 1`
     - Extract 8 lines of ASCII art from font file starting at calculated index
     - Store character's ASCII representation in 2D slice

7. **Horizontal Composition**
   - For each input line:
     - Process all characters in the line
     - For each of the 8 height levels:
       - Concatenate corresponding row from each character's ASCII art
       - Add newline after each complete row
8. **Multi-Line Processing**
   - Handle empty input lines by adding blank lines to output
   - Process each input line sequentially
   - Maintain proper line spacing between different input lines

9. **Output Generation**
   - Build final output string by concatenating all processed lines
   - Return complete ASCII art string
   - Print result to standard output in main function

10. **Error Handling**
    - Return empty string if:
      - Input is empty
      - Font file cannot be read
      - Unsupported characters found
    - Print error messages to console when appropriate

---
### Implementation Notes
- The program reads banner files from the file system
- Input is processed character by character
- Newlines in input create new lines in the output
- The program handles edge cases like empty strings and multiple consecutive newlines
- Characters are looked up in the banner file based on their ASCII position

## 🎯 Learning Objectives

This project helps you learn about:

- **Go file system (fs) API** - Reading and parsing files
- **String manipulation** - Processing and transforming text
- **Command-line interfaces** - Handling arguments and input
- **Data structures** - Efficient character mapping and lookup
- **Error handling** - Managing file errors and invalid input
- **Testing** - Writing unit tests for text processing


### Test Coverage Should Include:
- ✅ Empty string input
- ✅ Single character input
- ✅ Multi-word input
- ✅ Input with newlines
- ✅ Input with special characters
- ✅ Error cases (invalid banner files, etc.)
- ✅ Different banner styles


### Code Guidelines
- Follow Go standard formatting (`gofmt`)
- Add comments for complex logic
- Include tests for new features
- Update documentation as needed

## 📄 License

This project is for educational purposes as part of the ASCII Art curriculum.

## 🙏 Acknowledgments

- Project specifications provided by the educational curriculum
- Banner files provided as part of the project requirements
- Inspired by classic ASCII art generators

---

## 📚 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [ASCII Table Reference](https://www.asciitable.com/)
- [Go File System Package](https://pkg.go.dev/io/fs)

## 📞 Support

For issues or questions related to this implementation, please check:
1. The project documentation
2. Go standard library documentation
3. ASCII art format specifications

---