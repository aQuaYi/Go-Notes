package example

import (
	"os"
	"testing"

	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
)

func TestStubEnv(t *testing.T) {
	ast := assert.New(t)
	stubs := gostub.New()

	os.Setenv("SETTED_ENV", "SE")
	os.Unsetenv("NONE")

	stubs.SetEnv("NONE", "STUB")
	stubs.UnsetEnv("SETTED_ENV")

	ast.Equal("STUB", os.Getenv("NONE"), "没有把 NONE 变量打桩成 STUB")
	ast.Equal("", os.Getenv("SETTED_ENV"), "没有把 SETTED_ENV 打桩成 空")

	stubs.Reset()

	_, ok := os.LookupEnv("NONE")
	ast.False(ok, "NONE should be unset")

	ast.Equal("SE", os.Getenv("SETTED_ENV"), "SETTED_ENV 没有被还原成原先的值")
}
