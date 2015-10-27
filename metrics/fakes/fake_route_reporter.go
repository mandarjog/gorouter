// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/gorouter/metrics"
)

type FakeRouteReporter struct {
	CaptureRouteStatsStub        func(totalRoutes int, msSinceLastUpdate uint64)
	captureRouteStatsMutex       sync.RWMutex
	captureRouteStatsArgsForCall []struct {
		totalRoutes       int
		msSinceLastUpdate uint64
	}
}

func (fake *FakeRouteReporter) CaptureRouteStats(totalRoutes int, msSinceLastUpdate uint64) {
	fake.captureRouteStatsMutex.Lock()
	fake.captureRouteStatsArgsForCall = append(fake.captureRouteStatsArgsForCall, struct {
		totalRoutes       int
		msSinceLastUpdate uint64
	}{totalRoutes, msSinceLastUpdate})
	fake.captureRouteStatsMutex.Unlock()
	if fake.CaptureRouteStatsStub != nil {
		fake.CaptureRouteStatsStub(totalRoutes, msSinceLastUpdate)
	}
}

func (fake *FakeRouteReporter) CaptureRouteStatsCallCount() int {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	return len(fake.captureRouteStatsArgsForCall)
}

func (fake *FakeRouteReporter) CaptureRouteStatsArgsForCall(i int) (int, uint64) {
	fake.captureRouteStatsMutex.RLock()
	defer fake.captureRouteStatsMutex.RUnlock()
	return fake.captureRouteStatsArgsForCall[i].totalRoutes, fake.captureRouteStatsArgsForCall[i].msSinceLastUpdate
}

var _ metrics.RouteReporter = new(FakeRouteReporter)