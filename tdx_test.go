package tdx

import (
	"testing"
	"time"
)

func eq(a, b float32) bool {
	const epsilon float32 = 0.00001
	return ((a-b) < epsilon && (b-a) < epsilon)
}

func TestDecodeFile(t *testing.T) {
	cases := []struct {
		File     string
		Market   string
		Symbol   string
		BarSize  uint
		Time     time.Time
		Open     float32
		High     float32
		Low      float32
		Close    float32
		Volume   uint32
		Turnover float32
	}{
		{
			"testdata/sh600104.day", "XSHG", "600104", 1440,
			time.Date(2017, 6, 30, 15, 0, 0, 0, tz), 30.60, 31.11, 30.50, 30.50, 20422074, 630365184.00,
		},
		{
			"testdata/sh600104.lc1", "XSHG", "600104", 1,
			time.Date(2017, 6, 30, 9, 31, 0, 0, tz), 30.60, 30.69, 30.50, 30.55, 307500, 9407233.00,
		},
		{
			"testdata/sh600104.lc5", "XSHG", "600104", 5,
			time.Date(2017, 6, 30, 9, 35, 0, 0, tz), 30.60, 30.71, 30.50, 30.61, 776600, 23776554.00,
		},
		{
			"testdata/sh600104.5", "XSHG", "600104", 5,
			time.Date(2017, 6, 30, 9, 35, 0, 0, tz), 30.60, 30.71, 30.50, 30.50, 891800, 27297026.00,
		},
	}

	for _, c := range cases {
		dataset, err := DecodeFile(c.File)
		if err != nil {
			t.Error(err)
		}
		bar := dataset.Bars[0]

		if dataset.Market != c.Market || dataset.Symbol != c.Symbol || dataset.BarSize != c.BarSize {
			t.Errorf("unexpected dataset metadata (file: %s)\ngot: %s %s (%d min bars)\nwant: %s %s (%d min bars)\n",
				c.File, dataset.Market, dataset.Symbol, dataset.BarSize, c.Market, c.Symbol, c.BarSize)
		}

		if bar.Time() != c.Time {
			t.Errorf("unexpected timestamp (file: %s)\ngot: %s\nwant: %s\n", c.File, bar.Time(), c.Time)
		}

		if !eq(bar.Open(), c.Open) || !eq(bar.High(), c.High) || !eq(bar.Low(), c.Low) || !eq(bar.Close(), c.Close) {
			t.Errorf("unexpected OHLC (file: %s)\ngot: %.2f %.2f %.2f %.2f\nwant: %.2f %.2f %.2f %.2f\n", c.File,
				bar.Open(), bar.High(), bar.Low(), bar.Close(), c.Open, c.High, c.Low, c.Close)
		}

		if bar.Volume() != c.Volume || !eq(bar.Turnover(), c.Turnover) {
			t.Errorf("unexpected volume/turnover (file: %s)\ngot: %d / %.2f\nwant: %d / %.2f\n", c.File,
				bar.Volume(), bar.Turnover(), c.Volume, c.Turnover)
		}
	}
}
