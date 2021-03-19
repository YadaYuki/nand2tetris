package compilationengine

import (
	"jack/compiler/ast"
	"jack/compiler/tokenizer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let hoge = 111;
		let foobar = 838383;
		`
	jt := tokenizer.New(input)
	ce := New(jt)
	program := ce.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got=%d", len(program.Statements))
	}
	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"hoge"},
		{"foobar"},
	}
	for i, tt := range testCases {
		stmt := program.Statements[i]
		// fmt.Println(stmt.TokenLiteral())
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// TODO:Fix panic

	// if s.TokenLiteral() != "let" {
	// 	t.Errorf("s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
	// }
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got %T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'.got '%s'", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'.got '%s'", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}