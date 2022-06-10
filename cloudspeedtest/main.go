package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"

	"CloudflareSpeedTest/task"
	"CloudflareSpeedTest/utils"
)

var (
	version, versionNew string
)

func init() {
	var printVersion bool
	var help = `
CloudflareSpeedTest ` + version + ` Cloudflare CDN IP (IPv4+IPv6)！

Комманд ууд ：
    -n 200
	Хурд хэмжих хэлхээний тоо; хурдыг хэмжих тусам сул үзүүлэлттэй төхөөрөмж (чиглүүлэгч гэх мэт) илүү хурдан байх ёсгүй; (анхдагчаар 200-аас 1000 хүртэл)
    -t 4
	Хурд хэмжилтийн саатлын хугацаа; нэг IP саатлын хурд хэмжилтийн хугацаа, энэ нь 1 бол алдагдсан IP, TCP протоколыг шүүнэ; (өгөгдмөл 4)
    -tp 443
	Хурдны туршилтын портыг зааж өгнө үү; хурдны шалгалт/татаж авах хурдны туршилтыг хойшлуулах үед ашигладаг порт; (өгөгдмөл 443)
    -dn 10
	Татаж авах хурдны туршилтын тоо; саатлын хурдыг шалгаж, эрэмбэлсэний дараа хамгийн бага сааталаас татаж авах хурдны туршилтын тоо; (өгөгдмөл 10)
    -dt 10
	Татаж авах хурдыг хэмжих хугацаа; нэг IP-ийн татаж авах хурдыг хэмжих дээд хугацаа, нэгж: секунд; (өгөгдмөл 10)
    -tl 200
	Дундаж саатлын дээд хязгаар; зөвхөн заасан дундаж сааталаас доогуур гаралтын IP-г бусад дээд/доод хязгаартай тааруулж болно; (өгөгдмөл 9999 мс)
    -tll 40
	Дундаж саатлын доод хязгаар; зөвхөн гаралтын IP-г заасан дундаж саатлаас өндөр, хуурамч IP-г шүүх бусад дээд/доод хязгаартай тааруулж болно; (өгөгдмөл 0 мс)
    -sl 5
	Татаж авах хурдны доод хязгаар; зөвхөн заасан таталтын хурдаас өндөр IP гаралт хийх ба заасан тоо [-dn] цуглуулахад хурдны хэмжилт зогсох болно; (өгөгдмөл 0.00 MB/s)
    -p 10
	Үр дүнгийн тоог харуулах; хурд хэмжсэний дараа заасан үр дүнг шууд харуулах, хэрэв 0 бол үр дүнг харуулахгүйгээр гарах болно; (өгөгдмөл 10)
    -f ip.txt
	IP сегментийн өгөгдлийн файл; замд хоосон зай байгаа бол хашилт нэмнэ үү; бусад CDN IP сегментүүдийг дэмжинэ; (өгөгдмөл ip.txt)
    -o result.csv
	Үр дүнгийн файлыг бичнэ үү; хэрэв замд хоосон зай байгаа бол хашилт нэмнэ үү; утга хоосон байвал файлыг бичихгүй [-o ""]; (өгөгдмөл үр дүн.csv)
    -dd
	Татаж авах хурдны тестийг идэвхгүй болгох; идэвхгүй болгосон үед хурдны туршилтын үр дүнг сааталаар эрэмбэлнэ (өгөгдмөл нь татаж авах хурдаар эрэмблэгдсэн); (өгөгдмөл тохиргоонд идэвхтэй)
    -ipv6
	IPv6 хурдыг хэмжих горим; IP сегментийн өгөгдлийн файл нь зөвхөн IPv6 IP сегментүүдийг агуулж байгаа эсэхийг шалгаарай, програм хангамж нь IPv4+IPv6 хурдыг нэгэн зэрэг хэмжихийг дэмждэггүй; (өгөгдмөл IPv4)
    -allip
	Хурд хэмжих бүх IP; IP сегмент дэх IP бүрийн хурдыг хэмжих (зөвхөн IPv4 дэмжигддэг); (анхдагчаар IP сегмент бүрийг нэг IP-д санамсаргүй байдлаар шалгадаг)
    -v
	програмын хувилбарыг хэвлэх + хувилбарын шинэчлэлтийг шалгах
    -h
	тусламжийн зааврыг хэвлэх
	`
	var minDelay, maxDelay, downloadTime int
	flag.IntVar(&task.Routines, "n", 200, "测速线程数量")
	flag.IntVar(&task.PingTimes, "t", 4, "延迟测速次数")
	flag.IntVar(&task.TCPPort, "tp", 443, "指定测速端口")
	flag.IntVar(&maxDelay, "tl", 9999, "平均延迟上限")
	flag.IntVar(&minDelay, "tll", 0, "平均延迟下限")
	flag.IntVar(&downloadTime, "dt", 10, "下载测速时间")
	flag.IntVar(&task.TestCount, "dn", 10, "下载测速数量")
	flag.StringVar(&task.URL, "url", "https://cf.xiu2.xyz/Github/CloudflareSpeedTest.png", "下载测速地址")
	flag.BoolVar(&task.Disable, "dd", false, "禁用下载测速")
	flag.BoolVar(&task.IPv6, "ipv6", false, "启用IPv6")
	flag.BoolVar(&task.TestAll, "allip", false, "测速全部 IP")
	flag.StringVar(&task.IPFile, "f", "ip.txt", "IP 数据文件")
	flag.Float64Var(&task.MinSpeed, "sl", 0, "下载速度下限")
	flag.IntVar(&utils.PrintNum, "p", 10, "显示结果数量")
	flag.StringVar(&utils.Output, "o", "result.csv", "输出结果文件")
	flag.BoolVar(&printVersion, "v", false, "打印程序版本")
	flag.Usage = func() { fmt.Print(help) }
	flag.Parse()

	if task.MinSpeed > 0 && time.Duration(maxDelay)*time.Millisecond == utils.InputMaxDelay {
		fmt.Println("[小提示] 在使用 [-sl] 参数时，建议搭配 [-tl] 参数，以避免因凑不够 [-dn] 数量而一直测速...")
	}
	utils.InputMaxDelay = time.Duration(maxDelay) * time.Millisecond
	utils.InputMinDelay = time.Duration(minDelay) * time.Millisecond
	task.Timeout = time.Duration(downloadTime) * time.Second

	if printVersion {
		println(version)
		fmt.Println("检查版本更新中...")
		checkUpdate()
		if versionNew != "" {
			fmt.Printf("*** 发现新版本 [%s]！请前往 [https://github.com/XIU2/CloudflareSpeedTest] 更新！ ***", versionNew)
		} else {
			fmt.Println("当前为最新版本 [" + version + "]！")
		}
		os.Exit(0)
	}
}

