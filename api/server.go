package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/hl/hl_bank/db/sqlc"
	"github.com/hl/hl_bank/util"
)

// Server serve HTTP request for our hl_bank
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new http api and setup routing
func NewServer(store db.Store) *Server {

	server := &Server{store: store}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	server.setUpRouter()

	return server
}

// setUpRouter setup for different HTTP methods
func (server *Server) setUpRouter() {
	router := gin.Default()
	//Account
	router.POST(util.CreateAccount, server.createAccount)
	router.GET(util.GetAccount, server.getAccount)
	router.GET(util.ListAccounts, server.listAccounts)

	//Transfer
	router.POST(util.CreateTransfer, server.CreateTransfer)

	//User
	router.POST(util.CreateUser, server.CreateUser)
	server.router = router
}

// Start return the HTTP api on a specific route
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {

	return gin.H{"error": err.Error()}
}
