package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time1043/gvb-server/models/res"
)

const file = "models/res/err_code.json"

type ErrMap map[res.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}

	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)                    // map[1001:系统错误]
	fmt.Println(errMap[res.SettingsError]) // 系统错误
}
