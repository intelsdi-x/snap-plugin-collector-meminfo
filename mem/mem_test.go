// +build unit

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015-2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mem

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetStatsSuite struct {
	suite.Suite
	MockMemInfo                       string
	tot, buf, cache, free, slab, used uint64
}

func (gss *GetStatsSuite) SetupSuite() {
	memInfo = gss.MockMemInfo
	gss.tot = 1200
	gss.buf = 500
	gss.cache = 300
	gss.free = 100
	gss.slab = 100
	gss.used = gss.tot - (gss.free + gss.buf + gss.cache + gss.slab)
	createMockMemInfo(gss.tot, gss.buf, gss.cache, gss.free, gss.slab)
}

func (gss *GetStatsSuite) TearDownSuite() {
	removeMockMemInfo()
}

func TestGetStatsSuite(t *testing.T) {
	suite.Run(t, &GetStatsSuite{MockMemInfo: "mockMemInfo"})
}

func (gss *GetStatsSuite) TestGetStats() {
	Convey("Given memory info map", gss.T(), func() {
		stats := map[string]interface{}{}

		Convey("and mock memory info file created", func() {
			assert.Equal(gss.T(), "mockMemInfo", memInfo)
		})

		Convey("When reading memory statistics from file", func() {
			err := getStats(stats)

			Convey("Then no error is reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then proper memory metrics are retrived", func() {
				tot, ok := stats["MemTotal"].(uint64)
				So(ok, ShouldBeTrue)
				So(tot, ShouldEqual, gss.tot*1024)

				buf, ok := stats["Buffers"].(uint64)
				So(ok, ShouldBeTrue)
				So(buf, ShouldEqual, gss.buf*1024)

				cache, ok := stats["Cached"].(uint64)
				So(ok, ShouldBeTrue)
				So(cache, ShouldEqual, gss.cache*1024)

				free, ok := stats["MemFree"].(uint64)
				So(ok, ShouldBeTrue)
				So(free, ShouldEqual, gss.free*1024)

				slab, ok := stats["Slab"].(uint64)
				So(ok, ShouldBeTrue)
				So(slab, ShouldEqual, gss.slab*1024)

				used, ok := stats["MemUsed"].(uint64)
				So(ok, ShouldBeTrue)
				So(used, ShouldEqual, gss.used*1024)
			})
		})
	})
}

type MemPluginSuite struct {
	suite.Suite
	MockMemInfo                       string
	tot, buf, cache, free, slab, used uint64
}

func (mps *MemPluginSuite) TearDownSuite() {
	removeMockMemInfo()
}

func (mps *MemPluginSuite) SetupSuite() {
	memInfo = mps.MockMemInfo
	mps.tot = 2000
	mps.buf = 500
	mps.cache = 300
	mps.free = 100
	mps.slab = 100
	mps.used = mps.tot - (mps.free + mps.buf + mps.cache + mps.slab)
	createMockMemInfo(mps.tot, mps.buf, mps.cache, mps.free, mps.slab)

}

func (mps *MemPluginSuite) TestGetMetricTypes() {
	Convey("Given MemInfo plugin initialized", mps.T(), func() {
		memPlugin := New()

		Convey("When one wants to get list of available meterics", func() {
			mts, err := memPlugin.GetMetricTypes(plugin.PluginConfigType{})

			Convey("Then error should not be reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then list of metrics is returned", func() {
				So(len(mts), ShouldEqual, 12)

				namespaces := []string{}
				for _, m := range mts {
					namespaces = append(namespaces, strings.Join(m.Namespace(), "/"))
				}

				So(namespaces, ShouldContain, "intel/procfs/meminfo/Cached")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/Cached_perc")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/MemTotal")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/MemTotal_perc")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/MemFree")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/MemFree_perc")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/MemUsed")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/MemUsed_perc")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/Buffers_perc")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/Buffers")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/Slab_perc")
				So(namespaces, ShouldContain, "intel/procfs/meminfo/Slab")
			})
		})
	})
}

func (mps *MemPluginSuite) TestCollectMetrics() {
	Convey("Given memInfo plugin initlialized", mps.T(), func() {
		memPlugin := New()

		Convey("When one wants to get values for given metric types", func() {
			mTypes := []plugin.PluginMetricType{
				plugin.PluginMetricType{Namespace_: []string{"intel", "procfs", "meminfo", "Cached"}},
				plugin.PluginMetricType{Namespace_: []string{"intel", "procfs", "meminfo", "Cached_perc"}},
				plugin.PluginMetricType{Namespace_: []string{"intel", "procfs", "meminfo", "MemTotal"}},
				plugin.PluginMetricType{Namespace_: []string{"intel", "procfs", "meminfo", "MemUsed"}},
			}

			metrics, err := memPlugin.CollectMetrics(mTypes)

			Convey("Then no erros should be reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then proper metrics values are returned", func() {
				So(len(metrics), ShouldEqual, 4)

				stats := map[string]uint64{}
				for _, m := range metrics {
					n := strings.Join(m.Namespace(), "/")
					v, ok := m.Data().(uint64)
					if ok {
						stats[n] = v
					}
				}

				assert.Equal(mps.T(), len(metrics), len(stats))

				So(stats["intel/procfs/meminfo/Cached"], ShouldEqual, mps.cache*1024)
				So(stats["intel/procfs/meminfo/Cached_perc"], ShouldEqual, 100.0*mps.cache/mps.tot)
				So(stats["intel/procfs/meminfo/MemTotal"], ShouldEqual, mps.tot*1024)
				So(stats["intel/procfs/meminfo/MemUsed"], ShouldEqual, mps.used*1024)
			})

		})
	})
}

func TestMemPluginSuite(t *testing.T) {
	suite.Run(t, &MemPluginSuite{MockMemInfo: "mockMemInfo"})
}

func createMockMemInfo(tot, buf, cache, free, slab uint64) {
	content := fmt.Sprintf(
		"MemTotal: %d kb\nBuffers: %d kB\nCached: %d kB\nMemFree: %d kB\nSlab: %d kB",
		tot, buf, cache, free, slab)
	memInfoContent := []byte(content)
	f, err := os.Create(memInfo)
	if err != nil {
		panic(err)
	}
	f.Write(memInfoContent)
}

func removeMockMemInfo() {
	os.Remove(memInfo)
}
