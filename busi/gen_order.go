package busi

import (
	"fmt"
	"strconv"
	"time"
)

var sequence_ int64 = 0

func GenOrderId(uid int64) int64 {
	ts := time.Now().UnixNano() / 1e6
	orderId := ts << 22
	//fmt.Println(strconv.FormatInt(orderId, 2))
	//fmt.Println(strconv.FormatInt(0x3FF, 2))
	//fmt.Println(strconv.FormatInt(0x0000ff00, 2))
	//fmt.Println(strconv.FormatInt(uid&0x3FF, 2))
	//fmt.Println(strconv.FormatInt((uid&0x3FF)<<12, 2))
	orderId |= (uid & 0x3FF) << 12
	fmt.Println(strconv.FormatInt(orderId, 2))
	fmt.Println(orderId)
	orderId |= sequence_ & 0xFFF
	fmt.Println(orderId)
	return orderId
}
