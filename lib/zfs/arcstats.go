package zfs

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type ArcStats struct {
	rows map[string]*Row
}

type Row struct {
	Name  string
	Type  int64
	Value int64
}

func (a *ArcStats) Get(name string) *Row {
	v, ok := a.rows[name]
	if !ok {
		return nil
	}
	return v
}

func (a *ArcStats) SetRows(r map[string]*Row) {
	a.rows = r
}

func ParseArcStats(r io.Reader) (*ArcStats, error) {
	br := bufio.NewReader(r)
	// skip header and column names
	for i := 1; i <= 2; i++ {
		_, err := br.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read %d(st/nd) row of arcstats: %w", i, err)
		}
	}

	// splits rows into Rows struct
	rows := make(map[string]*Row)
	scanner := bufio.NewScanner(br)
	scanner.Split(bufio.ScanWords)

	row := &Row{}
	for scanner.Scan() {
		s := scanner.Text()

		if row.Name == "" {
			row = &Row{
				Name: s,
			}
			continue
		}

		var err error
		if row.Type == 0 {
			row.Type, err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("malformed type. s: %s, err: %w", s, err)
			}
			continue
		}

		row.Value, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("malformed value. s: %s, err: %w", s, err)
		}
		rows[row.Name] = row
		row = &Row{}
	}

	return &ArcStats{rows: rows}, nil
}
