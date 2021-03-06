package read_test

import (
	"context"
	"github.com/modern-go/parse"
	"github.com/modern-go/parse/read"
	"github.com/modern-go/test"
	"github.com/modern-go/test/must"
	"strings"
	"testing"
)

func TestAnyExcept1(t *testing.T) {
	t.Run("not found", test.Case(func(ctx context.Context) {
		src := must.Call(parse.NewSource,
			strings.NewReader("abcd"), make([]byte, 2))[0].(*parse.Source)
		must.Equal([]byte{'a', 'b', 'c', 'd'}, read.AnyExcept1(
			src, 'g'))
	}))
	t.Run("found", test.Case(func(ctx context.Context) {
		src := must.Call(parse.NewSource,
			strings.NewReader("abcd"), make([]byte, 2))[0].(*parse.Source)
		must.Equal([]byte{'a', 'b', 'c'}, read.AnyExcept1(
			src, 'd'))
	}))
}

func TestUntil1(t *testing.T) {
	t.Run("not found", test.Case(func(ctx context.Context) {
		src := must.Call(parse.NewSource,
			strings.NewReader("abcd"), make([]byte, 2))[0].(*parse.Source)
		must.Equal([]byte{'a', 'b', 'c', 'd'}, read.Until1(
			src, 'g'))
		must.NotNil(src.FatalError())
	}))
	t.Run("found", test.Case(func(ctx context.Context) {
		src := must.Call(parse.NewSource,
			strings.NewReader("abcd"), make([]byte, 2))[0].(*parse.Source)
		must.Equal([]byte{'a', 'b', 'c'}, read.Until1(
			src, 'd'))
		must.Nil(src.FatalError())
	}))
}

func TestAnyExcept2(t *testing.T) {
	t.Run("not found", test.Case(func(ctx context.Context) {
		src := must.Call(parse.NewSource,
			strings.NewReader("abcd"), make([]byte, 2))[0].(*parse.Source)
		must.Equal([]byte{'a', 'b', 'c', 'd'}, read.AnyExcept2(
			src, 'g', 'f'))
	}))
	t.Run("found", test.Case(func(ctx context.Context) {
		src := must.Call(parse.NewSource,
			strings.NewReader("abcd"), make([]byte, 2))[0].(*parse.Source)
		must.Equal([]byte{'a', 'b', 'c'}, read.AnyExcept2(
			src, 'd', 'f'))
	}))
}
