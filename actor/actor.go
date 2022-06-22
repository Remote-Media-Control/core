package actor

import (
	"log"
	"runtime"
	"time"

	pb "github.com/Remote-Media-Control/core/proto"

	"github.com/micmonay/keybd_event"
)

// keys from win, but maybe they will work somewhere else?
const (
	VK_MEDIA_NEXT_TRACK = 0xB0 + 0xFFF
	VK_MEDIA_PREV_TRACK = 0xB1 + 0xFFF
	VK_MEDIA_STOP       = 0xB2 + 0xFFF
	VK_MEDIA_PLAY_PAUSE = 0xB3 + 0xFFF
)

func PressKey(key pb.Key) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		log.Println("kb init die")
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	switch key {
	case pb.Key_NEXT:
		kb.SetKeys(VK_MEDIA_NEXT_TRACK)
	case pb.Key_PREV:
		kb.SetKeys(VK_MEDIA_PREV_TRACK)
	case pb.Key_PAUSE:
		kb.SetKeys(VK_MEDIA_PLAY_PAUSE)
	case pb.Key_PLAY:
		kb.SetKeys(VK_MEDIA_PLAY_PAUSE)
	case pb.Key_NONE:
	}

	kb.Launching()
}
