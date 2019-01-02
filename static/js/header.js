window.onload=function(){
    document.body.addEventListener('touchstart', function() {}, false);
}


var global = {
    input: {},
    lang: {},
    URL: "",
}

function UrlSearch() {
    var name,value;
    var str=location.href; //取得整个地址栏
    var num=str.indexOf("?")
    str=str.substr(num+1); //取得所有参数   stringvar.substr(start [, length ]

    var arr=str.split("&"); //各个参数放到数组里
    query = {}
    for(var i=0;i < arr.length;i++){
        num=arr[i].indexOf("=");
        if(num>0){
            name=arr[i].substring(0,num);
            value=arr[i].substr(num+1);
            query[name]=value;
        }
    }
    return query
}

var initGlobal = function () {
    query = UrlSearch()
    for (k in query){
        if (global.input[k] == undefined){
            global.input[k] = query[k]
        }
    }

    global.URL = window.location.href;

    if (global.URL.charAt(global.URL.length - 1) == "/"){
        global.URL = global.URL.substr(0, global.URL.length - 1)
    }

    var index = global.URL.lastIndexOf("/")
    global.URL = global.URL.substr(0, index)
}
initGlobal()

// 实现的大部分都需求的函数
// 格式化时间戳
function fmtStamp2dt(timestamp) {
    return timestamp ? phpjs.date("Y-m-d H:i:s", timestamp) : ""
}

function getDateTime() {
    return fmtStamp2dt(new Date().getTime()/1000)
}

// 将数据读入到编辑的form表单里 data=row form=$('#')
function loadData2Form(data, form) {
    $.each(data, function (k, e) {
        if (form.find('[name='+k+']').length != 0){
            if (form.find('input[name='+k+']').length != 0){
                if (form.find('input[name='+k+']').attr('type') == 'radio'){
                    form.find('input[name='+k+'][value='+e+']').attr('checked', true)
                }else {
                    form.find('[name='+k+']').val(e)
                }
            } else if (form.find('select[name='+k+']').length != 0 && form.find('select[name='+k+']').attr("multiple") != undefined){
                if (typeof(e) == "string"){
                    form.find('select[name='+k+']').val(e.split(","))
                }else{
                    form.find('select[name='+k+']').val(e)
                }
            } else {
                form.find('[name='+k+']').val(e)
            }
        }
    })
}

// 实现将form里面的数组数据 合并成,分割
function parseParamArray2String(formArray, keys) {

    var m = {}
    for (var k in formArray){
        var item = formArray[k]
        if (m[item.name] == undefined){
            m[item.name] = []
        }
        m[item.name].push(item.value)
    }

    var newArray = []
    for (var k in m){
        var value = m[k].join()
        newArray.push({
            name: k,
            value: value,
        })
    }
    return newArray
}


function parseJson2From(json, form, spacename) {
    // console.log(spacename)
    $.each(json, function (key, val) {
        switch (typeof val){
            case "object":
                var jspacename = spacename == "" ? key : spacename + ("\\." + key)
                parseJson2From(val, form, jspacename)
                break
            case "string":
            case "number":
            case "boolean":
                var ikey
                if (typeof key == "number"){
                    ikey =  spacename + "\\[\\]"
                    form.find("[name="+ikey+"]").eq(key).val(val)
                }else{
                    ikey = spacename == "" ? key : (spacename  + "\\." + key)
                    form.find("[name="+ikey+"]").val(val)
                }
                // console.log(ikey, val)
                break
        }
    })
}

function iframe_download_file(url)
{

    $('#update').empty()
    if(typeof(iframe_download_file.iframe)== "undefined")
    {
        var iframe = document.createElement("iframe");
        iframe_download_file.iframe = iframe;
        document.body.appendChild(iframe_download_file.iframe);
    }

    iframe_download_file.iframe.src = url;
    iframe_download_file.iframe.style.display = "none";
}

var formatterListItemData = function (data) {
    d = JSON.parse(data)
    content = ""
    $.each(d, function (k, v) {
        content += k +":" + JSON.stringify(v) +"<br/>"
    })
    return content
}

