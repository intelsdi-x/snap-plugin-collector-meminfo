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
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"

	"github.com/intelsdi-x/snap-plugin-utilities/config"
	"github.com/intelsdi-x/snap-plugin-utilities/ns"
)

const (
	// VENDOR namespace part
	pluginVendor = "intel"
	// CLASS namespace part
	fs = "procfs"
	// PLUGIN name namespace part
	pluginName = "meminfo"
	// VERSION of mem info plugin
	pluginVersion = 3
)

// New creates instance of mem info plugin
func New() *memPlugin {
	logger := log.New()
	return &memPlugin{logger: logger}
}

// Meta returns plugin meta data
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		pluginName,
		pluginVersion,
		plugin.CollectorPluginType,
		[]string{},
		[]string{plugin.SnapGOBContentType},
		plugin.ConcurrencyCount(1),
	)
}

// GetMetricTypes returns list of available metric types
// It returns error in case retrieval was not successful
func (mp *memPlugin) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	metricTypes := []plugin.MetricType{}
	metrics := &MemMetrics{}
	namespaces := []string{}
	ns.FromCompositionTags(metrics, strings.Join([]string{pluginVendor, fs, pluginName}, "/"), &namespaces)
	for _, namespace := range namespaces {
		metric := plugin.MetricType{
			Namespace_: core.NewNamespace(strings.Split(namespace, "/")...),
			Config_:    cfg.ConfigDataNode,
		}
		metricTypes = append(metricTypes, metric)
	}
	return metricTypes, nil
}

// CollectMetrics returns list of requested metric values
// It returns error in case retrieval was not successful
func (mp *memPlugin) CollectMetrics(metricTypes []plugin.MetricType) ([]plugin.MetricType, error) {
	metrics := []plugin.MetricType{}

	pathMemInfo, err := config.GetConfigItem(metricTypes[0], "proc_path")
	if err != nil {
		return nil, err
	}

	stats := &MemMetrics{}
	getStats(pathMemInfo.(string), stats)

	for _, metricType := range metricTypes {
		namespace := metricType.Namespace()
		if len(namespace.Strings()) < 4 {
			return nil, fmt.Errorf("Namespace length is too short (len = %d)", len(namespace.Strings()))
		}

		val := ns.GetValueByNamespace(stats, namespace.Strings()[3:])

		metric := plugin.MetricType{
			Namespace_: metricType.Namespace(),
			Data_:      val,
			Timestamp_: time.Now(),
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

// GetConfigPolicy returns config policy
// It returns error in case retrieval was not successful
func (mp *memPlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	rule, _ := cpolicy.NewStringRule("proc_path", false, "/proc")
	node := cpolicy.NewPolicyNode()
	node.Add(rule)
	cp.Add([]string{pluginVendor, fs, pluginName}, node)
	return cp, nil
}

type memPlugin struct {
	logger *log.Logger
}

func getStats(procPath string, metrics *MemMetrics) error {
	fh, err := os.Open(path.Join(procPath, "meminfo"))

	if err != nil {
		return err
	}
	defer fh.Close()

	var memSum uint64
	tmpStats := map[string]interface{}{}

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) < 2 {
			return fmt.Errorf("Wrong %v format", fields)
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

		tmpStats[formatName(name)] = value * unit
	}

	if tmpStats["mem_total"].(uint64) < memSum {
		return fmt.Errorf("More value mismatch! More used then available")
	}

	tmpStats["mem_used"] = tmpStats["mem_total"].(uint64) - memSum
	total := tmpStats["mem_used"].(uint64) + memSum
	stats := map[string]interface{}{}
	for stat, value := range tmpStats {
		stats[stat] = value
		percentage := stat + "_perc"
		stats[percentage] = 100.0 * float64(value.(uint64)) / float64(total)
	}

	err = mapstructure.Decode(stats, metrics)
	if err != nil {
		return err
	}

	return nil
}

func formatName(name string) string {
	uppers := []int{}
	name = strings.Replace(name, "(", "_", -1)
	name = strings.Replace(name, ")", "", -1)

	for i, c := range name {
		if unicode.IsUpper(c) {
			uppers = append(uppers, i)
		}
	}

	name = strings.ToLower(name)
	formatted := ""
	for i, _ := range uppers {
		if uppers[i] != 0 {
			bottom := uppers[i-1]
			up := uppers[i]
			prev := string(name[up-1])
			formatted += name[bottom:up]
			if up-bottom > 1 && !strings.Contains("1234567890_", prev) {
				formatted += "_"
			}
		}
	}

	return formatted + name[uppers[len(uppers)-1]:]
}
