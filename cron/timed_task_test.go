package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var job = mockJob{}

type mockJob struct{}

func (job mockJob) Run() {
	mockFunc()
}

func mockFunc() {
	time.Sleep(time.Second)
	fmt.Println("1s...")
}

func mockFunc2() {
	time.Sleep(time.Second)
	fmt.Println("2s...")
}

func TestNewTimerTask(t *testing.T) {
	tm := NewTimerTask()
	_tm := tm.(*timer)

	{
		_, err := tm.AddTaskByFunc("func", "@every 1s", mockFunc)
		assert.Nil(t, err)
		_, ok := _tm.taskList["func"]
		if !ok {
			t.Error("no find func")
		}
	}

	{
		_, err := tm.AddTaskByJob("job", "@every 1s", job)
		assert.Nil(t, err)
		_, ok := _tm.taskList["job"]
		if !ok {
			t.Error("no find job")
		}
	}

	{
		_, ok := tm.FindCron("func")
		if !ok {
			t.Error("no find func")
		}
		_, ok = tm.FindCron("job")
		if !ok {
			t.Error("no find job")
		}
		_, ok = tm.FindCron("none")
		if ok {
			t.Error("find none")
		}
	}
	{
		tm.Clear("func")
		_, ok := tm.FindCron("func")
		if ok {
			t.Error("find func")
		}
	}
}

func TestCronTaskInit(t *testing.T) {
	task := NewTimerTask()
	option := []cron.Option{cron.WithSeconds()}
	entryID, err := task.AddTaskByFunc("test", "*/4 * * * * *", mockFunc, option...)
	assert.Nil(t, err)
	assert.Equal(t, entryID, 1)
	time.Sleep(time.Second * 15)
	_, _ = task.AddTaskByFunc("test", "*/4 * * * * *", mockFunc2, option...)
	time.Sleep(time.Second * 15)
	task.Clear("test")
	time.Sleep(time.Second * 50)
}
