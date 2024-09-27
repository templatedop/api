package swagger

import (
	"encoding/json"
	"fmt"
	"os"

	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/templatedop/api/config"
	"github.com/templatedop/api/diutil/typlect"
)

type m map[string]any

type Docs m

func (d Docs) WithHost(h string) Docs {
	d["Host"] = h
	return d
}

const (
	refKey = "$ref"
)

func buildDocs(eds []EndpointDef, cfg *config.Config) *openapi3.T {
	dj := baseJSON(cfg)
	dj["definitions"] = buildDefinitions(eds)
	dj["paths"] = buildPaths(eds)

	var v2Doc openapi2.T
	data, err := json.Marshal(Docs(dj))
	if err != nil {
		return nil
	}

	if err := json.Unmarshal(data, &v2Doc); err != nil {
		return nil
	}
	v3Doc, err := openapi2conv.ToV3(&v2Doc)
	if err != nil {
		return nil
	}

	storeV3DocToFile(v3Doc)
	return v3Doc
}
func storeV3DocToFile(v3Doc *openapi3.T) error {
	v3DocJSON, err := json.MarshalIndent(v3Doc, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling v3Doc to JSON: %w", err)
	}

	if _, err := os.Stat("docs"); os.IsNotExist(err) {
		os.Mkdir("docs", os.ModePerm)
	}

	file, err := os.Create("./docs/v3Doc.json")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(v3DocJSON); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	return nil
}

var pathRegexp = regexp.MustCompile(`(\:[A-Za-z0-9_]*)`)

func toSwaggerPath(s string) string {
	return pathRegexp.ReplaceAllStringFunc(s, func(s string) string {
		return fmt.Sprintf("{%s}", s[1:])
	})
}

func baseJSON(cfg *config.Config) m {

	cfg.SetDefault("info.description", "")
	cfg.SetDefault("info.version", "1.1.0")
	cfg.SetDefault("info.title", "Application")
	cfg.SetDefault("info.terms", "http://swagger.io/terms/")
	cfg.SetDefault("info.email", "")
	of := cfg.Of("info")
	return m{
		"swagger": "2.0",
		"info": m{
			"description":    of.GetString("description"),
			"version":        of.GetString("version"),
			"title":          cfg.GetString("info.title"),
			"termsOfService": cfg.GetString("info.terms"),
			"contact": m{
				"email": cfg.GetString("info.email"),
			},
			"license": m{
				"name": "Apache 2.0",
				"url":  "http://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		"host":     "",
		"basePath": "/",
		"schemes":  []string{},
	}
}

func withDefinitionPrefix(s string) string {
	return fmt.Sprintf("#/definitions/%s", s)
}

func getPrimitiveType(t reflect.Type) m {
	if kp := t.Kind().String(); strings.HasPrefix(kp, "int") {
		return m{
			"type":   "integer",
			"format": kp,
		}
	}

	/* Added for other types compatability*/
	//added for uint64
	if kp := t.Kind().String(); strings.HasPrefix(kp, "uint64") {
		return m{
			"type":   "integer",
			"format": kp,
		}
	}

	//Add NullString
	k := t.Kind().String()

	if t.Kind() == reflect.Bool {
		k = "boolean"
	}

	return m{
		"type": k,
	}
}

func getPropertyField(t reflect.Type) m {

	if t == typlect.TypeNoParam {
		return m{"type": "string"}
	}

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t == typlect.TypeTime {
		b, _ := time.Now().MarshalJSON()
		return m{"type": "string", "example": strings.Trim(string(b), "\"")}
	}

	if t.Kind() == reflect.Struct {
		return m{
			refKey: withDefinitionPrefix(getNameFromType(t)),
		}
	}

	if t.Kind() == reflect.Slice {
		return arrayProperty(t)
	}

	return getPrimitiveType(t)
}

func arrayProperty(t reflect.Type) m {
	it := t.Elem()
	if it.Kind() == reflect.Pointer {
		it = it.Elem()
	}

	return m{
		"type":  "array",
		"items": getPropertyField(it),
	}
}

func getNameFromType(t reflect.Type) string {
	s := strings.ReplaceAll(t.Name(), "]", "")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "*", "")
	return strings.ReplaceAll(s, "[", "__")
}
