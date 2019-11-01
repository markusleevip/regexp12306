package service

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Ticket(t *testing.T) {

	content := "【铁路12306】订单EC19121xxx，李xx您已购11月1日Ab123次2车4F号,北京站08:47开，检票口：12B。"
	//content = "【铁路12306】候补订单已兑现成功，订单EC82731xxx，李xx您已购11月6日cc223次4车2C号,北京站19:56开，检票口：23。"
	TicketSave("123456",content)

}

func Test_DateTime(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))          // -->2017-04-17 17:21:33
	fmt.Println(time.Now().Format("2006年01月02日 15:04"))          // -->2017-04-17 17:21:33
	fmt.Println(time.Now().Format(yyyyMMddFormat))
	datetimeStr := strconv.Itoa(time.Now().Year())+"年10月12日17:46"
	ticketTime,err := time.Parse(yyyyMMddHHmm,datetimeStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(ticketTime)
	fmt.Println(ticketTime.Format(yyyyMMddFormat))


}