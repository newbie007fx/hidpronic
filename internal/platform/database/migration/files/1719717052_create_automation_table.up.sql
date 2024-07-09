CREATE TABLE IF NOT EXISTS automation (
    id SERIAL PRIMARY KEY,
    plant_id INTEGER NOT NULL,
    before_data jsonb NOT NULL,
    after_data jsonb NOT NULL,
    accuration FLOAT(32) NOT NULL,
    target_ppm FLOAT(32) NOT NULL,
    duration INTEGER NOT NULL,
    status VARCHAR(100) NOT NULL,
    triggered_at TIMESTAMP WITH TIME ZONE NOT NULL,
    finished_at TIMESTAMP WITH TIME ZONE NULL
);