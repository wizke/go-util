package util

import (
	"fmt"
	"math"
)

func FormatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%.0f B", bytes)
	}
	fBytes := float64(bytes)
	exp := math.Floor(math.Log(fBytes) / math.Log(unit))
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/math.Pow(unit, exp), units[int(exp)])
}
