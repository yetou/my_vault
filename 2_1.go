package main

import (

	"fmt"
)

type Author struct{
	Name string
	VIP bool
	Icon string
	Signature string
	Focus int
}
type Sl struct{
	Good int
	Tbi int
	Sc int
	zf int
}
type SP struct{
	own Author
	dz Sl
}
func main(){
	hkd := SP{
		own:Author{
			Name: "我跟白敬亭私奔了",
			VIP: false,
			Icon: "nil",
			Signature: "娱乐圈墙头草 哪边好看往哪倒听说偷视频的都被我抓去听白拉普唱rap了",
			Focus: 192,
		},
	dz:Sl{
		Good:124000,
		Tbi:20000,
		Sc:51000,
		zf:4407,
	},
	}
	fmt.Println(hkd)
}
