package swagger

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/templatedop/api/diutil/typlect"
	"github.com/templatedop/api/modules/server/response"
	"github.com/templatedop/api/util/slc"
)

func buildDefinitions(eds []EndpointDef) m {
	defs := make(m)

	for _, ed := range eds {

		buildModelDefinition(defs, ed.RequestType, true)
		buildModelDefinition(defs, ed.ResponseType, false)
		buildModelDefinition(defs, reflect.TypeOf(response.ResponseError{}), false)
	}

	return defs
}

func buildModelDefinition(defs m, t reflect.Type, isReq bool) {

	if t == typlect.TypeNoParam {
		return
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	var smr []string
	smp := m{}
	for i := 0; i < t.NumField(); i++ {

		var (
			f = t.Field(i)

			ft = f.Type
		)

		if ft.Kind() == reflect.Uint64 {
			ft = reflect.TypeOf(int(0))
		}
		if ft == reflect.TypeOf(sql.NullString{}) {
		}

		if ft != typlect.TypeTime && ft.Kind() == reflect.Struct {
			buildModelDefinition(defs, ft, isReq)
		}

		if ft.Kind() == reflect.Slice && ft.Elem().Kind() == reflect.Struct {
			buildModelDefinition(defs, ft.Elem(), isReq)
		}

		if !isReq || f.Tag.Get("json") != "" {
			if f.Tag.Get("json") == "-" {
				continue
			}
			if f.Type == reflect.TypeOf(sql.NullString{}) {
			}

			smp[getFieldName(f)] = getPropertyField(f.Type)

			if vts, ok := f.Tag.Lookup("validate"); isReq && ok {
				if slc.Contains(strings.Split(vts, ","), "required") {
					smr = append(smr, getFieldName(f))
				}
			}
		}

	}

	if len(smp) > 0 {
		mi := m{
			"type":       "object",
			"properties": smp,
		}

		if len(smr) > 0 {
			mi["required"] = smr
		}

		defs[getNameFromType(t)] = mi
	}
}

func getFieldName(f reflect.StructField) string {
	if tag := f.Tag.Get("json"); tag != "" {
		return strings.Split(tag, ",")[0] // ignore ',omitempty'
	}

	return f.Name
}
