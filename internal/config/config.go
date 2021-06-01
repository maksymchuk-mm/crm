package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHttpHost               = "0.0.0.0"
	defaultHttpPort               = "8000"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 2
	defaultLimiterTTL             = 10 * time.Minute
	defaultVerificationCodeLength = 8

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		HTTP        HTTPConfig
		Auth        AuthConfig
		CacheTTL    time.Duration `mapstructure:"ttl"`
		Limiter     LimiterConfig
	}
	PostgresConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		UserName string `mapstructure:"username"`
		Password string
		DbName   string `mapstructure:"dataBaseName"`
		SSL      bool   `mapstructure:"ssl"`
		Debug    bool   `mapstructure:"debug"`
	}

	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string
		VerificationCodeLength int           `mapstructure:"verificationCodeLength"`
		PasswordTTL            time.Duration `mapstructure:"passwordTTL"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}
	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
	LimiterConfig struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}
)

func Init(configsDir string) (*Config, error) {
	populateDefaults()
	if err := parseEnv(); err != nil {
		return nil, err
	}
	if err := parseConfigFile(configsDir, viper.GetString("env")); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func populateDefaults() {
	viper.SetDefault("http.host", defaultHttpHost)
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.maxHeaderBytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.readTimeout", defaultHttpRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHttpRWTimeout)
	viper.SetDefault("jwt.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("jwt.refreshTokenTTL", defaultRefreshTokenTTL)
	viper.SetDefault("jwt.verificationCodeLength", defaultVerificationCodeLength)
	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("jwt", &cfg.Auth.JWT); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("jwt.verificationCodeLength", &cfg.Auth.VerificationCodeLength); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("limiter", &cfg.Limiter); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("user.passwordTTL", &cfg.Auth.PasswordTTL); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.Password = viper.GetString("password")

	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signing_key")

	cfg.Environment = viper.GetString("env")
}

func parseEnv() error {
	if err := parsePostgresEnvVariables(); err != nil {
		return err
	}

	if err := parsePasswordFromEnv(); err != nil {
		return err
	}

	if err := parseJWTFromEnv(); err != nil {
		return err
	}

	if err := parsAppEnvFromEnv(); err != nil {
		return err
	}

	return nil
}

func parsePostgresEnvVariables() error {
	viper.SetEnvPrefix("postgres")
	return viper.BindEnv("password")
}

func parsePasswordFromEnv() error {
	viper.SetEnvPrefix("password")
	return viper.BindEnv("salt")
}

func parseJWTFromEnv() error {
	viper.SetEnvPrefix("jwt")
	return viper.BindEnv("signing_key")
}

func parsAppEnvFromEnv() error {
	viper.SetEnvPrefix("app")
	return viper.BindEnv("env")
}
