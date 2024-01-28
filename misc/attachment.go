package misc

type Attachment struct {
	Name       string `bson:"name" json:"name"`
	MimeType   string `bson:"mimetype" json:"mimetype"`
	URL        string `bson:"url" json:"url"`
	IsExternal bool   `bson:"isExternal" json:"isExternal"`
}

// IsExternal tells us if we currently manage the attachment
// on our own server or if it points to an external service
