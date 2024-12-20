package main

import "fmt"

type MusicPlayer interface {
	play(fileType string, fileName string)
}

type ExistPlayer struct {
}

func (*ExistPlayer) playMp3(fileName string) {
	fmt.Println("mp3:", fileName)
}

func (*ExistPlayer) playWma(fileName string) {
	fmt.Println("play:", fileName)
}

// 适配器
type Adaper struct {
	// 持有一个旧接口
	existPlayer ExistPlayer
}

func (a *Adaper) play(fileType string, fileName string) {
	fmt.Println("cc")
	switch fileType {
	case "mp3":
		a.existPlayer.playMp3(fileName)
	case "play":
		a.existPlayer.playWma(fileName)
	default:
		fmt.Println("暂时不支持此类型文件播放")
	}
}

func main() {
	ada := &Adaper{}
	ada.existPlayer.playMp3("wqqwqw")
	ada.play("mp3", "qwqw")
}
