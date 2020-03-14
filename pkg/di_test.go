package pkg

import (
	"testing"
)

type DPC string

func Test_container_Provide(t *testing.T) {
	container := NewContainer()
	container.Provide(
		func() DPC { return DPC("Hello") },
		func(dsn DPC) string { return "Parviz" },
	)

}

func Test_container_NoDependencies(t *testing.T) {
	container := NewContainer()
	container.Provide(
		func() DPC {return DPC("academy")},)
}

func TestNewMsg_Dep_Consumer(t *testing.T) {
	container := NewContainer()
	container.Provide(
		func() DPC {return DPC("Alif")},
		func(dsn DPC) string{return "academy"})
}

