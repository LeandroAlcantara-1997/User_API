package postgresconfig

const (
	ConfigPrefix = "POSTGRES_"
)

type Config struct {
	URL string `env:"URL, required"`
}
