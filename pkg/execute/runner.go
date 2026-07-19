package execute

import (
	"fmt"
	"os"

	"github.com/tom96da/sleepingknights/pkg/lexer"
	"github.com/tom96da/sleepingknights/pkg/token"
)

func executeScript(scriptPath string) ExitStatus {
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Printf("[Error] Failed to read script: %v\n", err)
		return ExitStatusCriticalException
	}

	l := lexer.New(string(content))
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		if tok.Type == token.ILLEGAL {
			fmt.Printf("[Error] Lexing failed at %d:%d: %s\n", tok.Line, tok.Column, tok.Literal)
			return ExitStatusCriticalException
		}
	}

	return ExitStatusSuccess
}

func compileAndExecute(_ string) ExitStatus {
	return ExitStatusNotImplemented
}
