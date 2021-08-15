package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	expJson "github.com/go-json-experiment/json"
	"github.com/jakha/lc/commands/bitree/levelorder/mocks"
	jsoniter "github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
	fastjson2 "github.com/valyala/fastjson"
)

func main() {
	easyJson()
}

func standartPack() {
	data, err := ioutil.ReadFile("generated.json")
	if err != nil {
		panic(err)
	}
	var objs = mocks.Objs{}
	start := time.Now()
	err = json.Unmarshal(data, &objs)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
}

func goExperementedJson() {
	data, err := ioutil.ReadFile("generated.json")
	if err != nil {
		panic(err)
	}

	var objs = mocks.Objs{}
	start := time.Now()
	err = expJson.Unmarshal(data, &objs)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
}

func ffencode() {
	data, err := ioutil.ReadFile("generated.json")
	if err != nil {
		panic(err)
	}
	var objs = mocks.Objs{}
	start := time.Now()
	err = ffjson.Unmarshal(data, &objs)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
}

func easyJson() {
	data, err := ioutil.ReadFile("generated.json")
	if err != nil {
		panic(err)
	}
	objsStr := mocks.Objs{}
	start := time.Now()
	err = objsStr.UnmarshalJSON(data)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
}

func fastjson() {
	data, err := ioutil.ReadFile("generated.json")
	if err != nil {
		panic(err)
	}
	start := time.Now()
	v, err := fastjson2.ParseBytes(data)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
	_ = v
}

func jsonIterator() {
	data, err := ioutil.ReadFile("generated.json")
	if err != nil {
		panic(err)
	}
	var objc = &mocks.Objs{}
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	start := time.Now()
	err = jsons.Unmarshal(data, objc)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
}
