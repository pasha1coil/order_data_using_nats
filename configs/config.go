package configs

type Config struct {
	Nats struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Cluster_id string `yaml:"cluster_id"`
		Client_id  string `yaml:"client_id"`
		Channel    string `yaml:"channel"`
	} `yaml:"nats"`
	Db struct {
		Uname      string `yaml:"Uname"`
		Pass       string `yaml:"Pass"`
		NameDB     string `yaml:"NameDB"`
		Host       string `yaml:"Host"`
		Port       string `yaml:"Port"`
		SSL        string `yaml:"SSL"`
		DriverName string `yaml:"DriverName"`
	} `yaml:"db"`
	Httpsrv struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"httpsrv"`
}
