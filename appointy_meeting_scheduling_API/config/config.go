package scheduling

type scheduling struct {
	DB *DBscheduling
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *scheduling {
	return &scheduling{
		DB: &DBscheduling{
			Dialect:  "MongoDB",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "Shristi",
			Password: "Deeplearning(16)",
			Name:     "appointy_meeting_scheduler",
			Charset:  "utf8",
		},
	}
}
