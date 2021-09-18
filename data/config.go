package data

type Config struct {
	Apps []App `yaml:"apps"`
}

type App struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}
