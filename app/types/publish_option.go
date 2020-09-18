package types

type PublishOption struct {
	Exchange    string
	Key         string
	Mandatory   bool
	Immediate   bool
	ContentType string
	Body        []byte
}
