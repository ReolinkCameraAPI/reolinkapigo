package network


type RtspHandler struct {
	Username string
	Password string
	Host     string
	Port     int
	Endpoint string
}

func NewRtspHandler(username string, password string, host string, endpoint string,
	options ...func(handler *RtspHandler) error) (*RtspHandler, error) {

	rtspHandler := &RtspHandler{
		Username: username,
		Password: password,
		Host:     host,
		Endpoint: endpoint,
	}

	for _, op := range options {
		err := op(rtspHandler)
		if err != nil {
			return nil, err
		}
	}

	return rtspHandler, nil
}

func RtspHandlerOptionPort(port int) func(rh *RtspHandler) error {
	return func(rh *RtspHandler) error {
		rh.Port = port
		return nil
	}
}
