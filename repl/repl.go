package repl

import (
	"bufio"
	"chiron/lexer"
	"chiron/token"
	"fmt"
	"io"
	"log"
	"os"
)

func start(input io.Reader, output io.Writer, mode string) {
	scanner := bufio.NewScanner(input)

	for {

		if mode == "interactive" {
			fmt.Print(PROMPT)
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		lexer := lexer.New(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(output, "%+v\n", tok)
		}
	}
}

func Load(args []string, username string) {

	if len(args) > 2 {
		log.Fatal("Usage: chiron <filename>")
	}

	if len(args) == 1 {
		fmt.Printf("Hello %s, welcome to Chiron!\n", username)
		start(os.Stdin, os.Stderr, INTERACTIVEMODE)
	} else {
		filename := args[1]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer func () {
			if err := file.Close(); err != nil {
				fmt.Fprintf(os.Stdout, "Error closing file: %v\n", err)
			}
		}()
		start(file, os.Stdout, FILEMODE)
	}
}
