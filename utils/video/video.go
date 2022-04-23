package video

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func videoCheck(filePath, resFilePath string) {
	hClient := http.Client{
		Timeout: time.Second * 5,
	}
	fmt.Println(hClient)
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fRes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fResStr := string(fRes)
	fList := strings.Split(fResStr, "\n")
	fmt.Println(len(fList))
	fCanList := []string{}
	for _, u := range fList {
		resp, err := hClient.Get(u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if resp.StatusCode == 200 {
			fCanList = append(fCanList, u)
		}
	}

	jsonByte, err := json.Marshal(fCanList)
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile(resFilePath, jsonByte, 0666)
}
