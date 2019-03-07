package object

import "testing"

func TestObjectType(t *testing.T) {
	tests := []struct {
		obj      Object
		expected ObjectType
	}{
		{&Integer{}, INTEGER_OBJ},
		{&String{}, STRING_OBJ},
		{&Boolean{}, BOOLEAN_OBJ},
		{&Null{}, NULL_OBJ},
		{&Array{}, ARRAY_OBJ},
		{&Hash{}, HASH_OBJ},
		{&Error{}, ERROR_OBJ},
		{&ReturnValue{}, RETURN_VALUE_OBJ},
		{&Function{}, FUNCTION_OBJ},
		{&Builtin{}, BUILTIN_OBJ},
	}

	for _, tt := range tests {
		if tt.obj.Type() != tt.expected {
			t.Errorf("wrong object type. expected=%s, got=%s",
				tt.obj.Type(), tt.expected)
		}
	}
}

func TestObjectInspect(t *testing.T) {
	tests := []struct {
		obj      Object
		expected string
	}{
		{&Integer{Value: -4}, "-4"},
		{&String{Value: "John Doe"}, "John Doe"},
		{&Boolean{Value: true}, "true"},
		{&Null{}, "null"},
		{
			&Array{
				Elements: []Object{
					&Integer{Value: 1},
					&String{Value: "aaa"},
					&Boolean{Value: true},
				}},
			"[1, aaa, true]",
		},
		// TODO: Hash
		{&Error{Message: "foobar"}, "ERROR: foobar"},
		// TODO: ReturnValue
		// TODO: Function
		{&Builtin{}, "builtin function"},
	}

	for _, tt := range tests {
		if tt.obj.Inspect() != tt.expected {
			t.Errorf("wrong object type. expected=%s, got=%s",
				tt.obj.Inspect(), tt.expected)
		}
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}
	if false1.HashKey() != false2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}
	if true1.HashKey() == false1.HashKey() {
		t.Errorf("booleans with different content have different hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	two1 := &Integer{Value: 2}
	two2 := &Integer{Value: 2}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	if two1.HashKey() != two2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	if one1.HashKey() == two1.HashKey() {
		t.Errorf("integers with different content have different hash keys")
	}
}

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have different hash keys")
	}
}
