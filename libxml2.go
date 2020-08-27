package main

import (
	"github.com/lestrrat-go/libxml2/types"
	"github.com/lestrrat-go/libxml2/dom"
	"reflect"
	"fmt"
	"strconv"
)

func isValidRune(ch rune) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= 'а' && ch <= 'я' {
		return true
	}
	if ch >= 'А' && ch <= 'Я' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}
	if ch == '_' {
		return true
	}
	return false
}

func parseTag(Tag string) ([]string, bool) {
	var dst = []string{}
	var prevI int
	var isValid = false
	tag := []rune(Tag)
	for i:=0; i<=len(tag); i++ {
		if i == len(tag) {
			if prevI == i {
				return nil, isValid
			}
			newStr := string(tag[prevI:i])
			dst = append(dst, newStr)
			break
		}
		if isValidRune(tag[i]) {
			continue
		} else if tag[i] == '>' {
			if prevI == i {
				return nil, isValid
			}
			newStr := string(tag[prevI:i])
			prevI = i + 1
			dst = append(dst, newStr)
		} else if tag[i] == ',' {
			if prevI == i {
				return nil, isValid
			}
			newStr := string(tag[prevI:i])
			dst = append(dst, newStr)
			break
		} else {
			return nil, isValid
		}
	}
	isValid = true
	return dst, isValid
}

func makeNodeDomTreeRecursion(doc *dom.Document, parentNode types.Element, i int, tagSlice []string, payLoad string) (types.Element, error) {
	// вызывающая функция обязана проверить len(tagSlice) > 0
	println("makeNodeDomTreeRecursion start")
	println("payload " + payLoad)
	if i == len(tagSlice) - 1 {
		newNode, err := doc.CreateElement(tagSlice[i])
		if err != nil {
			return parentNode, err
		}
		newNode.SetNodeValue(payLoad)
		err = parentNode.AddChild(newNode)
		return parentNode, err
	}
	newNode, err := doc.CreateElement(tagSlice[i])
	newNode, err = makeNodeDomTreeRecursion(doc, newNode, i + 1, tagSlice, payLoad)
	err = parentNode.AddChild(newNode)
	return parentNode, err
}

func reflectStructRecursion(val reflect.Value, doc *dom.Document, parentNode *types.Element) error {
	var valueField string
	println("reflectStructRecursion start")

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		// Нахожу тэг xml (если он есть) и в зависимости от его содержимого узнаю имя тэга
		tagName, ok := typeField.Tag.Lookup("xml")
		if !ok {
			tagName = typeField.Name	
		}
		// Разбиваю имя тэга на подстроки
		tagSlice, isValid := parseTag(tagName)
		if !isValid {
			tagSlice = []string{typeField.Name}
		}
		if len(tagSlice) < 1 {
			tagSlice = []string{typeField.Name}
		}

		// Тип поля Uint
		if typeField.Type.Kind() == reflect.Uint {
			valueField = strconv.Itoa(int(val.Field(i).Uint()))
		}
		// Тип поля Int
		if typeField.Type.Kind() == reflect.Int {
			valueField = strconv.Itoa(int(val.Field(i).Int()))
		}
		// Тип поля String
		if typeField.Type.Kind() == reflect.String {
			valueField = val.Field(i).String()
		}
		
		// if typeField.Type.Kind() == reflect.Slice {
		// 	UsrSlice := val.Field(i)
		// 	for i:=0; i<UsrSlice.Len(); i++ {
		// 		UsrItem := UsrSlice.Index(i)
		// 		if UsrItem.Type().Kind() != reflect.Struct {
		// 			break
		// 		}
		// 		// Применить рекурсию - обработка полей структуры
		// 		fmt.Printf("%#v\n%T\n%s\n%s\n", UsrItem, UsrItem, UsrItem.Type().Kind(), UsrItem.Field(1))
		// 	}
		// }

		// Если есть больше одного тэга, создаю дом дерево из тегов
		newNode, err := makeNodeDomTreeRecursion(doc, *parentNode, 0, tagSlice, valueField)
		if err != nil {
			return err
		}
		*parentNode = newNode
	}
	return nil
}

func MarshalLibXML2_tester(U *Users) ([]byte, error) {
	val := reflect.ValueOf(U).Elem()
	doc := dom.CreateDocument()
	mainNode, err := doc.CreateElement(`USERS_PACK`)
	err = reflectStructRecursion(val, doc, &mainNode)
	if err != nil {
		return nil, err
	}
	err = doc.SetDocumentElement(mainNode)
	if err != nil {
		return nil, err
	}
	return []byte(doc.String()), nil
}

func MarshalLibXML2(U *Users) ([]byte, error) { // аргументы к функции - reflect.Value, doc, *node, возвращает только ошибку, старшая функция сериализует накопленное в дом дереве
	val := reflect.ValueOf(U).Elem()
	doc := dom.CreateDocument()
	mainNode, err := doc.CreateElement(`USERS_PACK`)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%T have %d fields:\n", U, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		
		typeField := val.Type().Field(i)
		tagName, ok := typeField.Tag.Lookup("xml")
		if !ok || tagName == "-" {
			tagName = typeField.Name
		}
		if typeField.Type.Kind() == reflect.Uint {
			node, err := doc.CreateElement(tagName)
			if err != nil {
				return nil, err
			}
			valueField := strconv.Itoa(int(val.Field(i).Uint()))
			node.SetNodeValue(valueField)
			err = mainNode.AddChild(node)
		} else if typeField.Type.Kind() == reflect.Slice {
			UsrSlice := val.Field(i)
			for i:=0; i<UsrSlice.Len(); i++ {
				UsrItem := UsrSlice.Index(i)
				if UsrItem.Type().Kind() != reflect.Struct {
					break
				}
				// Применить рекурсию - обработка полей структуры
				fmt.Printf("%#v\n%T\n%s\n%s\n", UsrItem, UsrItem, UsrItem.Type().Kind(), UsrItem.Field(1))
			}
		}
	}
	err = doc.SetDocumentElement(mainNode)
	if err != nil {
		return nil, err
	}
	println(doc.String())
	return nil, nil
}








// 	doc := dom.CreateDocument()
// 	for i := 0; i < val.NumField(); i++ {
// 		// valueField := val.Field(i)
// 		typeField := val.Type().Field(i)
// 		if typeField.Tag != "-" {
// 			;
// 		}
// 		// typeField.Name,
// 		// 	typeField.Type.Kind(),
// 		// 	valueField,
// 		// 	typeField.Tag)

// 	}
// 	mainNode, err := doc.CreateElement(`PackageData`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	middle, err := doc.CreateElement("user")
// 	elem, err := doc.CreateElement("key")
// 	if err != nil {
// 		return nil, err
// 	}
// 	elem.SetNodeValue("value")
// 	err = middle.AddChild(elem)
// 	err = mainNode.AddChild(middle)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = doc.SetDocumentElement(mainNode)
// 	return nil, nil
// 	// return doc.String(), nil
// }
