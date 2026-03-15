package ttpl_test

import (
	"bytes"
	"html/template"
	"strings"
	"testing"

	"github.com/empijei/tst"
	"github.com/empijei/typ/ttpl"
)

func TestTpl(t *testing.T) {
	tst.Go(t)

	type Data struct {
		Name string
	}

	t.Run("Basic", func(t *testing.T) {
		tpl := ttpl.New[Data]("test")
		_, err := tpl.Parse("Hello {{.Name}}")
		tst.No(err, t)

		var buf bytes.Buffer
		err = tpl.Execute(&buf, Data{Name: "World"})
		tst.No(err, t)
		tst.Is("Hello World", buf.String(), t)
	})

	t.Run("Delims", func(t *testing.T) {
		tpl := ttpl.New[Data]("test").Delims("<<", ">>")
		_, err := tpl.Parse("Hello <<.Name>>")
		tst.No(err, t)

		var buf bytes.Buffer
		err = tpl.Execute(&buf, Data{Name: "World"})
		tst.No(err, t)
		tst.Is("Hello World", buf.String(), t)
	})

	t.Run("Funcs", func(t *testing.T) {
		tpl := ttpl.New[Data]("test").Funcs(template.FuncMap{
			"upper": strings.ToUpper,
		})
		_, err := tpl.Parse("Hello {{upper .Name}}")
		tst.No(err, t)

		var buf bytes.Buffer
		err = tpl.Execute(&buf, Data{Name: "world"})
		tst.No(err, t)
		tst.Is("Hello WORLD", buf.String(), t)
	})

	t.Run("Option", func(t *testing.T) {
		tpl := ttpl.New[Data]("test").Option("missingkey=error")
		_, err := tpl.Parse("Hello {{.Name}} {{.Missing}}")
		tst.No(err, t)

		var buf bytes.Buffer
		err = tpl.Execute(&buf, Data{Name: "world"})
		tst.Be(err != nil, t)
	})

	t.Run("ParseError", func(t *testing.T) {
		tpl := ttpl.New[Data]("test")
		_, err := tpl.Parse("Hello {{.Name")
		tst.Be(err != nil, t)
	})
}
