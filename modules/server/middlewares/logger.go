package middlewares

import (
	"github.com/templatedop/api/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(log *log.Logger) fiber.Handler {
	return logger.New(
		logger.Config{
			CustomTags: map[string]logger.LogFunc{
				"context": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
					cc := c.UserContext()
					return output.WriteString(cc.Value(RequestIDContextKey).(string))
				},
			},
			Format: "IP:${ip},Pid:${pid},Status:${status}, Method:${method},Path:${path},Latency:${latency}, Requestid:${context} ,BytesReceived:${bytesReceived}bytes, BytesSent:${bytesSent}bytes,Error:${error}\n",
			Output: log.ToZerolog(),
		})
}
