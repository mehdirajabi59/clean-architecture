package config

type Config struct {
	Database struct {
		Username string
		Password string
		Add      string
		DBName   string
	}
}

var c Config

func SetConfig() error {

	//TODO Set Config
	c = Config{}
	return nil
}
