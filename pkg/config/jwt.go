package config

type JwtConfig struct {
	JwtSecretKey           string
	JwtTokenExpireDuration int
}

func JwtConfigProvider() JwtConfig {
	var (
		err    error
		config = new(JwtConfig)
	)

	v, _ := ConfigProvider()

	if err = v.UnmarshalKey("jwt", &config); err != nil {
		return JwtConfig{}
	}

	return *config
}
