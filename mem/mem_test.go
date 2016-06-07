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
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"

	"github.com/intelsdi-x/snap/core/ctypes"
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
	gss.tot = 1200
	gss.buf = 500
	gss.cache = 300
	gss.free = 100
	gss.slab = 100
	gss.used = gss.tot - (gss.free + gss.buf + gss.cache + gss.slab)
	createMockMemInfo("meminfo", gss.tot, gss.buf, gss.cache, gss.free, gss.slab)
}

func (gss *GetStatsSuite) TearDownSuite() {
	removeMockMemInfo("meminfo")
}

func TestGetStatsSuite(t *testing.T) {
	suite.Run(t, &GetStatsSuite{MockMemInfo: "."})
}

func (gss *GetStatsSuite) TestGetStats() {
	Convey("Given memory statistics", gss.T(), func() {
		stats := &MemMetrics{}

		Convey("and mock memory info file created", func() {
			assert.Equal(gss.T(), ".", gss.MockMemInfo)
		})

		Convey("When reading memory statistics from file", func() {
			err := getStats(gss.MockMemInfo, stats)

			Convey("Then no error is reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then proper memory metrics are retrived", func() {
				So(stats.MemTotal, ShouldEqual, gss.tot*1024)
				So(stats.Buffers, ShouldEqual, gss.buf*1024)
				So(stats.Cached, ShouldEqual, gss.cache*1024)
				So(stats.MemFree, ShouldEqual, gss.free*1024)
				So(stats.Slab, ShouldEqual, gss.slab*1024)
				So(stats.MemUsed, ShouldEqual, gss.used*1024)
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
	removeMockMemInfo("meminfo")
}

func (mps *MemPluginSuite) SetupSuite() {
	mps.tot = 2000
	mps.buf = 500
	mps.cache = 300
	mps.free = 100
	mps.slab = 100
	mps.used = mps.tot - (mps.free + mps.buf + mps.cache + mps.slab)
	createMockMemInfo("meminfo", mps.tot, mps.buf, mps.cache, mps.free, mps.slab)

}

func (mps *MemPluginSuite) TestGetMetricTypes() {
	Convey("Given MemInfo plugin initialized", mps.T(), func() {
		memPlugin := New()

		Convey("When one wants to get list of available meterics", func() {
			cfg := plugin.NewPluginConfigType()
			cfg.AddItem("proc_path", ctypes.ConfigValueStr{mps.MockMemInfo})
			mts, err := memPlugin.GetMetricTypes(cfg)

			Convey("Then error should not be reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then list of metrics is returned", func() {
				So(len(mts), ShouldEqual, 108)

				namespaces := []string{}
				for _, m := range mts {
					namespaces = append(namespaces, m.Namespace().String())
				}

				So(namespaces, ShouldContain, "/intel/procfs/meminfo/cached")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/cached_perc")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/mem_total")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/mem_total_perc")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/mem_free")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/mem_free_perc")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/mem_used")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/mem_used_perc")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/buffers_perc")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/buffers")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/slab_perc")
				So(namespaces, ShouldContain, "/intel/procfs/meminfo/slab")
			})
		})
	})
}

func (mps *MemPluginSuite) TestCollectMetrics() {
	Convey("Given memInfo plugin initlialized", mps.T(), func() {
		memPlugin := New()

		Convey("When one wants to get values for given metric types", func() {
			cfg := plugin.NewPluginConfigType()
			cfg.AddItem("proc_path", ctypes.ConfigValueStr{mps.MockMemInfo})
			mTypes := []plugin.MetricType{
				plugin.MetricType{Namespace_: core.NewNamespace("intel", "procfs", "meminfo", "cached"), Config_: cfg.ConfigDataNode},
				plugin.MetricType{Namespace_: core.NewNamespace("intel", "procfs", "meminfo", "cached_perc"), Config_: cfg.ConfigDataNode},
				plugin.MetricType{Namespace_: core.NewNamespace("intel", "procfs", "meminfo", "mem_total"), Config_: cfg.ConfigDataNode},
				plugin.MetricType{Namespace_: core.NewNamespace("intel", "procfs", "meminfo", "mem_used"), Config_: cfg.ConfigDataNode},
			}

			metrics, err := memPlugin.CollectMetrics(mTypes)

			Convey("Then no erros should be reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then proper metrics values are returned", func() {
				So(len(metrics), ShouldEqual, 4)

				stats := map[string]interface{}{}
				for _, m := range metrics {
					n := m.Namespace().String()
					stats[n] = m.Data()
				}

				assert.Equal(mps.T(), len(metrics), len(stats))

				So(stats["/intel/procfs/meminfo/cached"].(uint64), ShouldEqual, mps.cache*1024)
				So(stats["/intel/procfs/meminfo/cached_perc"].(float64), ShouldEqual, 100.0*mps.cache/mps.tot)
				So(stats["/intel/procfs/meminfo/mem_total"].(uint64), ShouldEqual, mps.tot*1024)
				So(stats["/intel/procfs/meminfo/mem_used"].(uint64), ShouldEqual, mps.used*1024)
			})

		})
	})
}

func TestMemPluginSuite(t *testing.T) {
	suite.Run(t, &MemPluginSuite{MockMemInfo: "."})
}

func createMockMemInfo(memInfo string, tot, buf, cache, free, slab uint64) {
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

func removeMockMemInfo(memInfo string) {
	os.Remove(memInfo)
}
