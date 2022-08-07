package gen_test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/inspirer/textmapper/tm-go/gen"
	"github.com/inspirer/textmapper/tm-go/status"
	"github.com/inspirer/textmapper/tm-go/util/diff"
)

var grammars = []string{
	"../parsers/json/json.tm",
	"../parsers/simple/simple.tm",
	"../parsers/test/test.tm",
	"../../tm-parsers/tm/textmapper.tm",
	"../../tm-parsers/js/js.tm",
}

type mapWriter struct {
	files   []string
	content map[string]string
}

func (w *mapWriter) Write(filename, content string) error {
	w.files = append(w.files, filename)
	w.content[filename] = content
	return nil
}

func TestGenerate(t *testing.T) {
	for _, filename := range grammars {
		filename := filename
		t.Run(filename, func(t *testing.T) {
			w := &mapWriter{content: make(map[string]string)}
			err := gen.GenerateFile(filename, w, true /*compat*/)
			if err != nil {
				s := status.FromError(err)
				s.Sort()
				for _, err := range s {
					t.Errorf("GenerateFile() failed with %v", err)
				}
				return
			}

			for _, genfile := range w.files {
				content := w.content[genfile]
				if strings.HasSuffix(genfile, ".y") {
					// TODO compare final grammars
					continue
				}

				p := filepath.Join(filepath.Dir(filename), genfile)
				ondisk, err := ioutil.ReadFile(p)
				if err != nil {
					t.Errorf("ReadFile(%v) failed with %v", genfile, err)
					continue
				}
				if p == "../parsers/test/parser.go" {
					// TODO correctly handle semantic actions
					continue
				}
				t.Logf("comparing %v", p)
				if diff := diff.LineDiff(string(ondisk), content); diff != "" {
					t.Errorf("The on-disk content differs from the generated one.\n--- %v\n+++ %v (generated)\n%v", p, genfile, diff)
				}
			}
		})
	}
}
