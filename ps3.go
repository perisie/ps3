package ps3

type Ps3 interface {
	Put(key string, data []byte) error
	Get(key string) ([]byte, error)
}

const (
	aws_access_key_id     = "AWS_ACCESS_KEY_ID"
	aws_secret_access_key = "AWS_SECRET_ACCESS_KEY"
)
