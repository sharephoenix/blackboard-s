package handler

import (
	"blackboards/base/apibase"
	"blackboards/services/buglylog/api/logic"
	buglylogicinfo "blackboards/services/buglylog/api/type"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"io/ioutil"
	"net/http"
)

type CrashHandler struct {
	CrashLogic logic.CrashLogic
}

// 上传 crash 信息 csv
func (handler CrashHandler)PostCrashInfo(w http.ResponseWriter, r *http.Request) {
	rFile,_, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	byts, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(byts))
	reader := csv.NewReader(bufio.NewReader(rFile))
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
	err = handler.CrashLogic.UploadCrashInfo(infos)
	if err != nil {
		httpx.OkJson(w, apibase.ApiResponse{
			-1,
			"rpc 出错",
			infos,
		})
		return
	}
	// logic 调用 rpc 服务，写入数据库
	httpx.OkJson(w, apibase.ApiResponse{
		0,
		nil,
		infos,
	})
}

func (handler CrashHandler)PostCrashDetail(w http.ResponseWriter, r *http.Request) {
	rFile,_, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	byts, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(byts))
	reader := csv.NewReader(bufio.NewReader(rFile))
	infos := []buglylogicinfo.CrashDetail{}
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
		if len(colomn) >= 17 {
			info := buglylogicinfo.CrashDetail{
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
				colomn[11],
				colomn[12],
				colomn[13],
				colomn[14],
				colomn[15],
				colomn[16],
			}
			infos = append(infos, info)
		}
		number += 1
	}
	// logic 调用 rpc 服务，写入数据库
	err = handler.CrashLogic.UploadCrashDetail(infos)
	if err != nil {
		httpx.OkJson(w, apibase.ApiResponse{
			-1,
			"rpc 出错",
			infos,
		})
		return
	}
	httpx.OkJson(w, apibase.ApiResponse{
		0,
		nil,
		infos,
	})
}