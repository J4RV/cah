package cah

// Config stores the app's config
type Config struct {
	TemplatePath   string `yaml:"templatePath"`
	StaticPath     string `yaml:"staticPath"`
	ExpansionsPath string `yaml:"expansionsPath"`
	SQLiteDBPath   string `yaml:"sqliteDBPath"`
}
