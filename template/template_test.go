package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	var tmpl []Template
	var err error

	tmpl, err = Parse(`
author: hoge
`)
	assert.Nil(t, err)
	assert.Equal(t, len(tmpl), 1)
	assert.Equal(t, tmpl[0].String(), "author: hoge")

	tmpl, err = Parse(`
illegal: hoge
`)
	assert.NotNil(t, err)
	assert.Equal(t, len(tmpl), 0)

	tmpl, err = Parse(`
dir: hoge 755
`)
	assert.Nil(t, err)
	assert.Equal(t, len(tmpl), 1)
	assert.Equal(t, tmpl[0].String(), "dir: hoge(755)")

	tmpl, err = Parse(`
file: hoge 644
  <xml>
  </xml>
`)
	assert.Nil(t, err)
	assert.Equal(t, len(tmpl), 1)
	assert.Equal(t, tmpl[0].String(), "file: hoge(644)")
	content, ok := tmpl[0].GetContent()
	assert.Equal(t, ok, true)
	assert.Equal(t, content, "<xml>\n</xml>\n")

	tmpl, err = Parse(`
author: hoge
dir: hoge 755

file: hoge/fuga.html 644
  <html>
  </html>
`)
	assert.Nil(t, err)
	assert.Equal(t, len(tmpl), 3)
	assert.Equal(t, tmpl[0].String(), "author: hoge")
	assert.Equal(t, tmpl[1].String(), "dir: hoge(755)")
	assert.Equal(t, tmpl[2].String(), "file: hoge/fuga.html(644)")
	content, ok = tmpl[2].GetContent()
	assert.Equal(t, ok, true)
	assert.Equal(t, content, "<html>\n</html>\n")
}
