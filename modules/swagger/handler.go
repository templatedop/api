package swagger

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strings"

	//"github.com/getkin/kin-openapi/openapi2"
	//"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/templatedop/api/modules/server/common"
	"github.com/templatedop/api/modules/swagger/files"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/redirect"
)
//*openapi3.T
//func fiberWrapper(docs Docs) common.FiberAppWrapper {
	func fiberWrapper(v3Doc *openapi3.T) common.FiberAppWrapper {
	return func(a *fiber.App) *fiber.App {
	



		a.Use(
			newRedirectMiddleware(),
			newMiddleware(v3Doc),
		)
		return a
	}
}

func newRedirectMiddleware() fiber.Handler {
	return redirect.New(redirect.Config{
		Rules: map[string]string{
			"/":                        "/swagger/index.html",
			"/swagger":                 "/swagger/index.html",
			"/swagger.json":            "/swagger/docs.json",
			"/swagger/v1/swagger.json": "/swagger/docs.json",
		},
	})
}

// func (doc *openapi3.T) WithHost(host string) *openapi3.T {
//     doc.Servers = []*openapi3.Server{
//         {
//             URL: fmt.Sprintf("http://%s", host),
//         },
//     }
//     return doc
// }

// func (d *openapi3.T) WithHost(h string) *openapi3.T {
// 	d["Host"] = h
// 	return d
// }


func attachHostToV3Doc(doc *openapi3.T, host string) *openapi3.T {
    doc.Servers = []*openapi3.Server{
        {
            URL: fmt.Sprintf("http://%s", host),
        },
    }
    return doc
}

//func newMiddleware(docs Docs) fiber.Handler {
	func newMiddleware(v3Doc *openapi3.T) fiber.Handler {
	fscfg := filesystem.ConfigDefault
	fscfg.Root = http.FS(files.Files)
	fsmw := filesystem.New(fscfg)

	prefix := "/swagger"

	return func(c *fiber.Ctx) error {
		if c.Path() == "/swagger/docs.json" || c.Path() == "/swagger/docs.json/" {
			//return c.JSON(docs.WithHost(c.Hostname()))
			v3Doc=attachHostToV3Doc(v3Doc, c.Hostname())

			return c.JSON(v3Doc)
		}

		if strings.HasPrefix(c.Path(), prefix) {
			c.Path(strings.TrimPrefix(c.Path(), prefix))
			return fsmw(c)
		}

		return c.Next()
	}
}

