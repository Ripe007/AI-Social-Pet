package egret

import "time"

func InitEgretJob() {
	go NewTicker(5*time.Second, Draw)

}

// 定时器
func NewTicker(sat time.Duration, fun func()) {
	ticker := time.NewTicker(sat)
	defer ticker.Stop()

	for range ticker.C {
		fun()
	}
}
