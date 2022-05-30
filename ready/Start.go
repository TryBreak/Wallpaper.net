package ready

import (
	"fmt"
	"time"

	"Wallpaper.net/global/config"
	"Wallpaper.net/utils/bingData"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	mCycle.New(mCycle.Opt{
		Func:      GetBingImg,
		SleepTime: time.Hour * 4,
	}).Start()

	fmt.Println(config.BingImg)
}

func GetBingImg() {
	bingData.Start()
}
