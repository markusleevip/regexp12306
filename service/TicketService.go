package service

import (
	"fmt"
	"github.com/markusleevip/regexp12306/model"
	"regexp"
)

const (
	ticket12306ReCheck = `【铁路12306】.*`
	ticket12306ReStr = `【铁路12306】.*订单(.+)[，].*您已购([0-9]{1,2}[月][0-9]{1,2}[日])(.*)[,](.*[站])(.*)[开].+检票口：([0-9a-zA-Z]+)`
	ticket12306ReStr1 = `【铁路12306】.*订单(.+)[，].*您已购([0-9]{1,2}[月][0-9]{1,2}[日])(.*)[,](.*[站])(.*)[开]`

)
var (
	ticket12306ReMsgCheck = regexp.MustCompile(ticket12306ReCheck)
	ticket12306ReMsg = regexp.MustCompile(ticket12306ReStr)
	ticket12306ReMsg1 = regexp.MustCompile(ticket12306ReStr1)

	yyyyMMddFormat = "20060102"
	yyyyMMddHHmm = "2006年1月2日15:04"
)

func TicketCheck(content string ) bool {
	return ticket12306ReMsgCheck.FindString(content) != ""
}

func TicketSave(openid, content string) bool {
	checkTicket := ticket12306ReMsgCheck.FindString(content)
	result := false
	fmt.Println("checkTicket=",checkTicket)
	if checkTicket != "" {
		resultContent := ticket12306ReMsg.FindAllStringSubmatch(content, -1)
		if len(resultContent) == 0 {
			resultContent = ticket12306ReMsg1.FindAllStringSubmatch(content, -1)
		}
		fmt.Println(len(resultContent))
		fmt.Println(resultContent)
		for i := 0; i < len(resultContent); i++ {
			subContent := resultContent[i]
			fmt.Println(len(subContent))
			for j := 0; j < len(subContent); j++ {
				fmt.Println(j, ",", subContent[j])
			}
			ticket := model.Ticket{}
			ticket.Openid=openid
			if len(subContent) >= 6 {
				fmt.Println(subContent[1])
				ticket.OrderId 			= subContent[1]
				ticket.DepartDate 		= subContent[2]
				ticket.TicketInfo 		= subContent[3]
				ticket.DeaprtStation	= subContent[4]
				ticket.DepartTime		= subContent[5]
			}
			if len(subContent) == 7{
				if ticket.TicketInfo != "" {
					ticket.TicketCheck = subContent[6]
				}
			}
			ticket.Save()
		}

		result = true
	}else{
		fmt.Println("No Ticket...")

	}
	return result
}