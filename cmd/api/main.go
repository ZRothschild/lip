package main

import (
	"context"
	"github.com/ZRothschild/ldp/app/company/companyRepo"
	"github.com/ZRothschild/ldp/app/user/userRepo"
	"github.com/ZRothschild/ldp/app/userBindCompany/userBindCompanyRepo"
	"github.com/ZRothschild/ldp/cmd/api/opt"
	"github.com/ZRothschild/ldp/gen/company"
	"github.com/ZRothschild/ldp/gen/login"
	"github.com/ZRothschild/ldp/gen/register"
	"github.com/ZRothschild/ldp/gen/user"
	"github.com/ZRothschild/ldp/infrastr/conf"
	"github.com/ZRothschild/ldp/infrastr/mysql"
	companySrv "github.com/ZRothschild/ldp/server/company"
	loginSrv "github.com/ZRothschild/ldp/server/login"
	registerSrv "github.com/ZRothschild/ldp/server/register"
	userSrv "github.com/ZRothschild/ldp/server/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"
	"strconv"
)

/*
这个是支持 opts 中间键
1. grpc server 创建监听
2. http 请求创建
3. grpc client 创建
4. 当请求http http转发到grpc client 由他发送请求到 grpc server

or

当然也可以
1. grpc server 创建监听
2. http 请求创建
3. 当请求http http转发到 grpc server
*/

func run() error {
	var (
		err error
		l   net.Listener
		s   = grpc.NewServer()
		mux = runtime.NewServeMux(
			runtime.WithMiddlewares(),
			runtime.WithErrorHandler(opt.CustomHTTPErrorHandler),
			runtime.WithForwardResponseRewriter(opt.CustomForwardResponseRewriter),
			runtime.WithOutgoingHeaderMatcher(opt.CustomOutgoingHeaderMatcher),
			runtime.WithOutgoingTrailerMatcher(opt.CustomOutgoingTrailerMatcher),
		)
		db   = mysql.NewDb(conf.Conf)
		addr = ":" + strconv.Itoa(conf.Conf.GrpcPort)
	)

	if l, err = net.Listen("tcp", addr); err != nil {
		log.Printf("Failed to listen: %v", err)
		return err
	}

	companyR := companyRepo.NewCompanyRepo(db)
	userR := userRepo.NewUserRepo(db)
	userBindCompanyR := userBindCompanyRepo.NewUserBindCompanyRepo(db)

	registerS := registerSrv.NewRegisterServer(userR, companyR, userBindCompanyR)
	loginS := loginSrv.NewLoginServer(userR)

	userS := userSrv.NewUserServer(userR)
	// 这里是注册grpc服务器服务，http 可以不注册
	user.RegisterUserServiceServer(s, userS)
	login.RegisterLoginServiceServer(s, loginS)
	register.RegisterRegisterServiceServer(s, registerS)
	company.RegisterCompanyServiceServer(s, companySrv.NewCompanyServer())

	go func() {
		log.Fatalln(s.Serve(l))
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// http 请求直接转发到 ServiceServer
	if err = user.RegisterUserServiceHandlerServer(ctx, mux, userS); err != nil {
		log.Printf("Failed RegisterUserServiceHandlerServer: %v", err)
		return err
	}

	if err = login.RegisterLoginServiceHandlerServer(ctx, mux, loginS); err != nil {
		log.Printf("Failed RegisterLoginServiceHandlerServer: %v", err)
		return err
	}

	if err = register.RegisterRegisterServiceHandlerServer(ctx, mux, registerS); err != nil {
		log.Printf("Failed RegisterRegisterServiceHandlerServer: %v", err)
		return err
	}

	if err = company.RegisterCompanyServiceHandlerServer(ctx, mux, companySrv.NewCompanyServer()); err != nil {
		log.Printf("Failed RegisterCompanyServiceHandlerServer: %v", err)
		return err
	}

	// http 请求先转发到 ServiceClient 再由 ServiceClient 转发到 ServiceServer 这里会多一层中间件

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err = user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		log.Printf("Failed RegisterUserServiceHandlerFromEndpoint: %v", err)
		return err
	}

	if err = login.RegisterLoginServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		log.Printf("Failed RegisterLoginServiceHandlerFromEndpoint: %v", err)
		return err
	}

	if err = register.RegisterRegisterServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		log.Printf("Failed RegisterRegisterServiceHandlerFromEndpoint: %v", err)
		return err
	}

	if err = company.RegisterCompanyServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		log.Printf("Failed RegisterCompanyServiceHandlerFromEndpoint: %v", err)
		return err
	}
	
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(conf.Conf.HttpPort),
		Handler: mux, // logRequestBody(allowCORS(mux)),
	}
	return srv.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
