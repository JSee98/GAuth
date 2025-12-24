package db

type DBConfig struct {
	Driver          string
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

func (cfg *DBConfig) Validate() []string {
	var errors []string
	if cfg.Driver == "" {
		errors = append(errors, "DB Driver is required")
	}
	if cfg.Host == "" {
		errors = append(errors, "DB Host is required")
	}
	if cfg.Port <= 0 {
		errors = append(errors, "DB Port must be greater than 0")
	}
	if cfg.User == "" {
		errors = append(errors, "DB User is required")
	}
	if cfg.Password == "" {
		errors = append(errors, "DB Password is required")
	}
	if cfg.Name == "" {
		errors = append(errors, "DB Name is required")
	}
	return errors
}
