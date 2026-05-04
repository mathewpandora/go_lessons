package main

import (
	"errors"
	"fmt"
)

// композиция - вложенные структуры

type Engine struct {
	Started    bool
	HorsePower int
}

// если хоти что то изменить в структуре принимавем указатель
func (e *Engine) Start() error {
	if e.Started {
		return errors.New("Already Start")
	}
	e.Started = true
	return nil
}

type CarStruct struct {
	Engine Engine
	Model  string
}

func (c *CarStruct) Start() error {
	if err := c.Engine.Start(); err != nil {
		return fmt.Errorf("Engine start: %w", err)
	}
	//start other service
	return nil
}

func (c CarStruct) Drive() error {
	if !c.Engine.Started {
		return errors.New("enine not started")
	}
	return nil
}

func main() {

	car := CarStruct{
		Engine: Engine{
			Started:    false,
			HorsePower: 23,
		},
		Model: "BMW",
	}

	if err := car.Start(); err != nil {
		fmt.Errorf("Error %w", err)
	}

	fmt.Println(car, "Car Started")
}
