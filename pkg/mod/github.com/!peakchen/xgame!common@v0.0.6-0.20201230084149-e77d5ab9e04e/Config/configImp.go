package Config

/*
	common load config module.
*/
type TConfig struct {
	data interface{}   //config data
	obj  ICommonConfig //config obj
}

/*
	all config module inherit interface.
*/
type ICommonConfig interface {
	ComfireAct(data interface{}) (errlist []string)
	DataRWAct(data interface{}) (errlist []string)
}

/*
	comfire data right or not before load config.
*/
func (this *TConfig) Before() (errlist []string) {
	errlist = this.obj.ComfireAct(this.data)
	return
}

/*
	if data right, then load data to related config data..
*/
func (this *TConfig) After() (errlist []string) {
	errlist = this.obj.DataRWAct(this.data)
	return
}
