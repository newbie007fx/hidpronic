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

	value = fis.Inference(22, 95, 1001)
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

	value := fis.Inference(26, 104, 1350)
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

	value := fis.Inference(21, 102, 2550)
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
