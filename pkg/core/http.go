package core

import (
	"gameSrv/pkg/log"
	"reflect"
	"strings"
)

type RpcReq struct {
	ServiceName string
	MethodName  string
	PayLoad     string
}

type AresMethod struct {
	MethodName string
	CallFun    reflect.Method
	ParamsType reflect.Type
	Param1     interface{}
}

func (aresMethod *AresMethod) Invoke(param interface{}) []reflect.Value {
	reValue := reflect.ValueOf(param)
	parmams := []reflect.Value{reflect.ValueOf(aresMethod.Param1), reValue}
	return aresMethod.CallFun.Func.Call(parmams)
}

type ServiceMethods struct {
	ServiceName    string
	AresMethodsMap map[string]*AresMethod
}

type HttpMethodInterface interface {
	GetCallFun(serviceName string, methodName string) *AresMethod
	RegisterController(controller interface{})
}

type HttpMethodWrap struct {
	ServiceMethodsMap map[string]*ServiceMethods
}

func (core *HttpMethodWrap) HttpInit() {
	core.ServiceMethodsMap = map[string]*ServiceMethods{}
}

func (core *HttpMethodWrap) RegisterController(controller interface{}) {
	rtype := reflect.TypeOf(controller)
	methods := rtype.NumMethod()
	log.Infof("============register service = %s   methods count = %d ============== begin", rtype.Elem().Name(), methods)
	for i := 0; i < methods; i++ {
		method := rtype.Method(i)
		var serviceName = strings.ToLower(rtype.Elem().Name())
		var methodName = strings.ToLower(method.Name)
		log.Infof("---------- method_name = %s   service_name = %s", methodName, serviceName)
		var val *ServiceMethods
		var ok bool
		if val, ok = core.ServiceMethodsMap[serviceName]; !ok {
			val = &ServiceMethods{}
			core.ServiceMethodsMap[serviceName] = val
		}
		aresMethod := AresMethod{}
		aresMethod.CallFun = method
		aresMethod.Param1 = controller
		aresMethod.MethodName = methodName
		//method params
		mt := method.Type
		numIn := mt.NumIn()
		for j := 0; j < numIn; j++ {
			param := mt.In(j)
			aresMethod.ParamsType = param.Elem()
		}
		val.addMethod(&aresMethod)
	}
	log.Infof("============register service = %s  methods count =  %d  ============== end", rtype.Elem().Name(), methods)
}

func (srviceMethods *ServiceMethods) addMethod(aresMethod *AresMethod) {
	if srviceMethods.AresMethodsMap == nil {
		srviceMethods.AresMethodsMap = map[string]*AresMethod{}
	}
	srviceMethods.AresMethodsMap[aresMethod.MethodName] = aresMethod
}

func (core *HttpMethodWrap) GetCallFun(serviceName string, methodName string) *AresMethod {
	var val *ServiceMethods
	var ok bool
	if val, ok = core.ServiceMethodsMap[serviceName]; !ok {
		return nil
	}
	if val.AresMethodsMap == nil {
		return nil
	}
	return val.AresMethodsMap[methodName]
}
