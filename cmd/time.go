package cmd

import (
	"github.com/eighteen20/goTools/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "",
		"需要计算的时间，有效单位为时间戳或已经格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "",
		"持续时间，有效时间单位为‘ns’, 'us'(or 'μs'), 'ms', 's', 'm', 'h'")
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

/*
 子命令： 获取当前时间
*/
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		// 子命令：当前时间
		nowTime := timer.GetNowTime()
		log.Printf("输出结果： %s,  %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

/*
 子命令： 推算时间
 处理三种时间格式： 时间戳、年-月-日、年-月-日 时:分:秒
*/
var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		// 当前时间
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTime = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			//currentTime, err = time.Parse(layout, calculateTime)
			location, _ := time.LoadLocation("Asia/Shanghai")
			currentTime, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(currentTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果： %s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
}
