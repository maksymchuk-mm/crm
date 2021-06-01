package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	type env struct {
		postgresHost     string
		postgresPort     string
		postgresDbName   string
		postgresUserName string
		postgresPassword string
		postgresSSL      bool
		passwordTTL      time.Duration

		jwtSigningKey string
		passwordSalt  string
		host          string
		appEnv        string
	}
	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("POSTGRES_PASSWORD", env.postgresPassword)
		os.Setenv("PASSWORD_SALT", env.passwordSalt)
		os.Setenv("HTTP_HOST", env.host)
		os.Setenv("JWT_SIGNING_KEY", env.jwtSigningKey)
		os.Setenv("APP_ENV", env.appEnv)
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				path: "fixtures",
				env: env{
					postgresHost:     "localhost",
					postgresPort:     "5432",
					postgresDbName:   "test",
					postgresUserName: "postgres",
					postgresPassword: "postgres",
					postgresSSL:      false,
					jwtSigningKey:    "key",
					host:             "localhost",
					appEnv:           "local",
					passwordSalt:     "salt",
					passwordTTL:      time.Minute * 5,
				},
			},
			want: &Config{
				Environment: "local",
				Postgres: PostgresConfig{
					Host:     "localhost",
					Port:     "5432",
					UserName: "postgres",
					Password: "postgres",
					SSL:      false,
					DbName:   "test",
				},
				HTTP: HTTPConfig{
					Host:               "localhost",
					Port:               "8000",
					ReadTimeout:        time.Second * 10,
					WriteTimeout:       time.Second * 10,
					MaxHeaderMegabytes: 1,
				},
				Auth: AuthConfig{
					JWT: JWTConfig{
						AccessTokenTTL:  time.Minute * 15,
						RefreshTokenTTL: time.Minute * 60,
						SigningKey:      "key",
					},
					PasswordSalt:           "salt",
					VerificationCodeLength: 10,
					PasswordTTL:            time.Minute * 5,
				},
				CacheTTL: time.Second * 3600,
				Limiter: LimiterConfig{
					RPS:   10,
					Burst: 2,
					TTL:   time.Minute * 10,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)
			got, err := Init(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}
