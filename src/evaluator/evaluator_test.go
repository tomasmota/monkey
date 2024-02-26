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

// INTS
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

// BOOLEANS
func TestEvalBooleanExpression(t *testing.T) {
    tests := []struct {
        input string
        expected bool
    }{
        {"true", true},
        {"false", false},
    }
    for _, tt := range tests {
        testBooleanObject(t, testEval(tt.input), tt.expected)
    }
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
    result, ok := obj.(*object.Boolean)
    if !ok {
        t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
        return false
    }
    if result.Value != expected {
        t.Errorf("object has wrong value. got=%t, want=%t", obj, expected)
        return false
    }
    return true
}
