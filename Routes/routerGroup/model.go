package routerGroup

type apiGroup struct {
	User iUserGroup
}

var RoutesGroupCollection = apiGroup{
	User: newRouterGroup(),
}
