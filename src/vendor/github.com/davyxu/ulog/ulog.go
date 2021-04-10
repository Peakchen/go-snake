package ulog

type Fields map[string]interface{}

var (
	TextTimeFormat = "2006-01-02 15:04:05"

	// Elastic时间标准ISO8601
	JsonTimeFormat = "2006-01-02T15:04:05.000Z"
)

type Formatter interface {
	Format(*Entry) ([]byte, error)
}