var popBox = {
    success: function (msg) {
        layer.msg(msg, {icon: 1})
    },
    fail: function (msg) {
        layer.msg(msg, {icon: 2})
    },
    frame: function (url) {
        var _index = layer.open({
            type: 2,
            // title: '表单',
            content: url,
            maxmin: true,
            area: ['550px', '550px'],
            // btn: ['确定', '取消'],
            // yes: function(index, layero){
            //     layero.find('.form-submit-btn').click()
            // }
        });
    },
    form: function (el, layui, fields) {
        var content = $(el).html()
        var _index = layer.open({
            type: 1,
            title: '表单',
            content: content,
            maxmin: true,
            area: ['550px', '550px'],
            btn: ['确定', '取消'],
            yes: function(index, layero){
                layero.find('.form-submit-btn').click()
            }
        });


        FORM_FIELDS.forEach(function (value) {
            if (value.element == 'select' && value.options != ""){
                loadSelectorOptions('select[name=' + value.name + ']', value.options)
            }
        });

        layui.form.render()
        var laydate = layui.laydate
        laydate.render({elem: '.lay-date'})
        laydate.render({elem: '.lay-datetime', type: 'datetime'})
        return _index
    }
}

var request = {
    callback: function (res, next) {
        if (res.error != undefined && res.error != "") {
            popBox.fail(res.error)
        }
        next(res)
    },

    post: function (url, data, callback) {
        $.post(url, data,  (res) => {
             this.callback(res, callback)
        });
    },

    put: function (url, data, callback) {
        $.ajax({
            type: "PUT",
            data: data,
            url: url,
            success: (res) => {
                this.callback(res, callback)
            },
        });
    },

    options: function (url, data, callback) {
        $.ajax({
            type: "OPTIONS",
            data: data,
            url: url,
            success: (res) => {
                this.callback(res, callback)
            },
        });
    },

    del: function (url, data, callback) {
        $.ajax({
            type: "DELETE",
            dataType: "json",
            contentType:"application/json",
            data: JSON.stringify(data),
            url: url,
            success: (res) => {
                this.callback(res, callback)
            },
        });
    }
}


String.format = function(src){
    if (arguments.length == 0) return null;
    var args = Array.prototype.slice.call(arguments, 1);
    return src.replace(/\{(\d+)\}/g, function(m, i){
        return args[i];
    });
};

(function($){
    $.fn.serializeJson=function(){
        var serializeObj={};
        $(this.serializeArray()).each(function(){
            serializeObj[this.name]=this.value;
        });
        return serializeObj;
    };
})(jQuery);


var renderForm = function (FORM_FIELDS) {
    var container = $('#formBox')
    var FORM_TEMPLDATE = container.html();
    var ITEM_TEMPLATE = '<div class="layui-form-item">\n' +
        '                <label class="layui-form-label">__LABEL__</label>\n' +
        '                <div class="layui-input-block">\n' +
        '                    __FIELD__' +
        '                </div>\n' +
        '            </div>'
    var fields = '';
    FORM_FIELDS.forEach(function (v, i) {
        var item = ITEM_TEMPLATE;
        var field = '';
        v.attrs["class"] = v.attrs["class"] || '';
        switch (v.element){
            case 'select':
                field = '<select name="__NAME__" __ATTRS__ >__OPTIONS__</select>'
                var options = ''
                for (var k in (v.options || [])){
                    var op = v.options[k]
                    options += '<option value=" ' + op.key + ' "> ' + op.value + '</option>\n';
                }
                field = field.replace("__OPTIONS__", options)
                break
            case 'textarea':
                v.attrs["class"] += ' layui-textarea'
                field = '<textarea name="__NAME__" __ATTRS__></textarea>';
                break;
            default:
                field = '<input name="__NAME__" __ATTRS__>'
                break
        }

        var att = "";
        v.attrs['class'] += ' layui-input'
        for (var key in v.attrs) {
            var val = v.attrs[key]
            att += ' ' + key + '="' + val + '" '
        }
        field = field.replace("__NAME__", v.name)
        field = field.replace("__ATTRS__", att)
        item = item.replace("__LABEL__", v.label)
        item = item.replace("__FIELD__", field)
        fields += item + "\n";
    })
    html = FORM_TEMPLDATE.replace("__FORM_FIELDS__", fields);
    container.html(html);
}

