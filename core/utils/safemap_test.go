package utils

import "testing"

var safeMap *RadicalMap

func TestNewRadicalMap(t *testing.T) {
	safeMap = NewRadicalMap()
	if safeMap == nil {
		t.Fatal("expected to return non-nil RadicalMap", "got", safeMap)
	}
}

func TestSet(t *testing.T) {
	safeMap = NewRadicalMap()
	if ok := safeMap.Set("astaxie", 1); !ok {
		t.Error("expected", true, "got", false)
	}
}

func TestReSet(t *testing.T) {
	safeMap := NewRadicalMap()
	if ok := safeMap.Set("astaxie", 1); !ok {
		t.Error("expected", true, "got", false)
	}
	// set diff value
	if ok := safeMap.Set("astaxie", -1); !ok {
		t.Error("expected", true, "got", false)
	}

	// set same value
	if ok := safeMap.Set("astaxie", -1); ok {
		t.Error("expected", false, "got", true)
	}
}

func TestCheck(t *testing.T) {
	if exists := safeMap.Check("astaxie"); !exists {
		t.Error("expected", true, "got", false)
	}
}

func TestGet(t *testing.T) {
	if val := safeMap.Get("astaxie"); val.(int) != 1 {
		t.Error("expected value", 1, "got", val)
	}
}

func TestDelete(t *testing.T) {
	safeMap.Delete("astaxie")
	if exists := safeMap.Check("astaxie"); exists {
		t.Error("expected element to be deleted")
	}
}

func TestItems(t *testing.T) {
	safeMap := NewRadicalMap()
	safeMap.Set("astaxie", "hello")
	for k, v := range safeMap.Items() {
		key := k.(string)
		value := v.(string)
		if key != "astaxie" {
			t.Error("expected the key should be astaxie")
		}
		if value != "hello" {
			t.Error("expected the value should be hello")
		}
	}
}

func TestCount(t *testing.T) {
	if count := safeMap.Count(); count != 0 {
		t.Error("expected count to be", 0, "got", count)
	}
}
