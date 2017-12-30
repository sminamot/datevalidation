package datevalidation

import (
	"testing"
	"time"
)

func setNow(t time.Time) {
	timeNow = func() time.Time { return t }
}

func TestValidateDate(t *testing.T) {
	type test struct {
		description string
		format      string
		targetDate  string
		allowFuture bool
	}

	testDate, _ := time.Parse("20060102", "20170101")
	setNow(testDate)

	normalTests := []test{
		{"正常_01", "20060102", "20170101", true},
		{"正常_02", "20060102", "20170101", false},
		{"正常_03", "2006-Jan-02(Mon)", "2017-Jan-01(Sun)", true},
		{"正常_04", "2006-Jan-02(Mon)", "2017-Jan-01(Sun)", false},
		{"未来日を許可する場合の未来日_01", "20060102", "20180101", true},
		{"未来日を許可する場合の未来日_02", "2006-Jan-02(Mon)", "2018-Jan-01(Mon)", true},
	}

	for i, test := range normalTests {
		if ret := validateDate(test.format, test.targetDate, test.allowFuture); ret != nil {
			t.Errorf("validateDate test failed(Expetcted that error not occured) #%d(%s)", i, test.description)
		}
	}

	errorTests := []test{
		{"存在しない日付_01", "20060102", "20170132", true},
		{"存在しない日付_02", "20060102", "20170132", false},
		{"フォーマットが不正_01", "20060102", "2017-Jan-01(Sun)", true},
		{"フォーマットが不正_02", "20060102", "2017-Jan-01(Sun)", false},
		{"曜日が異なる", "2006-Jan-02(Mon)", "2017-Jan-01(Mon)", true},
		{"曜日が異なる", "2006-Jan-02(Mon)", "2017-Jan-01(Mon)", false},
		{"未来日を禁止する場合の未来日_01", "20060102", "20180101", false},
		{"未来日を禁止する場合の未来日_02", "2006-Jan-02(Mon)", "2018-Jan-01(Mon)", false},
	}

	for i, test := range errorTests {
		if ret := validateDate(test.format, test.targetDate, test.allowFuture); ret == nil {
			t.Errorf("validateDate test failed(Expetcted that error occured) #%d(%s)", i, test.description)
		}
	}
}
