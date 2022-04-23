package xf

type Url struct {
	Host   string
	Path   string
	Schema string
}


type XfOcrReqBody struct {
	Header    XfOcrReqHeader    `json:"header"`
	Parameter XfOcrReqParameter `json:"parameter"`
	Payload   XfOcrReqPayload   `json:"payload"`
}

type XfOcrReqHeader struct {
	AppID  string `json:"app_id"`
	Status int    `json:"status"`
}

type Result struct {
	Encoding string `json:"encoding"`
	Compress string `json:"compress"`
	Format   string `json:"format"`
}

type Sf8E6Aca1 struct {
	Category string `json:"category"`
	Result   Result `json:"result"`
}

type XfOcrReqParameter struct {
	Sf8E6Aca1 Sf8E6Aca1 `json:"sf8e6aca1"`
}

type Sf8E6Aca1Data1 struct {
	Encoding string `json:"encoding"`
	Status   int    `json:"status"`
	Image    string `json:"image"`
}

type XfOcrReqPayload struct {
	Sf8E6Aca1Data1 Sf8E6Aca1Data1 `json:"sf8e6aca1_data_1"`
}
