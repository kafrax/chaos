package chaos

import "github.com/kafrax/logx"

func Recover(msg interface{}){
	if r:=recover();r!=nil{
		logx.Warn(msg)
	}
}