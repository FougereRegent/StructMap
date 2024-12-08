package main

// this file is used to suscribe struct to generate mapper

type dtoMaps struct {
	Source      interface{}
	destination interface{}
}

type suscriber struct {
	structsMaps map[string]dtoMaps
}

func NewSusciber() suscriber {
	return suscriber{}
}

func (s *suscriber) AddStructMapp(source, destination interface{}) error {
	return nil
}
