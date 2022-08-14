package main

import "strconv"

type PhoneInterface interface {
	PrintPhoneInfo() string
}
type CellPhone struct {
	phoneModel, phoneColor, phonePrice string
	phoneProcessor                     string
	whereReleased                      string
	warantyYear                        int
}

//todo ->  write the first letter of functions inside of the interface with CAPITAL LETTER
func (c CellPhone) PrintPhoneInfo() string {
	resultPhoneInfo := "\nPhone model : " + c.phoneModel + "\nPhone Color : " + c.phoneColor +
		"\nPhone price : " + c.phonePrice + "\nPhone Processor : " + c.phoneProcessor +
		"\nWhere manufactured : " + c.whereReleased + "\nWaranty year : " + strconv.Itoa(c.warantyYear)
	return resultPhoneInfo

}
