package epub

//Ncx OPS/toc.ncx
type Ncx struct {
	Points   []NavPoint `xml:"navMap>navPoint" json:"points"`
	PageList PageList   `xml:"pageList" json:"page_list"`
}

//NavPoint nav point
type NavPoint struct {
	Text        string     `xml:"navLabel>text" json:"text"`
	Content     Content    `xml:"content" json:"content"`
	Points      []NavPoint `xml:"navPoint" json:"points"`
	PlayerOrder int        `xml:"playOrder" json:"play_order"`
}

//Content nav-point content
type Content struct {
	Src string `xml:"src,attr" json:"src"`
}

// PageList page list
type PageList struct {
	PageTarget []PageTarget `xml:"pageTarget" json:"page_targets"`
	Class      string       `xml:"class" json:"class"`
	ID         string       `xml:"id" json:"id"`
}

// PageTarget page target
type PageTarget struct {
	Text      string  `xml:"navLabel>text" json:"text"`
	Value     string  `xml:"value" json:"value"`
	Type      string  `xml:"type" json:"type"`
	PlayOrder int     `xml:"playOrder" json:"play_order"`
	Content   Content `xml:"content" json:"content"`
}
