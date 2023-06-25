### Writing and Interpreter in Go
This reposetory is the summation of my progress of building an interpreter of the Monkey Programming Language using the book: **Writing an Interpreter in Go** by Thorston Hall. I have some experience in C, minimal experience using Git, and I have never progrmamed in the Go Programming Language before. This will be a learning experience for me. If you have any comments or recommendations to help me, feel free to email me at cea57@cornell.edu. Thank you and enjoy!

Interpreters take source code and evaluate it so that it can be executed later. 
The output of this book will be a "tree-walking interpreter."

## Lexer

The lexer takes in source code as an input and outputs the tokens that represent the source code. It goes through its input then outputs the next token it recognizes.

## Parser
The parser is a software component that takes input data (frequently text) and builds a data structure that gives structural representation of the input(parse tree, abstract syntax tree, etc.), checking for correct syntax in the process. The parser is preceded by the lexer that creates tokens from the sequence of the input characters.