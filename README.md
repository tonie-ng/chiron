# Chiron
A lightweight interpreter based on Thorsten Ball's Writing an interpreter in Go.

## Installation.
To install Chiron, follow these steps:

Clone the Chiron repository to your local machine:
```bash
git clone https://github.com/tonie-ng/chiron.go.git
```

Navigate to the Chiron directory:
```bash
cd chiron
```

Build the Chiron interpreter:
```bash
go build
```
This will create an executable file named chiron in the current directory.

Optionally, move the chiron executable to a directory in your system's PATH for easier access:
```bash
sudo mv chiron /usr/local/bin/
```
Now, Chiron is ready to be used on your system.

## Usage
The compiler can be used in both interactive mode and non-interactive mode.
1. To use in interactive mode, just run the executable without any arguments like this
```bash
./chiron
```
2. To use in non-interactive mode, provide a file with the `.ch` extension. Like this
```bash
./chiron test.ch
```
This will evaluate the file and run the necessary command.

## Todo
- [x] Implement the lexer.
- [ ] Implement the parser.
- [ ] Implement evaluation.
- [ ] Add more tests
