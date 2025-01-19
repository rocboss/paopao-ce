// Copyright 2025 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"context"
	"fmt"
	"sync"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

var (
	resource          *sdkresource.Resource
	initResourcesOnce sync.Once
)

func InitTelemetry() (shutdownFn func(), err error) {
	s := loggerOtlpSetting
	var (
		tp *sdktrace.TracerProvider
		mp *sdkmetric.MeterProvider
		lp *sdklog.LoggerProvider
	)

	if tp, err = initTracerProvider(s); err != nil {
		return
	}
	if mp, err = initMeterProvider(s); err != nil {
		return
	}
	if lp, err = initLoggerProvider(s); err != nil {
		return
	}
	// Starting runtime instrumentation
	if err = runtime.Start(); err != nil {
		shutdownFn = func() {
			fmt.Printf("skip shutdown by step of init telemetry occurs error:%s", err)
		}
		return
	}
	shutdownFn = func() {
		ctx := context.Background()
		tp.Shutdown(ctx)
		mp.Shutdown(ctx)
		lp.Shutdown(ctx)
	}
	return
}

func initResource() *sdkresource.Resource {
	initResourcesOnce.Do(func() {
		extraResources, _ := sdkresource.New(
			context.Background(),
			sdkresource.WithContainer(),
			sdkresource.WithHost(),
			sdkresource.WithSchemaURL(semconv.SchemaURL),
			sdkresource.WithAttributes(semconv.ServiceNameKey.String("jswhatx")),
		)
		resource, _ = sdkresource.Merge(
			sdkresource.Default(),
			extraResources,
		)
	})
	return resource
}

func initTracerProvider(s *loggerOtlponf) (*sdktrace.TracerProvider, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(s.Endpoint),
		otlptracegrpc.WithHeaders(newOtlpHeaders(s.Authorization, s.Organization, s.TraceStream)),
	}
	if s.Insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	exporter, err := otlptracegrpc.New(context.Background(), opts...)
	if err != nil {
		return nil, fmt.Errorf("otlp trace grpc exporter failed: %w", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(initResource()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func initMeterProvider(s *loggerOtlponf) (*sdkmetric.MeterProvider, error) {
	opts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint(s.Endpoint),
		otlpmetricgrpc.WithHeaders(newOtlpHeaders(s.Authorization, s.Organization, s.MetricStream)),
	}
	if s.Insecure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}
	exporter, err := otlpmetricgrpc.New(context.Background(), opts...)
	if err != nil {
		return nil, fmt.Errorf("otlp metric grpc exporter failed: %w", err)
	}
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter)),
		sdkmetric.WithResource(initResource()),
	)
	otel.SetMeterProvider(mp)
	return mp, nil
}

func initLoggerProvider(s *loggerOtlponf) (*sdklog.LoggerProvider, error) {
	opts := []otlploggrpc.Option{
		otlploggrpc.WithEndpoint(s.Endpoint),
		otlploggrpc.WithHeaders(newOtlpHeaders(s.Authorization, s.Organization, s.LogStream)),
	}
	if s.Insecure {
		opts = append(opts, otlploggrpc.WithInsecure())
	}
	exporter, err := otlploggrpc.New(context.Background(), opts...)
	if err != nil {
		return nil, fmt.Errorf("otlp log grpc exporter failed: %w", err)
	}
	processor := sdklog.NewBatchProcessor(exporter)
	provider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(processor),
		sdklog.WithResource(initResource()),
	)
	global.SetLoggerProvider(provider)
	return provider, nil
}

func newOtlpHeaders(auth, org, stream string) map[string]string {
	return map[string]string{
		"Authorization": auth,
		"organization":  org,
		"stream-name":   stream,
	}
}
