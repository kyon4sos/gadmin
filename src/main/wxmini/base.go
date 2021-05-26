package wxmini

type WxRepErr struct {
	Errcode int `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	err error
}

func (w *WxRepErr) Error() string {
	return w.Errmsg
}
func (w *WxRepErr) Unwrap() error {
	return w.err
}
