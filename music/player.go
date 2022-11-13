package music

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("不支持的类型：", mtype)
		return
	}
	p.Play(source)
}