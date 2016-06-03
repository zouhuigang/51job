//Post数据所在地
package pika

import (
	// "github.com/PuerkitoBio/goquery"
	"net/url"
	// "strings"
	"51job/util"
)

var userName = "sunteng"
var password = "sunteng@2015"
var postdata = url.Values{
	"ctmName": []string{
		"舜飞",
	},
	"userName": []string{
		userName,
	},
	"password": []string{
		password,
	},
}

/*
hidValue
KEYWORDTYPE#0*LASTMODIFYSEL#5*AGE#|*WORKYEAR#0|99*AREA#030200*TOPDEGREE#|*KEYWORD#PHP
hidWhere
00#0#0#0|99|20151124|20160524|99|99|99|99|99|000000|030200|99|99|99|0000|99|99|99|00|0000|99|99|99|0000
|99|99|00|99|99|99|99|99|99|99|99|99|000000|0|0|0000|99#%BeginPage%#%EndPage%#PHP
*/

func SetPostData(searchargs url.Values, keyword string, address string, addressid string, kind string) url.Values {
	searchargs.Set("MainMenuNew1$CurMenuID", "MainMenuNew1_imgResume|sub4")
	searchargs.Set("__EVENTTARGET", "ctrlSerach$btnConditionQuery")
	searchargs.Set("pagerBottom$txtGO", "1")
	searchargs.Set("ctrlSerach$AREA$Text", address)
	searchargs.Set("ctrlSerach$AREA$Value", addressid)
	searchargs.Set("cbxColumns$0", "AGE")
	searchargs.Set("cbxColumns$1", "WORKYEAR")
	searchargs.Set("cbxColumns$14", "LASTUPDATE")
	searchargs.Set("cbxColumns$2", "SEX")
	searchargs.Set("cbxColumns$4", "AREA")
	searchargs.Set("cbxColumns$8", "TOPMAJOR")
	searchargs.Set("cbxColumns$9", "TOPDEGREE")
	searchargs.Set("ctrlSerach$KEYWORD", keyword)
	searchargs.Set("ctrlSerach$KEYWORDTYPE", kind)
	searchargs.Set("ctrlSerach$LASTMODIFYSEL", "5")
	searchargs.Set("ctrlSerach$WorkYearFrom", "0")
	searchargs.Set("ctrlSerach$WorkYearTo", "99")
	hidv := "KEYWORDTYPE#" + kind + "*LASTMODIFYSEL#5*AGE#|*WORKYEAR#0|99*AREA#" + addressid + "*TOPDEGREE#|*KEYWORD#" + keyword
	searchargs.Set("hidValue", hidv)
	hidW := "00#0#0#0|99|" + util.SixMonthAgo() + "|" + util.TodayDate() + "|99|99|99|99|99|000000|"
	hidW = hidW + addressid + "|99|99|99|0000|99|99|99|00|0000|99|99|99|0000|99|99|00|99|99|99|99|99|99|99|99|99|000000|0|0|0000|99"
	hidW = hidW + "#%BeginPage%#%EndPage%#" + keyword
	searchargs.Set("hidWhere", hidW)
	return searchargs
}
