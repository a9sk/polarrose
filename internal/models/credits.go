package models

var Watermark = []string{"a9sk  (Emiliano Rizzonelli)", "github.com/a9sk/polarrose"}

func GetWatermark() string {
	return Watermark[0] + " source: " + Watermark[1]
}
