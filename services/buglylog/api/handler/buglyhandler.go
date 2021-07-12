package handler

import (
	buglylogicinfo "blackboards/services/buglylog/api/type"
	"encoding/csv"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"io/ioutil"
	"net/http"
)

type Requestn struct {
	File []byte `form:"file"`
}

// 上传 crash 信息 csv
func PostCrashInfo(w http.ResponseWriter, r *http.Request) {
	var req Requestn
	err := httpx.ParseForm(r, &req)
	if err != nil {
		return
	}
	byts, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(byts))
	reader := csv.NewReader(r.Body)
	infos := []buglylogicinfo.CrashInfo{}
	number := 0

	for {
		colomn, err := reader.Read()
		if err != nil {
			fmt.Println("error", err.Error())
			break
		}
		if number == 0 {
			number += 1
			continue
		}
		for _, line := range colomn {
			fmt.Println(line)
		}
		if len(colomn) >= 11 {
			info := buglylogicinfo.CrashInfo{
				colomn[0],
				colomn[1],
				colomn[2],
				colomn[3],
				colomn[4],
				colomn[5],
				colomn[6],
				colomn[7],
				colomn[8],
				colomn[9],
				colomn[10],
			}
			infos = append(infos, info)
		}
		number += 1
	}
	httpx.OkJson(w, infos)
}