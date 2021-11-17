package main

import (
	config2 "blackboards/demo/sqldemosql/config"
	"blackboards/demo/sqldemosql/model"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	io "io/ioutil"
	"strings"
)
var (
	studentClassRelationFieldNames          = builderx.FieldNames(&model.UserModelTable{})
	studentClassRelationRows                = strings.Join(studentClassRelationFieldNames, ",")
)


func main() {

	config := config2.DemoMysql{
		"root:qwer1234@(127.0.0.1:3306)/pingyin",
		config2.Table{
			"job",
			"pingyin",
		},
	}

	btys, err := io.ReadFile("./pingyin.json")
	if err != nil {
		fmt.Println("读取文件失败：", err.Error())
		return
	}
	var pingyinjson map[string]string
	err = json.Unmarshal(btys, &pingyinjson)

	if err != nil {
		fmt.Println("解析json失败：", err.Error())
		return
	}

	dytMysql := sqlx.NewMysql(config.DataSource)
	for key, value := range pingyinjson {
		pingyin := key
		for _, char := range strings.Split(value, "") {
			hanzi := char
			sql := `insert into ` + config.TableName.Pingyin + `(pingyin, hanzi)` + ` values ('` + pingyin + `','` + hanzi + `')`
			_, err = dytMysql.Exec(sql)
			if err != nil {
				fmt.Println(sql)
				fmt.Println("error: ", pingyin, hanzi, err.Error())
			}
		}
	}

}

func test0() {
	config := config2.DemoMysql{
		"root:qwer1234@(127.0.0.1:3306)/myjob",
		config2.Table{
			"job",
			"pingyin",
		},
	}
	dytMysql := sqlx.NewMysql(config.DataSource)
	var users model.UserModelTable
	sqlStr := `select ` + studentClassRelationRows +  ` from ` + config.TableName.Myjob
	err := dytMysql.QueryRow(&users, sqlStr)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println(string(data))
}


