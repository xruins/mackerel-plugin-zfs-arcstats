package main

import (
	"flag"
	"log"
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/xruins/mackerel-plugin-zfs-arcstats/lib/zfs"
)

var graphdef map[string]mp.Graphs = map[string]mp.Graphs{
	"arcstats": {
		Label: "ARC Hits/Misses",
		Unit:  "integer",
		Metrics: []mp.Metrics{
			{Name: "hits", Label: "ARC Hits", Diff: true},
			{Name: "misses", Label: "ARC Misses", Diff: true},
			{Name: "l2_hits", Label: "L2ARC Hits", Diff: true},
			{Name: "l2_misses", Label: "L2ARC Misses", Diff: true},
		},
	},
}

type ZFSArcStatsPlugin struct {
	file     string
	prefix   string
	enableL2 bool
}

func (m *ZFSArcStatsPlugin) MetricKeyPrefix() string {
	return m.prefix
}

func (m *ZFSArcStatsPlugin) FetchMetrics() (map[string]float64, error) {
	fp, err := os.Open(m.file)
	if err != nil {
		log.Fatalf("failed to open the file for ZFS arcstats. err: %s", err)
	}
	defer fp.Close()

	stats, err := zfs.ParseArcStats(fp)
	if err != nil {
		log.Fatalf("failed to parse ZFS arcstats. err: %s", err)
	}

	fields := []string{"hits", "misses"}
	if m.enableL2 {
		fields = append(fields, "l2_hits", "l2_misses")
	}
	ret := make(map[string]float64, len(fields))
	for _, f := range fields {
		row := stats.Get(f)
		ret[f] = float64(row.Value)
	}

	return ret, nil
}

func (m *ZFSArcStatsPlugin) GraphDefinition() map[string]mp.Graphs {
	return graphdef
}

func main() {
	optFile := flag.String("f", "/proc/spl/kstat/zfs/arcstats", "Path to the file of ZFS arcstats")
	optTempfile := flag.String("tempfile", "", "Tempfile name")
	optMetricKeyPrefix := flag.String("metric-key-prefix", "zfs", "Metric Key Prefix")
	optEnableL2Metrics := flag.Bool("enable-l2-metrics", false, "enable metrics for L2ARC")
	flag.Parse()

	p := &ZFSArcStatsPlugin{
		file:     *optFile,
		prefix:   *optMetricKeyPrefix,
		enableL2: *optEnableL2Metrics,
	}
	h := mp.NewMackerelPlugin(p)
	h.Tempfile = *optTempfile
	h.Run()
}
