package handler

import "github.com/gin-gonic/gin"

const {
	authorizationHeader = "Authorization"
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, )
	}

}