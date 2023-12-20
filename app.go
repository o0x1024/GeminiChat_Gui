package main

import (
	"changeme/global"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"strings"

	"gopkg.in/yaml.v3"
)

// App struct
type App struct {
	ctx context.Context
}

type GeminiResp struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason  string `json:"finishReason"`
		Index         int    `json:"index"`
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"candidates"`
	PromptFeedback struct {
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"promptFeedback"`
}

//	type Message struct {
//		Id     int
//		Sender string
//		Text   string
//	}

type RetrueMsg struct {
	Code    int
	Message string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	//读取配置文件
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	gminiPath := usr.HomeDir + "\\.gemini"
	if !PathExists(gminiPath) {
		err = os.Mkdir(gminiPath, 0666)
		if err != nil {
			panic(err)
		}
	}
	yamlPath := gminiPath + "\\gemini.yaml"

	//有就读取，没有就初始化一个文件
	if !PathExists(yamlPath) {
		global.Conf.APIKEY = ""
		global.Conf.Proxy = ""
		data, err := yaml.Marshal(global.Conf)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(yamlPath, data, 0666)
		if err != nil {
			panic(err)
		}

	} else {
		content, err := os.ReadFile(yamlPath)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(content, &global.Conf)
		if err != nil {
			panic(err)
		}
	}

}

func (b *App) shutdown(ctx context.Context) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	gminiPath := usr.HomeDir + "\\.gemini"
	yamlPath := gminiPath + "\\gemini.yaml"
	data, err := yaml.Marshal(global.Conf)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(yamlPath, data, 0666)
	if err != nil {
		panic(err)
	}
}

func (a *App) SyncConf(apikey string, proxy string) {
	global.Conf.APIKEY = apikey
	global.Conf.Proxy = proxy
}

func (a *App) GetConf() global.Config {
	return global.Conf
}

func (a *App) SendMessage(question string, apiKey string, proxyStr string) RetrueMsg {

	fmt.Println(question, "  ", apiKey, "  ", proxyStr)
	rm := RetrueMsg{}

	gurl := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key=" + apiKey
	data := fmt.Sprintf(`{"contents": [{"parts": [{"text": "%s"}]}]}`, question)

	req, _ := http.NewRequest("POST", gurl, strings.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		fmt.Println(err)
		rm.Code = 201
		return rm
	}
	fmt.Println("1111")

	client := &http.Client{}
	if proxyStr != "" {
		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
	}
	fmt.Println("2222")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		rm.Code = 201
		return rm
	}
	fmt.Println("4444")

	if resp.StatusCode == 200 {
		fmt.Println("666")

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			rm.Code = 201
			return rm
		}
		fmt.Println("xxxx")

		geminiresp := GeminiResp{}
		err = json.Unmarshal(body, &geminiresp)
		if err != nil {
			fmt.Println(err)
			rm.Code = 201
			return rm
		}
		fmt.Println(geminiresp)

		if len(geminiresp.Candidates) > 0 {
			if len(geminiresp.Candidates[0].Content.Parts) > 0 {
				rm.Code = 200
				rm.Message = geminiresp.Candidates[0].Content.Parts[0].Text
			}

		}

		return rm
	} else {
		fmt.Println("777")

		rm.Code = 201
		return rm
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
