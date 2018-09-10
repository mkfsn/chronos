package jobs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mkfsn/chronos/clients"

	"github.com/PuerkitoBio/goquery"
)

type aburiyaJob struct {
	date   string
	url    string
	state  map[string]string
	target string
}

func newAburiyaJob() *aburiyaJob {
	date := "20181028"
	format := "https://www.hotpepper.jp/CSP/imrCourseCalendar/ajaxCourseFullDateStock?SP=J001128046&RDT=%s&CN=&CB=&RTM=&RPN=&_=1536414798956"
	target := "18:00"

	return &aburiyaJob{
		date:   date,
		url:    fmt.Sprintf(format, date),
		state:  make(map[string]string),
		target: target,
	}
}

func (a *aburiyaJob) Name() string {
	return "Aburiya" + a.date
}

func (a *aburiyaJob) Interval() time.Duration {
	return time.Minute * 30
}

func (a *aburiyaJob) Run() error {
	doc, err := goquery.NewDocument(a.url)
	if err != nil {
		return err
	}

	// rows = div.find("table.courseDetailReserveSeatList tbody tr")
	rows := doc.Find("table.courseDetailReserveSeatList tbody tr")

	// [{pq(r).find("td.cellVisitTime").text(): pq(r).find("td.taC").text()} for r in rows if pq(r).text().find("18:00") != -1]
	result := make(map[string]string)
	rows.Each(func(i int, row *goquery.Selection) {
		if strings.Contains(row.Text(), a.target) {
			time := row.Find("td.cellVisitTime").Text()
			seat := row.Find("td.taC").Text()
			result[time] = seat
		}
	})

	if a.updateState(result) {
		// update to somewhere
		b, err := json.Marshal(a.state)
		if err != nil {
			return err
		}
		clients.Slack.Send(string(b))
	}

	return nil
}

func (a *aburiyaJob) updateState(state map[string]string) bool {
	l1, l2 := len(a.state), len(state)
	if l1 == l2 && (l1 == 0 || l2 == 0) {
		return false
	} else if !a.compareState(state) {
		return false
	}

	a.state = state
	return true
}

func (a *aburiyaJob) compareState(state map[string]string) bool {
	for key, value := range a.state {
		if v, ok := state[key]; !ok || v != value {
			return true
		}
	}

	for key, value := range state {
		if v, ok := a.state[key]; !ok || v != value {
			return true
		}
	}

	return false
}
