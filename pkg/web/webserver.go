package web

import (
	"fmt"
	"gameSrv/pkg/core"
	"gameSrv/pkg/log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Webapp struct {
	HttpMethod *core.HttpMethodWrap
}

func NewHttpServer() *Webapp {
	httpMethodController := &core.HttpMethodWrap{}
	httpMethodController.HttpInit()
	return &Webapp{HttpMethod: httpMethodController}
}

func (webApp *Webapp) WebAppStart(port int32) {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.0"})
	r.POST("/", func(c *gin.Context) {
		//inner rpc
		var rpcReq core.RpcReq
		if err := c.ShouldBind(&rpcReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		serviceName := strings.ToLower(rpcReq.ServiceName)
		methodName := strings.ToLower(rpcReq.MethodName)
		aresMethod := webApp.HttpMethod.GetCallFun(serviceName, methodName)
		if aresMethod == nil {
			c.JSON(403, "server or method not found")
			return
		}

		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		json.UnmarshalFromString(rpcReq.PayLoad, &realParam)
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})

	r.POST("/:serviceName/:methodName", func(c *gin.Context) {
		//inner restful or single server restful
		serviceName := strings.ToLower(c.Params.ByName("serviceName"))
		methodName := strings.ToLower(c.Params.ByName("methodName"))
		aresMethod := webApp.HttpMethod.GetCallFun(serviceName, methodName)
		if aresMethod == nil {
			c.JSON(403, "server or method not found")
			return
		}
		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		if err := c.ShouldBind(&realParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})

	r.POST("/game/:serviceName/:methodName", func(c *gin.Context) {
		//by proxy such as nginx

		serviceName := c.Params.ByName("serviceName")
		methodName := c.Params.ByName("methodName")
		aresMethod := webApp.HttpMethod.GetCallFun(serviceName, methodName)
		if aresMethod == nil {
			c.JSON(403, "server or method not found")
			return
		}
		param := reflect.New(aresMethod.ParamsType)
		realParam := param.Interface()
		if err := c.ShouldBind(&realParam); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ret := aresMethod.Invoke(realParam)
		c.JSON(200, ret[0].Elem().Interface())
	})
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Infof("start webserver:", addr)
	r.Run(addr)
}
