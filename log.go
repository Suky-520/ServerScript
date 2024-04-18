//log是ss内置的日志对象

package ss

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	fileLogger    zerolog.Logger
	consoleLogger zerolog.Logger
)

type logConfig struct {
	Out        int    `json:"out"`
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxSize"`
	MaxBackups int    `json:"maxBackups"`
	MaxAge     int    `json:"maxAge"`
	Compress   bool   `json:"compress"`
}

func init() {
	consoleLogger = zerolog.New(newConsoleWriter()).With().Timestamp().Logger()
	// 设置全局日志级别
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// 设置时间格式
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	//默认日志
	log.Logger = consoleLogger
	setGlobalObject("log", initLog())
}

func initLog() JavaScript {
	c := NewDefaultObject("log")
	c.AddProperty("info", logMsg("info"))
	c.AddProperty("debug", logMsg("debug"))
	c.AddProperty("warn", logMsg("warn"))
	c.AddProperty("error", logMsg("error"))
	c.AddProperty("fatal", logMsg("fatal"))
	c.AddProperty("config", logConfigFun())
	return c
}

func jsToAny(args []JavaScript) []any {
	a := make([]any, len(args))
	for i, v := range args {
		a[i] = v
	}
	return a
}
func logConfigFun() JavaScript {
	return NewDefaultFunction("config", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewUndefined()
		}
		if _, ok := args[0].(*Object); ok {
			config := initLogConfig(args)
			SetLevel(config.Level)
			if config.Out == 1 {
				ConfigFileLog(config.Filename, config.MaxSize, config.MaxBackups, config.MaxAge, config.Compress)
			}
			SetOutPut(config.Out)
		}
		return NewUndefined()
	})
}

func logMsg(level string) JavaScript {
	return NewDefaultFunction(level, func(args []JavaScript, ctx *Context) JavaScript {
		m := fmt.Sprint(jsToAny(args)...)
		switch level {
		case "info":
			JsInfo(m, ctx)
		case "debug":
			JsDebug(m, ctx)
		case "warn":
			JsWarn(m, ctx)
		case "error":
			JsError(m, ctx)
		case "fatal":
			JsFatal(m, ctx)
		default:
			JsInfo(m, ctx)
		}
		return NewUndefined()
	})
}

func newConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		FormatLevel: func(i interface{}) string {
			// 将级别转换为字符串
			var l string
			if ll, ok := i.(string); ok {
				l = ll
			}
			// 根据级别设置颜色
			switch l {
			case "debug":
				return "\033[36mDEBUG\033[0m" // 蓝绿色
			case "info":
				return "\033[32mINFO \033[0m" // 绿色
			case "warn":
				return "\033[33mWARN \033[0m" // 黄色
			case "error":
				return "\033[31mERROR\033[0m" // 红色
			case "fatal":
				return "\033[31mFATAL\033[0m" // 红色
			case "panic":
				return "\033[31mPANIC\033[0m" // 红色
			default:
				return "\033[37m" + l + "\033[0m" // 白色
			}
		},
	}
}

// SetOutPut 设置日志输出位置,0console,1file
func SetOutPut(out int) {
	if out == 1 {
		log.Logger = fileLogger
	} else {
		log.Logger = consoleLogger
	}
}

// ConfigFileLog 配置日志文件
func ConfigFileLog(filename string, maxSize int, maxBackups int, maxAge int, compress bool) {
	fileLogger = zerolog.New(&lumberjack.Logger{
		Filename:   filename,   // 日志文件的位置
		MaxSize:    maxSize,    // 文件最大大小(MB)
		MaxBackups: maxBackups, // 保留旧文件的最大个数
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   compress,   // 是否压缩/归档旧文件
	}).With().Timestamp().Logger()
}

// SetTimeFormat 设置时间格式
func SetTimeFormat(format string) {
	// 设置时间格式
	zerolog.TimeFieldFormat = format
}

// SetLevel 设置日志级别
func SetLevel(level string) {
	l, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Error().Err(err).Msg("解析日志级别失败")
	} else {
		zerolog.SetGlobalLevel(l)
	}
}
func JsDebug(msg string, ctx *Context) {
	if ctx != nil {
		src, file := ctx.GetSrcMap()
		caller := fmt.Sprintf("%s:%d", file, src.StartLine)
		log.Debug().Str("caller", caller).Msg(msg)
	} else {
		log.Debug().Msg(msg)
	}
}
func JsWarn(msg string, ctx *Context) {
	if ctx != nil {
		src, file := ctx.GetSrcMap()
		caller := fmt.Sprintf("%s:%d", file, src.StartLine)
		log.Warn().Str("caller", caller).Msg(msg)
	} else {
		log.Warn().Msg(msg)
	}
}
func JsInfo(msg string, ctx *Context) {
	if ctx != nil {
		src, file := ctx.GetSrcMap()
		caller := fmt.Sprintf("%s:%d", file, src.StartLine)
		log.Info().Str("caller", caller).Msg(msg)
	} else {
		log.Info().Msg(msg)
	}
}
func JsError(msg string, ctx *Context) {
	if ctx != nil {
		src, file := ctx.GetSrcMap()
		if src != nil {
			caller := fmt.Sprintf("%s:%d", file, src.StartLine)
			log.Error().Str("caller", caller).Msg(msg)
		} else {
			log.Error().Msg(msg)
		}
	} else {
		log.Error().Msg(msg)
	}
}
func JsFatal(msg string, ctx *Context) {
	if ctx != nil {
		src, file := ctx.GetSrcMap()
		caller := fmt.Sprintf("%s:%d", file, src.StartLine)
		log.Fatal().Str("caller", caller).Msg(msg)
	} else {
		log.Fatal().Msg(msg)
	}
}
func NativeDebug(msg string) {
	log.Debug().Caller().Msg(msg)
}
func NativeWarn(msg string) {
	log.Warn().Caller().Msg(msg)
}
func NativeInfo(msg string) {
	log.Info().Caller().Msg(msg)
}
func NativeError(msg string) {
	log.Error().Caller().Msg(msg)
}
func NativeFatal(msg string) {
	log.Fatal().Caller().Msg(msg)
}
func initLogConfig(args []JavaScript) logConfig {
	if len(args) > 0 {
		config := logConfig{}
		e := json.Unmarshal([]byte(JsToJson(args[0])), &config)
		if e != nil {
			panic(NewTypeError(e.Error()))
		}
		if config.Out != 0 && config.Out != 1 {
			config.Out = 0
		}
		if config.Out == 1 && config.Filename == "" {
			panic(NewError("Error", "日志文件名不能为空"))
		}
		return config
	}
	return logConfig{
		Out:   0,
		Level: "debug",
	}
}
