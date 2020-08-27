package main

import (
	"testing"
)

func TestStream(t *testing.T) {

	t.Run("Marshal 0 user", func(t_ *testing.T) {
		expected := []byte("<USERS_PACK><id_organization>68</id_organization><users></users></USERS_PACK>")
		U := InitUsersStruct(0)
		got, err := MarshalStream(U)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = byteComparator(expected, got)
		if err != nil {
			t_.Errorf(RED_BG + "Marshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Marshal success" + NO_COLOR)
	})

	t.Run("Marshal 1 user", func(t_ *testing.T) {
		expected := []byte("<USERS_PACK><id_organization>68</id_organization><users>" +
			"<user><ID>0</ID><Login>skinny</Login><user_data_1>User Data 1 Lorem " +
			"ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor " +
			"incididunt ut labore et dolore magna aliqua.</user_data_1>" +
			"<user_data_2>User Data 2 Ut enim ad minim veniam, quis nostrud exercitation " +
			"ullamco laboris nisi ut aliquip ex ea commodo consequat.</user_data_2>" +
			"<user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate " +
			"velit esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user>" +
			"</users></USERS_PACK>")
		U := InitUsersStruct(1)
		got, err := MarshalStream(U)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = byteComparator(expected, got)
		if err != nil {
			t_.Errorf(RED_BG + "Marshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Marshal success" + NO_COLOR)
	})

	t.Run("Marshal 3 users", func(t_ *testing.T) {
		expected := []byte("<USERS_PACK><id_organization>68</id_organization><users>" +
			"<user><ID>0</ID><Login>skinny</Login><user_data_1>User Data 1 Lorem " +
			"ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor " +
			"incididunt ut labore et dolore magna aliqua.</user_data_1>" +
			"<user_data_2>User Data 2 Ut enim ad minim veniam, quis nostrud exercitation " +
			"ullamco laboris nisi ut aliquip ex ea commodo consequat.</user_data_2>" +
			"<user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate " +
			"velit esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user>" +
			"<user><ID>1</ID><Login>skinny</Login><user_data_1>User Data 1 Lorem " +
			"ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor " +
			"incididunt ut labore et dolore magna aliqua.</user_data_1>" +
			"<user_data_2>User Data 2 Ut enim ad minim veniam, quis nostrud exercitation " +
			"ullamco laboris nisi ut aliquip ex ea commodo consequat.</user_data_2>" +
			"<user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate " +
			"velit esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user>" +
			"<user><ID>2</ID><Login>skinny</Login><user_data_1>User Data 1 Lorem " +
			"ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor " +
			"incididunt ut labore et dolore magna aliqua.</user_data_1>" +
			"<user_data_2>User Data 2 Ut enim ad minim veniam, quis nostrud exercitation " +
			"ullamco laboris nisi ut aliquip ex ea commodo consequat.</user_data_2>" +
			"<user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate " +
			"velit esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user>" +
			"</users></USERS_PACK>")
		U := InitUsersStruct(3)
		got, err := MarshalStream(U)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = byteComparator(expected, got)
		if err != nil {
			t_.Errorf(RED_BG + "Marshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Marshal success" + NO_COLOR)
	})

	t.Run("Unmarshal 0 user", func(t_ *testing.T) {
		xmlSrc := "<USERS_PACK><id_organization>68</id_organization></USERS_PACK>"
		Uexpected := InitUsersStruct(0)
		Ugot, err := UnmarshalStream([]byte(xmlSrc))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = UserComparator(Uexpected, Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Unmarshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Unmarshal success" + NO_COLOR)
	})

	t.Run("Unmarshal 1 user", func(t_ *testing.T) {
		xmlSrc := "<USERS_PACK><id_organization>68</id_organization><user><ID>0</ID><Login>skinny</Login>" +
			"<user_data_1>User Data 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.</user_data_1><user_data_2>User Data 2 Ut enim " +
			"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." +
			"</user_data_2><user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate velit " +
			"esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user></USERS_PACK>"
		Uexpected := InitUsersStruct(1)
		Ugot, err := UnmarshalStream([]byte(xmlSrc))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = UserComparator(Uexpected, Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Unmarshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Unmarshal success" + NO_COLOR)
	})

	t.Run("Unmarshal 3 users", func(t_ *testing.T) {
		xmlSrc := "<USERS_PACK><id_organization>68</id_organization><user><ID>0</ID><Login>skinny</Login>" +
			"<user_data_1>User Data 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.</user_data_1><user_data_2>User Data 2 Ut enim " +
			"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." +
			"</user_data_2><user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate velit " +
			"esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user><user><ID>1</ID><Login>skinny</Login>" +
			"<user_data_1>User Data 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.</user_data_1><user_data_2>User Data 2 Ut enim " +
			"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." +
			"</user_data_2><user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate velit " +
			"esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user><user><ID>2</ID><Login>skinny</Login>" +
			"<user_data_1>User Data 1 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.</user_data_1><user_data_2>User Data 2 Ut enim " +
			"ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." +
			"</user_data_2><user_data_3>User Data 3 Duis aute irure dolor in reprehenderit in voluptate velit " +
			"esse cillum dolore eu fugiat nulla pariatur.</user_data_3></user></USERS_PACK>"
		Uexpected := InitUsersStruct(3)
		Ugot, err := UnmarshalStream([]byte(xmlSrc))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = UserComparator(Uexpected, Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Unmarshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Unmarshal success" + NO_COLOR)
	})

	t.Run("Marshal unmarshal 0 users", func(t_ *testing.T) {
		Uexpected := InitUsersStruct(0)
		xmlMid, err := MarshalStream(Uexpected)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		Ugot, err := UnmarshalStream([]byte(xmlMid))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		xmlMid, err = MarshalStream(Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		Ugot, err = UnmarshalStream([]byte(xmlMid))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = UserComparator(Uexpected, Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Unmarshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Marshal Unmarshal success" + NO_COLOR)
	})

	t.Run("Marshal unmarshal 5 users", func(t_ *testing.T) {
		Uexpected := InitUsersStruct(5)
		xmlMid, err := MarshalStream(Uexpected)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		Ugot, err := UnmarshalStream([]byte(xmlMid))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		xmlMid, err = MarshalStream(Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Error marshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		Ugot, err = UnmarshalStream([]byte(xmlMid))
		if err != nil {
			t_.Errorf(RED_BG + "Error unmarshal Stream: " + err.Error() + NO_COLOR)
			return
		}
		err = UserComparator(Uexpected, Ugot)
		if err != nil {
			t_.Errorf(RED_BG + "Unmarshal Stream missmatch: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Marshal Unmarshal success" + NO_COLOR)
	})
}
