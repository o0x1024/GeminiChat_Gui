package global

type Config struct {
	APIKEY string `yaml:"apikey"`
	Proxy  string `yaml:"proxy"`
}

var (
	Conf = Config{}
)
