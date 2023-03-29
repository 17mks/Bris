package service

import (
	"strings"
	"testing"
)

func TestFileObjNameParse(t *testing.T) {
	fileObjName := "bfr/user/肺围手术期症状量表(PSA-Lung)-1638421045419540480.zip"
	split := strings.SplitN(fileObjName, "/", 2)

	for i, s := range split {
		t.Log("i = ", i, ", s = ", s)
	}
}
