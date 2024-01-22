package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()
    return Eval(program)
}

func TestEvalIntegerExpression(t *testing.T) {
    tests := []struct {
        input string
        expected int64
    }{
        {"5", 5},
        {"10", 10},
    }

    for _, tt := range tests {
        testIntegerObject(t, testEval(tt.input), tt.expected)
    }
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool{
    intObj, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("expected obj to be of type object.Integer. got=%T", intObj)
        return false
    }
    if intObj.Value != expected {
        t.Errorf("object has wrong value. got=%d, want=%d", intObj.Value, expected)
        return false
    }
    return true
}
