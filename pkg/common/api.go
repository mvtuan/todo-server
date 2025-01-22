package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiStatusType string
type statusType struct {
	Ok                  apiStatusType
	Invalid             apiStatusType
	Unauthorized        apiStatusType
	NotFound            apiStatusType
	InternalServerError apiStatusType
}

var APIStatus = statusType{
	Ok:                  "OK",
	Invalid:             "INVALID",
	Unauthorized:        "UNAUTHORIZED",
	NotFound:            "NOT_FOUND",
	InternalServerError: "INTERNAL_SERVER_ERROR",
}

type APIResponse struct {
	Status    apiStatusType `json:"status,omitempty"`
	Data      interface{}   `json:"data,omitempty"`
	Message   string        `json:"message,omitempty"`
	ErrorCode string        `json:"errorCode,omitempty"`
	RootCause error         `json:"-"`
}

func Respond(c *gin.Context, res *APIResponse) {
	switch res.Status {
	case APIStatus.Ok:
		c.JSON(http.StatusOK, res)
	case APIStatus.Invalid:
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
	case APIStatus.Unauthorized:
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
	case APIStatus.NotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, res)
	case APIStatus.InternalServerError:
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
	}
}
