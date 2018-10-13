package contents

import (
	"net/url"
	"testing"
	"time"
)

type dummy struct {
	Text string
	Time time.Time
}

func (d *dummy) Type() string {
	return ""
}

func (d *dummy) Title() string {
	return d.Text
}

func (d *dummy) Date() time.Time {
	return d.Time
}

func (d *dummy) Outline() string {
	return ""
}

func (d *dummy) ImageURL() url.URL {
	return url.URL{}
}

func (d *dummy) URL() url.URL {
	return url.URL{}
}

func TestNew(t *testing.T) {
	test := []*dummy{
		{
			Text: "ABC",
			Time: time.Now(),
		},
		{
			Text: "DEF",
			Time: time.Now(),
		},
	}

	c, err := New(test)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(c)
	c.Reverse()
	t.Log(c)
}

func TestReverse(t *testing.T) {
	even := []*dummy{
		{
			Text: "ABC",
			Time: time.Now(),
		},
		{
			Text: "DEF",
			Time: time.Now(),
		},
	}
	var c Contents
	for _, v := range even {
		c = append(c, v)
	}
	c.Reverse()
	t.Log(c[0].Title(), c[1].Title())
}

func TestSortByDate(t *testing.T) {
	even := []*dummy{
		{
			Text: "ABC",
			Time: time.Date(2009, time.November, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			Text: "DEF",
			Time: time.Date(2009, time.October, 0, 0, 0, 0, 0, time.UTC),
		},
	}
	var c Contents
	for _, v := range even {
		c = append(c, v)
	}
	c.SortByDate()
	t.Log(c[0].Title(), c[1].Title())
}