var PageInit = function (URL, TABLE_COLS, FORM_FIELDS, call) {
    renderForm(FORM_FIELDS)
    ajaxOptionsFields(FORM_FIELDS)
    const TABLE_ID = "Id"
    var formBox;
    layui.use(['form','table','laydate'], function(){
        var form = layui.form;
        var table = layui.table;
        const active = {
            refreshOper: function () {
                table.reload(TABLE_ID)
            },
            addOper: function () {
                formBox = popBox.form('#formBox', layui, FORM_FIELDS)
            },
            editOper: function () {
                var status = table.checkStatus('Id')
                var data = status.data
                if (data.length == 0) {
                    layer.msg("未选择任何行")
                    return
                }
                if (data.length > 1) {
                    layer.msg("只能选择编辑一行")
                    return
                }
                formBox = popBox.form('#formBox', layui, FORM_FIELDS)
                form.val("form", data[0]);
                // layer.msg(JSON.stringify(data))
            },
            delOper: function () {
                var status = table.checkStatus('Id')
                var data = status.data
                if (data.length == 0) {
                    layer.msg("未选择任何行")
                    return
                }
                layer.confirm("是否确定删除", function (index) {
                    var ids = []
                    data.forEach( (d, i) => {
                        ids.push(d.Id)
                });
                    request.del(URL, ids, function (res) {
                        if (res.error == "") {
                            popBox.success("成功删除: " + res.data.id.join(", "))
                        }
                        table.reload(TABLE_ID)
                    })
                })
            }
        }

        form.verify({
            length: function (value, item) {
                var min = parseInt($(item).attr('lay-min'))
                var max = parseInt($(item).attr('lay-max'))
                if (min > 0 && value.length < min){
                    return "长度最小需要: " + min
                }
                if (max > 0 && value.length > max) {
                    return "长度最大为: " + max
                }
            }
        });

        //监听提交
        form.on('submit(formBtn)', function(data){
            if (data.field.Id == undefined || parseInt(data.field.Id) == 0){
                request.post(URL, data.field, function (res) {
                    if (res.error == "") {
                        ajaxOptionsFields(FORM_FIELDS)
                        popBox.success("添加成功")
                        active.refreshOper()
                        layer.close(formBox)
                    }
                })
            } else {
                request.put(URL, data.field, function (res) {
                    if (res.error == "") {
                        ajaxOptionsFields(FORM_FIELDS)
                        popBox.success("编辑成功")
                        active.refreshOper()
                        layer.close(formBox)
                    }
                })
            }
            return false;
        });

        var cols = [
            //表头
            // 固定头
            {type: "checkbox", fixed: 'left'},
            // 固定ID
            {field: 'Id', title: 'ID', width:80, sort: true, fixed: 'left'},
            // field => name, title => label
        ];

        cols = cols.concat(TABLE_COLS)
        //第一个实例
        table.render({
            elem: '#table',
            url: URL, //数据接口
            page: true, //开启分页
            toolbar: '#toolbar',
            defaultToolbar: ['filter', 'print', 'exports'],
            loading: true,
            id: TABLE_ID,
            cols: [cols]
        });

        table.on('toolbar(t)', function (obj) {
            var event = obj.event
            active[event] ? active[event].call(this): '';
        })

        call()
    });
}

// op string
var getOptionsByFuncStr = function (op) {
    console.log("load option:", op)
    var options = [];
    try {
        var opt = eval(op)
        options = typeof opt == "function" ? opt(): opt;
    } catch(err){
        console.error("function " + op + " is not defined")
    }
    return options
}

var loadSelectorOptions = function (elem, op) {
    var html = '';
    var options;
    if (typeof op == "object"){
        options = op || [];
    } else {
        options = getOptionsByFuncStr(op) || [];
    }
    options.forEach(function (value) {
        html += '<option value="'+value.key+'">'+value.value+'</option>';
    })
    $(elem).html(html);
}

var ajaxOptionsFields = function (FORM_FIELDS) {
    FORM_FIELDS.forEach(function (value) {
        if (value.element == 'select' && value.options_url != undefined && value.options_url != ''){
            request.options(value.options_url, {}, function (res) {
                value.options = [{
                    "key": "",
                    "value": "",
                }].concat(res);
            })
        }
    })
}