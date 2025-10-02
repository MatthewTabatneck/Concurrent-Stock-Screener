package tickers

import (
	"encoding/csv"
	"io"
	"strings"
)

func LoadtickersCSV(r io.Reader) ([]string, error) {
	cr := csv.NewReader(r)
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{}, len(records))
	for _, rec := range records {
		if len(rec) == 0 {
			continue
		}
		sym := strings.TrimSpace(rec[0])
		if sym == "" {
			continue
		}
		sym = strings.ToUpper(sym)
		seen[sym] = struct{}{} // set membership (zero-size value)
	}

	// Emit final []string (order is arbitrary)
	out := make([]string, 0, len(seen))
	for sym := range seen {
		out = append(out, sym)
	}
	return out, nil
}
