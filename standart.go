package main

import (
	"encoding/xml"
)

func UnmarshalStandart(in []byte) (*Users, error) {
	U := &Users{}
	err := xml.Unmarshal([]byte(in), U)
	return U, err
}

func MarshalStandart(Users *Users) ([]byte, error) {
	ret, err := xml.Marshal(Users)
	return ret, err
}
