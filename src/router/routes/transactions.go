package routes

import (
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/controllers"
)

var transactionsRoutes = []Route{
	{
		URI:      "/transactions/{UUID}",
		Method:   http.MethodPost,
		Handler:  controllers.CreateTransaction,
		NeedAuth: true,
	},
	{
		URI:      "/transactions/{UUID}",
		Method:   http.MethodGet,
		Handler:  controllers.FetchTransactions,
		NeedAuth: true,
	},
}
