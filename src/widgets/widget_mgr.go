package widgets

import "github.com/astaxie/beego"

/**
 * @brief 插件管理.
 * 在main里面执行, 对程序进行初始化工作或者需要在协程里同时执行的
 * @author legend.Xie
 * @date 2018-10-29 11:08:14
 * @function GetWidgetMgr() 获取管理器单例
 * @function RegisterWidgets() 请在此方法内填写注册的插件. 有顺序
 */

/** Widget
 * @brief 插件接口 定义插件需要实现此接口
 * @function Init() bool 接口的初始化方法
 * @function Daemon() 接口是否是goroutine执行
 * @function Run()  接口的开始执行的方法 如果非goroutine执行返回false则可以打断并退出程序
 */
type Widget interface{
	Init() bool
	Daemon() bool
	Run() bool
}

type WidgetMgr struct{
	widgets []Widget
}

var widgetMgr *WidgetMgr

func (this *WidgetMgr) Init(){
	beego.Debug("WidgetMgr Init")
	this.widgets = []Widget{}
	RegisterWidgets()
}

func (this *WidgetMgr) Register(widget Widget)  {
	this.widgets = append(this.widgets, widget)
}

func (this *WidgetMgr) Run() bool {
	
	for _, w := range this.widgets {
		if !w.Init() { return false }
		if w.Daemon() {
			go w.Run()
		} else {
			if !w.Run() {
				return false
			}
		}
	}
	return true
}

func init() {
	widgetMgr = new(WidgetMgr)
	widgetMgr.Init()
}

func GetWidgetMgr() *WidgetMgr  {
	return widgetMgr
}

func RegisterWidgets()  {
	// #TODO 放置注册的插件. 可定义顺序
	// widgetMgr.Register(...)
	 widgetMgr.Register(quickBuilder)
}