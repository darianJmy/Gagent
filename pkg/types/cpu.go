package types

type CPU struct {
	Number int
}

type Network struct {
	Name string
	IP   string `json:"ip,omitempty"`
}

type System struct {
	UUID string
}

type Resp struct {
	Code int
	Msg  interface{}
}
