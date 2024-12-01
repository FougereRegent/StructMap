package main

// this file is used to suscribe struct to generate mapper

type suscriber struct {
}

func NewSusciber() suscriber {
	return suscriber{}
}

func (s *suscriber) AddStructMapp(source, destination interface{}) error {
	return nil
}
