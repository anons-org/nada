package Utils_test

import (
	Utils "nada/src/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	rootPath := "/Volumes/Develop/anons-org/nada"
	assert.Equal(t, Utils.RootPath, rootPath, "根路径不对")
}
