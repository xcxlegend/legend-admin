$('.datetime').datetimepicker({
    format: 'yyyy-mm-dd hh:ii:00',
    weekStart: 1,
    language: 'zh-CN',
    autoclose: true,
    minuteStep: 1,
    todayBtn: true,
});


var doPost = function (url, params, callback) {
    $.post(url, params, function (data) {
        layer.msg(data.info)
        if (callback != undefined){
            callback()
        }
    })
}

var postDataFormatter = function (data) {
    if (data.status != undefined && !data.status){
        layer.msg(data.info)
        return data
    }
    return data.data || {}
    // var newData = []
    // try {
    //     var js = JSON.parse(data.data)
    //     console.log(js)
    //     return js
    // }catch(e){
    //     console.log(e)
    //     return {}
    // }
}