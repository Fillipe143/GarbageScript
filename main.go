package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type BlockType int

const (
	IF BlockType = iota
	WHILE
	FUNC
)

type Lexer struct {
	source []byte
	offset int
	addr   int
}

type Block struct {
	blockType BlockType
	addr      int
	addr2     int
	offset    int
	isOpened  bool
}

var blocks []Block

var escapeChars = map[rune]rune{
	'\\': '\\',
	'"':  '"',
	'n':  '\n',
	't':  '\t',
	'r':  '\r',
	'\'': '\'',
}

func main() {
	filename := "example.gb"
	name := strings.Split(filename, ".")[0]
	source := readFile(filename)

	lexer := Lexer{source: source}
	instructions := lexer.compile(filename)

	err := os.WriteFile(name+".asm", instructions, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("nasm", "-felf64", name+".asm", "-o", name+".o").Run()
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("ld", name+".o", "-o", name).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(filename string) []byte {
	source, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return source
}

func (l *Lexer) compile(filename string) []byte {
	var instructions bytes.Buffer
	var functions bytes.Buffer
	var functionList []string

	fileList := []string{filename}
	numberOfFunctions := 0
	strCounter := 0

	var variables []string

	var data bytes.Buffer
	data.WriteString("section .data\n")

	var bss bytes.Buffer
	bss.WriteString("section .bss\ntmp resq 1000\ni resq 1\nvars resq ")

	instructions.WriteString("section .text\nglobal _start\nprint:\nmov r9, -3689348814741910323\nsub rsp, 40\nmov BYTE [rsp+31], 10\nlea rcx, [rsp+30]\n.L2:\nmov rax, rdi\nlea r8, [rsp+32]\nmul r9\nmov rax, rdi\nsub r8, rcx\nshr rdx, 3\nlea rsi, [rdx+rdx*4]\nadd rsi, rsi\nsub rax, rsi\nadd eax, 48\nmov BYTE [rcx], al\nmov rax, rdi\nmov rdi, rdx\nmov rdx, rcx\nsub rcx, 1\ncmp rax, 9\nja  .L2\nlea rax, [rsp+32]\nmov edi, 1\nsub rdx, rax\nxor eax, eax\nlea rsi, [rsp+32+rdx]\nmov rdx, r8\nmov rax, 1\nsyscall\nadd rsp, 40\nret\n_start:\n")
	for l.offset < len(l.source) {
		var writer *bytes.Buffer
		if numberOfFunctions > 0 {
			writer = &functions
		} else {
			writer = &instructions
		}
		switch l.source[l.offset] {
		case 'x':
			l.offset++
			m := 1
			if l.source[l.offset] == '-' {
				m = -1
				l.offset++
			}
			if unicode.IsDigit(rune(l.source[l.offset])) {
				n := l.readNumber() * m
				writer.WriteString(fmt.Sprintf("push %d\n", n))
				continue
			}
			if l.source[l.offset] == '"' {
				l.offset++
				s := l.readString()
				data.WriteString("str_" + strconv.Itoa(strCounter) + " db ")
				for i := 0; i < len(s); i++ {
					data.WriteString(fmt.Sprintf("0x%X", s[i]))
					if i < len(s)-1 {
						data.WriteString(", ")
					}
				}

				data.WriteRune('\n')
				data.WriteString("str_len_" + strconv.Itoa(strCounter) + " equ $ - str_" + strconv.Itoa(strCounter) + "\n")
				writer.WriteString("push str_len_" + strconv.Itoa(strCounter) + "\n")
				writer.WriteString("push str_" + strconv.Itoa(strCounter) + "\n")
				strCounter++
				continue
			}
			if l.source[l.offset] == '\'' {
				l.offset++
				c := l.readChar()
				writer.WriteString(fmt.Sprintf("push 0x%X\n", c))
				continue
			}

			log.Fatalf("Invalid instruction at offset %d", l.offset+1)
		case 'y':
			l.offset++
			writer.WriteString("pop rax\npush rax\npush rax\n")
		case 'z':
			l.offset++
			if l.offset < len(l.source) && l.source[l.offset] == '3' {
				l.offset++
				writer.WriteString("pop rax\npop rdi\npop rbx\npush rax\npush rdi\npush rbx\n")
			} else {
				writer.WriteString("pop rax\npop rdi\npush rax\npush rdi\n")
			}
		case 't':
			l.offset++
			writer.WriteString("pop rax\npop rdi\npush rdi\npush rax\npush rdi\n")
		case 'm':
			l.offset++
			var name string
			if l.offset >= len(l.source) || l.source[l.offset] != '\'' {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}

			l.offset++
			for l.offset < len(l.source) && l.source[l.offset] != '\'' {
				name += string(l.source[l.offset])
				l.offset++
			}

			if l.offset >= len(l.source) {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}
			l.offset++
			i := len(variables) * 16
			if slices.Contains(variables, name) {
				i = slices.Index(variables, name) * 16
				variables[i] = name
			} else {
				variables = append(variables, name)
			}
			writer.WriteString("pop rsi\npush rsi\nmov rax, 9\nxor rdi, rdi\nmov rdx, 3\nmov r10, 34\n mov r8, -1\nxor r9, r9\nsyscall\n")
			writer.WriteString("mov QWORD [vars+" + strconv.Itoa(i) + "], rax\n")
			writer.WriteString("pop rax\nmov QWORD [vars+" + strconv.Itoa(i+8) + "], rax\n")
		case 'p':
			l.offset++
			var name string
			if l.offset >= len(l.source) || l.source[l.offset] != '\'' {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}

			l.offset++
			for l.offset < len(l.source) && l.source[l.offset] != '\'' {
				name += string(l.source[l.offset])
				l.offset++
			}

			if l.offset >= len(l.source) {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}
			l.offset++

			i := slices.Index(variables, name)
			if i == -1 {
				log.Fatalf("Undefined variable '%s' at offset %d", name, l.offset+1)
			}

			i *= 16
			writer.WriteString("mov rax, ")
			writer.WriteString("[vars+")
			writer.WriteString(strconv.Itoa(i))
			writer.WriteString("]\n")
			writer.WriteString("push rax\n")
		case 'u':
			l.offset++
			var name string
			if l.offset >= len(l.source) || l.source[l.offset] != '\'' {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}

			l.offset++
			for l.offset < len(l.source) && l.source[l.offset] != '\'' {
				name += string(l.source[l.offset])
				l.offset++
			}

			if l.offset >= len(l.source) {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}
			l.offset++

			i := slices.Index(variables, name)
			if i == -1 {
				log.Fatalf("Undefined variable '%s' at offset %d", name, l.offset+1)
			}

			i *= 16
			writer.WriteString("mov rax, ")
			writer.WriteString("[vars+")
			writer.WriteString(strconv.Itoa(i + 8))
			writer.WriteString("]\n")
			writer.WriteString("push rax\n")
		case 'v':
			l.offset++
			var name string
			if l.offset >= len(l.source) || l.source[l.offset] != '\'' {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}

			l.offset++
			for l.offset < len(l.source) && l.source[l.offset] != '\'' {
				name += string(l.source[l.offset])
				l.offset++
			}

			if l.offset >= len(l.source) {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}
			l.offset++

			i := slices.Index(variables, name)
			if i == -1 {
				log.Fatalf("Undefined variable '%s' at offset %d", name, l.offset+1)
			}

			i *= 16
			writer.WriteString("mov rax, ")
			writer.WriteString("[vars+")
			writer.WriteString(strconv.Itoa(i))
			writer.WriteString("]\n")
			writer.WriteString("xor rbx, rbx\nmov QWORD rbx, [rax]\npush rbx\n")
		case 's':
			l.offset++
			size := "QWORD"
			register := "rbx"
			if l.offset+1 < len(l.source) {
				sizes := map[byte]string{
					'b': "BYTE",
					'w': "WORD",
					'd': "DWORD",
					'q': "QWORD",
				}

				registers := map[byte]string{
					'b': "bl",
					'w': "bx",
					'd': "ebx",
					'q': "rbx",
				}

				s, ok := sizes[l.source[l.offset]]
				if ok {
					size = s
					register = registers[l.source[l.offset]]
					l.offset++
				}
			}

			writer.WriteString("pop rbx\npop rax\nmov " + size + " [rax], " + register + "\n")
		case 'l':
			l.offset++
			size := "QWORD"
			register := "rbx"
			if l.offset+1 < len(l.source) {
				sizes := map[byte]string{
					'b': "BYTE",
					'w': "WORD",
					'd': "DWORD",
					'q': "QWORD",
				}

				registers := map[byte]string{
					'b': "bl",
					'w': "bx",
					'd': "ebx",
					'q': "rbx",
				}

				s, ok := sizes[l.source[l.offset]]
				if ok {
					size = s
					register = registers[l.source[l.offset]]
					l.offset++
				}
			}

			writer.WriteString("pop rax\nxor rbx, rbx\nmov " + size + " " + register + ", [rax]\npush rbx\n")
		case 'r':
			l.offset++
			if len(blocks) == 0 {
				log.Fatalf("No function for instruction 'r' at offset %d", l.offset)
			}

			functions.WriteString("mov rcx, [i]\nsub rcx, 8\nmov [i], rcx\nmov rbx, [tmp+rcx]\npush rbx\n")
			functions.WriteString("ret\n")

		case 'i':
			l.offset++
			if l.offset >= len(l.source) || l.source[l.offset] != '"' {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}

			l.offset++
			file := l.readString()
			if slices.Contains(fileList, file) {
				log.Fatalf("Duplicate import at offset %d", l.offset+1)
			}

			fileList = append(fileList, file)

			content := readFile(file)
			p1 := append(l.source[:l.offset-1], '\n')
			p2 := append(content, l.source[l.offset:]...)
			l.source = append(p1, p2...)
		case '$':
			args := []string{"rax", "rdi", "rsi", "rdx", "r10", "r8", "r9"}
			l.offset++
			n := l.readNumber()
			if n <= 0 || n > len(args) {
				log.Fatalf("Invalid numbe of arguments at offset %d", l.offset+1)
			}
			for i := 0; i < n; i++ {
				writer.WriteString(fmt.Sprintf("pop %s\n", args[i]))
			}
			writer.WriteString("syscall\n")
			writer.WriteString("push rax\n")
		case '+':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nadd rax, rdi\npush rax\n")
		case '-':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nsub rax, rdi\npush rax\n")
		case '/':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nxor rdx, rdx\ndiv rdi\npush rdx\npush rax\n")
		case '*':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nmul rdi\npush rax\n")
		case '=':
			l.offset++
			writer.WriteString("mov rcx, 0\nmov rdx, 1\npop rax\npop rdi\ncmp rax, rdi\ncmove rcx, rdx\npush rcx\n")
		case '<':
			l.offset++
			writer.WriteString("mov rcx, 0\nmov rdx, 1\npop rax\npop rdi\ncmp rax, rdi\ncmovl rcx, rdx\npush rcx\n")
		case '>':
			l.offset++
			writer.WriteString("mov rcx, 0\nmov rdx, 1\npop rax\npop rdi\ncmp rax, rdi\ncmovg rcx, rdx\npush rcx\n")
		case '!':
			l.offset++
			writer.WriteString("mov rcx, 0\nmov rdx, 1\npop rax\ncmp rax, rcx\ncmove rcx, rdx\npush rcx\n")
		case '&':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nand rax, rdi\npush rax\n")
		case '|':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nor rax, rdi\npush rax\n")
		case '^':
			l.offset++
			writer.WriteString("pop rax\npop rdi\nxor rax, rdi\npush rax\n")
		case '.':
			l.offset++
			writer.WriteString("pop rax\n")
		case '@':
			l.offset++
			writer.WriteString("pop rdi\ncall print\n")
		case '{':
			l.offset++
			if len(blocks) > 0 {
				block := blocks[len(blocks)-1]
				if block.blockType == WHILE && !block.isOpened {
					blocks[len(blocks)-1].isOpened = true
					blocks[len(blocks)-1].addr2 = block.addr
					blocks[len(blocks)-1].addr = l.addr
					writer.WriteString("pop rax\ncmp rax, 0\nje .addr_")
					writer.WriteString(strconv.Itoa(l.addr))
					writer.WriteString("\n")
					l.addr++
					continue
				}
			}
			blocks = append(blocks, Block{IF, l.addr, -1, l.offset, true})
			writer.WriteString("pop rax\ncmp rax, 0\nje .addr_")
			writer.WriteString(strconv.Itoa(l.addr))
			writer.WriteString("\n")
			l.addr++
		case ']':
			l.offset++
			if len(blocks) == 0 {
				log.Fatalf("Unmatched ']' at offset %d", l.offset)
			}
			block := blocks[len(blocks)-1]

			writer.WriteString("jmp .addr_")
			writer.WriteString(strconv.Itoa(l.addr))
			writer.WriteString("\n")

			writer.WriteString(".addr_")
			writer.WriteString(strconv.Itoa(block.addr))
			writer.WriteString(":\n")
			blocks[len(blocks)-1].addr = l.addr
			blocks[len(blocks)-1].offset = l.offset
			l.addr++
		case '[':
			l.offset++
			blocks = append(blocks, Block{WHILE, l.addr, -1, l.offset, false})
			writer.WriteString(".addr_")
			writer.WriteString(strconv.Itoa(l.addr))
			writer.WriteString(":\n")
			l.addr++
		case '#':
			l.offset++
			var name string
			for l.offset < len(l.source) && isAlpha(l.source[l.offset]) {
				name += string(l.source[l.offset])
				l.offset++
			}
			if l.source[l.offset] != '%' {
				log.Fatalf("Unclosed %% at offset %d", l.offset)
			}
			l.offset++
			if len(name) == 0 {
				log.Fatalf("Invalid name at offset %d", l.offset+1)
			}

			for l.offset < len(l.source) && unicode.IsSpace(rune(l.source[l.offset])) {
				l.offset++
			}

			if l.offset >= len(l.source) || l.offset+2 >= len(l.source) || l.source[l.offset] != '(' {
				log.Fatalf("Invalid function at offset %d", l.offset+1)
			}
			l.offset++
			if slices.Contains(functionList, name) {
				log.Fatalf("Function '%s' already defined at offset %d", name, l.offset+1)
			}
			blocks = append(blocks, Block{FUNC, l.addr, -1, l.offset, false})
			functions.WriteString("_func_")
			functions.WriteString(name)
			functions.WriteString(":\n")
			functions.WriteString("pop rbx\nmov rcx, [i]\nmov [tmp+rcx], rbx\nadd rcx, 8\nmov [i], rcx\n")
			functionList = append(functionList, name)
			numberOfFunctions++
		case '%':
			l.offset++
			var name string
			for l.offset < len(l.source) && isAlpha(l.source[l.offset]) {
				name += string(l.source[l.offset])
				l.offset++
			}
			if l.source[l.offset] != '#' {
				log.Fatalf("Unclosed #  at offset %d", l.offset)
			}
			l.offset++
			if len(name) == 0 {
				log.Fatalf("Invalid name at offset %d", l.offset+1)
			}

			if !slices.Contains(functionList, name) {
				log.Fatalf("Undefined function at offset %d", l.offset+1)
			}

			writer.WriteString("call _func_")
			writer.WriteString(name)
			writer.WriteString("\n")
		case ';':
			l.offset++
			if len(blocks) == 0 {
				log.Fatalf("Unmatched ';' at offset %d", l.offset)
			}
			block := blocks[len(blocks)-1]
			blocks = blocks[:len(blocks)-1]

			if block.blockType == FUNC {
				numberOfFunctions--
				functions.WriteString("mov rcx, [i]\nsub rcx, 8\nmov [i], rcx\nmov rbx, [tmp+rcx]\npush rbx\n")
				functions.WriteString("ret\n")
				continue
			}
			if block.blockType == WHILE {
				writer.WriteString("jmp .addr_")
				writer.WriteString(strconv.Itoa(block.addr2))
				writer.WriteString("\n")
			}

			writer.WriteString(".addr_")
			writer.WriteString(strconv.Itoa(block.addr))
			writer.WriteString(":\n")
		case ',':
			l.offset++
			for l.offset < len(l.source) && l.source[l.offset] != '\n' && l.source[l.offset] != ',' {
				l.offset++
			}
			l.offset++
		case '\'':
			l.offset++
			var name string
			for l.offset < len(l.source) && l.source[l.offset] != '\'' {
				name += string(l.source[l.offset])
				l.offset++
			}

			if l.offset >= len(l.source) {
				log.Fatalf("Invalid instruction at offset %d", l.offset+1)
			}
			l.offset++

			i := slices.Index(variables, name)
			if i == -1 {
				log.Fatalf("Undefined variable '%s' at offset %d", name, l.offset+1)
			}

			i *= 16
			writer.WriteString("mov rax, ")
			writer.WriteString("[vars+")
			writer.WriteString(strconv.Itoa(i + 8))
			writer.WriteString("]\n")
			writer.WriteString("push rax\n")

			writer.WriteString("mov rax, ")
			writer.WriteString("[vars+")
			writer.WriteString(strconv.Itoa(i))
			writer.WriteString("]\n")
			writer.WriteString("push rax\n")
		default:
			if unicode.IsSpace(rune(l.source[l.offset])) {
				l.offset++
				continue
			}
			log.Fatalf("Invalid instruction at offset %d", l.offset+1)
		}
	}
	if len(blocks) > 0 {
		log.Fatalf("Unmatched '%c' at offset %d\n", l.source[blocks[0].offset-1], blocks[0].offset)
	}
	instructions.WriteString("mov rax, 60\nxor rdi, rdi\nsyscall\n")
	bss.WriteString(strconv.Itoa(len(variables) * 2))
	bss.WriteString("\n")

	return append(data.Bytes(), append(bss.Bytes(), append(instructions.Bytes(), functions.Bytes()...)...)...)
}

func (l *Lexer) readNumber() int {
	var buf bytes.Buffer
	for l.offset < len(l.source) && unicode.IsDigit(rune(l.source[l.offset])) {
		buf.WriteRune(rune(l.source[l.offset]))
		l.offset++
	}
	if len(buf.Bytes()) == 0 {
		log.Fatalf("Invalid number at offset %d", l.offset+1)
	}
	n, err := strconv.Atoi(buf.String())
	if err != nil {
		log.Fatalf("Invalid number at offset %d", l.offset+1)
	}
	return n
}

func (l *Lexer) readString() string {
	var buf bytes.Buffer
	lastChar := '"'
	for l.offset < len(l.source) && (l.source[l.offset] != '"' || lastChar == '\\') {
		if l.source[l.offset] == '\\' {
			if l.offset+1 >= len(l.source) {
				log.Fatalf("Invalid string at offset %d", l.offset+1)
			}
			c, ok := escapeChars[rune(l.source[l.offset+1])]
			if !ok {
				log.Fatalf("Invalid escape character at offset %d", l.offset+1)
			}
			buf.WriteRune(c)
			l.offset += 2
			continue
		}
		buf.WriteRune(rune(l.source[l.offset]))
		lastChar = rune(l.source[l.offset])
		l.offset++
	}

	if len(buf.Bytes()) == 0 {
		log.Fatalf("Invalid string at offset %d", l.offset+1)
	}
	l.offset++
	return buf.String()
}

func (l *Lexer) readChar() string {
	var buf bytes.Buffer
	lastChar := '\''
	for l.offset < len(l.source) && (l.source[l.offset] != '\'' || lastChar == '\\') {
		if l.source[l.offset] == '\\' {
			if l.offset+1 >= len(l.source) {
				log.Fatalf("Invalid char at offset %d", l.offset+1)
			}
			c, ok := escapeChars[rune(l.source[l.offset+1])]
			if !ok {
				log.Fatalf("Invalid escape character at offset %d", l.offset+1)
			}
			buf.WriteRune(c)
			l.offset += 2
			continue
		}
		buf.WriteRune(rune(l.source[l.offset]))
		lastChar = rune(l.source[l.offset])
		l.offset++
	}

	if len(buf.Bytes()) == 0 {
		log.Fatalf("Invalid char at offset %d", l.offset+1)
	}
	l.offset++
	if len(buf.Bytes()) > 1 {
		log.Fatalf("Invalid char at offset %d", l.offset+1)
	}
	return buf.String()
}

func isAlpha(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_'
}
