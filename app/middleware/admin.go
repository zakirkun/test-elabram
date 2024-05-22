package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zakirkun/test-elabram/app/domain/types"
	"github.com/zakirkun/test-elabram/internal/pkg/config"
	"github.com/zakirkun/test-elabram/internal/pkg/jwt"
	"github.com/zakirkun/test-elabram/internal/utils"
)

func AdminMiddleware(ctx *gin.Context) {
	var header types.GeneralHeader
	var token string
	if err := ctx.ShouldBindHeader(&header); err != nil {

		ctx.JSON(http.StatusUnauthorized, utils.SetErrorResponse("ERROR", "Unauthorized", uuid.NewString()))
		ctx.Abort()
		return
	}

	bearerSplit := strings.Split(header.Authorization, " ")

	if len(bearerSplit) == 0 {
		token = bearerSplit[0]
	} else {
		token = bearerSplit[1]
	}

	_jwt := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))

	ok, err := _jwt.ValidateToken(token)
	if err != nil {

		ctx.JSON(http.StatusUnauthorized, utils.SetErrorResponse("ERROR", "Unauthorized", uuid.NewString()))
		ctx.Abort()
		return
	}

	if !ok {

		ctx.JSON(http.StatusUnauthorized, utils.SetErrorResponse("ERROR", "Unauthorized", uuid.NewString()))
		ctx.Abort()
		return
	}

	extract, _ := _jwt.ParseToken(token)
	_acc_type := extract["acc_type"].(string)
	id := extract["id"].(string)

	if _acc_type != "admin" {

		ctx.JSON(http.StatusUnauthorized, utils.SetErrorResponse("ERROR", "Unauthorized", uuid.NewString()))
		ctx.Abort()
		return
	}

	ctx.Set("_id", id)
	ctx.Next()
}
