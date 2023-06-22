package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Route{
	// SELECT All
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Function:    controllers.SelectPosts,
		AuthRequire: false,
	},
	// SELECT One
	{
		URI:         "/post/{postId}",
		Method:      http.MethodGet,
		Function:    controllers.SelectPost,
		AuthRequire: true,
	},
	// SEARCH *
	{
		URI:         "/post",
		Method:      http.MethodGet,
		Function:    controllers.SearchPost,
		AuthRequire: true,
	},
	// INSERT
	{
		URI:         "/post/{postId}",
		Method:      http.MethodPost,
		Function:    controllers.InsertPost,
		AuthRequire: false,
	},
	// UPDATE
	{
		URI:         "/post/{postId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePost,
		AuthRequire: true,
	},
	// DELETE
	{
		URI:         "/post/{postId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeletePost,
		AuthRequire: true,
	}}
