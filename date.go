// Date 内置日期对象,不兼容JavaScript中的Date规范。

package ss

import (
	"strings"
	"time"
	_ "time/tzdata"
)

type Date struct {
	Data time.Time
}

var _ JavaScript = &Date{}

func init() {
	setGlobalObject("Date", initDate())
}

func initDate() JavaScript {
	d := NewDefaultClass("Date", newDate())
	return d
}

func GoToDate(time time.Time) *Date {
	date := &Date{Data: time}
	return date
}

func (d *Date) parse() JavaScript {
	return NewDefaultFunction("parse", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return NewNaN()
		}
		t, e := parseDate(JsToString(args[0]))
		if e != nil {
			return NewNaN()
		}
		return GoToNumber(float64(t.Unix() * 1000))
	})
}

func newDate() func(args []JavaScript, ctx *Context) JavaScript {
	return func(args []JavaScript, ctx *Context) JavaScript {
		var t time.Time
		if len(args) == 0 {
			t = time.Now()
		} else if len(args) == 1 {
			var e error
			tp := JsToType(args[0])
			if tp.Type == StringType {
				t, e = parseDate(args[0].(*String).Data)
			} else if tp.Type == NumberType {
				t, e = parseTimeStamp(args[0].(*Number).Data)
			}
			if e != nil {
				panic(NewError("Error", e.Error()))
			}
		} else {
			var e error
			t, e = time.Parse(format(JsToString(args[1])), JsToString(args[0]))
			if e != nil {
				panic(NewError("Error", e.Error()))
			}
		}
		return GoToDate(t)
	}
}

func dateFormat(d *Date) JavaScript {
	return NewDefaultFunction("format", func(args []JavaScript, ctx *Context) JavaScript {
		if len(args) == 0 {
			return GoToString(d.Data.Format(time.RFC3339))
		} else {
			return GoToString(d.Data.Format(format(JsToString(args[0]))))
		}
	})
}

func dateGetDate(d *Date) JavaScript {
	return NewDefaultFunction("getDate", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Day()))
	})
}

func dateGetFullYear(d *Date) JavaScript {
	return NewDefaultFunction("getFullYear", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Year()))
	})
}

func dateGetMonth(d *Date) JavaScript {
	return NewDefaultFunction("getMonth", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Month()))
	})
}

func dateGetDay(d *Date) JavaScript {
	return NewDefaultFunction("getDay", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Weekday()))
	})
}

func dateGetHours(d *Date) JavaScript {
	return NewDefaultFunction("getHours", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Hour()))
	})
}

func dateGetMinutes(d *Date) JavaScript {
	return NewDefaultFunction("getMinutes", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Minute()))
	})
}
func dateGetSeconds(d *Date) JavaScript {
	return NewDefaultFunction("getSeconds", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Second()))
	})
}

func dateGetMilliseconds(d *Date) JavaScript {
	return NewDefaultFunction("getMilliseconds", func(args []JavaScript, ctx *Context) JavaScript {
		return GoToNumber(float64(d.Data.Nanosecond() / 1000000))
	})
}

func dateGetTimezoneOffset(d *Date) JavaScript {
	return NewDefaultFunction("getTimezoneOffset", func(args []JavaScript, ctx *Context) JavaScript {
		_, offset := d.Data.Zone()
		timezoneOffset := offset / 60
		return GoToNumber(float64(-timezoneOffset))
	})
}

func dateGetTime(d *Date) JavaScript {
	return NewDefaultFunction("getTime", func(args []JavaScript, ctx *Context) JavaScript {
		unixMillis := d.Data.UnixNano() / 1000000
		return GoToNumber(float64(unixMillis))
	})
}
func format(format string) string {
	format = strings.ReplaceAll(format, "yyyy", "2006")
	format = strings.ReplaceAll(format, "MM", "01")
	format = strings.ReplaceAll(format, "dd", "02")
	format = strings.ReplaceAll(format, "HH", "15")
	format = strings.ReplaceAll(format, "mm", "04")
	format = strings.ReplaceAll(format, "ss", "05")
	return format
}

func parseDate(date string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
	}
	var parseErr error
	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t, nil
		}
		parseErr = err
	}
	return time.Now(), parseErr
}

func parseTimeStamp(timeStamp float64) (time.Time, error) {
	var seconds int64
	var nanoseconds int64
	if timeStamp > 1e10 { // 假设大于这个值的是毫秒
		seconds = int64(timeStamp / 1000)
		nanoseconds = int64((timeStamp/1000 - float64(seconds)) * 1e9)
	} else {
		seconds = int64(timeStamp)
		nanoseconds = int64((timeStamp - float64(seconds)) * 1e9)
	}
	t := time.Unix(seconds, nanoseconds)
	return t, nil
}

func (d *Date) GetName() string {
	return "date"
}
func (d *Date) GetProperty(name string) JavaScript {
	switch name {
	case "format":
		return dateFormat(d)
	case "getDate":
		return dateGetDate(d)
	case "getMonth":
		return dateGetMonth(d)
	case "getFullYear":
		return dateGetFullYear(d)
	case "getDay":
		return dateGetDate(d)
	case "getHours":
		return dateGetHours(d)
	case "getMinutes":
		return dateGetMinutes(d)
	case "getSeconds":
		return dateGetSeconds(d)
	case "getMilliseconds":
		return dateGetMilliseconds(d)
	case "getTimezoneOffset":
		return dateGetTimezoneOffset(d)
	case "getTime":
		return dateGetTime(d)
	case "toString":
		return NewDefaultFunction("toString", func(args []JavaScript, ctx *Context) JavaScript {
			return GoToString(d.ToString())
		})
	default:
		return NewUndefined()
	}
}
func (d *Date) ToString() string {
	return d.Data.Format("2006-01-02 15:04:05")
}
func (d *Date) Print() string {
	return PurpleStyle + "[object date]" + EndStyle + d.Data.Format("2006-01-02 15:04:05")
}
func (d *Date) ToNumber() float64 {
	return float64(d.Data.Unix())
}
func (d *Date) Calculate(v JavaScript, operator Operate) JavaScript {
	if operator == Plus {
		return GoToString(d.ToString() + JsToString(v))
	}
	t := d.Data.UnixNano() / 1000000
	var t2 int64
	switch c := v.(type) {
	case *Date:
		t2 = c.Data.UnixNano() / 1000000
	default:
		return GoToBoolean(false)
	}
	switch operator {
	case LessThan:
		return GoToBoolean(t < t2)
	case MoreThan:
		return GoToBoolean(t > t2)
	case Equals:
		return GoToBoolean(t == t2)
	case NotEquals:
		return GoToBoolean(t != t2)
	case IdentityEquals:
		return GoToBoolean(t == t2)
	case IdentityNotEquals:
		return GoToBoolean(t != t2)
	default:
		return GoToBoolean(false)
	}
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Data.Format("2006-01-02 15:04:05") + "\""), nil
}
