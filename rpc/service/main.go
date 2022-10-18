package main

import (
   "encoding/json"
   "fmt"
   "net"
   "net/http"
   "reflect"
   "strings"
)

type ControllerMapType map[string]reflect.Value

var RegisterControllerContainer ControllerMapType

type Calc struct {

}

func (c *Calc)Add(a, b int) int {
   total := a + b

   return total
}

type Company struct {
   Name string
   Addr string
}

type Employee struct {
   Name string
   Company Company
}

type PrintResult struct {
   Info string
   Err error
}

func RpcPrintln(employee Employee) {
   /**
   客户端
   1、建立链接，TCP、HTTP
   2、将employee对象序列化成json字符串 交给服务器去打印 序列化
   3、发送json字符串
   4、等待服务器发送结果
   5、将服务器返回的数据解析成PrintResult  反序列化
   服务端
   1、监听80端口
   2、读取二进制的json数据
   3、对数据反序列化为Employee对象
   4、处理业务逻辑
   5、将处理的结果PrintResult序列化为json二进制数据
   6、将数据返回
    */


}


type Handl struct {


}

func (h Handl) ServeHTTP(res http.ResponseWriter,req *http.Request)()  {
   //body, err := ioutil.ReadAll(req.Body)
   query := req.URL.Query()
   //if err != nil {
   //   panic(err)
   //}
   //fmt.Println(body)
   fmt.Println(query)

   employee := Employee{
      Name: "Body",
      Company: Company{
         Name: "Mooc",
         Addr: "BeiJing",
      },
   }

   e, err := json.Marshal(employee)
   if err != nil {
     panic(err)
   }

   pr := &PrintResult{
      Info: string(e),
      Err:  nil,
   }
   p, err := json.Marshal(pr)
   if err != nil {
      panic(err)
   }

   write, err := res.Write(p)
   if err != nil {
      panic(err)
   }
   fmt.Println(write)

   //err = req.ParseForm()
   fmt.Println(req.URL.Path,req.URL.RawQuery)

   rpc := strings.Trim(req.URL.Path,"/")
   pathName := strings.ToUpper(rpc)
   firstName := pathName[0:1]
   funcName := firstName + rpc[1:]
   fmt.Println(rpc,pathName,funcName,firstName)
   //funcName()


   //注册路由
   var calc Calc

   ruterMap := make(ControllerMapType,0)

   vf := reflect.ValueOf(&calc)

   vft := vf.Type()
   mNum := vf.NumMethod()

   //vf.MethodByName("TMethod").Type().In()


   fmt.Printf("there are %v methods in calc params\n",mNum)

   //首先需要注册类，之后再遍历注册类中的每个方法
   //当有请求的时候，可以使用便利来进行循环匹配，如果匹配到就进行远程过程调用，如果未找到就报错
   for i := 0; i < mNum; i++ {
      mName := vft.Method(i).Name
      fmt.Println("index:",i," MethodName:",mName)
      ruterMap[mName] = vf.Method(i)
      fmt.Println("Method ",mName,"need Params num is :",vft.Method(i).Type.NumIn()) //返回func类型的参数个数，如果不是函数
      /*
      vType:=reflect.TypeOf(add)
      numIn:=vType.NumIn() //返回func类型的参数个数，如果不是函数，将会panic
      addIn:=make([]reflect.Type，numIn)
      for i:=0;i<numIn;i++{
         addIn[i]=vType.In(i) //返回func类型的第i个参数的类型，如非函数或者i不在[0, NumIn())内将会panic
         fmt.Println(addIn[i])
      }
      */
      /**

      透视函数类型，需要以下方法：

      reflect.Type.NumIn()：获取函数参数个数；
      reflect.Type.In(i)：获取第i个参数的reflect.Type；
      reflect.Type.NumOut()：获取函数返回值个数；
      reflect.Type.Out(i)：获取第i个返回值的reflect.Type。
       */
   }

   //演示
   testA := 1
   testB := 2
   //创建带调用方法时需要传入的参数列表
   parms := []reflect.Value{
      reflect.ValueOf(testA),
      reflect.ValueOf(testB),
   }
   //parm1 := []reflect.Value{reflect.ValueOf(testB)}
   //parms =  append(parms,parm1...)
   //使用方法名字符串调用指定方法
   call := ruterMap[funcName].Call(parms)

   for k,v := range  call{
      fmt.Println(k,v.Type(),v)
   }

   //创建带调用方法时需要传入的参数列表
   parms = []reflect.Value{
      reflect.ValueOf(5),
      reflect.ValueOf(6),
   }
   //parm1 = []reflect.Value{reflect.ValueOf(&testB)}
   //parms =  append(parms,parm1...)
   //使用方法名字符串调用指定方法
   call = ruterMap[funcName].Call(parms)
   for k,v := range  call{
      fmt.Println(k,v.Type(),v)
   }
   //可见，testA、testB的值被进行了计算
   //fmt.Println("testStr:", testA)
}


func Server() {

   l, err := net.Listen("tcp", ":80")

   if err != nil {
      panic(err)
   }

   http.Handle("/add", Handl{

   })

   err = http.Serve(l, nil)

   fmt.Println(err)
}

func main() {

   Server()
   
   //fmt.Println("i am service")
   fmt.Println(Employee{
      Name:    "Body",
      Company: Company{
         Name: "Mooc",
         Addr: "BeiJing",
      },
   })
}
