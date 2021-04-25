/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	searchFlag string
	commands   = map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "search",
	Aliases: []string{"sr"},
	Run: func(cmd *cobra.Command, args []string) {
		err := search(url.QueryEscape(args[0]), searchFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchFlag, "platform", "p", "google", "search platform")
}

func search(keyword, platform string) error {
	var uri string
	// 谷歌搜索
	if platform == "google" || platform == "gg" {
		uri = fmt.Sprintf("https://www.google.com/search?q=%s", keyword)
	}
	// 百度搜索
	if platform == "baidu" || platform == "bd" {
		uri = fmt.Sprintf("https://www.baidu.com/s?wd=%s", keyword)
	}
	// 微信搜索
	if platform == "wechat" || platform == "wx" {
		uri = fmt.Sprintf("https://weixin.sogou.com/weixin?type=2&query=%s", keyword)
	}
	// 知乎搜索
	if platform == "zhihu" || platform == "zh" {
		uri = fmt.Sprintf("https://www.zhihu.com/search?type=content&q=%s", keyword)
	}
	// 掘金搜索
	if platform == "juejin" || platform == "jj" {
		uri = fmt.Sprintf("https://juejin.im/search?query=%s&type=all", keyword)
	}
	// csdn
	if platform == "csdn" || platform == "cs" {
		uri = fmt.Sprintf("https://so.csdn.net/so/search?q=%s", keyword)
	}

	if uri == "" {
		return fmt.Errorf("invalid platform")
	}
	return Open(uri)
}

func Open(uri string) error {
	fmt.Println(uri)
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	cmd := exec.Command(run, uri)
	fmt.Println(cmd)
	return cmd.Start()
}
