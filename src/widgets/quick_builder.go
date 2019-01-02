package widgets

import (
	"github.com/astaxie/beego"
	"os"
	"flag"
	"fmt"
	. "common"
	_ "models"
	"models"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"strings"
	"reflect"
	"github.com/astaxie/beego/orm"
)


/** QuickBuilder
 * @brief 快速构建工具.
 * 使用model和JSON快速的创建. controller, model, router, view文件.
 * @author legend.Xie
 * @date 2018-10-29 11:08:14
 */
type QuickBuilder struct{
	model string
	force bool
	files []string
}

var quickBuilder *QuickBuilder

func init()  {
	quickBuilder = new(QuickBuilder)
}

func (this *QuickBuilder) Init() bool {
	beego.Debug("QuickBuilder init")
	return true
}

func (this *QuickBuilder) Daemon() bool{
	return false
}

func (this *QuickBuilder) Run() bool{
	beego.Debug("QuickBuilder run")
	// flag处理

	// gen: gen config
	// build: build file
	beego.Debug(os.Args)

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		os.Args = append(os.Args[0:1], os.Args[2:]...)
		if cmd == "build"{
			var (
				op string
				//c string
				h bool
				m string
				force bool
				files string
			)
			flag.StringVar( &op,  "op", "gen", "operation `gen` gen conf json file. `make` start make all file")
			//flag.StringVar(&c, "c","", "use conf file to make all")
			flag.StringVar(&m, "m","", "make from model in models, model file name should be lower like `user.go`, model struct should be captal like `User` ")
			flag.BoolVar( &h, "h", false, "help")
			flag.BoolVar( &force, "f", false, "force drop table")
			flag.StringVar( &files, "files", "", "files inclucd `view`, `controller`, `model`, `view`, use `,` to join")
			flag.Usage = this.usage
			flag.Parse()
			if h {
				flag.Usage()
			}

			m = strings.ToLower(m)

			this.model = m
			this.force = force
			this.files = strings.Split(files, ",")

			if op == "make" {
				this.make()
			} else {
				this.gen()
			}
			return false
		}
	}
	return true
}

// 生成所有文件
func (this *QuickBuilder) make() bool {
	modelName := this.model
	force := this.force
	if !this.check(modelName) {
		return false
	}

	if force {
		if err := this.DropTable(modelName); err != nil {
			beego.Error("drop table ", modelName, " error:", err)
			return false
		}
		// drop table
	}

	models.Syncdb()
	beego.Debug(fmt.Sprintf(`start make model: %v`, modelName))
	jsonName := modelName + ".json"
	filename := filepath.Join(PATH_JSON, jsonName)
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		beego.Error(fmt.Sprintf(`json file %v not exist`, filename))
		return false
	}

	jsonContent, err := ioutil.ReadFile(filename)
	if err != nil {
		beego.Error(fmt.Sprintf(`read json file:%v err:%v`, filename, err))
		return false
	}

	conf := &ConfigJSON{}
	if err := json.Unmarshal(jsonContent, conf); err != nil {
		beego.Error(fmt.Sprintf(`parse json file:%v err:%v`, filename, err))
		return false
	}

	buildData := this.genConfigStruct(conf)

	if len(this.files) == 0 {
		this.files = []string{"model", "controller", "view", "router"}
	}

	for _, f := range this.files {

		switch f {
		case "model":
			// make model
			if err := this.makeModel(buildData); err != nil {
				return false
			}
		case "controller":
			// make controller
			if err := this.makeController(buildData); err != nil {
				return false
			}
		case "router":
			// make router
			if err := this.makeRouter(buildData); err != nil {
				return false
			}
		case "view":
			// make view
			if err := this.makeView(buildData); err != nil {
				return false
			}
		default:
			break
		}
	}

	return true
}

func (this *QuickBuilder) DropTable(modelName string) error {
	_, err := orm.NewOrm().Raw(fmt.Sprintf(`DROP TABLE IF EXISTS %s`, modelName)).Exec()
	return err
}

// 将model转成配置json对应的表单配置
func (this *QuickBuilder) renderFormJSON(obj interface{}) []*ConfigForm  {
	form := []*ConfigForm{}
	objT := reflect.TypeOf(obj)
	//beego.Debug("obj type:", obj, objT)
	objT = objT.Elem()

	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		//beego.Debug("obj field:", i, fieldT)
		this.renderFormJSONField(&form, fieldT)
		//tags := strings.Split(fieldT.Tag.Get("form"), ",")
	}
	return form
}

