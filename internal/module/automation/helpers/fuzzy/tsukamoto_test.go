package fuzzy_test

import (
	"hidroponic/internal/module/automation/helpers/fuzzy"
	"testing"
)

// fuzzy turn on

func TestInferenceTemperatureLowNutritionLowWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(20, 90, 1000)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}

	value = fis.Inference(19, 89, 999)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}

	value = fis.Inference(22, 93, 1001)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value above threshold", value)
	}
}
func TestInferenceTemperatureOptimalNutritionLowWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 90, 1000)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}

	value = fis.Inference(23, 85, 995)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}

	value = fis.Inference(27, 95, 1100)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionLowWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(35, 94, 1400)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureLowNutritionOptimalWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(22, 101, 1010)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionOptimalWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(26, 102, 1150)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionOptimalWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(30, 100, 1000)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureLowNutritionHighWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(18, 109, 998)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionHighWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 110, 1000)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionHighWaterLow(t *testing.T) {
	fis := getFis()

	value := fis.Inference(30, 110, 1000)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureLowNutritionLowWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(21, 88, 2400)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionLowWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 90, 2500)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionLowWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(30, 90, 2500)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionOptimalWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(30, 100, 2500)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureLowNutritionHighWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(17, 115, 2200)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionHighWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 110, 2500)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionHighWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(29, 109, 2600)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureLowNutritionLowWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(23, 93, 3100)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionLowWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 90, 3000)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

func TestInferenceTemperatureHighNutritionLowWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(31, 93, 2999)
	t.Log("fuzzy value is ", value)
	if value < fuzzy.THRESHOLD {
		t.Error("value below threshold")
	}
}

// fuzzy turn off
func TestInferenceTemperatureLowNutritionOptimalWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(21, 101, 2500)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionOptimalWaterMedium(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 100, 2500)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureLowNutritionOptimalWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(16, 99, 3010)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionOptimalWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(24, 99, 3100)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureHighNutritionOptimalWaterHigh(t *testing.T) {
	fis := getFis()

	for i := float32(2799); i <= 3000; i++ {
		value := fis.Inference(25, 97, i)
		t.Log("fuzzy value is ", value)
		if value > fuzzy.THRESHOLD {
			t.Error("value above threshold")
		}
	}
}

func TestInferenceTemperatureLowNutritionHighWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(21, 110, 3000)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureOptimalNutritionHighWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(25, 110, 3000)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestInferenceTemperatureHighNutritionHighWaterHigh(t *testing.T) {
	fis := getFis()

	value := fis.Inference(28, 111, 2850)
	t.Log("fuzzy value is ", value)
	if value > fuzzy.THRESHOLD {
		t.Error("value above threshold")
	}
}

func TestAbsolutValue(t *testing.T) {
	fis := getFis()

	var data map[string][]float32 = map[string][]float32{
		"rendah sedikit dingin":   {90, 1000, 20, 3},
		"rendah sedikit optimal":  {90, 1000, 25, 3},
		"rendah sedikit panas":    {90, 1000, 30, 3},
		"optimal sedikit dingin":  {100, 1000, 20, 3},
		"optimal sedikit optimal": {100, 1000, 25, 3},
		"optimal sedikit panas":   {100, 1000, 30, 3},
		"tinggi sedikit dingin":   {110, 1000, 20, 3},
		"tinggi sedikit optimal":  {110, 1000, 20, 3},
		"tinggi sedikit panas":    {110, 1000, 30, 3},
		"rendah sedang dingin":    {90, 2500, 20, 3},
		"rendah sedang optimal":   {90, 2500, 25, 3},
		"rendah sedang panas":     {90, 2500, 30, 3},
		"optimal sedang dingin":   {100, 2500, 20, 1},
		"optimal sedang optimal":  {100, 2500, 25, 1},
		"optimal sedang panas":    {100, 2500, 30, 3},
		"tinggi sedang dingin":    {110, 2500, 20, 3},
		"tinggi sedang optimal":   {110, 2500, 25, 3},
		"tinggi sedang panas":     {110, 2500, 30, 3},
		"rendah banyak dingin":    {90, 3000, 20, 3},
		"rendah banyak optimal":   {90, 3000, 25, 3},
		"rendah banyak panas":     {90, 3000, 30, 3},
		"optimal banyak dingin":   {100, 3000, 20, 1},
		"optimal banyak optimal":  {100, 3000, 25, 1},
		"optimal banyak panas":    {100, 3000, 30, 1},
		"tinggi banyak dingin":    {110, 3000, 20, 1},
		"tinggi banyak optimal":   {110, 3000, 25, 1},
		"tinggi banyak panas":     {110, 3000, 30, 1},
	}

	for k, v := range data {
		value := fis.Inference(v[2], v[0], v[1])
		t.Logf("%s: fuzzy value is %f", k, value)
		if value != v[3] {
			t.Error("invalid value")
		}
	}
}

func TestScopeValue(t *testing.T) {
	fis := getFis()

	var nutrition, volume, temperature float32 = 100, 2500, 20

	for range 11 {
		value := fis.Inference(temperature, nutrition, volume)
		t.Logf("%f, %f, %f : fuzzy value is %f", nutrition, volume, temperature, value)

		nutrition += 1
		volume += 50
		temperature += 0.5
	}
}

func getFis() fuzzy.TsukamotoFIS {
	return fuzzy.TsukamotoFIS{
		LowNutritionTarget:      90,
		OptimalNutritionTarget:  100,
		HighNutritionTarget:     110,
		LowTemperatureValue:     20,
		OptimalTemperatureValue: 25,
		HighTemperatureValue:    30,
		LowWaterVolume:          1000,
		MediumWaterVolume:       2500,
		HighWaterVolume:         3000,
	}
}
