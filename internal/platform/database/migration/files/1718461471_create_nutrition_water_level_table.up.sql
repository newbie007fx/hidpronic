CREATE TABLE IF NOT EXISTS nutrition_water_levels (
  id SERIAL PRIMARY KEY, 
  plant_id INTEGER NOT NULL, 
  value FLOAT(32) NOT NULL, 
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE nutrition_water_levels_temp (LIKE nutrition_water_levels INCLUDING ALL);