// 根据model渲染页面的form元素
func (this *QuickBuilder) renderFormJSONField(form *[]*ConfigForm, fieldT reflect.StructField) error {
	if fieldT.Type.Kind() == reflect.Ptr && fieldT.Type.Elem().Kind() == reflect.Struct || fieldT.Type.Kind() == reflect.Struct {
		//*form = append((*form), this.renderFormJSON(fieldT.Type)...)
		return nil
	}
	tagForm := fieldT.Tag.Get("form")
	//beego.Debug("tags:", fieldT, tagForm)
	if tagForm == "" {
		return nil
	}
	tags := strings.Split(fieldT.Tag.Get("form"), ",")
	if len(tags) == 0 {
		return nil
	}
	name := fieldT.Tag.Get("json")
	if name == "-" || name == ""{
		name = tags[0]
	}
	c := &ConfigForm{
		Name: name,
		Element: "input",
		Options: "",
		Attrs: map[string]string{
			"autocomplete": "off",
			"lay-verify": "",
		},
	}
	if len(tags) > 1{
		t := tags[1]
		if t == "select" {
			c.Element = "select"
			c.Attrs["lay-search"] = ""
		} else if t == "textarea"{
			c.Element = "textarea"
		} else if t == "password" {
			c.Attrs["type"] = "password"
		} else if t == "date" {
			c.Attrs["class"] = "lay-date"
		} else if t == "datetime" {
			c.Attrs["class"] = "lay-datetime"
		} else {
			c.Attrs["type"] = "text"
		}
	}
	if len(tags) > 2 {
		l := tags[2]
		c.Label = l
	}
	if c.Element == "select" && len(tags) > 3 {
		os := strings.Split(tags[3], ":")
		if len(os) > 1 {
			if os[0] == "OptionsUrl" {
				c.OptionsUrl = os[1]
			} else {
				c.Options = os[1]
			}
		}
	}

	if (fieldT.Type.Kind() == reflect.Int || fieldT.Type.Kind() == reflect.Int64) && c.Element == "input" {
		c.Attrs["lay-verify"] += "|number"
	}

	verifys := strings.Split(fieldT.Tag.Get("valid"), ";")
	if len(verifys) > 0 {
		hasLength := false
		for _, v := range verifys {
			if v == "Required" {
				c.Attrs["lay-verify"] += "|required"
			} else {
				i := strings.Split(v, "(")
				if len(i) > 1{
					ii := strings.Split(i[1], ")")
					value := ii[0]
					if i[0] == "MinSize" {
						if !hasLength{
							c.Attrs["lay-verify"] += "|length"
							hasLength = true
						}
						c.Attrs["lay-min"] = value
					} else if i[0] == "MaxSize" {
						if !hasLength{
							c.Attrs["lay-verify"] += "|length"
							hasLength = true
						}
						c.Attrs["lay-max"] = value
					}
				}
			}
		}
	}
	*form = append((*form), c)
	return nil
}


func (this *QuickBuilder) makeView(conf *ConfigStruct) error  {
	_, err := os.Stat(conf.ViewTempFile)
	if err != nil && os.IsNotExist(err) {
		beego.Error(fmt.Sprintf(`view template file %v not exist`, conf.ViewTempFile))
		return err
	}

	content, err := ioutil.ReadFile(conf.ViewTempFile)
	if err != nil {
		beego.Error(fmt.Sprintf(`read view template file %v err:%v`, conf.ViewTempFile, err))
		return err
	}

	tips := []string{}
	for _, t := range conf.JSON.Tips {
		tips = append(tips, "<p>"+t+"</p>")
	}

	newContent := strings.Replace(string(content), "__TIPS__", strings.Join(tips, "\n"), -1)
	newContent = strings.Replace(newContent, "__TITLE__", conf.JSON.Title, -1)

	buttons := []string{}
	if _, ok := conf.JSON.HtmlOper["add"]; ok {
		buttons = append(buttons, fmt.Sprintf(`<button class="layui-btn layui-btn-normal table-oper-btn" lay-event="addOper">添加</button>`))
	}
	if _, ok := conf.JSON.HtmlOper["edit"]; ok {
		buttons = append(buttons, fmt.Sprintf(`<button class="layui-btn layui-btn-normal table-oper-btn" lay-event="editOper">编辑</button>`))
	}
	if _, ok := conf.JSON.HtmlOper["del"]; ok {
		buttons = append(buttons, fmt.Sprintf(`<button class="layui-btn layui-btn-danger table-oper-btn" lay-event="delOper">删除</button>`))
	}

	newContent = strings.Replace(newContent, "__BUTTONS__", strings.Join(buttons, "\n		"), -1)

	table := []map[string]interface{}{}

	for _, f := range conf.JSON.Form{
		table = append(table, map[string]interface{}{
			"field": f.Name,
			"title": f.Label,
		})
	}
	tableStr, _ := json.MarshalIndent(&table, "", "	")
	newContent = strings.Replace(newContent, "__TABLE_COLS__", string(tableStr), -1)

	form, _ := json.MarshalIndent(conf.JSON.Form, "", "	")
	newContent = strings.Replace(newContent, "__JS_FORM_FIELDS__", string(form), -1)
	newContent = strings.Replace(newContent, "__ROUTER__", conf.RoutePrefix, -1)

	dir := filepath.Dir(conf.ViewFile)
	os.MkdirAll(dir, 0777)

	if err := ioutil.WriteFile(conf.ViewFile, []byte(newContent), 0777); err != nil {
		beego.Error(fmt.Sprintf(`write view file %v err:%v`, conf.ViewFile, err))
		return err
	}
	beego.Debug("create view file success:", conf.ViewFile)

	return nil
}

