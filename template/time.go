package template

import (
	"fmt"
	"html/template"
	"time"
)

const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func timeSince(then time.Time) string {
	now := time.Now()
	lbl := "前"

	diff := now.Unix() - then.Unix()
	if then.After(now) {
		lbl = "今から"
		diff = then.Unix() - now.Unix()
	}

	switch {
	case diff <= 0:
		return "たった今"
	case diff <= 2:
		return "1秒前"
	case diff < 1*Minute:
		return fmt.Sprintf("%d秒%s", diff, lbl)
	case diff < 2*Minute:
		return fmt.Sprintf("1分%s", lbl)
	case diff < 1*Hour:
		return fmt.Sprintf("%d分%s", diff/Minute, lbl)
	case diff < 2*Hour:
		return fmt.Sprintf("1時間%s", lbl)
	case diff < 1*Day:
		return fmt.Sprintf("%d時間%s", diff/Hour, lbl)
	case diff < 2*Day:
		return fmt.Sprintf("1日%s", lbl)
	case diff < 1*Week:
		return fmt.Sprintf("%d日%s", diff/Day, lbl)
	case diff < 2*Week:
		return fmt.Sprintf("1週間%s", lbl)
	case diff < 1*Month:
		return fmt.Sprintf("%d週間%s", diff/Week, lbl)
	case diff < 2*Month:
		return fmt.Sprintf("1ヶ月%s", lbl)
	case diff < 1*Year:
		return fmt.Sprintf("%dヶ月%s", diff/Month, lbl)
	case diff < 2*Year:
		return fmt.Sprintf("1年%s", lbl)
	default:
		return fmt.Sprintf("%d年%s", diff/Year, lbl)
	}
}

func timeSinceColor(then time.Time) string {
	now := time.Now()

	diff := now.Unix() - then.Unix()
	if then.After(now) {
		diff = then.Unix() - now.Unix()
	}

	switch {
	case diff <= 0:
		return "red"
	case diff <= 2:
		return "red"
	case diff < 1*Minute:
		return "red"
	case diff < 2*Minute:
		return "red"
	case diff < 1*Hour:
		return "salmon"
	case diff < 2*Hour:
		return "salmon"
	case diff < 1*Day:
		return "chocolate"
	case diff < 2*Day:
		return "chocolate"
	case diff < 1*Week:
		return "darkred"
	case diff < 2*Week:
		return "darkred"
	case diff < 1*Month:
		return "olivedrab"
	case diff < 2*Month:
		return "olivedrab"
	case diff < 1*Year:
		return "steelblue"
	case diff < 2*Year:
		return "steelblue"
	default:
		return "slategrey"
	}
}


func TimeSince(t time.Time) template.HTML {
	return template.HTML(
		fmt.Sprintf(
			`<span class="time-since" style="color: %s;" title="%s">%s</span>`,
			timeSinceColor(t), t.Format("2006/01/02 15:04"), timeSince(t),
		),
	)
}

func ToUnix(t time.Time) int64 {
	return t.Unix()
}

func ToDay(t time.Time) int {
	return t.Day()
}
func ToMonth(t time.Time) time.Month {
	return t.Month()
}
func ToYear(t time.Time) int {
	return t.Year()
}

func Datenow(format string) string {
	return time.Now().Add(time.Duration(9) * time.Hour).Format(format)
}
