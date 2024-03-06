package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/hl/hl_bank/db/sqlc"
	"github.com/hl/hl_bank/util"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"fullName" binding:"required"`
}

type createUserResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	FullName          string    `json:"fullName"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) CreateUser(ctx *gin.Context) {

	fmt.Println("createUser...")

	var req createUserRequest

	fmt.Println("Req", req)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashPassword, err := util.HashPassword(req.Password)

	// fmt.Println("hashPassword", hashPassword, err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
		FullName: req.FullName,
	}

	user, err := server.store.CreateUser(ctx, args)

	// fmt.Println("user is:", user)

	if err != nil {
		var pgError *pgconn.PgError

		if errors.As(err, &pgError) {
			switch db.UniqueKeyViolation {
			case db.UniqueKeyViolation:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return

			}
		}
	}

	ur := createUserResponse{
		Username:          user.Username,
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, ur)

}
