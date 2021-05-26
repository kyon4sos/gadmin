package exception

type ErrParam struct {
	Code int
	Msg  string
}

func (e *ErrParam) Error() string {
	return string(e.Msg)
}

func NewError(code int,msg string) *ErrParam{
	err :=&ErrParam{code, msg}
	return err
}