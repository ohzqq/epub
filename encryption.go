package epub

//Encryption encruption.xml
type Encryption struct {
	EncryptedData []EncryptedData `xml:"EncryptedData" json:"encrypted_data"`
}

type EncryptedData struct {
	EncryptionMethod EncryptionMethod `xml:"EncryptionMethod"`
	KeyInfo          KeyInfo          `xml:"KeyInfo"`
	CipherData       CipherData       `xml:"CipherData"`
}

type EncryptionMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type KeyInfo struct {
	Resource string `xml:"resource,cdata"`
}

type CipherData struct {
	CipherReference CipherReference `xml:"CipherReference"`
}

type CipherReference struct {
	URI string `xml:"URI,attr"`
}
