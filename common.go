package spold2

// TextAndImage is a list of text, imageUri and variable elements.
type TextAndImage struct {
	Texts []Text `xml:"text,omitempty"`
	// TODO: images
}

// Text is a free text used to describe the current section.
// Its index attribute is used to enforce a specific order
// of text and image elements.
type Text struct {
	Index int    `xml:"index,attr"`
	Value string `xml:",chardata"`
}