func main() {
	go checkUpdate()    // 检查版本更新
	task.InitRandSeed() // 置随机数种子

	fmt.Printf("# XIU2/CloudflareSpeedTest %s \n\n", version)

	// 开始延迟测速
	pingData := task.NewPing().Run().FilterDelay()
	// 开始下载测速
	speedData := task.TestDownloadSpeed(pingData)
	utils.ExportCsv(speedData)
	speedData.Print(task.IPv6)

	if versionNew != "" {
		fmt.Printf("\n*** 发现新版本 [%s]！请前往 [https://github.com/XIU2/CloudflareSpeedTest] 更新！ ***\n", versionNew)
	}
	endPrint()
}

func endPrint() {
	if utils.NoPrintResult() {
		return
	}
	if runtime.GOOS == "windows" { // 如果是 Windows 系统，则需要按下 回车键 或 Ctrl+C 退出（避免通过双击运行时，测速完毕后直接关闭）
		fmt.Printf("按下 回车键 或 Ctrl+C 退出。")
		var pause int
		fmt.Scanln(&pause)
	}
}

// 检查更新
func checkUpdate() {
	timeout := 10 * time.Second
	client := http.Client{Timeout: timeout}
	res, err := client.Get("https://api.xiu2.xyz/ver/cloudflarespeedtest.txt")
	if err != nil {
		return
	}
	// 读取资源数据 body: []byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// 关闭资源流
	defer res.Body.Close()
	if string(body) != version {
		versionNew = string(body)
	}
}
