package api

import "C"
import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/hl/hl_bank/db/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
	Balance  int64  `json:"balance"`
}

func (server *Server) createAccount(c *gin.Context) {
	var request createAccountRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    request.Owner,
		Currency: request.Currency,
		Balance:  request.Balance,
	}
	account, err := server.store.CreateAccount(c, arg)

	if err != nil {
		fmt.Println("err is*****:", err.(*pgconn.PgError).Detail)

		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {

			switch pgError.Code {
			case db.ForeignKeyViolation, db.UniqueKeyViolation:
				c.JSON(http.StatusForbidden, errorResponse(err))
				return

			}
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(c *gin.Context) {
	var request getAccountRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(c, request.ID)

	if err != nil {
		fmt.Println("getAccountErr", err)
		if errors.Is(err, db.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

	}

	c.JSON(http.StatusOK, account)

}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{

		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
