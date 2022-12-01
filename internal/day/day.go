package day

import (
	"fmt"
	"os"
)

type InputKind string

const (
	InputKindExample InputKind = "example"
	InputKindReal    InputKind = "real"
)

type Day struct {
	Runner Runnable
	Name   string
	Kind   InputKind
}

type Runnable interface {
	Init(path string) error
	Step1() (int, error)
	Step2() (int, error)
}

func (d *Day) InputPath() string {
	path := fmt.Sprintf("%s/input", os.Getenv("PROJECT_PATH"))
	if d.Kind == InputKindExample {
		return fmt.Sprintf("%s/%s/example.txt", path, d.Name)
	}

	return fmt.Sprintf("%s/%s/real.txt", path, d.Name)
}

func (d *Day) Init() error {
	return d.Runner.Init(d.InputPath())
}

func (d *Day) Step1() (int, error) {
	return d.Runner.Step1()
}

func (d *Day) Step2() (int, error) {
	return d.Runner.Step2()
}
