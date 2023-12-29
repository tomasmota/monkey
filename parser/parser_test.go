package parser

import (
	"fmt"
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

func TestIntegerLiteral(t *testing.T) {
	input := `5;`

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

    literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expression is not *ast.IntegerLiteral. got=%T", stmt)
	}
    if literal.Value != 5 {
        t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
    }
    if literal.TokenLiteral() != "5" {
        t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
    }
}

func TestParsingPrefixExpressions(t *testing.T) {
    prefixTests:= []struct {
        input string
        operator string
        integerValue int64
    }{
        {"!15", "!", 15},
        {"-15", "-", 15},
    }

    for _, tt := range prefixTests {
        l := lexer.New(tt.input)
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

        exp, ok := stmt.Expression.(*ast.PrefixExpression)
        if !ok {
            t.Errorf("expression is not *ast.PrefixExpression. got=%T", stmt)
        }
        if exp.Operator != tt.operator {
            t.Errorf("exp.Operator not %s. got=%s", tt.operator, exp.Operator)
        }
        if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
            return
        }
    }
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
    il, ok := exp.(*ast.IntegerLiteral)
    if !ok {
        t.Errorf("exp is not *ast.IntegerLiteral. got=%T", exp)
        return false
    }
    if il.Value != value {
        t.Errorf("expected il.Value to be %d. got=%d", value, il.Value)
        return false
    }
    if il.TokenLiteral() != fmt.Sprintf("%d", value) {
        t.Errorf("expected il.TokenLiteral to be %d. got=%s", value, il.TokenLiteral())
        return false
    }
    return true
}
