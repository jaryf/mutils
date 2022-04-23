package xf

type XfOcrRes struct {
	Header  XfOcrResHeader  `json:"header"`
	Payload XfOcrResPayload `json:"payload"`
}

type XfOcrResHeader struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Sid     string `json:"sid"`
}

type XfOcrResPayload struct {
	Result XfOcrResPayloadResult `json:"result"`
}

type XfOcrResPayloadResult struct {
	Compress string `json:"compress"`
	Encoding string `json:"encoding"`
	Format   string `json:"format"`
	Text     string `json:"text"`
}

// XfOcrResPage 页面集合
type XfOcrResPage struct {
	Category string  `json:"category"` // 附加信息
	Version  string  `json:"version"`  // 引擎版本号
	Pages    []Pages `json:"pages"`    // 页面集合
}

type Coord struct {
	X int `json:"x"` // 文本行坐标4个顶点x轴的位置信息
	Y int `json:"y"` // 文本行坐标4个顶点y轴的位置信息
}

type Words struct {
	Content string  `json:"content"` // 识别结果文本
	Conf    float64 `json:"conf"`    // 置信度，取值范围[0-1]
	Coord   []Coord `json:"coord"`   // 单词坐标，记录4个顶点位置
}

type CenterPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type WordUnits struct {
	Content     string      `json:"content"`      // 字符（中文单字，英文单个字母）
	Conf        float64     `json:"conf"`         // 置信度，取值范围[0-1]
	Coord       []Coord     `json:"coord"`        // 单字坐标，记录4个顶点位置
	CenterPoint CenterPoint `json:"center_point"` // 单字中心点的坐标
}

type Lines struct {
	Coord     []Coord     `json:"coord"`      // 文本行坐标，记录4个顶点位置
	Exception int         `json:"exception"`  // 正常返回0 异常返回-1
	Words     []Words     `json:"words"`      // 单词集合
	Conf      float64     `json:"conf"`       // 置信度，取值范围[0-1]
	WordUnits []WordUnits `json:"word_units"` // 单字集合
	Angle     float64     `json:"angle"`      // 文本行的旋转角度
}

type Pages struct {
	Lines     []Lines `json:"lines"`     // 文本行集合
	Exception int     `json:"exception"` // 正常返回 0 异常返回 -1
	Angle     float64 `json:"angle"`     // 图像的旋转角度
	Height    int     `json:"height"`    // 图片高
	Width     int     `json:"width"`     // 图片宽
}
