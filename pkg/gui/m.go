package gui

import (
	"fmt"
	"ip-proxy/pkg/config"
	phttp "ip-proxy/pkg/http"
	"log"
	"time"

	"github.com/lxn/walk"

	. "github.com/lxn/walk/declarative"
)

func Start() {
	var Mw *walk.MainWindow
	var number *walk.Label
	var ipAC *walk.Label
	var ipPW *walk.Label
	var netDeviceName *walk.Label
	var sStatus *walk.Label
	var ip *walk.Label
	var port *walk.Label
	var version *walk.Label

	animal := Animal{
		Number:        config.Number,
		IpAC:          config.IpAC,
		IpPW:          config.IpPW,
		NetDeviceName: config.NetDeviceName}

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(3))
			if config.ProxyRun {
				sStatus.SetText(fmt.Sprintln(" 当前代理服务器状态：运行中..."))
			} else {
				sStatus.SetText(fmt.Sprintln(" 当前代理服务器状态：关闭"))
			}

			ip.SetText(fmt.Sprintf(" 当前公网IP：%s", config.Ip))
			port.SetText(fmt.Sprintf(" 当前代理端口号：%d", config.Port))
		}
	}()

	if err := (MainWindow{
		AssignTo: &Mw,
		Title:    "代理终端",
		MinSize:  Size{320, 240},
		Size:     Size{400, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			PushButton{
				Text: "启动代理",
				OnClicked: func() {
					if config.ProxyRun {
						log.Println("当前服务器正在启动中，请先关闭服务器...")
					} else {
						go func() {
							config.CommandChan <- config.StartCommand
						}()
					}
				},
			},
			PushButton{
				Text: "关闭代理",
				OnClicked: func() {
					go func() {
						config.CommandChan <- config.ShutdownCommand
					}()
				},
			},
			PushButton{
				Text: "重新拨号",
				OnClicked: func() {
					go func() {
						config.CommandChan <- config.ShutdownThenStartCommand
					}()
				},
			},
			PushButton{
				Text: "编辑代理服务器信息",
				OnClicked: func() {
					if cmd, err := RunAnimalDialog(Mw, &animal); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						number.SetText(fmt.Sprintf(" 设备唯一号：%s", animal.Number))
						ipAC.SetText(fmt.Sprintf(" PPPOE拨号账号：%s", animal.IpAC))
						ipPW.SetText(fmt.Sprintf(" PPPOE拨号密码：%s", animal.IpPW))
						netDeviceName.SetText(fmt.Sprintf(" 网络设备名：%s", animal.NetDeviceName))

						config.Number = animal.Number
						config.IpAC = animal.IpAC
						config.IpPW = animal.IpPW
						config.NetDeviceName = animal.NetDeviceName
					}
				},
			},
			Label{
				AssignTo: &version,
				Text:     fmt.Sprintf(" 客户端版本号：%s", config.Version),
			},
			Label{
				AssignTo: &ip,
				Text:     fmt.Sprintf(" 当前公网IP：%s", ""),
			},
			Label{
				AssignTo: &port,
				Text:     fmt.Sprintf(" 当前代理端口号：%s", ""),
			},
			Label{
				AssignTo: &sStatus,
				Text:     fmt.Sprintf(" 当前代理服务器状态：%s", "关闭"),
			},
			Label{
				AssignTo: &number,
				Text:     fmt.Sprintf(" 设备唯一号：%s", animal.Number),
			},
			Label{
				AssignTo: &ipAC,
				Text:     fmt.Sprintf(" PPPOE拨号账号：%s", animal.IpPW),
			},
			Label{
				AssignTo: &ipPW,
				Text:     fmt.Sprintf(" PPPOE拨号密码：%s", animal.IpPW),
			},
			Label{
				AssignTo: &netDeviceName,
				Text:     fmt.Sprintf(" 网络设备名：%s", animal.NetDeviceName),
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	lv, err := NewLogView(Mw)
	if err != nil {
		log.Fatal(err)
	}

	//writers := []io.Writer{
	//	lv,
	//	os.Stderr,
	//	os.Stdout}
	//fileAndStdoutWriter := io.MultiWriter(writers...)
	log.SetOutput(lv)

	Mw.Run()
}

type Animal struct {
	S             *phttp.Server
	T             *time.Ticker
	Rt            *time.Ticker
	Number        string
	IpAC          string
	IpPW          string
	NetDeviceName string
}

func RunAnimalDialog(owner walk.Form, animal *Animal) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         "编辑代理服务器信息",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "animal",
			DataSource:     animal,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{

					Label{
						Text: "设备唯一号:",
					},
					LineEdit{
						Text: Bind("Number"),
					},

					Label{
						Text: "PPPOE拨号账号:",
					},
					LineEdit{
						Text: Bind("IpAC"),
					},

					Label{
						Text: "PPPOE拨号密码:",
					},
					LineEdit{
						Text: Bind("IpPW"),
					},
					Label{
						Text: "网络设备名:",
					},
					LineEdit{
						Text: Bind("NetDeviceName"),
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}