func (this *QuickBuilder) makeRouter(conf *ConfigStruct) error {

	_, err := os.Stat(conf.RouteTempateFile)
	if err != nil && os.IsNotExist(err) {
		beego.Error(fmt.Sprintf(`router template file %v not exist`, conf.RouteTempateFile))
		return err
	}

	content, err := ioutil.ReadFile(conf.RouteTempateFile)
	if err != nil {
		beego.Error(fmt.Sprintf(`read router template file %v err:%v`, conf.RouteTempateFile, err))
		return err
	}
	newContent := strings.Replace(string(content), "//##", "", -1)
	newContent = strings.Replace(newContent, `// __ROUTER__ => path`, "", -1)
	newContent = strings.Replace(newContent, `// __CONTROLLER__ => controller`, "", -1)
	newContent = strings.Replace(newContent, "__ROUTER__", conf.RoutePrefix , -1)
	newContent = strings.Replace(newContent, "__CONTROLLER__", conf.ModelName + "Controller" , -1)

	if err := ioutil.WriteFile(conf.RouterFile, []byte(newContent), 0666); err != nil {
		beego.Error(fmt.Sprintf(`write router file %v err:%v`, conf.RouterFile, err))
		return err
	}
	beego.Debug("create router file success:", conf.RouterFile)
	return nil


	return nil
}



func (this *QuickBuilder) makeController(conf *ConfigStruct) error {
	_, err := os.Stat(conf.ControllerTempFile)
	if err != nil && os.IsNotExist(err) {
		beego.Error(fmt.Sprintf(`controller template file %v not exist`, conf.ControllerTempFile))
		return err
	}

	content, err := ioutil.ReadFile(conf.ControllerTempFile)
	if err != nil {
		beego.Error(fmt.Sprintf(`read controller template file %v err:%v`, conf.ControllerTempFile, err))
		return err
	}
	newContent := strings.Replace(string(content), "T__CONTROLLERNAME__", conf.ModelName, -1)
	newContent = strings.Replace(newContent, "T__MODELTEMPLDATE__", conf.ModelName, -1)
	newContent = strings.Replace(newContent, "__VIEW_INDEX__", filepath.Join(conf.JSON.Model, "index.html"), -1)

	if err := ioutil.WriteFile(conf.ControllerFile, []byte(newContent), 0666); err != nil {
		beego.Error(fmt.Sprintf(`write controller file %v err:%v`, conf.ControllerFile, err))
		return err
	}
	beego.Debug("create controller file success:", conf.ControllerFile)

	return nil
}

func (this *QuickBuilder) makeModel(conf *ConfigStruct) error {

	_, err := os.Stat(conf.ModelTempFile)
	if err != nil && os.IsNotExist(err) {
		beego.Error(fmt.Sprintf(`model template file %v not exist`, conf.ModelTempFile))
		return err
	}

	content, err := ioutil.ReadFile(conf.ModelTempFile)
	if err != nil {
		beego.Error(fmt.Sprintf(`read model template file %v err:%v`, conf.ModelTempFile, err))
		return err
	}
	newContent := strings.Replace(string(content), "type T__MODELTEMPLDATE__ struct {Base}", "", 1)
	newContent = strings.Replace(newContent, "T__MODELTEMPLDATE__", conf.ModelName, -1)

	if err := ioutil.WriteFile(conf.ModelFuncFile, []byte(newContent), 0666); err != nil {
		beego.Error(fmt.Sprintf(`write model file %v err:%v`, conf.ModelFuncFile, err))
		return err
	}
	beego.Debug("create model file success:", conf.ModelFuncFile)
	return nil
}

