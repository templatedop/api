package swagger

import (
	"database/sql"
	//"fmt"
	"reflect"
	"strings"

	//"database/sql"

	"github.com/templatedop/api/diutil/typlect"
	"github.com/templatedop/api/modules/server/response"
	"github.com/templatedop/api/util/slc"
)

func buildDefinitions(eds []EndpointDef) m {
	defs := make(m)

	//fmt.Println("eds:",eds)

	for _, ed := range eds {
		// fmt.Println("ED endpoint:", ed.Endpoint)
		// fmt.Println("ED RequestType:", ed.RequestType)
		// fmt.Println("ED ResponseType:", ed.ResponseType)
		// fmt.Println("ED Method:", ed.Method)
		// fmt.Println("ED Group:", ed.Group)
		// fmt.Println("ED Name:", ed.Name)

		buildModelDefinition(defs, ed.RequestType, true)
		buildModelDefinition(defs, ed.ResponseType, false)
		buildModelDefinition(defs, reflect.TypeOf(response.ResponseError{}), false)
	}

	return defs
}

func buildModelDefinition(defs m, t reflect.Type, isReq bool) {
	// fmt.Println("Starting of buildModelDefinition")
	// fmt.Println("defs: ", defs)
	// fmt.Println("t: ", t)
	// fmt.Println("t kind: ", t.Kind())
	// fmt.Println("t Name: ", t.Name())

	// fmt.Println("isReq: ", isReq)
	// fmt.Println("Ends of buildModelDefinition")

	if t == typlect.TypeNoParam {
		return
	}

	//fmt.Println("t NumIn: ", t.NumIn())
	//fmt.Println("t Numout: ", t.NumOut())

	if t.Kind() == reflect.Slice {
		//fmt.Println("t elem: ", t.Elem())
		t = t.Elem()
	}

	if t.Kind() == reflect.Pointer {
		//fmt.Println("t elem: ", t.Elem())
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	// typeMapping := map[reflect.Type]reflect.Type{
    //     reflect.TypeOf(sql.NullString{}): reflect.TypeOf(""),
    //     reflect.TypeOf(sql.NullInt64{}):  reflect.TypeOf(0),
    //     reflect.TypeOf(sql.NullFloat64{}): reflect.TypeOf(0.0),
    //     reflect.TypeOf(sql.NullBool{}):   reflect.TypeOf(true),
    // }

	var smr []string
	smp := m{}
	for i := 0; i < t.NumField(); i++ {

		//fmt.Println("i: ", i)
		//fmt.Println("t.NumField(): ", t.NumField())

		var (
			f = t.Field(i)

			ft = f.Type
		)
		// if basicType, ok := typeMapping[ft]; ok {
        //     fmt.Println("Converting special type: ", f)
        //     ft = basicType
        // }

		if ft.Kind() == reflect.Uint64 {
			//fmt.Println("came inside uint64: ", f)
			ft = reflect.TypeOf(int(0))
		}
		if ft == reflect.TypeOf(sql.NullString{}) {
			//fmt.Println("came inside Nulstring: ", f)
			//ft = reflect.TypeOf(string)
		}

		// build subtype definitions
		if ft != typlect.TypeTime && ft.Kind() == reflect.Struct {
			buildModelDefinition(defs, ft, isReq)
		}

		if ft.Kind() == reflect.Slice && ft.Elem().Kind() == reflect.Struct {
			buildModelDefinition(defs, ft.Elem(), isReq)
		}

		if !isReq || f.Tag.Get("json") != "" {
			//fmt.Println("FieldName: ", getFieldName(f))
			//fmt.Println("fname: ", f.Name)
			// fmt.Println("ftype: ",f.Type)
			// fmt.Println("ftype kind: ",f.Type.Kind())
			// fmt.Println("ftype name: ",f.Type.Name())
			//fmt.Println("Tag: ", f.Tag.Get("json"))
			if f.Tag.Get("json") == "-" {
				continue
			}
			//fmt.Println("f type:", f.Type)
			if f.Type == reflect.TypeOf(sql.NullString{}) {
				//fmt.Println("Name inside nullstring:", f.Name)
				//fmt.Println("came inside Nulstring: ", f)
				//f.Type = reflect.TypeOf("")
				//f.Name = "string"
				
				//fmt.Println("After changing type inside Nullstring: ", f)
			}

			smp[getFieldName(f)] = getPropertyField(f.Type)

			if vts, ok := f.Tag.Lookup("validate"); isReq && ok {
				if slc.Contains(strings.Split(vts, ","), "required") {
					smr = append(smr, getFieldName(f))
				}
			}
		}

		//fmt.Println("f:", f, "ft:", ft)
	}

	if len(smp) > 0 {
		mi := m{
			"type":       "object",
			"properties": smp,
		}

		if len(smr) > 0 {
			mi["required"] = smr
		}

		//fmt.Println("getNameFromType(t): ", getNameFromType(t))

		defs[getNameFromType(t)] = mi
	}
}

func getFieldName(f reflect.StructField) string {
	if tag := f.Tag.Get("json"); tag != "" {
		return strings.Split(tag, ",")[0] // ignore ',omitempty'
	}

	return f.Name
}
