CREATE TABLE IF NOT EXISTS container_configs (
  id VARCHAR(100) PRIMARY KEY, 
  name VARCHAR (155) NOT NULL, 
  sensor_gap FLOAT(32) NOT NULL,
  height FLOAT(32) NOT NULL, 
  bottom_area FLOAT(32) NOT NULL, 
  top_area FLOAT(32) NOT NULL,
  volume FLOAT(32) NOT NULL 
);