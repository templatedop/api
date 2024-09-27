package swagger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/getkin/kin-openapi/openapi3"
)

func generatejson(v3 *openapi3.T) {
	file, err := os.Open("./docs/v3Doc.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	traverseAndReplaceRefs(jsonParsed, jsonParsed)
	replaceDataType(jsonParsed, "NullString", "string")
	wrap200Responses(jsonParsed)
	nullStringPath := "components.schemas.NullString"
	if jsonParsed.ExistsP(nullStringPath) {
		jsonParsed.DeleteP(nullStringPath)
	}

	err = ioutil.WriteFile("./docs/resolved_swagger.json", []byte(jsonParsed.StringIndent("", "  ")), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func replaceDataType(container *gabs.Container, targetType, newType string) {
	fmt.Println("container details: ", container)

	if container.Exists("schemas") {
		fmt.Println("schema is: ", container.Path("schema"))
	}

	children, _ := container.ChildrenMap()

	for key, child := range children {
		if key == "properties" && child.Exists(targetType) {
			container.Delete("properties") // Remove "properties" key
			container.Set(newType, "type") // Set "type" to the new type (e.g., "string")
		} else if key == targetType {
			container.Set(newType, key)
		} else if key == "type" && child.Data().(string) == targetType {
			container.Set(newType, "type")
		} else if child != nil {

			replaceDataType(child, targetType, newType)
		}
	}
}

func traverseAndReplaceRefs(container, root *gabs.Container) {
	children, _ := container.ChildrenMap()

	for key, child := range children {
		if key == "$ref" {
			refPath := child.Data().(string)
			if schema := resolveSchema(refPath, root); schema != nil {
				container.Merge(schema)
				container.Delete("$ref")
			}
		} else if child != nil {
			traverseAndReplaceRefs(child, root)
		}
	}
}

func resolveSchema(refPath string, root *gabs.Container) *gabs.Container {
	if strings.HasPrefix(refPath, "#/components/schemas/") {
		// Extract the schema name
		schemaName := refPath[len("#/components/schemas/"):]
		resolved := root.Path(fmt.Sprintf("components.schemas.%s", schemaName))
		if resolved.Exists() {
			return resolved
		}
	}
	return nil
}

func wrap200Responses(container *gabs.Container) {
	paths := container.Path("paths")
	if paths == nil {
		fmt.Println("No paths found in the Swagger document.")
		return
	}

	pathsMap, _ := paths.ChildrenMap()
	for _, pathData := range pathsMap {
		methods, _ := pathData.ChildrenMap()

		for _, methodData := range methods {
			responses := methodData.Path("responses")
			if responses != nil {
				response200 := responses.Path("200")
				if response200.Exists() {
					existingSchema := response200.Path("content.application/json.schema")
					if existingSchema.Exists() {

						successResponse := map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"success": map[string]interface{}{
									"type":    "boolean",
									"example": true,
								},
								"message": map[string]interface{}{
									"type":    "string",
									"example": "success",
								},
								"data": existingSchema.Data(),
							},
						}

						response200.Set(successResponse, "content", "application/json", "schema")
					}
				}
			}
		}
	}
}
