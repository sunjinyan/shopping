package svc

const HelloServiceName = "svc/HelloService"


type HelloService struct {

}

//可以使用反射的方式来将所有服务注册到map中
func (h *HelloService)Hello	(input string,output *string) error  {
	*output = "Hello " + input

	return  nil
}
