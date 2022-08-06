package pkg

type LinuxApp struct {
}

func (l *LinuxApp) Do() {
	ERInfoApi.Do()
	l.Syslog()
}

func (l *LinuxApp) Syslog() {

}
