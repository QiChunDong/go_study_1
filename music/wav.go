package music

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("播放 WAV 音乐：", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.progress += 10
	}

	fmt.Println("WAV 音乐 播放完毕：", source)
}