package fuzzy

import (
	"slices"
)

const RESULT_MIN_VALUE float32 = 0
const RESULT_MAX_VALUE float32 = 1

type TsukamotoFIS struct {
	LowNutritionTarget      float32
	OptimalNutritionTarget  float32
	HighNutritionTarget     float32
	LowTemperatureValue     float32
	OptimalTemperatureValue float32
	HighTemperatureValue    float32
	LowWaterVolume          float32
	MediumWaterVolume       float32
	HighWaterVolume         float32
}

func (f TsukamotoFIS) CalculateOptimalTemperature(temperature float32) float32 {
	if temperature >= f.LowTemperatureValue && temperature <= f.OptimalTemperatureValue {
		return (temperature - f.LowTemperatureValue) / (f.OptimalTemperatureValue - f.LowTemperatureValue)
	} else if temperature >= f.OptimalTemperatureValue && temperature <= f.HighTemperatureValue {
		return (f.HighTemperatureValue - temperature) / (f.HighTemperatureValue - f.OptimalTemperatureValue)
	}

	return 0
}

func (f TsukamotoFIS) CalculateHighTemperature(temperature float32) float32 {
	if temperature >= f.HighTemperatureValue {
		return 1
	} else if temperature >= f.OptimalTemperatureValue && temperature <= f.HighTemperatureValue {
		return (temperature - f.OptimalTemperatureValue) / (f.HighTemperatureValue - f.OptimalTemperatureValue)
	}

	return 0
}

func (f TsukamotoFIS) CalculateLowNutrition(nutritionTarget float32) float32 {
	if nutritionTarget <= f.LowNutritionTarget {
		return 1
	} else if nutritionTarget >= f.LowNutritionTarget && nutritionTarget <= f.OptimalNutritionTarget {
		return (f.HighNutritionTarget - nutritionTarget) / (f.OptimalNutritionTarget - f.LowNutritionTarget)
	}

	return 0
}

func (f TsukamotoFIS) CalculateOptimalNutrition(nutritionTarget float32) float32 {
	if nutritionTarget >= f.LowNutritionTarget && nutritionTarget <= f.OptimalNutritionTarget {
		return (nutritionTarget - f.LowNutritionTarget) / (f.OptimalNutritionTarget - f.LowNutritionTarget)
	} else if nutritionTarget >= f.OptimalNutritionTarget && nutritionTarget <= f.HighNutritionTarget {
		return (f.HighNutritionTarget - nutritionTarget) / (f.HighNutritionTarget - f.OptimalNutritionTarget)
	}

	return 0
}

func (f TsukamotoFIS) CalculateHighNutrition(nutritionTarget float32) float32 {
	if nutritionTarget >= f.HighNutritionTarget {
		return 1
	} else if nutritionTarget >= f.OptimalNutritionTarget && nutritionTarget <= f.HighNutritionTarget {
		return (nutritionTarget - f.OptimalNutritionTarget) / (f.HighNutritionTarget - f.OptimalNutritionTarget)
	}

	return 0
}

func (f TsukamotoFIS) CalculateLowWaterVolume(waterVolume float32) float32 {
	if waterVolume <= f.LowWaterVolume {
		return 1
	} else if waterVolume >= f.LowWaterVolume && waterVolume <= f.MediumWaterVolume {
		return (f.HighWaterVolume - waterVolume) / (f.MediumWaterVolume - f.LowWaterVolume)
	}

	return 0
}

func (f TsukamotoFIS) CalculateMediumWaterVolume(waterVolume float32) float32 {
	if waterVolume >= f.LowWaterVolume && waterVolume <= f.MediumWaterVolume {
		return (waterVolume - f.LowWaterVolume) / (f.MediumWaterVolume - f.LowWaterVolume)
	} else if waterVolume >= f.MediumWaterVolume && waterVolume <= f.HighWaterVolume {
		return (f.HighWaterVolume - waterVolume) / (f.HighWaterVolume - f.MediumWaterVolume)
	}

	return 0
}

func (f TsukamotoFIS) CalculateHighWaterVolume(waterVolume float32) float32 {
	if waterVolume >= f.HighWaterVolume {
		return 1
	} else if waterVolume >= f.MediumWaterVolume && waterVolume <= f.HighWaterVolume {
		return (waterVolume - f.MediumWaterVolume) / (f.HighWaterVolume - f.MediumWaterVolume)
	}

	return 0
}

func (f TsukamotoFIS) Inference(temperature, nutrition, volume float32) float32 {
	rules := [][]float32{
		{f.CalculateOptimalTemperature(temperature), f.CalculateLowNutrition(nutrition), f.CalculateLowWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateLowNutrition(nutrition), f.CalculateLowWaterVolume(volume)},
		{f.CalculateOptimalTemperature(temperature), f.CalculateOptimalNutrition(nutrition), f.CalculateLowWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateOptimalNutrition(nutrition), f.CalculateLowWaterVolume(volume)},
		{f.CalculateOptimalTemperature(temperature), f.CalculateHighNutrition(nutrition), f.CalculateLowWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateHighNutrition(nutrition), f.CalculateLowWaterVolume(volume)},

		{f.CalculateOptimalTemperature(temperature), f.CalculateLowNutrition(nutrition), f.CalculateMediumWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateLowNutrition(nutrition), f.CalculateMediumWaterVolume(volume)},
		{f.CalculateOptimalTemperature(temperature), f.CalculateOptimalNutrition(nutrition), f.CalculateMediumWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateOptimalNutrition(nutrition), f.CalculateMediumWaterVolume(volume)},
		{f.CalculateOptimalTemperature(temperature), f.CalculateHighNutrition(nutrition), f.CalculateMediumWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateHighNutrition(nutrition), f.CalculateMediumWaterVolume(volume)},

		{f.CalculateOptimalTemperature(temperature), f.CalculateLowNutrition(nutrition), f.CalculateHighWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateLowNutrition(nutrition), f.CalculateHighWaterVolume(volume)},
		{f.CalculateOptimalTemperature(temperature), f.CalculateOptimalNutrition(nutrition), f.CalculateHighWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateOptimalNutrition(nutrition), f.CalculateHighWaterVolume(volume)},
		{f.CalculateOptimalTemperature(temperature), f.CalculateHighNutrition(nutrition), f.CalculateHighWaterVolume(volume)},
		{f.CalculateHighTemperature(temperature), f.CalculateHighNutrition(nutrition), f.CalculateHighWaterVolume(volume)},
	}

	var numerator, denominator float32
	for _, rule := range rules {
		alphaPredicate := slices.Min(rule)
		zRule := alphaPredicate - RESULT_MIN_VALUE/RESULT_MAX_VALUE
		numerator += alphaPredicate * zRule
		denominator += alphaPredicate
	}

	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}
