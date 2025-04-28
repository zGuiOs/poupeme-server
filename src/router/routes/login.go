package routes

import (
	"net/http"

	"github.com/zGuiOs/poupeme-server/src/controllers"
)

var loginRoutes = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Handler:  controllers.Login,
	NeedAuth: false,
}
