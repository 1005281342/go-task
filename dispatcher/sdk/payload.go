package sdk

type Payload struct {
	ServiceName string      `json:"service_name"` // 目标服务名
	MethodName  string      `json:"method_name"`  // 目标方法名
	Body        interface{} `json:"body"`
}
