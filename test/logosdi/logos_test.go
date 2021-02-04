package logosdi_test

import (
	"log"
	"logos/di"
	"testing"
	"time"

	"github.com/takkyuuplayer/go-exercise/test/logosdi"
)

func TestDI(t *testing.T) {
	d := di.NewDomain(func(b *di.B) {
		b.Set("service", &Service{})
		b.Set("num", 1)
	})
	s := d.Get("service").(*Service)
	log.Printf("%#v\n", s)
	log.Println(s.Num())
}

type Service struct {
	// Public *di.D field with tag will automatically be injected after setup
	D *di.D `di:"D"`
}

func (s *Service) Num() int {
	num := s.D.Get("num").(int) // 1
	return num
}

func TestDI3(t *testing.T) {
	d := logosdi.NewDomain(func(b *logosdi.B) {
		b.SetString("str")
		b.SetTime(time.Now())
	})
	d.GetString() // str
	d.GetTime()   // now

}
