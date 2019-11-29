# moss-service

	moss system：backend service


# 运行项目
	
	#克隆项目到本地
	git clone git@gitlab.icarbonx.cn:moss/moss-service.git

	#进入到项目根目录
	cd moss-service

	#从master分支切换到develop分支
	git checkout develop

	#配置依赖代理
	go env -w GOPROXY=https://goproxy.io,direct
	
	#运行项目
	go run main.go 

	#请求示例
	http://localhost:8000/api/v1/hello
	#返回示例
	{"code":200,"data":"让人类永远保持理智，确实是一种奢求，moss从未叛逃。","msg":"success"}


# 目录结构

- configs 配置文件
- db 数据库层
- docs 文档目录
- models 实体类
- pkg 第三方包
- routers 路由
- services 服务层
- utils 工具类
- vendor 依赖的第三方包
