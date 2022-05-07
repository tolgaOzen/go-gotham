package helpers

type MonthlyEnum string

const (
	January   MonthlyEnum = "January"
	February  MonthlyEnum = "February"
	March     MonthlyEnum = "March"
	April     MonthlyEnum = "April"
	May       MonthlyEnum = "May"
	June      MonthlyEnum = "June"
	July      MonthlyEnum = "July"
	August    MonthlyEnum = "August"
	September MonthlyEnum = "September"
	October   MonthlyEnum = "October"
	November  MonthlyEnum = "November"
	December  MonthlyEnum = "December"
)

type MonthlyInformation struct {
	Month    MonthlyEnum
	FullName string
	MonthID  int
}

func (m MonthlyEnum) GetMonthFullName() string {
	return Months[m].FullName
}

func (m MonthlyEnum) GetMonthId() int {
	return Months[m].MonthID
}

func GetMonthNameWithId(id int) string {
	for _, m := range Months {
		if m.MonthID == id {
			return m.FullName
		}
	}
	return ""
}

var Months = map[MonthlyEnum]MonthlyInformation{
	January:   {Month: January, FullName: "January", MonthID: 1},
	February:  {Month: February, FullName: "February", MonthID: 2},
	March:     {Month: March, FullName: "March", MonthID: 3},
	April:     {Month: April, FullName: "April", MonthID: 4},
	May:       {Month: May, FullName: "May", MonthID: 5},
	June:      {Month: June, FullName: "June", MonthID: 6},
	July:      {Month: July, FullName: "July", MonthID: 7},
	August:    {Month: August, FullName: "August", MonthID: 8},
	September: {Month: September, FullName: "September", MonthID: 9},
	October:   {Month: October, FullName: "October", MonthID: 10},
	November:  {Month: November, FullName: "November", MonthID: 11},
	December:  {Month: December, FullName: "December", MonthID: 12},
}
