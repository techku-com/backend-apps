package routerGroup

type apiGroup struct {
	User  iUserGroup
	Order iOrderGroup
	Home  iHomeGroup
}

var RoutesGroupCollection = apiGroup{
	User:  newUserRouterGroup(),
	Order: newOrderRouterGroup(),
	Home:  newHomeRouterGroup(),
}