// 生成配置文件
func (this *QuickBuilder) gen() bool {
	modelName := this.model
	// check model
	if !this.check(modelName) {
		return false
	}

	jsonName := modelName + ".json"

	filename := filepath.Join(PATH_JSON, jsonName)

	model, _ := models.ModelCache[modelName]

	content, _ := json.MarshalIndent(&ConfigJSON{
		Model: modelName,
		HtmlOper: map[string]bool{"add":true, "edit":true, "del":true},
		Form: this.renderFormJSON(model),
		RouterPrefix: "/admin",
		Tips: []string{"默认说明"},
	}, "", "	")

	if err := ioutil.WriteFile(filename, content, 0777); err != nil {
		beego.Error(fmt.Sprintf(`write json file %v err:%v`, filename, err))
		return false
	}

	beego.Warn(fmt.Sprintf(`json file has created, please use "./main --op=make -m=%v" to make files`, modelName))
	return true
}

func (this *QuickBuilder) check(m string) bool  {
	_, ok := models.ModelCache[m]
	if !ok {
		beego.Error("model name <", m, "> not exist")
		return false
	}
	/* // 屏蔽检查已存在
	// 检查controller view router是否存在
	controller_file := filepath.Join(PATH_CONTROLLERS, m + ".go")
	view_file := filepath.Join(PATH_VIEWS, m, "index.html")
	router_file := filepath.Join(PATH_ROUTERS, "router_" + m + ".go")

	_, err := os.Stat(controller_file)
	if err == nil {
		beego.Error(fmt.Sprintf(`controller file %v has exist`, controller_file))
		return false
	}

	_, err = os.Stat(view_file)
	if err == nil {
		beego.Error(fmt.Sprintf(`view file %v has exist`, view_file))
		return false
	}

	_, err = os.Stat(router_file)
	if err == nil {
		beego.Error(fmt.Sprintf(`router file %v has exist`, router_file))
		return false
	}
	*/
	return true
}

func (this *QuickBuilder) genConfigStruct(confJson *ConfigJSON) *ConfigStruct {
	config := &ConfigStruct{
		JSON: confJson,
		ModelName: CamelString(confJson.Model),
		ControllerFile: filepath.Join(PATH_CONTROLLERS, confJson.Model + ".go"),
		ControllerTempFile: filepath.Join(PATH_CONTROLLERS, "controller_template.go"),
		RouterFile: filepath.Join(PATH_ROUTERS, "router_auto_" + confJson.Model + ".go"),
		RoutePrefix: confJson.RouterPrefix + "/" + confJson.Model,
		RouteTempateFile:filepath.Join(PATH_ROUTERS, "router_template.go"),
		ViewFile: filepath.Join(PATH_VIEWS, confJson.Model, "index.html"),
		ViewTempFile: filepath.Join(PATH_VIEWS, "quick_builder", "view_template.html"),
		ModelFuncFile: filepath.Join(PATH_MODELES, confJson.Model + "_func.go"),
		ModelTempFile: filepath.Join(PATH_MODELES, "model_func_template.go"),
	}

	return config

}


func (this *QuickBuilder) usage() {
	fmt.Fprintf(os.Stderr, `quickBuilder version: 1.0.0
Usage: ./main  [-op operation "gen" or "make"] [-c confile]

Options:
`)
	flag.PrintDefaults()
}

type ConfigForm struct{
	Label string `json:"label"`
	Name string `json:"name"`
	Options interface{} `json:"options"`
	OptionsUrl string `json:"options_url"`
	Element string `json:"element"`
	Attrs map[string]string `json:"attrs"`
}

//type FormOption struct{
//	Key string `json:"key"`
//	Value string `json:"value"`
//}
type ConfigJSON struct{
	Title string `json:"title"`
	Model string `json:"model"`
	RouterPrefix string `json:"router_prefix"`
	HtmlOper map[string]bool `json:"html_oper"`
	Form []*ConfigForm `json:"form"`
	Tips []string `json:"tips"`
}

type ConfigStruct struct{
	JSON *ConfigJSON
	ModelName string
	ControllerFile string
	ControllerTempFile string
	RouterFile string
	RoutePrefix string
	RouteTempateFile string
	ViewFile string
	ViewTempFile string
	ModelFuncFile string
	ModelTempFile string
}

const (
	ROUTER_TEMPLATE = `beego.Router("__ROUTER__", &controllers.__CONTROLLER__{})`
)