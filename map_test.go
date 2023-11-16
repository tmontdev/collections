package collections

import (
	"fmt"
	"testing"
)

func TestMap_IsEmpty(t *testing.T) {
	var m Map[string, any]
	isEmpty := m.IsEmpty()
	if m.IsNotEmpty() {
		t.Error("map should be empty")
	}
	m = Map[string, any]{}
	isEmpty = m.IsEmpty()
	if !isEmpty {
		t.Error("map should be empty")
	}
	m.Set("Key", "Value")
	if m.IsEmpty() {
		t.Error("map should not be empty")
	}
}

func TestMap_Length(t *testing.T) {
	var m Map[string, any]

	if m.Length() != 0 {
		t.Error("Length should be zero")
	}
	m = map[string]any{}

	if m.Length() != 0 {
		t.Error("Length should be zero")
	}
	m.Set("key", "value")
	if m.Length() != 1 {
		t.Error("Length should be one")
	}
}

func TestMap_Every(t *testing.T) {
	m := Map[string, any]{}
	m.Set("string1", "string1")
	m.Set("int1", 1)
	m.Set("string2", "string2")
	predicate := func(key string, value any) bool {
		_, is := value.(string)
		return is
	}
	if !m.Every(predicate) {
		strs := m.Where(predicate)
		if !strs.Every(predicate) {
			t.Error("All elements should be string")
		}
	}

}

func TestMap_None(t *testing.T) {
	m := Map[string, any]{}
	predicate := func(key string, value any) bool {
		_, is := value.(string)
		return !is
	}
	if m.Where(predicate).IsNotEmpty() {
		t.Error("any predicate on empty map should be empty, but should not crash")
	}
	if m.RemoveWhere(predicate).IsNotEmpty() {
		t.Error("any predicate on empty map should be empty, but should not crash")
	}

	m.Set("string1", "string1")
	m.Set("int1", 1)
	m.Set("string2", "string2")

	if m.Some(predicate) {
		m.RemoveWhere(predicate)
	}

	if !m.None(predicate) {
		t.Error("All elements should be string")
	}
}

func TestMap_Has(t *testing.T) {
	m := Map[int, float64]{
		1: 1.24,
		2: 1.25,
		3: 1.26,
	}
	has := m.Has(3)
	value := m.Get(3)

	dvalue, dhas := m.Access(3)
	if !has || !dhas {
		t.Error("should have 3")
	}
	if value != 1.26 || dvalue != 1.26 {
		t.Error("3 should be 1.23")
	}
	fmt.Println(value, has)
}

func TestMap_Access(t *testing.T) {
	m := Map[int, float64]{
		1: 1.24,
		2: 1.25,
		3: 1.26,
	}
	c := m.Clone()
	for mk, mv := range m {
		cv, has := c.Access(mk)
		if !has || cv != mv {
			t.Error("clone should have the same keys and values of original")
		}
	}
}
