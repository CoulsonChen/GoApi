package config

type DbConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	Sslmode  string
}

func DbConfigProvider() DbConfig {
	var (
		err    error
		config = new(DbConfig)
	)

	v, _ := ConfigProvider()

	if err = v.UnmarshalKey("db", &config); err != nil {
		return DbConfig{}
	}

	return *config
}
