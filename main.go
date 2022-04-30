package main

import (
	"flag"
	"log"
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin"
	"github.com/xruins/mackerel-plugin-zfs-arcstats/lib/zfs"
)

var graphdef map[string]mp.Graphs = map[string]mp.Graphs{
	"zfs.arcstats": {
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
	file string
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

	fields := []string{"hits", "misses", "l2_hits", "l2_misses"}
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
	optFile := flag.String("f", "/proc/spl/kstat/zfs/arcstats", "path to the file of ZFS arcstats")
	flag.Parse()

	p := &ZFSArcStatsPlugin{
		file: *optFile,
	}
	h := mp.NewMackerelPlugin(p)
	h.Run()
}
