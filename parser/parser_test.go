package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		t.Errorf("parser error: %q", err)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 123123;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, *p)

	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("expected s to be an *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("setStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("setStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("setStmt.Name.TokenLiteral not '%s', got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 1123123;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, *p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		rs, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("expected s to be an *ast.ReturnStatement. got=%T", rs)
			continue
		}

		if rs.TokenLiteral() != "return" {
			t.Errorf("rs.TokenLiteral not 'return', got=%s", rs.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, *p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("expected program.Statements[0] to be an *ast.ExpressionStatement. got=%T", stmt)
	}

    ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Errorf("expression is not and Identifier. got=%T", stmt)
	}
    if ident.Value != "foobar" {
        t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
    }
    if ident.TokenLiteral() != "foobar" {
        t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar", ident.Value)
    }
}
