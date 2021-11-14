package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"math/rand"
	"towelong/mogu/service"
	"towelong/mogu/utils"
	"github.com/joho/godotenv"
)

func main() {
	// load enviroment secret key
	godotenv.Load()

	a := rand.Intn(180)
	fmt.Println(a)
	time.Sleep(time.Second * time.Duration(a))

	moguding_wjk := service.NewMoGuService()
	token_wjk, userId_wjk := moguding_wjk.MoGuLogin(os.Getenv("ACCOUNT_wjk"), os.Getenv("PASSWORD_wjk"))
	planID_wjk := moguding_wjk.GetPlanID(token_wjk, userId_wjk)
                fmt.Println("planID是"+planID_wjk)
	fmt.Println("token是"+token_wjk)
	isSuccess_wjk, types_wjk := moguding_wjk.SignIn(token_wjk, planID_wjk, userId_wjk)
	title_wjk, message_wjk := utils.EnumToMsg(types_wjk)
	
	if !isSuccess_wjk {
		service.SendMessage_wjk(title_wjk, message_wjk)
		log.Fatal(title_wjk)
	}
	fmt.Println("订阅消息发送前的title"+title_wjk)
	fmt.Println("订阅消息发送前的message"+message_wjk)
	if service.SendMessage_wjk(title_wjk, message_wjk) {
		fmt.Println("wjk打卡成功")
	}

	// 写周记
	isTrue_wjk, weekType_wjk := moguding_wjk.WeeklyDiary(token_wjk, planID_wjk)
	fmt.Println(isTrue_wjk)
	fmt.Println(weekType_wjk)
	headline_wjk, msg_wjk := utils.EnumToMsg(weekType_wjk)
	if !isTrue_wjk && weekType_wjk == utils.NOWEEK {
		fmt.Println(msg_wjk)
	} else {
		if service.SendMessage_wjk(headline_wjk, msg_wjk) {
			fmt.Println("wjk周记已提交")
		}
	}

	b := rand.Intn(180)
	fmt.Println(b)
	time.Sleep(time.Second * time.Duration(b))
	moguding_xh := service.NewMoGuService()
	token_xh, userId_xh := moguding_xh.MoGuLogin(os.Getenv("ACCOUNT_xh"), os.Getenv("PASSWORD_xh"))
	planID_xh := moguding_xh.GetPlanID(token_xh, userId_xh)
                fmt.Println("planID是"+planID_xh)
	fmt.Println("token是"+token_xh)
	isSuccess_xh, types_xh := moguding_xh.SignIn_xh(token_xh, planID_xh, userId_xh)
	title_xh, message_xh := utils.EnumToMsg(types_xh)
	
	if !isSuccess_xh {
		service.SendMessage_xh(title_xh, message_xh)
		log.Fatal(title_xh)
	}
	fmt.Println("订阅消息发送前的title"+title_xh)
	fmt.Println("订阅消息发送前的message"+message_xh)
	if service.SendMessage_xh(title_xh, message_xh) {
		fmt.Println("xh打卡成功")
	}

	// 写周记
	isTrue_xh, weekType_xh := moguding_xh.WeeklyDiary(token_xh, planID_xh)
	fmt.Println(isTrue_xh)
	fmt.Println(weekType_xh)
	headline_xh, msg_xh := utils.EnumToMsg(weekType_xh)
	if !isTrue_xh && weekType_xh == utils.NOWEEK {
		fmt.Println(msg_xh)
	} else {
		if service.SendMessage_xh(headline_xh, msg_xh) {
			fmt.Println("xh周记已提交")
		}
	}

	c := rand.Intn(180)
	fmt.Println(c)
	time.Sleep(time.Second * time.Duration(c))
	moguding_drt := service.NewMoGuService()
	token_drt, userId_drt := moguding_drt.MoGuLogin(os.Getenv("ACCOUNT_drt"), os.Getenv("PASSWORD_drt"))
	planID_drt := moguding_drt.GetPlanID(token_drt, userId_drt)
                fmt.Println("planID是"+planID_drt)
	fmt.Println("token是"+token_drt)
	isSuccess_drt, types_drt := moguding_drt.SignIn(token_drt, planID_drt, userId_drt)
	title_drt, message_drt := utils.EnumToMsg(types_drt)
	
	if !isSuccess_drt {
		service.SendMessage_drt(title_drt, message_drt)
		log.Fatal(title_drt)
	}
	fmt.Println("订阅消息发送前的title"+title_drt)
	fmt.Println("订阅消息发送前的message"+message_drt)
	if service.SendMessage_drt(title_drt, message_drt) {
		fmt.Println("drt打卡成功")
	}

	// 写周记
	isTrue_drt, weekType_drt := moguding_drt.WeeklyDiary(token_drt, planID_drt)
	fmt.Println(isTrue_drt)
	fmt.Println(weekType_drt)
	headline_drt, msg_drt := utils.EnumToMsg(weekType_drt)
	if !isTrue_drt && weekType_drt == utils.NOWEEK {
		fmt.Println(msg_drt)
	} else {
		if service.SendMessage_drt(headline_drt, msg_drt) {
			fmt.Println("drt周记已提交")
		}
	}

	d := rand.Intn(180)
	fmt.Println(d)
	time.Sleep(time.Second * time.Duration(d))
	moguding_zzx := service.NewMoGuService()
	token_zzx, userId_zzx := moguding_zzx.MoGuLogin(os.Getenv("ACCOUNT_zzx"), os.Getenv("PASSWORD_zzx"))
	planID_zzx := moguding_zzx.GetPlanID(token_zzx, userId_zzx)
                fmt.Println("planID是"+planID_zzx)
	fmt.Println("token是"+token_zzx)
	isSuccess_zzx, types_zzx := moguding_zzx.SignIn(token_zzx, planID_zzx, userId_zzx)
	title_zzx, message_zzx := utils.EnumToMsg(types_zzx)
	
	if !isSuccess_zzx {
		service.SendMessage_zzx(title_zzx, message_zzx)
		log.Fatal(title_zzx)
	}
	fmt.Println("订阅消息发送前的title"+title_zzx)
	fmt.Println("订阅消息发送前的message"+message_zzx)
	if service.SendMessage_zzx(title_zzx, message_zzx) {
		fmt.Println("zzx打卡成功")
	}

	// 写周记
	isTrue_zzx, weekType_zzx := moguding_zzx.WeeklyDiary(token_zzx, planID_zzx)
	fmt.Println(isTrue_zzx)
	fmt.Println(weekType_zzx)
	headline_zzx, msg_zzx := utils.EnumToMsg(weekType_zzx)
	if !isTrue_zzx && weekType_zzx == utils.NOWEEK {
		fmt.Println(msg_zzx)
	} else {
		if service.SendMessage_zzx(headline_zzx, msg_zzx) {
			fmt.Println("zzx周记已提交")
		}
	}
		s := rand.Intn(180)
	fmt.Println(s)
	time.Sleep(time.Second * time.Duration(d))
	moguding_jzx := service.NewMoGuService()
	token_jzx, userId_jzx := moguding_jzx.MoGuLogin(os.Getenv("ACCOUNT_jzx"), os.Getenv("PASSWORD_jzx"))
	planID_jzx := moguding_jzx.GetPlanID(token_jzx, userId_jzx)
                fmt.Println("planID是"+planID_jzx)
	fmt.Println("token是"+token_jzx)
	isSuccess_jzx, types_jzx := moguding_jzx.SignIn(token_jzx, planID_jzx, userId_jzx)
	title_jzx, message_jzx := utils.EnumToMsg(types_jzx)
	
	if !isSuccess_jzx {
		service.SendMessage_jzx(title_jzx, message_jzx)
		log.Fatal(title_jzx)
	}
	fmt.Println("订阅消息发送前的title"+title_jzx)
	fmt.Println("订阅消息发送前的message"+message_jzx)
	if service.SendMessage_jzx(title_jzx, message_jzx) {
		fmt.Println("jzx打卡成功")
	}

	// 写周记
	isTrue_jzx, weekType_jzx := moguding_jzx.WeeklyDiary(token_jzx, planID_jzx)
	fmt.Println(isTrue_jzx)
	fmt.Println(weekType_jzx)
	headline_jzx, msg_jzx := utils.EnumToMsg(weekType_jzx)
	if !isTrue_jzx && weekType_jzx == utils.NOWEEK {
		fmt.Println(msg_jzx)
	} else {
		if service.SendMessage_jzx(headline_jzx, msg_jzx) {
			fmt.Println("jzx周记已提交")
		}
	}
	
}
