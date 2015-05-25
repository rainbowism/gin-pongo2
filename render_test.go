package render

import (
	"io/ioutil"
	"testing"

	"github.com/flosch/pongo2"
)

var renderProd *PongoProduction
var renderDebug *PongoDebug

func init() {
	renderProd = NewProduction("template_tests")
	renderDebug = NewDebug("template_tests")
	renderProd.Instance("index.tpl", Context{})
}

func BenchmarkProduction(b *testing.B) {
	r := renderProd.Instance("index.tpl", pongo2.Context{"data": "test data"}).(Pongo)
	for n := 0; n < b.N; n++ {
		ctx := r.Data.(pongo2.Context)
		err := r.Template.ExecuteWriterUnbuffered(ctx, ioutil.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDebug(b *testing.B) {
	r := renderDebug.Instance("index.tpl", pongo2.Context{"data": "test data"}).(Pongo)
	for n := 0; n < b.N; n++ {
		ctx := r.Data.(pongo2.Context)
		err := r.Template.ExecuteWriterUnbuffered(ctx, ioutil.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProductionWrapper(b *testing.B) {
	r := renderProd.Instance("index.tpl", Context{"data": "test data"}).(Pongo)
	for n := 0; n < b.N; n++ {
		ctx := pongo2.Context(r.Data.(Context))
		err := r.Template.ExecuteWriterUnbuffered(ctx, ioutil.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDebugWrapper(b *testing.B) {
	r := renderDebug.Instance("index.tpl", Context{"data": "test data"}).(Pongo)
	for n := 0; n < b.N; n++ {
		ctx := pongo2.Context(r.Data.(Context))
		err := r.Template.ExecuteWriterUnbuffered(ctx, ioutil.Discard)
		if err != nil {
			b.Fatal(err)
		}
	}
}
