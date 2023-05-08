package Iris

import (
	"github.com/kataras/iris/v12"
	"testing"
)

func Test1(t *testing.T) {
	//创建一个app结构体指针对象
	app := iris.New()

	//输出html
	// 请求方式: GET
	// 访问地址: http://localhost:8080/
	app.Handle("GET", "/", func(ctx iris.Context) {
		// ctx.HTML返回一个html页面，
		ctx.HTML("<h1>Hello World！！</h1>")
	})
	//输出字符串
	// 类似于 app.Handle("GET", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		// ctx.WriteString将向请求方返回一个字符串
		ctx.WriteString("pong")
	})
	//输出json
	// 请求方式: GET
	// 请求地址: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		// ctx表示返回的结果，ctx.JSON即为返回一个json字符串
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	//8080 监听端口 ，如果未发现服务则返回系统自定义的错误
	app.Run(iris.Addr(":7070"), iris.WithoutServerError(iris.ErrServerClosed))
}

func Test2(t *testing.T) {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<b>Hello!</b>")
	})
	// [...]
	//我们可以用这种方法单独定义我们的配置项
	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	//也可以使用app.run的第二个参数，第二个参数的类型同app.Configure()参数一致。
	app.Run(iris.Addr(":9090"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
	//通过多参数配置 但是上面两种方式是我们最推荐的
	// 我们使用With+配置项名称 如WithCharset("UTF-8") 其中就是With+ Charset的组合
	//app.Run(iris.Addr(":9090"), iris.WithoutStartupLog, iris.WithCharset("UTF-8"))
	//当使用app.Configure(iris.WithoutStartupLog, iris.WithCharset("UTF-8"))设置配置项时
	//需要app.run()面前使用
}
