package main

import (
	"testing"
)

func TestNewSyntaxANodes(t *testing.T) {
	var o1 = NewSyntaxANodes("@RCV_INS_ID_CD = 01000000")

	if o1.SyntaxNodes.MatchString("01000000") != true {
		t.Error("MatchString fail")
	}

	if o1.SyntaxNodes.MatchString("010000") != false {
		t.Error("MatchString fail")
	}

	var o2 = NewSyntaxANodes("@RCV_INS_ID_CD in {01000000, 01000001}")

	if o2.SyntaxNodes.MatchString("01000000") != true || o2.SyntaxNodes.MatchString("01000001") != true {
		t.Error("MatchString fail")
	}

	if o2.SyntaxNodes.MatchString("aaaa") != false {
		t.Error("MatchString fail")
	}

	var o3 = NewSyntaxANodes("@RCV_INS_ID_CD ex {01000000, 01000001}")

	if o3.SyntaxNodes.MatchString("0100000") != true {
		t.Error("MatchString fail")
	}

	var o4 = NewSyntaxANodes("@requestUrl head /sso/login")

	if o4.SyntaxNodes.MatchString("/sso/login") != false || o4.SyntaxNodes.MatchString("/sso/login/") != true {
		t.Error("MatchString fail")
	}

	var o5 = NewSyntaxANodes("@requestUrl tail  /sso/login")

	if o5.SyntaxNodes.MatchString("/sso/login") != false || o5.SyntaxNodes.MatchString("s/sso/login") != true {
		t.Error("MatchString fail")
	}

	var o6 = NewSyntaxANodes("@requestUrl vague   /sso/login")

	if o6.SyntaxNodes.MatchString("/sso/login") != false || o6.SyntaxNodes.MatchString("s/sso/login/sfd") != true {
		t.Error("MatchString fail")
	}

	var o7 = NewSyntaxANodes("@ORDER_TYPE sub[0,2] in {02}  ")

	if o7.SyntaxNodes.MatchString("02333") != true || o7.SyntaxNodes.MatchString("04") != false {
		t.Error("MatchString fail")
	}

	var res = NewSyntaxANodes("@RCV_INS_ID_CD = 01000000|FWD_INS_ID_CD = 010100000")
	if res.SyntaxNodes.MatchString("01000000") != true {
		t.Error("MatchString fail")
	}

	var res1 = NewSyntaxANodes("@FWD = 00000000 | @RCV in {12345678,87654321} | (@ISS = 1222222 | @ISS = 888888883 | (@ISS = 222222 & @ISS in {222222  }  )  ) ")

	if res1.SyntaxNodes.MatchString("222222") != true {
		t.Error("MatchString fail")
	}

	var res2 = NewSyntaxANodes("@FWD = 00000000 | @RCV in {12345678,87654321} | (@ISS = 1222222 | @ISS = 888888883 | (@ISS = 222222 & @ISS head 222  ) )")

	if res2.SyntaxNodes.MatchString("222222") != true {
		t.Error("MatchString fail")
	}

	var res3 = NewSyntaxANodes("@FWD tail 1 | @RCV in {123456781,87654321} | (@ISS = 1222222 | @ISS = 888888883 | (@ISS = 222222 & @ISS head 222  ) )")

	if res3.SyntaxNodes.MatchString("87654321") != true {
		t.Error("MatchString fail")
	}
}
