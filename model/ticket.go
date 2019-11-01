package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var (
	PreTicketKey ="ticket:%s:%s:%s"		// ticket:openid:yyyyMMdd:orderid
	PreTicketGetDay = "ticket:%s:%s"
	YyyyMMddFormat = "20060102"
	yyyyMMddHHmm = "2006年1月2日15:04"
)
type Ticket struct {
	Openid			string
	OrderId			string
	// 发车日期
	DepartDate		string
	DepartTime 		string
	// 发车车站
	DeaprtStation	string

	// 车次及座席
	TicketInfo		string
	// 检票口
	TicketCheck		string

}

func (t *Ticket) Save() error {
	datetimeStr := strconv.Itoa(time.Now().Year())+"年"+t.DepartDate+t.DepartTime
	ticketTime,err := time.Parse(yyyyMMddHHmm,datetimeStr)
	if err != nil {
		return err
	}
	key := fmt.Sprintf(PreTicketKey,t.Openid,ticketTime.Format(YyyyMMddFormat),t.OrderId)
	fmt.Println("key=",key)
	if key != "" {
		obj, err := json.Marshal(t)
		if err != nil {
			return err
		}
		fmt.Println(string(obj))
	}
	return nil
}

