package main

// this file is used to suscribe struct to generate mapper

type dtoMaps struct {
	Source      interface{}
	Destination interface{}
}

type suscriber struct {
	structsMaps map[string]dtoMaps
}

func NewSusciber() suscriber {
	return suscriber{
		structsMaps: make(map[string]dtoMaps),
	}
}

func (s *suscriber) AddStructMapp(source, destination interface{}) *suscriber {
	dict := s.structsMaps
	keyName := hash(source, destination)
	if !s.checkIfKeyExist(keyName) {
		dict[keyName] = dtoMaps{
			Source:      source,
			Destination: destination,
		}
		return s
	}
	panic("Cannot insert twice same object")
}

func (s *suscriber) checkIfKeyExist(key string) bool {
	dict := s.structsMaps
	_, ok := dict[key]
	return ok
}
