package learning_adapter

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet() string
}

func Greet() string {
	return "hello, world"
}

func GreetWorld(t testing.TB, greeter Greeter) {
	got := greeter.Greet()
	assert.Equal(t, got, "hello, world")
}

type GreetAdpater func() string

func (g GreetAdpater) Greet() string {
	return g()
}

func TestGreet(t *testing.T) {
	// GreetWorld(t, Greet)  // 错误写法
	GreetWorld(t, GreetAdpater(Greet))
}
