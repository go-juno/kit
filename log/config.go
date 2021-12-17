package log

type Level string

const (
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
	ErrorLevel Level = "error"
	WarnLevel  Level = "warn"
	InfoLevel  Level = "info"
	DebugLevel Level = "debug"
	TraceLevel Level = "trace"
)

type loggerConfig struct {
	Path []string // 配置文件路径
	Name string   // 文件名称
	Type string   // 文件类型,默认 yaml
	File string   // 配置文件路径,设置后,将覆盖 Path,Name,Type 等配置
}

var defaultConfigPath = []string{".", "./conf", "./etc", "./configs", "/etc/log"}
var defaultConfigName = "logger"
var defaultConfigType = "yaml"
