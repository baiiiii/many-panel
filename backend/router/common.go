package router

func commonGroups() []CommonRouter {
	return []CommonRouter{
		&BaseRouter{},
		&MpHostRouter{},
		&SettingRouter{},
	}
}
