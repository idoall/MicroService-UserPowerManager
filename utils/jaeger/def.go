// Jaeger 是Uber推出的一款开源分布式追踪系统，兼容OpenTracing API。
// https://www.jaegertracing.io/
package jaeger

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"

	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	globalopentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// 初始化 Jaeger 服务
// @param service 服务名称
func InitJaeger(service string) (globalopentracing.Tracer, io.Closer) {

	var agentHostPort string

	if os.Getenv("JAEGER_AGENTHOSTPORT") != "" {
		agentHostPort = os.Getenv("JAEGER_AGENTHOSTPORT")
	} else {
		agentHostPort = utils.TConfig.String("JaegerServices::AgentHostPort")
	}

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: agentHostPort,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	globalopentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

// 写入一个标签
func StartSpan(ctx context.Context, spanName string) (context.Context, globalopentracing.Span) {
	// 是否开启 Jaeger
	if enableJaeget, err := utils.TConfig.Bool("JaegerServices::Enable"); err != nil {
		inner.Mlogger.Fatal(err)
	} else if enableJaeget {
		//定义一个开始标签
		ctx, span, err := startSpanFromContext(ctx, spanName)
		if err != nil {
			inner.Mlogger.Fatal(err)
		}
		return ctx, span
		//---
	}
	return nil, nil
}

// StartSpanFromContext returns a new span with the given operation name and options. If a span
// is found in the context, it will be used as the parent of the resulting span.
func startSpanFromContext(ctx context.Context, name string, opts ...opentracing.StartSpanOption) (context.Context, opentracing.Span, error) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}

	// copy the metadata to prevent race
	md = metadata.Copy(md)

	// find trace in go-micro metadata
	if spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md)); err == nil {
		opts = append(opts, opentracing.ChildOf(spanCtx))
	}

	// find span context in opentracing library
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
	}

	sp := opentracing.GlobalTracer().StartSpan(name, opts...)

	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md)); err != nil {
		return nil, nil, err
	}

	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	return ctx, sp, nil
}
