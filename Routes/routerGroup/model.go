package routerGroup

type apiGroup struct {
	User  iUserGroup
	Order iOrderGroup
}

var RoutesGroupCollection = apiGroup{
	User:  newUserRouterGroup(),
	Order: newOrderRouterGroup(),
}
