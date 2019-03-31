package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func printToday(r Response) {
	fmt.Println("日期:", "今天")
	fmt.Println("湿度:", r.Data.Shidu)
	fmt.Println("PM2.5:", r.Data.Pm25)
	fmt.Println("PM10:", r.Data.Pm10)
	fmt.Println("空气质量:", r.Data.Quality)
	fmt.Println("温馨提示:", r.Data.Ganmao)
	fmt.Println("====================================")
}

func printYesterday(r Response) {
	fmt.Println("日期:", r.Data.Yesterday.Ymd)
	fmt.Println("温度:", r.Data.Yesterday.Low, r.Data.Yesterday.High)
	fmt.Println("风量:", r.Data.Yesterday.Fx, r.Data.Yesterday.Fl)
	fmt.Println("天气:", r.Data.Yesterday.Type)
	fmt.Println("温馨提示:", r.Data.Yesterday.Notice)
	fmt.Println("====================================")
}

func printFurther(r Response) {
	for _, item := range r.Data.Forecast {
		fmt.Println("日期:", item.Ymd)
		fmt.Println("温度:", item.Low, item.High)
		fmt.Println("风量:", item.Fx, item.Fl)
		fmt.Println("天气:", item.Type)
		fmt.Println("温馨提示:", item.Notice)
		fmt.Println("====================================")
	}
}

func Print(day string, r Response) {
	fmt.Println("城市:", r.CityInfo.City)
	fmt.Println("====================================")
	if day == "今天" {
		printToday(r)
	} else if day == "昨天" {
		printYesterday(r)
	} else if day == "预测" {
		printFurther(r)
	} else {
		printYesterday(r)
		printFurther(r)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "weather-cli"
	app.Usage = "天气预报小程序"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "city, c",
			Usage: "城市中文名",
		},
		cli.StringFlag{
			Name:  "day, d",
			Value: "全部",
			Usage: "可选: 今天, 昨天, 预测",
		},
	}

	app.Action = func(c *cli.Context) error {
		city := c.String("city")
		day := c.String("day")

		if len(city) == 0 {
			fmt.Println("必须填写查询城市, 用法 [-c|--city] <城市名>")
			return nil
		}
		var body, err = Request(city)
		if err != nil {
			fmt.Printf("err was %v", err)
			return nil
		}

		var r Response
		err = json.Unmarshal([]byte(body), &r)
		if err != nil {
			fmt.Printf("\nError message: %v", err)
			return nil
		}
		if r.Status != 200 {
			fmt.Printf("获取天气API出现错误, %s", r.Message)
			return nil
		}
		Print(day, r)
		return nil
	}
	app.Run(os.Args)
}
