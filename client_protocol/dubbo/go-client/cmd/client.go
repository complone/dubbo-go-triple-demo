package cmd

import (
	"context"
	hessian "github.com/apache/dubbo-go-hessian2"
)

type UserProvider struct {
	GetContext func(ctx context.Context)
}

type ContextContent struct {
	Path              string
	InterfaceName     string
	DubboVersion      string
	LocalAddr         string
	RemoteAddr        string
	UserDefinedStrVal string
	CtxStrVal         string
	CtxIntVal         int64
}

func (c *ContextContent) JavaClassName() string {
	return "org.apache.duboo.ContextContent"
}

var userProvider = new(UserProvider)

func main() {
	hessian.RegisterPOJO(&ContextContent{})

}
