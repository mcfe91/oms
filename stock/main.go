package main

import (
	"context"
	"net"
	"time"

	common "github.com/mcfe91/commons"
	"github.com/mcfe91/commons/broker"
	"github.com/mcfe91/commons/discovery"
	"github.com/mcfe91/commons/discovery/consul"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	serviceName = "stock"
	grpcAddr    = common.EnvString("GRPC_ADDR", "localhost:2002")
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
	amqpUser    = common.EnvString("RABBITMQ_USER", "guest")
	amqpPass    = common.EnvString("RABBITMQ_PASS", "guest")
	amqpHost    = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort    = common.EnvString("RABBITMQ_PORT", "5672")
	jaegerAddr  = common.EnvString("JAEGER_HOST", "localhost:4318")
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	if err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr); err != nil {
		logger.Fatal("failed to set global tracer", zap.Error(err))
	}

	registry, err := consul.NewRegistry(consulAddr)
	if err != nil {
		panic(err)
	}

	instanceID := discovery.GenerateInstanceID(serviceName)
	ctx := context.Background()
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				logger.Fatal("failed to health check", zap.Error(err))
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	svcWithTelemetry := NewTelemetryMiddleware(svc)

	NewGRPCHandler(grpcServer, svcWithTelemetry, nil)

	amqpConsumer := NewConsumer(svc)
	go amqpConsumer.Listen(ch)

	logger.Info("starting gRPC server", zap.String("port", grpcAddr))

	if err := grpcServer.Serve(l); err != nil {
		logger.Fatal("failed to serve gRPC", zap.Error(err))
	}
}
