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
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	"github.com/intelsdi-x/snap-plugin-utilities/ns"
)

const (
	// PLUGIN name
	PluginName = "meminfo"
	// VERSION of mem info plugin
	PluginVersion = 4

	// VENDOR namespace part
	nsVendor = "intel"
	// CLASS namespace part
	nsClass = "procfs"
	// PLUGIN name namespace part
	nsType = "meminfo"
)

// procPath source of data for metrics
var procPath = "/proc"

// prefix in metric namespace
var prefix = []string{nsVendor, nsClass, nsType}

// New creates instance of mem info plugin
func New() *memPlugin {
	return &memPlugin{
		proc_path: procPath,
	}
}

// Function to check properness of configuration parameter
// and set plugin attribute accordingly
func (mp *memPlugin) setProcPath(cfg plugin.Config) error {
	procPath, err := cfg.GetString("proc_path")
	if err != nil {
		return err
	}

	if procPath == "" {
		return errors.New("Provided `proc_path` is empty. Cannot get meminfo stats.")
	}

	procPathStats, err := os.Stat(procPath)
	if err != nil {
		return err
	}
	if !procPathStats.IsDir() {
		return errors.New(fmt.Sprintf("%s is not a directory", procPath))
	}
	// set provided `proc_path`
	mp.proc_path = procPath

	return nil
}

// GetMetricTypes returns list of available metric types
// It returns error in case retrieval was not successful
func (mp *memPlugin) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	err := mp.setProcPath(cfg)
	if err != nil {
		return nil, err
	}

	mts := []plugin.Metric{}
	metrics := &MemMetrics{}
	namespaces := []string{}

	ns.FromCompositionTags(metrics, strings.Join(prefix, "/"), &namespaces)
	for _, namespace := range namespaces {
		metric := plugin.Metric{
			Namespace: plugin.NewNamespace(strings.Split(namespace, "/")...),
		}
		mts = append(mts, metric)
	}
	return mts, nil
}

// CollectMetrics returns list of requested metric values
// It returns error in case retrieval was not successful
func (mp *memPlugin) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	err := mp.setProcPath(mts[0].Config)
	if err != nil {
		return nil, err
	}

	metrics := []plugin.Metric{}
	stats := &MemMetrics{}
	if err = getStats(mp.proc_path, stats); err != nil {
		log.Error("Cannot get meminfo statistics")
		return nil, err
	}

	for _, mt := range mts {
		namespace := mt.Namespace.Strings()
		if len(namespace) < 4 {
			return nil, fmt.Errorf("Namespace length is too short (len = %d)", len(namespace))
		}

		val := ns.GetValueByNamespace(stats, namespace[3:])

		metric := plugin.Metric{
			Namespace: mt.Namespace,
			Data:      val,
			Timestamp: time.Now(),
			Version:   PluginVersion,
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

// GetConfigPolicy returns config policy
func (mp *memPlugin) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule(prefix, "proc_path", false, plugin.SetDefaultString("/proc"))
	return *policy, nil
}

type memPlugin struct {
	proc_path string
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

	if err := mapstructure.Decode(stats, metrics); err != nil {
		log.Error("Cannot decode meminfo stats structure")
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
