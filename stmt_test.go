package sqlu_test

import (
	"github.com/digitive/sqlu"
	"testing"
)

func TestStmt(t *testing.T) {
	sql, args := sqlu.Update("people").Set("age", 18).Where("name=?", "John Smith").Build()
	if sql != "UPDATE `people` SET `age`=? WHERE name=?" {
		t.Errorf("Update() failed, got: %s", sql)
	}
	if len(args) != 2 || args[0] != 18 || args[1] != "John Smith" {
		t.Errorf("Update() failed, got: %v", args)
	}
}
