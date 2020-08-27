package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
)

// Эта функция парсит входящий xml в структуру Users без особой валидации (!!!)
// То есть я вообще не слежу за закрывающимися тегами
func UnmarshalStream(xmlData []byte) (*Users, error) {
	var U Users
	var uintData uint
	var userData User
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	for {
		token, err := decoder.Token()
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			break
		}
		if token == nil {
			return nil, errors.New("token is equal nil")
		}
		switch tokItem := token.(type) {
		case xml.StartElement:
			if tokItem.Name.Local == "id_organization" {
				if err := decoder.DecodeElement(&uintData, &tokItem); err != nil {
					return nil, err
				}
				U.IDOrganization = uintData
			}
			if tokItem.Name.Local == "user" {
				if err := decoder.DecodeElement(&userData, &tokItem); err != nil {
					return nil, err
				}
				U.Users = append(U.Users, userData)
			}
		}
	}
	return &U, nil
}

// Эта функция распечатает построчно все поля, которые ей попадутся
func StreamPrinter(xmlData []byte) error {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	for {
		token, err := decoder.Token()
		if err != nil && err != io.EOF {
			return err
		} else if err == io.EOF {
			break
		}
		if token == nil {
			return errors.New("token is equal nil")
		}
		fmt.Printf("%#v\n", token)
	}
	return nil
}

// bytes.Buffer - объект со множеством методов записи и чтения из него. Имеет метод Write(p []byte) (n int, err error)
// xml.NewEncoder принимает интерфейс Writer с единственным методом Write(p []byte) (n int, err error)
func MarshalStream(U *Users) ([]byte, error) {
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)
	err := encoder.Encode(U) // нужно просто скармливать в этот метод новую инфу, она по идее добавляется в связанный список
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil // а вот тут инфа за один проход конвертируется в нужный тип данных
}
