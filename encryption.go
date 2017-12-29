package epub

//Encryption encryption.xml
type Encryption struct {
	EncryptedData []EncryptedData `xml:"EncryptedData" json:"encrypted_data"`
}

//EncryptedData provides encryption information
type EncryptedData struct {
	EncryptionMethod     EncryptionMethod     `xml:"EncryptionMethod"`
	KeyInfo              KeyInfo              `xml:"KeyInfo"`
	CipherData           CipherData           `xml:"CipherData"`
	EncryptionProperties []EncryptionProperty `xml:"EncryptionProperties>EncryptionProperty"`
}

// EncryptionProperty provides encryption compression information
type EncryptionProperty struct {
	Compression Compression `xml:"Compression"`
}

// Compression provides encryption compression details
type Compression struct {
	Method         string `xml:"Method,attr"`
	OriginalLength string `xml:"OriginalLength,attr"`
}

// EncryptionMethod provides the encryption algorithm
type EncryptionMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// KeyInfo provides the encryption key details
type KeyInfo struct {
	Resource string `xml:",chardata"`
}

// CipherData provides the encryption cipher information
type CipherData struct {
	CipherReference CipherReference `xml:"CipherReference"`
}

// CipherReference provides the encryption
type CipherReference struct {
	URI string `xml:"URI,attr"`
}
