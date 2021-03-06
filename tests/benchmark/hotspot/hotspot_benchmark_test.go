package hotspot

import (
	"log"
	"math"
	"testing"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/alibaba/sentinel-golang/tests/benchmark"
)

const (
	resReject      = "resReject"
	resConcurrency = "resConcurrency"
	resThrottling  = "resThrottling"
)

func init() {
	benchmark.InitSentinel()
	rule1 := &hotspot.Rule{
		Resource:      resConcurrency,
		MetricType:    hotspot.Concurrency,
		ParamIndex:    0,
		Threshold:     math.MaxInt64,
		DurationInSec: 0,
	}
	rule2 := &hotspot.Rule{
		Resource:        resReject,
		MetricType:      hotspot.QPS,
		ControlBehavior: hotspot.Reject,
		ParamIndex:      0,
		Threshold:       math.MaxInt64,
		BurstCount:      0,
		DurationInSec:   1,
	}
	rule3 := &hotspot.Rule{
		Resource:          resThrottling,
		MetricType:        hotspot.QPS,
		ControlBehavior:   hotspot.Throttling,
		ParamIndex:        0,
		Threshold:         math.MaxInt64,
		BurstCount:        0,
		DurationInSec:     1,
		MaxQueueingTimeMs: 0,
	}
	_, err := hotspot.LoadRules([]*hotspot.Rule{rule1, rule2, rule3})
	if err != nil {
		panic(err)
	}
}

func doCheck(res string) {
	if se, err := sentinel.Entry(res); err == nil {
		se.Exit()
	} else {
		log.Fatalf("Block err: %s", err.Error())
	}
}

func Benchmark_Concurrency_Concurrency4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resConcurrency)
		}
	})
}

func Benchmark_Concurrency_Concurrency8(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resConcurrency)
		}
	})
}

func Benchmark_Concurrency_Concurrency16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(16)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resConcurrency)
		}
	})
}

func Benchmark_QPS_Reject_Concurrency4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resConcurrency)
		}
	})
}

func Benchmark_QPS_Reject_Concurrency8(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resReject)
		}
	})
}

func Benchmark_QPS_Reject_Concurrency16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(16)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resReject)
		}
	})
}

func Benchmark_QPS_Throttling_Concurrency4(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(4)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resThrottling)
		}
	})
}

func Benchmark_QPS_Throttling_Concurrency8(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resThrottling)
		}
	})
}

func Benchmark_QPS_Throttling_Concurrency16(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.SetParallelism(16)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			doCheck(resThrottling)
		}
	})
}
