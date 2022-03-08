package handler

import (
	"fmt"
	"strings"

	"github.com/DeanThompson/ginpprof"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"jzsg.com/mca/core/server/config"
	"jzsg.com/mca/core/utils"
)

func createRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.POST("/login", Login) //登录

	web := v1.Group("web")

	{
		web.Use(utils.TokenAuthMiddleware())
		web.POST("/register", Register)                    //超管注册操作员
		web.PATCH("/user/pass", PassReset)                 //超管重置密码
		web.POST("/user/role", CreateRole)                 //超管创建角色
		web.POST("/user/user_role", CreateUserRole)        //超管给用户绑定一角色
		web.POST("/user/role/role_users", AddUsersForRole) //超管给某个角色添加多用户
		web.POST("/user/role/user_roles", AddRolesForUser) //超管给某个角色添加多用户

		wf := web.Group("workflow")
		wf.POST("/", CreateCertWF)              //超管创建申请流程
		wf.POST("/node/role", CreateWfNodeRole) //创建流程节点操作角色
	}

	//v1.POST("/report", handler.Report)
	//v1.POST("/syncfile", handler.SyncFile)          //从非密区接收校验后的密钥
	//v1.POST("/sync/single", handler.SyncSingleFile) //从非密区接收校验后的密钥
	//v1.POST("/sync/area", handler.SyncCryptoArea) //从非密区接收校验后的区块
	//
	//v1.POST("/relay", handler.RelayCrypto) //接收密钥中继请求
	//
	//v1.GET("/balance", handler.GetCryptoBalance)
	//v1.GET("/balance/detail", handler.GetCryptoBalanceDetail)
	//
	////v1.POST("/consistency/check", handler.CheckConsistency) //单文件校验
	////v1.POST("/block/check", handler.BlockCheck) //密钥区校验
	//v1.POST("/consistency/area", handler.SingleAreaCheck)
	//v1.POST("/consistency/mix", handler.MixInfoCheck)
	////v1.POST("/area/check", handler.AreaCheck) //密钥区校验
	//v1.POST("/crypto/mix", handler.MixCrypto) //密钥杂糅操作
	//
	//v1.PUT("/loglevel", handler.SetLogLevel) //设置日志level
	//
	//if config.GetLocalConfig().Level == 6 {
	//	v1.POST("/encrypt", handler.EncryptBytes) //执行加密算法
	//	v1.POST("/decrypt", handler.DecryptBytes) //执行解密算法
	//}
	//
	//if config.GetLocalConfig().Level != 6 {
	//	//v1.POST("/report", handler.Report) //接收密钥上报请求
	//}

	return r
}

func InitRouter() error {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = config.GetDefaultLogWriter()

	r := createRouter()
	ginpprof.Wrapper(r) // for debug

	//log.Fatal()
	listenPort := viper.GetInt("local.port")
	if listenPort == 0 {
		listenPort = 8888
	}
	logger.Infof("The listen port is %d", listenPort)
	server := endless.NewServer(fmt.Sprintf(":%d", listenPort), r)

	//// save pid file
	//server.BeforeBegin = func(add string) {
	//	pid := syscall.Getpid()
	//	fmt.Println("Actual pid is ", pid)
	//	pidFile := "server.pid"
	//	if utils.CheckFileIsExist(pidFile) {
	//		os.Remove(pidFile)
	//	}
	//	if err := ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", pid)), 0666); err != nil {
	//		fmt.Printf("Api server write pid file failed! err:%v", err)
	//	}
	//}

	err := server.ListenAndServe()
	if err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			fmt.Println(err)
		} else {
			fmt.Printf("Api server start failed! err:%v", err)
			panic(err)
		}
		return err
	}
	return nil
}
