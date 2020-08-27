package main

import (
	"encoding/xml"
	"errors"
	"fmt"
)

const (
	YELLOW   = "\033[33m"
	RED_BG   = "\033[41;30m"
	GREEN_BG = "\033[42;30m"
	NO_COLOR = "\033[m"
)

type Users struct {
	XMLName        xml.Name `xml:"USERS_PACK"`
	IDOrganization uint     `xml:"id_organization"`
	Users          []User   `xml:"users>user,omitempty"`
}

type User struct {
	Id            uint   `xml:"ID"`
	Login         string `xml:"Login"`
	Pass          string `xml:"-"`
	EncryptedPass string `xml:"-"`
	UserData1     string `xml:"user_data_1,omitempty"`
	UserData2     string `xml:"user_data_2,omitempty"`
	UserData3     string `xml:"user_data_3,omitempty"`
	UserData4     string `xml:"user_data_4,omitempty"`
}

func InitUsersStruct(nbr uint) *Users {
	var Users Users
	var User User

	Users.IDOrganization = 68
	for i := uint(0); i < nbr; i++ {
		User.Id = i
		User.Login = "skinny"
		User.Pass = "password"
		User.EncryptedPass = "asdkjczaSADHSAK=="
		User.UserData1 = "User Data 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
		User.UserData2 = "User Data 2 Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
		User.UserData3 = "User Data 3 Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."

		Users.Users = append(Users.Users, User)
	}
	return &Users
}

func UserComparator(Users1 *Users, Users2 *Users) error {
	if Users1 == nil {
		return errors.New("Users1 is equal nil !!!")
	}
	if Users2 == nil {
		return errors.New("Users2 is equal nil !!!")
	}
	if Users1.IDOrganization != Users2.IDOrganization {
		return errors.New("IDOrganization missmatch")
	}
	if len(Users1.Users) != len(Users2.Users) {
		return errors.New(fmt.Sprintf("Amount of User structs missmatch - %d != %d", len(Users1.Users), len(Users2.Users)))
	}
	for i := 0; i < len(Users1.Users); i++ {
		if Users1.Users[i].Id != Users2.Users[i].Id ||
			Users1.Users[i].Login != Users2.Users[i].Login ||
			Users1.Users[i].UserData1 != Users2.Users[i].UserData1 ||
			Users1.Users[i].UserData2 != Users2.Users[i].UserData2 ||
			Users1.Users[i].UserData3 != Users2.Users[i].UserData3 ||
			Users1.Users[i].UserData4 != Users2.Users[i].UserData4 {
			return errors.New(fmt.Sprintf("User #%d missmatch\n%#v\n%#v", i, Users1.Users[i], Users2.Users[i]))
		}
	}
	return nil
}

func stringComparator(str1 string, str2 string) error {
	if str1 != str2 {
		return errors.New(fmt.Sprintf("strings missmatch: %s != %s", str1, str2))
	}
	return nil
}

func byteComparator(expected []byte, got []byte) error {
	if expected == nil {
		return errors.New("expected slice byte is nil")
	}
	if got == nil {
		return errors.New("got slice byte is nil")
	}
	if len(expected) != len(got) {
		return errors.New(fmt.Sprintf("slice byte size missmatch: %d != %d %s %s", len(expected), len(got), string(expected), string(got)))
	}
	if string(expected) != string(got) {
		return errors.New(fmt.Sprintf("slice byte missmatch: %s != %s", string(expected), string(got)))
	}
	return nil
}

func main() {

	// dst, isValid := parseTag("a>b>c,omitempty")
	// fmt.Println(dst, isValid)


	// U := InitUsersStruct(2)

	// xmlData, err := MarshalLibXML2_tester(U)
	// if err != nil {
	// 	println("Error marshal libxml2: " + err.Error())
	// 	return
	// }
	// println(string(xmlData))
}
