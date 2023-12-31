package record

const (
	OperationOperateCreateFolder     = 1  // 创建文件夹
	OperationOperateCreateAPI        = 2  // 创建接口
	OperationOperateCreatePlan       = 4  // 创建计划
	OperationOperateCreateScene      = 5  // 创建场景
	OperationOperateUpdateFolder     = 6  // 修改文件夹
	OperationOperateUpdateAPI        = 7  // 修改接口
	OperationOperateUpdatePlan       = 9  // 修改计划
	OperationOperateUpdateScene      = 10 // 修改场景
	OperationOperateClonePlan        = 11 // 克隆计划
	OperationOperateDeleteReport     = 12 // 删除报告
	OperationOperateDeleteScene      = 13 // 删除场景
	OperationOperateDeletePlan       = 14 // 删除计划
	OperationOperateRunScene         = 15 // 运行场景
	OperationOperateRunPlan          = 16 // 运行计划
	OperationOperateSavePreinstall   = 17 // 新建预设配置
	OperationOperateUpdatePreinstall = 18 // 修改并保存预设配置
	OperationOperateDeletePreinstall = 19 // 删除预设配置
	OperationOperateRunSceneCase     = 20 // 调试场景用例
	OperationOperateUpdateSceneCase  = 21 // 修改场景用例
	OperationOperateSaveEnv          = 22 // 新建环境
	OperationOperateUpdateEnv        = 23 // 编辑环境
	OperationOperateDeleteEnv        = 24 // 删除环境
	OperationOperateCopyEnv          = 25 // 克隆环境
	OperationOperateDeleteEnvService = 26 // 删除环境服务
	OperationOperateDeleteFolder     = 27 // 删除文件夹
	OperationOperateDeleteApi        = 28 // 删除接口
	OperationOperateDeleteTestCase   = 30 // 删除场景用例
	OperationOperateExecPlan         = 31 // 执行计划
	OperationOperateDebugApi         = 32 // 调试接口
	OperationOperateDebugScene       = 33 // 调试场景
	OperationOperateCreateTestCase   = 34 // 创建场景用例
	OperationLogCreateSql            = 35 // 创建sql语句
	OperationLogUpdateSql            = 36 // 修改sql语句
	OperationLogDeleteSql            = 37 // 删除sql语句
	OperationLogCreateDatabase       = 38 // 新建数据库
	OperationLogDeleteDatabase       = 39 // 删除数据库
	OperationLogUpdateDatabase       = 40 // 修改数据库
	OperationLogCreateService        = 41 // 新建服务
	OperationLogUpdateService        = 42 // 修改服务
	OperationLogCreateTcp            = 43 // 创建websocket
	OperationLogUpdateTcp            = 44 // 修改websocket
	OperationLogDeleteTcp            = 45 // 删除websocket
	OperationLogCreateWebsocket      = 46 // 创建websocket
	OperationLogUpdateWebsocket      = 47 // 修改websocket
	OperationLogDeleteWebsocket      = 48 // 删除websocket
	OperationLogCreateMqtt           = 49 // 创建mqtt
	OperationLogUpdateMqtt           = 50 // 修改mqtt
	OperationLogDeleteMqtt           = 51 // 删除mqtt
	OperationLogCreateDubbo          = 52 // 创建dubbo
	OperationLogUpdateDubbo          = 53 // 修改dubbo
	OperationLogDeleteDubbo          = 54 // 删除dubbo

	OperationOperateCreateMockFolder = 101 // 创建mock文件夹
	OperationOperateCreateMockAPI    = 102 // 创建mock接口
	OperationOperateUpdateMockFolder = 103 // 修改mock文件夹
	OperationOperateUpdateMockAPI    = 104 // 修改mock接口
	OperationOperateDebugMockApi     = 105 // 调试mock接口
	OperationOperateDeleteMockFolder = 106 // 删除mock文件夹
	OperationOperateDeleteMockApi    = 107 // 删除mock接口
	OperationOperateMockTarget       = 108 // mock接口保存至测试对象
	OperationOperateMockSyncApi      = 109 // mock接口同步测试对象
)
