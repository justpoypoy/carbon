package carbon

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Parse 将标准格式时间字符串解析成 Carbon 实例
func (c Carbon) Parse(value string) Carbon {
	if c.Error != nil {
		return c
	}

	layout := DateTimeFormat

	if value == "" || value == "0" || value == "0000-00-00 00:00:00" || value == "0000-00-00" || value == "00:00:00" {
		return c
	}

	if len(value) == 10 && strings.Count(value, "-") == 2 {
		layout = DateFormat
	}

	if strings.Index(value, "T") == 10 {
		layout = RFC3339Format
	}

	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		switch len(value) {
		case 8:
			layout = ShortDateFormat
		case 14:
			layout = ShortDateTimeFormat
		}
	}

	c.Time, c.Error = c.parseByLayout(value, layout)

	return c
}

// Parse 将标准格式时间字符串解析成 Carbon 实例(默认时区)
func Parse(value string) Carbon {
	return SetTimezone(Local).Parse(value)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例
func (c Carbon) ParseByFormat(value string, format string) Carbon {
	if c.Error != nil {
		return c
	}
	layout := format2layout(format)
	return c.ParseByLayout(value, layout)
}

// ParseByFormat 将特殊格式时间字符串解析成 Carbon 实例(默认时区)
func ParseByFormat(value string, format string) Carbon {
	return SetTimezone(Local).ParseByFormat(value, format)
}

// ParseByLayout 将布局时间字符串解析成 Carbon 实例
func (c Carbon) ParseByLayout(value string, layout string) Carbon {
	if c.Error != nil {
		return c
	}
	c.Time, c.Error = c.parseByLayout(value, layout)

	return c
}

// parseByLayout 通过布局模板解析
func (c Carbon) parseByLayout(value string, layout string) (time.Time, error) {
	if c.Loc == nil {
		c.Loc, _ = time.LoadLocation(Local)
	}

	tt, err := time.ParseInLocation(layout, value, c.Loc)
	if err != nil {
		err = errors.New("the value \"" + value + "\" can't parse string as time")
	}
	return tt, err
}


// ParseByLayout 将布局时间字符串解析成 Carbon 实例(默认时区)
func ParseByLayout(value string, layout string) Carbon {
	return SetTimezone(Local).ParseByLayout(value, layout)
}