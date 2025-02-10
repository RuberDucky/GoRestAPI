package config

func GetDBURL() string {
	dbConfig := AppConfig.Database
	return "host=" + dbConfig.Host +
		" user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.DBName +
		" port=" + dbConfig.Port +
		" sslmode=disable"
} 