package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	msgColor = "ed5c4c"
	// msgColor = "ffffff"
)

func main() {
	// color()
	// test()
	do()
	var color string
	for range time.NewTicker(time.Second).C {
		color = robotgo.GetPixelColor(877, 852)
		log.Println(color)
		if color == msgColor {
			if err := do(); err != nil {
				log.Println("do error:", err.Error())
			}
			time.Sleep(time.Minute)
		}
	}

}

func do() error {
	cmd := exec.Command("sh", "./run.sh")
	b, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Println(string(b))
	log.Println("do end ...")
	return nil
}

func color() {
	color := robotgo.GetPixelColor(877, 852) //获取坐标100, 200的颜色
	log.Println("color----", color, "-----------------")
	// ed5c4c
}
func mouse() {
	for range time.NewTicker(time.Second).C {
		x, y := robotgo.GetMousePos() //获取鼠标坐标位置
		log.Println("pos:", x, y)
		// 877 852
	}
}

func test() {
	//键盘控制
	robotgo.TypeStr("Hello World") //输入Hello World
	robotgo.KeyTap("enter")        //按下enter键
	robotgo.KeyTap("a", "control")
	robotgo.KeyTap("h", "command") //隐藏窗口

	robotgo.KeyTap("i", "alt", "command")
	//按下"i", "alt", "command"组合键
	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)
	//按下"i", "alt", "command"组合键

	robotgo.KeyTap("w", "command") //关闭窗口
	robotgo.KeyTap("m", "command") //最小化窗口
	robotgo.KeyTap("f1", "control")
	robotgo.KeyTap("a", "control")
	robotgo.KeyToggle("a", "down") //切换a键
	robotgo.KeyToggle("a", "down", "alt")
	robotgo.KeyToggle("a", "down", "alt", "command")
	robotgo.KeyToggle("enter", "down")
	robotgo.TypeStr("en")

	//鼠标控制
	robotgo.Move(100, 200)                   //移动鼠标到100, 200位置
	robotgo.Click()                          //鼠标左键单击
	robotgo.Click("right", false)            //右键单击
	robotgo.Click("left", true)              //左键双击
	robotgo.Scroll(10, 0)                    //向上滚动鼠标
	robotgo.Toggle("down", "right")          //鼠标右键切换
	robotgo.MoveSmooth(100, 200)             //平滑移动鼠标到100, 200
	robotgo.MoveSmooth(100, 200, 1.0, 100.0) //设置平滑移动速度
	x, y := robotgo.GetMousePos()            //获取鼠标坐标位置
	log.Println("pos:", x, y)
	if x == 456 && y == 586 {
		log.Println("mouse...", "586")
	}

	robotgo.Toggle("up")
	robotgo.Move(x, y)
	robotgo.Move(100, 200)

	for i := 0; i < 1080; i += 1000 {
		log.Println(i)
		robotgo.Move(800, i)
	}
	//屏幕控制
	// robotgo.CaptureScreen()
	bit_map := robotgo.CaptureScreen()
	// log.Println("CaptureScreen...", bit_map)
	// robotgo.FreeBitmap(bit_map)
	// gbit_map := robotgo.BCaptureScreen() //获取屏幕位图
	// log.Println("Capture_Screen...", gbit_map.Width)

	sx, sy := robotgo.GetScreenSize() //获取屏幕width和height
	log.Println("...", sx, sy)

	color := robotgo.GetPixelColor(100, 200) //获取坐标100, 200的颜色
	log.Println("color----", color, "-----------------")

	color2 := robotgo.GetPixelColor(10, 20) //获取坐标10, 20的颜色
	log.Println("color---", color2)

	// Bitmap
	abit_map := robotgo.CaptureScreen() //获取全屏位图
	log.Println("a...", abit_map)

	// bit_map := robotgo.CaptureScreen(100, 200, 30, 40)
	//获取100, 200, 30, 40的位图
	// log.Println("CaptureScreen...", bit_map)
	// log.Println("...", bit_map.Width, bit_map.BytesPerPixel)

	fx, fy := robotgo.FindBitmap(bit_map) //查找位图
	log.Println("FindBitmap------", fx, fy)

	bit_pos := robotgo.GetPortion(bit_map, 10, 10, 11, 10) //截取位图
	log.Println(bit_pos)

	bit_str := robotgo.TostringBitmap(bit_map) //Tostring位图
	log.Println("bit_str...", bit_str)

	// sbit_map := robotgo.BitmapFromstring(bit_str, 2)
	// log.Println("...", sbit_map)

	robotgo.SaveBitmap(bit_map, "test.png") //保存位图为图片
	robotgo.SaveBitmap(bit_map, "test31.tif", 1)
	robotgo.Convert("test.png", "test.tif") //转换位图图片格式

	open_bit := robotgo.OpenBitmap("test.tif") //打开图片位图
	log.Println("open...", open_bit)

	//全局监听事件
	log.Println("---请按v键---")
	eve := robotgo.AddEvent("v")

	if eve {
		log.Println("---你按下v键---", "v")
	}

	log.Println("---请按k键---")
	keve := robotgo.AddEvent("k")
	if keve {
		log.Println("---你按下k键---", "k")
	}

	log.Println("---请按鼠标左键---")
	mleft := robotgo.AddEvent("mleft")
	if mleft {
		log.Println("---你按下左键---", "mleft")
	}

	// mright := robotgo.AddEvent("mright")
	// if mright == 0 {
	//  log.Println("---你按下右键---", "mright")
	// }

	// robotgo.LStop()

	//窗口
	abool := robotgo.ShowAlert("hello", "robotgo") //弹出窗口
	if abool {
		log.Println("ok@@@", "确认")
	}
	robotgo.ShowAlert("hello", "robotgo", "确认", "取消")
	// robotgo.GetPID()
	mdata := robotgo.GetActive() //获取当前窗口
	hwnd := robotgo.GetHandle()  //获取当前窗口hwnd
	log.Println("hwnd---", hwnd)
	title := robotgo.GetTitle() //获取当前窗口标题
	log.Println("title-----", title)
	robotgo.CloseWindow()    //关闭当前窗口
	robotgo.SetActive(mdata) //SetActive窗口
}
