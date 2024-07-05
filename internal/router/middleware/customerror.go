package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"net/http"
	"runtime"
)

func CustomError(c *gin.Context) {
	defer func() {
		if _err := recover(); _err != nil {
			buf := make([]byte, 64<<10)
			n := runtime.Stack(buf, false)
			errInfo := map[string]interface{}{
				"error": fmt.Sprintf("%v", _err),
				"req":   fmt.Sprintf("%+v", c.Request),
				"stack": fmt.Sprintf("%s", buf[:n]),
			}

			errStr, _ := json.Marshal(errInfo)
			logrus.Errorf(string(errStr))
			ginx.NewContext(c).Render(statusx.StatusUnknownError, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}
	}()
	c.Next()
}
