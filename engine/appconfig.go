package engine

var appName, appVersion string

// Returns name of the application
func GetAppName() string {
	return appName
}

func SetAppName(name string) {
	appName = name
}

// Returns version of the application
func GetAppVersion() string {
	return appVersion
}

func SetAppVersion(version string) {
	appVersion = version
}
