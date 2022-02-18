/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	gxlog "github.com/dubbogo/gost/log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	hessian "github.com/apache/dubbo-go-hessian2"
)

type User struct {
	ID   string    `json:"id,omitempty"`
	Code int64     `json:"code,omitempty"`
	Name string    `json:"name,omitempty"`
	Age  int32     `json:"age,omitempty"`
	Time time.Time `json:"time,omitempty"`
}

type UserProvider struct {
	GetUserByName func(ctx context.Context, name string) (rsp *User, err error)
}

//type ContextContent struct {
//	Path              string
//	InterfaceName     string
//	DubboVersion      string
//	LocalAddr         string
//	RemoteAddr        string
//	UserDefinedStrVal string
//	CtxStrVal         string
//	CtxIntVal         int64
//}

//func (u *UserProvider) GetContext(ctx context.Context) (*ContextContent, error) {
//	ctxAtta := ctx.Value(constant.DubboCtxKey("attachment")).(map[string]interface{})
//	userDefinedval := ctxAtta["user-defined-value"].(*ContextContent)
//	gxlog.CInfo("get user defined struct:%#v", userDefinedval)
//	rsp := ContextContent{
//		Path:              ctxAtta["path"].(string),
//		InterfaceName:     ctxAtta["interface"].(string),
//		DubboVersion:      ctxAtta["dubbo"].(string),
//		LocalAddr:         ctxAtta["local-addr"].(string),
//		RemoteAddr:        ctxAtta["remote-addr"].(string),
//		UserDefinedStrVal: userDefinedval.InterfaceName,
//		CtxIntVal:         ctxAtta["int-value"].(int64),
//		CtxStrVal:         ctxAtta["string-value"].(string),
//	}
//	gxlog.CInfo("rsp:%#v", rsp)
//	return &rsp, nil
//}

func (c *User) JavaClassName() string {
	return "com.dubbogo.pixiu.UserService"
}

var userProvider = new(UserProvider)

func init() {
	//config.SetConsumerService(userProvider)
	hessian.RegisterPOJO(&User{ID: "100004", Name: "currentUser002"})
}

var (
	runTime = 0.0
)

// need to setup environment variable "CONF_CONSUMER_FILE_PATH" to "conf/client.yml" before run
func main() {

	//hessian.RegisterPOJO(&ContextContent{})

	//path := "/Users/windwheel/Documents/gitrepo/dubbo-test-samples/context/triple/go-server/conf/dubbogo.yml"
	//
	//if err := config.Load(config.WithPath(path)); err != nil {
	//	panic(err)
	//}
	gxlog.CInfo("\n\n\nstart to test dubbo")

	url := "http://localhost:8883/api/v1/test-dubbo/UserProvider/com.dubbogo.pixiu.UserService?group=dubbo-test&version=1.0.0&method=GetUserByName"
	data := "{\"types\":\"string\",\"values\":\"tc\" }"

	//ctx := context.Background()
	//tpsNum,_ :=strconv.Atoi(os.Getenv("tps"))
	parallel, _ := strconv.Atoi(os.Getenv("parallel"))

	bytesLen := len(data)

	i := 0
	startTime := time.Now()

	successCounter := 0
	totalCounter := 0
	for i = 0; i < parallel; {

		client := &http.Client{Timeout: 5 * time.Second}
		req, err := http.NewRequest("POST", url, strings.NewReader(data))
		req.Header.Add("Content-Type", "application/json")
		resp, _ := client.Do(req)
		if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
			successCounter++
		}
		totalCounter++
		endTime := time.Since(startTime)
		runTime += endTime.Seconds()
		if err != nil {

		}
	}

	fmt.Println("该包体的请求长度: %s", bytesLen)
	fmt.Println("吞吐量为: %s", successCounter)
	fmt.Println("平均调用时长: %s", runTime/float64(successCounter))

	//go_client.NewStressTestConfigBuilder().SetTPS(tpsNum).SetDuration("1h").SetParallel(parallel).Build().Start(func() {
	//
	//	startTime := time.Now()
	//	client := &http.Client{Timeout: 5 * time.Second}
	//	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	//	req.Header.Add("Content-Type", "application/json")
	//	client.Do(req)
	//
	//	if err!=nil {
	//
	//	}
	//	ecpliseTime := time.Since(startTime)
	//	fmt.Println("该包体的请求调用时长: %s",bytesLen)
	//	fmt.Println("调用时长: %s",ecpliseTime)
	//})
	//atta := make(map[string]interface{})
	//
	//atta["string-value"] = "string-demo"
	//atta["int-value"] = 1231242
	//atta["user-defined-value"] = &ContextContent{InterfaceName: "test.interface.name"}
	//reqContext := context.WithValue(context.Background(), constant.DubboCtxKey("attachment"), atta)
	//
	//rspContent, err := userProvider.GetContext(reqContext)
	//if err != nil {
	//	gxlog.CError("error: %v\n", err)
	//	os.Exit(1)
	//	return
	//}
	//gxlog.CInfo("response result: %+v\n", rspContent)
}
