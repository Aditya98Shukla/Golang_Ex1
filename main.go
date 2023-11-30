package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	// "strconv"
	"strings"
  scheme "ps/lscpu/scheme"
)

func main(){
	cmd := exec.Cmd{}
	//cmd.Path = "D:\\Git\\usr\\bin\\ls.exe"
	//cmd.Path = "D:\\Downloads\\go1.21.0\\bin\\go"
	cmd.Path = "/usr/bin/lscpu"
  cmd.Args = []string{}

	res, err := cmd.Output()
	if err != nil {
		panic(err)
	}
  //res_str = This is the output in string format
	res_str := string(res)

	res_str_line := strings.Split(res_str, "\n")
  output := make(map[string]interface{})
	for _, line := range res_str_line {
		if line==""{
			continue
		}
    //fmt.Println(i,line)
    outs := strings.Split(line,":")
		if len(outs) == 2{
      //fmt.Println(outs[0],outs[1])
			output[outs[0]] = strings.Trim(outs[1]," ")
		}else{
			continue
		}
	}

	jsonData, err := json.Marshal(output)
	if err != nil{
		panic(err)
	}
	fmt.Println("--------------------------The JSON String Response is shown below:----------------------\n",string(jsonData))

  var result = &scheme.Response{}

	err = json.Unmarshal(jsonData, result)
	if err != nil {
		panic(err)
	}
  fmt.Println("--------------------------The JSON Response to Golang Struct is shown below:----------------------")
	fmt.Printf("%+v",result)
}
