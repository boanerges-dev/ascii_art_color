# ASCII Art Colorizer

A lightweight command-line application written in Go that converts text into ASCII art banners and applies ANSI color formatting to either the entire output or specific substrings.

This project reads characters from a standard ASCII font file (`standard.txt`) and renders them as multi-line banner text directly in the terminal.

## Features

* Generate ASCII art from text input
* Apply color to the entire banner output
* Highlight specific substrings within the generated ASCII art
* Support for multi-line input using `\n`
* Multiple ANSI color options
* Simple and lightweight CLI interface
* Built with pure Go (no external dependencies)

---

## Supported Colors

| Color   |
| ------- |
| black   |
| red     |
| green   |
| yellow  |
| blue    |
| magenta |
| cyan    |
| white   |
| orange  |

If an unsupported color is provided, the application automatically falls back to **red**.

---

## Project Structure

```text
.
├── main.go
├── standard.txt
├── README.md
└── go.mod
```

### Files

| File           | Description                                 |
| -------------- | ------------------------------------------- |
| `main.go`      | Application entry point and rendering logic |
| `standard.txt` | ASCII banner font definitions               |
| `README.md`    | Project documentation                       |

---

## Requirements

* Go 1.20 or newer
* Terminal supporting ANSI escape sequences

---

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/ascii-art-colorizer.git
cd ascii-art-colorizer
```

Run the application:

```bash
go run .
```

Or build an executable:

```bash
go build -o ascii-art-colorizer
```

Execute:

```bash
./ascii-art-colorizer
```

---

## Usage

```bash
go run . [OPTION] [STRING]
```

### Usage Output

```text
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

---

## Examples

### 1. Color Entire Text (Default Red)

```bash
go run . "Hello"
```

Result:

```text
HELLO
```

Rendered as ASCII art in red.

---

### 2. Specify a Color

```bash
go run . --color=blue "Hello"
```

The entire ASCII banner will be displayed in blue.

---

### 3. Highlight a Substring

```bash
go run . world "hello world"
```

The substring `world` will be highlighted using the default color (red).

---

### 4. Specify Color and Substring

```bash
go run . --color=green world "hello world"
```

Only the ASCII-art representation of `world` will be colored green.

---

### 5. Multi-line Input

```bash
go run . --color=cyan "Hello\nWorld"
```

Output will render two separate ASCII-art lines.

---

## How It Works

### Argument Parsing

The application supports four valid execution modes:

#### Mode 1: Text Only

```bash
go run . "hello"
```

* Uses default color: red
* Colors the entire output

#### Mode 2: Color + Text

```bash
go run . --color=blue "hello"
```

* Uses specified color
* Colors the entire output

#### Mode 3: Substring + Text

```bash
go run . world "hello world"
```

* Uses default red color
* Colors only the matching substring

#### Mode 4: Color + Substring + Text

```bash
go run . --color=green world "hello world"
```

* Uses specified color
* Colors only the matching substring

---

## Font Rendering

Each printable ASCII character is stored inside `standard.txt`.

The rendering process:

1. Read the font file.
2. Locate the character pattern.
3. Extract the corresponding 8-line banner representation.
4. Apply ANSI color formatting where required.
5. Assemble the final output.

---

## Error Handling

### Invalid Arguments

When invalid arguments are supplied:

```bash
go run . invalid arguments here
```

The program displays:

```text
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

### Missing Font File

If `standard.txt` cannot be found:

```text
Error reading standard.txt
```

---

## ANSI Color Support

The application uses ANSI escape sequences:

```go
"\033[31m" // red
"\033[32m" // green
"\033[34m" // blue
```

Colors are reset after every rendered ASCII segment using:

```go
"\033[0m"
```

---

## Performance

* Time Complexity: O(n × h)

Where:

* `n` = number of characters
* `h` = banner height (8 rows)

Memory usage remains minimal due to efficient string construction using `strings.Builder`.

---

## Future Improvements

* Multiple font support
* RGB and hexadecimal color support
* Background color rendering
* Gradient text effects
* Custom font loading
* Export to text files
* Cross-platform color detection

---

## Author

**JOHN SUNDAY OLUWASHOMI**

Software Developer | Go Developer

GitHub: https://github.com/boanerges-dev

---

## License

This project is released under the MIT License.

Feel free to use, modify, and distribute it in accordance with the license terms.
