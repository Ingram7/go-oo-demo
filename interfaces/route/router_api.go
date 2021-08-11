package route

func (router *Router) Init() {
	//router.initCaptcha()
	//router.initUser()
}

//func (router *Router) initCaptcha() {
//	ctr := controller.NewCaptcha(repository.NewCaptcha(router.db))
//	captcha := router.engine.Group("/captcha")
//	{
//		captcha.GET("/send", router.wrapper(ctr.Send))
//	}
//}
//
//func (router *Router) initUser() {
//	ctr := controller.NewUser(repository.NewUser(router.db), repository.NewCaptcha(router.db))
//	user := router.engine.Group("/user")
//	{
//		user.GET("/login", router.wrapper(ctr.Login))
//	}
//}
