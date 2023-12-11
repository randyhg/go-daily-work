package log

import "fmt"

const (
	TextBlack = iota + 30
	TextRed
	TextYellow
)

func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}

func SetMsgColor(level int, msg string) string {
	switch level {
	case debugLevel:
		//return SetColor(msg, 0, 0, TextWhite)
		return msg
	case releaseLevel:
		//return SetColor(msg, 0, 0, TextWhite)
		return msg
	case errorLevel:
		return SetColor(msg, 0, 0, TextRed)
	case fatalLevel:
		return SetColor(msg, 0, 0, TextRed)
	default:
		//return SetColor(msg, 0, 0, TextWhite)
		return msg
	}
}
