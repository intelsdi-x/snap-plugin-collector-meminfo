/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
)

const (
	// VENDOR namespace part
	VENDOR = "intel"
	// OS namespace part
	OS = "linux"
	// PLUGIN name namespace part
	PLUGIN = "meminfo"
	// VERSION of mem info plugin
	VERSION = 1
)

var memInfo = "/proc/meminfo"

type memPlugin struct {
	stats map[string]interface{}
	host  string
}

// New creates instance of mem info plugin
func New() *memPlugin {
	fh, err := os.Open(memInfo)

	if err != nil {
		return nil
	}
	defer fh.Close()

	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	mp := &memPlugin{stats: map[string]interface{}{}, host: host}
	return mp
}

func getStats(stats map[string]interface{}) error {
	fh, err := os.Open(memInfo)

	if err != nil {
		return err
	}
	defer fh.Close()

	var memSum uint64
	tmpStats := map[string]uint64{}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) < 2 {
			return fmt.Errorf("Wrong %s format", memInfo)
		}

		name := fields[0]
		name = name[:len(name)-1]

		var unit uint64 = 1024
		if len(fields) == 3 {
			switch fields[2] {
			case "MB":
				unit *= 1024
			case "GB":
				unit *= 1024 * 1024
			}
		}

		value, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return err
		}

		switch name {
		case "MemFree", "Buffers", "Cached", "Slab":
			memSum += value * unit
		}

		tmpStats[name] = value * unit
	}

	if tmpStats["MemTotal"] < memSum {
		return fmt.Errorf("More value mismatch! More used then available")
	}

	tmpStats["MemUsed"] = tmpStats["MemTotal"] - memSum

	total := tmpStats["MemUsed"] + memSum
	for stat, value := range tmpStats {
		percentage := stat + "_perc"
		stats[stat] = value
		stats[percentage] = 100.0 * value / total
	}

	return nil
}

// GetMetricTypes returns list of available metric types
// It returns error in case retrieval was not successful
func (mp *memPlugin) GetMetricTypes(_ plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	metricTypes := []plugin.PluginMetricType{}
	if err := getStats(mp.stats); err != nil {
		return nil, err
	}
	for stat := range mp.stats {
		metricType := plugin.PluginMetricType{Namespace_: []string{VENDOR, OS, PLUGIN, stat}}
		metricTypes = append(metricTypes, metricType)
	}
	return metricTypes, nil
}

// CollectMetrics returns list of requested metric values
// It returns error in case retrieval was not successful
func (mp *memPlugin) CollectMetrics(metricTypes []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	metrics := []plugin.PluginMetricType{}
	getStats(mp.stats)
	for _, metricType := range metricTypes {
		ns := metricType.Namespace()
		if len(ns) < 4 {
			return nil, fmt.Errorf("Namespace length is too short (len = %d)", len(ns))
		}
		stat := ns[3]
		val, ok := mp.stats[stat]
		if !ok {
			return nil, fmt.Errorf("Requested stat %s is not available!", stat)
		}
		metric := plugin.PluginMetricType{
			Namespace_: ns,
			Data_:      val,
			Source_:    mp.host,
			Timestamp_: time.Now(),
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil

}

// GetConfigPolicy returns config policy
// It returns error in case retrieval was not successful
func (mp *memPlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	return cpolicy.New(), nil
}
