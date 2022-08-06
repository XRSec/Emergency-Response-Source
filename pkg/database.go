package pkg

import "os"

type Syslog struct{}

type Linux struct{}

var LinuxSyslog = []string{
	"/var/log/syslog",
	"",
}

type MacOS struct{}

type Windows struct{}

type IOT struct{}

type EREnvApp struct{}

var EREnvApi EREnvApp

//var SystemOS string

func (e *EREnvApp) Get(arg string) string { return os.Getenv(arg) }

func (e *EREnvApp) Set(arg string, value string) {
	err := os.Setenv(arg, value)
	if err != nil {
		ERErrorsApi.Do(err)
	}
}
