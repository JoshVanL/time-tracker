package timesheet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type Timesheet struct {
	filepath string

	stamps []*Stamp
	Log    *logrus.Entry
}

type JSONStamp struct {
	Title   string `json:"title"`
	Action  string `json:"action"`
	Comment string `json:"comment"`

	Duration  int       `json:"duration"`
	Timestamp time.Time `json:"timestamp"`
}

type Stamp struct {
	title   string
	action  string
	comment string

	duration  int
	timestamp time.Time
}

func New() (t *Timesheet, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	dir = fmt.Sprintf("%s%s", dir, "/test")

	stamp := &Stamp{
		title:     "this is a title",
		action:    "this is a action",
		comment:   "this is a comment",
		duration:  5,
		timestamp: time.Now(),
	}

	stamps := []*Stamp{stamp}

	return &Timesheet{
		filepath: dir,
		stamps:   stamps,
	}, nil

}

func (t *Timesheet) WriteConfig() error {
	_, err := os.Create(t.FilePath())
	if err != nil {
		return err
	}

	var stamps []*JSONStamp
	for _, s := range t.Stamps() {
		jsonStamp := &JSONStamp{
			Title:     s.Title(),
			Action:    s.Action(),
			Comment:   s.Comment(),
			Duration:  s.Duration(),
			Timestamp: s.TimeStamp(),
		}
		stamps = append(stamps, jsonStamp)
	}

	data, err := json.Marshal(stamps)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(t.FilePath(), data, 0666)
}

func (t *Timesheet) ReadConfig() error {
	if err := t.checkFilePath(); err != nil {
		return nil
	}

	data, err := ioutil.ReadFile(t.FilePath())
	if err != nil {
		return fmt.Errorf("failed to read data from file: %v", err)
	}

	var stamps []*Stamp
	if err := json.Unmarshal(data, &stamps); err != nil {
		return fmt.Errorf("failed to unmarshal data from file: %v", err)
	}

	fmt.Printf("%+v", stamps[0])

	return nil
}

func (t *Timesheet) checkFilePath() error {
	if _, err := os.Stat(t.FilePath()); os.IsNotExist(err) {
		if _, err := os.Create(t.FilePath()); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (t *Timesheet) FilePath() string {
	return t.filepath
}

func (t *Timesheet) Stamps() []*Stamp {
	return t.stamps
}

func (s *Stamp) Title() string {
	return s.title
}

func (s *Stamp) Action() string {
	return s.action
}

func (s *Stamp) Comment() string {
	return s.comment
}

func (s *Stamp) Duration() int {
	return s.duration
}

func (s *Stamp) TimeStamp() time.Time {
	return s.timestamp
}
