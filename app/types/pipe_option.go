package types

type PipeOption struct {
	PublishID string `yaml:"publish_id"`
	AckID     string `yaml:"ack_id"`
	NackID    string `yaml:"nack_id"`
}
