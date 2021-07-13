package main

import (
	config2 "blackboards/demo/sqldemosql/config"
	"blackboards/demo/sqldemosql/model"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"strings"
)
var (
	studentClassRelationFieldNames          = builderx.FieldNames(&model.UserModelTable{})
	studentClassRelationRows                = strings.Join(studentClassRelationFieldNames, ",")
)


func main() {

	config := config2.DemoMysql{
		"root:qwer1234@(127.0.0.1:3306)/myjob",
			config2.Table{
			"job",
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


