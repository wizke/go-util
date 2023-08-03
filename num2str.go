package util

import (
	"fmt"
	"math"
)

func FormatBytes(bytes uint64) string {
	fBytes := float64(bytes)
	const unit = 1024
	if fBytes < unit {
		return fmt.Sprintf("%.0f B", fBytes)
	}
	exp := math.Floor(math.Log(fBytes) / math.Log(unit))
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/math.Pow(unit, exp), units[int(exp)])
}
