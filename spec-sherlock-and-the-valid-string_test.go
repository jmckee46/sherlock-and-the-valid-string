package main

import "testing"

func TestSherlockAndTheValidString1(t *testing.T) {
	s := "aabbcd"

	valid := isValid(s)

	if valid != "NO" {
		t.Errorf("got %s instead of NO", valid)
	}
}

func TestSherlockAndTheValidString2(t *testing.T) {
	s := "aabbccddeefghi"

	valid := isValid(s)

	if valid != "NO" {
		t.Errorf("got %s instead of NO", valid)
	}
}

func TestSherlockAndTheValidString3(t *testing.T) {
	s := "abc"

	valid := isValid(s)

	if valid != "YES" {
		t.Errorf("got %s instead of YES", valid)
	}
}

func TestSherlockAndTheValidString4(t *testing.T) {
	s := "abcc"

	valid := isValid(s)

	if valid != "YES" {
		t.Errorf("got %s instead of YES", valid)
	}
}

func TestSherlockAndTheValidString5(t *testing.T) {
	s := "abccc"

	valid := isValid(s)

	if valid != "NO" {
		t.Errorf("got %s instead of NO", valid)
	}
}

func TestSherlockAndTheValidString6(t *testing.T) {
	s := "a"

	valid := isValid(s)

	if valid != "YES" {
		t.Errorf("got %s instead of YES", valid)
	}
}

func TestSherlockAndTheValidStringTestCase7(t *testing.T) {
	s := "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd"

	valid := isValid(s)

	if valid != "YES" {
		t.Errorf("got %s instead of YES", valid)
	}
}
