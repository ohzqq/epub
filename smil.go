package epub

// SMIL or Synchronized Multimedia Integration Language, wrapper
type SMIL struct {
	Body Body `xml:"body"`
}

// Body provides the SMIL body element
type Body struct {
	TextRef string `xml:"textref,attr"`
	Seq     []Seq  `xml:"seq"`
	Par     []Par  `xml:"par"`
}

// Seq provides sequential SMIL elements
type Seq struct {
	TextRef string `xml:"textref,attr"`
	Par     []Par  `xml:"par"`
	Seq     []Seq  `xml:"seq"`
}

// Par provides parallel SMIL elements
type Par struct {
	Text  Text  `xml:"text"`
	Audio Audio `xml:"audio"`
}

// Text provides the source text
type Text struct {
	Src string `xml:"src,attr"`
}

// Audio provides the audio element details
type Audio struct {
	Src       string `xml:"src,attr"`
	ClipBegin string `xml:"clipBegin,attr"`
	ClipEnd   string `xml:"clipEnd,attr"`
}
