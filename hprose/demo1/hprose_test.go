package demo1

import (
	"github.com/hprose/hprose-golang/v3/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/plugins/log"
	"github.com/stretchr/testify/assert"
	"net"
	"sync"
	"testing"
	"time"
)

type StudentService struct {
	students []Student
	lock     sync.RWMutex
}

type Student struct {
	ID       int
	Name     string
	Birthday time.Time
	Grade    int
	Class    int
}

type Person struct {
	Name     string
	Birthday time.Time
}

func TestHproseServer(t *testing.T) {
	service := rpc.NewService()
	service.Codec = rpc.NewServiceCodec(rpc.WithDebug(true))
	service.AddInstanceMethods(&StudentService{})
	server, err := net.Listen("tcp", "127.0.0.1:8412")
	assert.NoError(t, err)
	err = service.Bind(server)
	assert.NoError(t, err)
	time.Sleep(60 * time.Second)
	server.Close()
}

type Proxy struct {
	Add        func(s ...Student) error
	Get        func(id int) (*Student, error)
	GetStudent func(id int) Student `name:"get"`
	GetPerson  func(id int) Person  `name:"get"`
	Delete     func(id int) error
}

func TestHproseClient(t *testing.T) {
	client := rpc.NewClient("tcp://127.0.0.1/")
	client.Use(log.Plugin)
	var proxy Proxy
	client.UseService(&proxy)
	s1 := Student{
		ID:       1,
		Name:     "张三",
		Birthday: time.Date(2008, 11, 23, 0, 0, 0, 0, time.Local),
		Grade:    6,
		Class:    1,
	}
	s2 := Student{
		ID:       2,
		Name:     "李四",
		Birthday: time.Date(2013, 12, 11, 0, 0, 0, 0, time.Local),
		Grade:    1,
		Class:    2,
	}
	err := proxy.Add(s1, s2)
	assert.NoError(t, err)
	var student *Student
	student, err = proxy.Get(1)
	assert.Equal(t, s1, *student)
	assert.NoError(t, err)
	student, err = proxy.Get(2)
	assert.Equal(t, s2, *student)
	assert.NoError(t, err)
	err = proxy.Delete(2)
	assert.NoError(t, err)
	student, err = proxy.Get(2)
	assert.Nil(t, student)
	assert.NoError(t, err)
	s := proxy.GetStudent(1)
	assert.Equal(t, s1, s)
	p := proxy.GetPerson(1)
	assert.Equal(t, "张三", p.Name)
	assert.Equal(t, time.Date(2008, 11, 23, 0, 0, 0, 0, time.Local), p.Birthday)
}
