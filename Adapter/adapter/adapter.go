package adapter

import "DesignModel/Adapter/oldAPI"

type PrintBanner struct {
	banner *oldAPI.Banner
}

func NewPrintBanner(str string) *PrintBanner {
	return &PrintBanner{banner: oldAPI.NewBanner(str)}
}

func (pb *PrintBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
