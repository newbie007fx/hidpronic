CREATE TABLE IF NOT EXISTS installation_configs (
  id SERIAL PRIMARY KEY, 
  nutrition_ppm FLOAT(32) NOT NULL,
  raw_water_ppm FLOAT(32) NOT NULL,
  fuzzy_nutrition_water_level_percent FLOAT(32) NOT NULL,
  fuzzy_water_temperature_percent FLOAT(32) NOT NULL,
  fuzzy_nutrition_water_volume_low FLOAT(32) NOT NULL,
  fuzzy_nutrition_water_volume_medium FLOAT(32) NOT NULL,
  fuzzy_nutrition_water_volume_high FLOAT(32) NOT NULL
);