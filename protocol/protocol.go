package protocol

const (
	Auth = byte(1)
)

func AuthProtocol() []byte {
	return []byte{Auth}
}
