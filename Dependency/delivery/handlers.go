package delivery

import (
	"aqary/usecase"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Usecase usecase.UsecaseMethods
}

type HandlerMethods interface {
	RateLimitCheck(*gin.Context)
}

func InitHandlers(usecase usecase.UsecaseMethods) Handlers {
	return Handlers{
		Usecase: usecase,
	}

}

func (h Handlers) RateLimitCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Resource created successfully"})
}
