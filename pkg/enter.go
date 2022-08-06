package pkg

type ERApp struct {
	ERExecApp
	ERErrorsApp
	ERInfoApp
	EREnvApp
	LinuxApp
}

var ERApi ERApp

type ERErrorsApp struct{}

var ERErrors []string
var ERErrorsApi ERErrorsApp

func (e *ERErrorsApp) Do(err error) {
	ERErrors = append(ERErrors, err.Error())
}
