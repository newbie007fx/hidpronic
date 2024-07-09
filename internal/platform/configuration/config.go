package configuration

type Config struct {
	App        ConfigApp         `mapstructure:"app"`
	Server     ConfigServer      `mapstructure:"server"`
	Database   ConfigDatabase    `mapstructure:"database"`
	CorsConfig CorsConfiguration `mapstructure:"cors_config"`
	MqttConfig ConfigMqtt        `mapstructure:"mqtt"`
	JWT        JWTConfiguration  `mapstructure:"jwt"`
}

type ConfigApp struct {
	Env      string `mapstructure:"env"`
	LogLevel string `mapstructure:"log_level"`
}

type ConfigServer struct {
	Port int `mapstructure:"port"`
}

type ConfigDatabase struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
}

type ConfigMqtt struct {
	Protocol string `mapstructure:"protocol"`
	Port     int    `mapstructure:"port"`
	Broker   string `mapstructure:"broker"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type CorsConfiguration struct {
	AllowedHeaders []string `mapstructure:"allowed_headers"`
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
}

type JWTConfiguration struct {
	Secret string `mapstructure:"secret"`
}
