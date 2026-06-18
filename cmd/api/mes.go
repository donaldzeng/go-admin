package api

import (
	mesRouter "mes-app/app/mes/router"
	otherRouter "mes-app/app/other/router"
	jobsRouter "mes-app/app/jobs/router"
)

func init() {
	// 注册 mes-app 业务路由
	AppRouters = append(AppRouters,
		mesRouter.InitRouter,
		otherRouter.InitRouter,
		jobsRouter.InitRouter,
	)
}
