package utils

import "os"

// EnvironmentVariable represents the name of an environment variable used to configure Photoview
type EnvironmentVariable string

// General options
const (
	EnvDevelopmentMode           EnvironmentVariable = "PHOTOVIEW_DEVELOPMENT_MODE"
	EnvServeUI                   EnvironmentVariable = "PHOTOVIEW_SERVE_UI"
	EnvUIPath                    EnvironmentVariable = "PHOTOVIEW_UI_PATH"
	EnvMediaCachePath            EnvironmentVariable = "PHOTOVIEW_MEDIA_CACHE"
	EnvFaceRecognitionModelsPath EnvironmentVariable = "PHOTOVIEW_FACE_RECOGNITION_MODELS_PATH"
)

// Network related
const (
	EnvListenIP    EnvironmentVariable = "PHOTOVIEW_LISTEN_IP"
	EnvListenPort  EnvironmentVariable = "PHOTOVIEW_LISTEN_PORT"
	EnvAPIEndpoint EnvironmentVariable = "PHOTOVIEW_API_ENDPOINT"
	EnvUIEndpoint  EnvironmentVariable = "PHOTOVIEW_UI_ENDPOINT"
)

// Database related
const (
	EnvDatabaseDriver EnvironmentVariable = "PHOTOVIEW_DATABASE_DRIVER"
	EnvMysqlURL       EnvironmentVariable = "PHOTOVIEW_MYSQL_URL"
	EnvPostgresURL    EnvironmentVariable = "PHOTOVIEW_POSTGRES_URL"
	EnvSqlitePath     EnvironmentVariable = "PHOTOVIEW_SQLITE_PATH"
)

// GetName returns the name of the environment variable itself
func (v EnvironmentVariable) GetName() string {
	return string(v)
}

// GetValue returns the value of the environment
func (v EnvironmentVariable) GetValue() string {
	return os.Getenv(string(v))
}

// ShouldServeUI whether or not the "serve ui" option is enabled
func ShouldServeUI() bool {
	return EnvServeUI.GetValue() == "1"
}

// DevelopmentMode describes whether or not the server is running in development mode,
// and should thus print debug informations and enable other features related to developing.
func DevelopmentMode() bool {
	return EnvDevelopmentMode.GetValue() == "1"
}

// UIPath returns the value from where the static UI files are located if SERVE_UI=1
func UIPath() string {
	if path := EnvUIPath.GetValue(); path != "" {
		return path
	}

	return "./ui"
}
