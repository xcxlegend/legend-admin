

var URL = "/gm/gamer";
var gid = 0;

var data; // gamer主信息



var spec_options = [
];

$.each(specs, function (id, val) {
    var spec = {
        id: id,
        name: val.speciality_describe,
        property: val.property_id
    }
    spec_options.push(spec)
})



$(function() {
    //创建添加服务器窗口
    $("#dialog").dialog({
        modal: true,
        resizable: true,
        top: 150,
        closed: true,
        buttons: [{
            text: '保存',
            iconCls: 'icon-save',
            handler: function() {
                $("#form_mail").form('submit', {
                    url: '/gm/mail/send_to_gamer',
                    onSubmit: function() {
                        return $("#form_mail").form('validate');
                    },
                    success: function(r) {
                        var r = $.parseJSON(r);
                        if (r.status) {
                            $("#dialog").dialog("close");
                            doSearch($('#gamer_gid').val())
                        } else {
                            vac.alert(r.info);
                        }
                    }
                });
            }
        }, {
            text: '取消',
            iconCls: 'icon-cancel',
            handler: function() {
                $("#dialog").dialog("close");
            }
        }]
    });



    $("#dialog_player").dialog({
        modal: true,
        resizable: true,
        top: 150,
        closed: true,
        buttons: [{
            text: '保存',
            iconCls: 'icon-save',
            handler: function() {
                $("#form_player").form('submit', {
                    url: '/gm/gamer/add_player',
                    onSubmit: function() {
                        return $("#form_player").form('validate');
                    },
                    success: function(r) {
                        var r = $.parseJSON(r);
                        if (r.status) {
                            $("#dialog_player").dialog("close");
                            doSearch($('#gamer_gid').val())
                        } else {
                            vac.alert(r.info);
                        }
                    }
                });
            }
        }, {
            text: '取消',
            iconCls: 'icon-cancel',
            handler: function() {
                $("#dialog_player").dialog("close");
            }
        }]
    });




    $("#dialog_player_spec").dialog({
        modal: true,
        resizable: true,
        top: 150,
        closed: true,
        buttons: [{
            text: '保存',
            iconCls: 'icon-save',
            handler: function() {
                $("#form_player_spec").form('submit', {
                    url: '/gm/gamer/update_player',
                    onSubmit: function() {
                        return $("#form_player_spec").form('validate');
                    },
                    success: function(r) {
                        var r = $.parseJSON(r);
                        if (r.status) {
                            $("#dialog_player_spec").dialog("close");
                            doSearch($('#gamer_gid').val())
                        } else {
                            vac.alert(r.info);
                        }
                    }
                });
            }
        }, {
            text: '取消',
            iconCls: 'icon-cancel',
            handler: function() {
                $("#dialog_player_spec").dialog("close");
            }
        }]
    });






});

function mailBox() {
    var id = $('#mail_gid').val();
    if (id == ""){
        alert("请先搜索用户");
        return
    }
    $('#form_mail').form('clear');
    $('#mail_gid').val(id);
    $('#dialog').dialog('open');
}

function playerSpe() {
    console.log(data)
    $('#form_player_spec').form('clear');
    $('#form_player_spec_action').val('spec');
    var row = $("#datagrid_players").datagrid("getSelected");
    $('#form_player_spec_pid').val(row.id)
    $('input[name=gid]').val(data.gamer.id)

    $.each(data.players, function (i, e) {
        if (e.id == row.id && e.speciality != undefined){
            console.log("e:", e, e.speciality)
            $.each(e.speciality, function (i, s) {
                $('#spec_'+s.slotId).combobox('setValue', s.specialityId)
            })
        }
    })


    $('#dialog_player_spec').dialog('open');
}


var omail_items = $('#mail_items')
$.each(items, function (i, e) {
    omail_items.append('<option value="'+i+'">'+e.item_name+ '('+i+')'+ '</option>')
})

omail_items.change(function () {
    var _this = $(this)
//        console.log(_this.val());
    var attaches = $('input[name=attaches]').val()
    var num = $('input[name=attaches_num]').val()
    if (attaches == ""){
        attaches = _this.val()
        num = 1
    }else{
        attaches += ","+_this.val()
        num += ",1"
    }
    $('input[name=attaches]').val(attaches)
    $('input[name=attaches_num]').val(num)
})



var form_players_selector = $('#form_players_selector')
$.each(SDplayers, function (i, e) {
    form_players_selector.append('<option value="'+i+'">'+e.name+ '('+i+')'+ '</option>')
})

var formatter_fields = ["diamond", "exper", "money"]

function loadPlayerGrid(players) {

    $.each(players, function (k, e) {
//            console.log(k, e)
        if (SDplayers[e.id] != undefined){
            players[k]['name'] = SDplayers[e.id]['name']
        }else{
            players[k]['name'] = 'unkown'
        }
    })

//        console.log(players)

    $("#datagrid_players").datagrid({
        title: 'Players',
        data: players,
        fitColumns: true,
        striped: true,
        rownumbers: true,
        singleSelect: true,
        idField: 'id',
        pagination: false,
        columns: [
            [
                {
                    field: 'id',
                    title: 'ID',
                    width: 100,
                    align: 'center',
                }, {
                field: 'name',
                title: 'Name',
                width: 100,
                align: 'center',
            },{
                field: 'level',
                title: 'Level',
                width: 30,
                align: 'center',
                editor: 'numberbox'
            }, {
                field: 'experience',
                title: 'Experience',
                width: 40,
                align: 'center',
            }
            ]
        ],
        onAfterEdit: function(index, data, changes) {
            if (vac.isEmpty(changes)) {
                return;
            }
            // changes.id = data.id;
            data.gid = gid
            vac.ajax(URL + '/update_player', data, 'POST', function(r) {
                if (!r.status) {
                    vac.alert(r.info);
                } else {
                    $("#datagrid_players").datagrid("reload");
                }
            })
        },
        onDblClickRow: function(index, row) {
            editPlayerRow();
        },

    });
}
var editPlayerIndex
function editPlayerRow() {
    var index = $("#datagrid_players").datagrid("getSelected")
    if (!index) {
        vac.alert("请选择要编辑的行");
        return;
    }
    var vacindex = vac.getindex("datagrid_players"); //
    if (editPlayerIndex != vacindex) {
        if (editPlayerIndex == undefined) {
            $('#datagrid_players').datagrid('selectRow', vacindex).datagrid('beginEdit', vacindex);
        } else {
            $("#datagrid_players").datagrid("cancelEdit", editIndex);
            // $('#dg').datagrid('selectRow', editIndex);
            $('#datagrid_players').datagrid('beginEdit', vacindex);
        }
        editIndex = vacindex;
    }
    // $('#datagrid').datagrid('beginEdit', vac.getindex("datagrid"));
}

function saveplayer(index) {
    if (!$("#datagrid_players").datagrid("getSelected")) {
        vac.alert("请选择要保存的行");
        return;
    }
    $('#datagrid_players').datagrid('endEdit', vac.getindex("datagrid_players"));
    editIndex = undefined
}


function allplayer() {
    if (!confirm("确定?")){
        return
    }
    $.post(URL+"/set_all_player", {gid: gid}, function (r) {
        if (!r.status) {
            vac.alert(r.info);
        } else {
            $("#datagrid_players").datagrid("reload");
        }
    })
}

function refreshplayer() {
    $("#datagrid_players").datagrid("reload");
}

function loadPlayerGoods(goods) {

    $.each(goods, function (k, e) {
        if (items[e.id] != undefined){
            goods[k]['name'] = items[e.id]['item_name']
        }else{
            goods[k]['name'] = 'unkown'
        }
    })
//        console.log(goods)
    $("#datagrid_goods").datagrid({
        title: 'Pack',
        data: goods,
        fitColumns: true,
        striped: true,
        rownumbers: true,
        singleSelect: true,
        idField: 'id',
        pagination: false,
        columns: [
            [
                {
                    field: 'id',
                    title: 'ID',
                    width: 100,
                    align: 'center',
                }, {
                field: 'name',
                title: 'Name',
                width: 100,
                align: 'center',
            },{
                field: 'number',
                title: 'Number',
                width: 30,
                align: 'center',
            }
            ]
        ],
    });
}

function loadPlayerMail(mails) {

    $.each(mails, function (i, e) {
        if (e.attachments != undefined){
            e.attachments_text = ""
            $.each(e.attachments, function (k, att) {
                if (items[att.id] != undefined){
                    var name = items[att.id]['item_name']
                }else{
                    var name = 'unkown'
                }
                e.attachments_text += name + '('+att.id+') * ' + att.number + "<br/>";
            })
        }
    })

    $("#datagrid_mail").datagrid({
        title: 'Mail',
        data: mails,
        fitColumns: true,
        striped: true,
        rownumbers: true,
        singleSelect: true,
        idField: 'id',
        pagination: false,
        columns: [
            [
                {
                    field: 'id',
                    title: 'ID',
                    width: 100,
                    align: 'center',
                },{
                field: 'theme',
                title: 'Theme',
                width: 100,
                align: 'center',
            },{
                field: 'msg',
                title: 'Msg',
                width: 100,
                align: 'center',
            },{
                field: 'state',
                title: 'State',
                width: 100,
                align: 'center',
            },{
                field: 'attachmentState',
                title: 'AttachmentState',
                width: 100,
                align: 'center',
            },{
                field: 'time',
                title: 'Time',
                width: 100,
                align: 'center',
                formatter: function (timestamp) {
                    return phpjs.date("Y-m-d H:i:s", timestamp)
                }
            },{
                field: 'attachments_text',
                title: '附件信息',
                width: 100,
                align: 'center',
            }
            ]
        ],
    });
}

function searchCallback() {
    $('#form_gamer').show()
    var gamers = data.gamer

    // format player/pack/mail -> map
    data.playerMap = {}
    data.goodsMap  = {}
    if (data.players.length > 0) {
        $.each(data.players, function (i, e) {
            data.playerMap[e.id] = e
        })
    }
    if  (data.goods.length > 0){
        $.each(data.goods, function (i, e) {
            data.goodsMap[e.id] = e
        })
    }

//        console.log(data)


    if (gamers.id == undefined){
        $.messager.alert("alert","没有该玩家信息","alert")
        return
    }
    $('#mail_gid').val(data.gamer.id)
    gid = data.gamer.id
    $.each(formatter_fields, function (i, e) {
        gamers[e] = gamers[e]['number']
    })
    $.each(gamers.timeRecord, function (k, e) {
        gamers['timeRecord'+k] = phpjs.date("Y-m-d H:i:s", e)
    })

    if (gamers.numberRecord != undefined){
        $.each(gamers.numberRecord, function (k, e) {
            gamers['numberRecord'+k] = e
        })
    }

//                console.log(gamers)
    $('#form_gamer').form('load', gamers)

//                var players = data.players
    loadPlayerGrid(data.players)
    loadPlayerGoods(data.goods)
    loadPlayerMail(data.mails)

}

function doSearchByName(value) {
    data = null;
    if (value == undefined || value == null){
        value = gid
    }
    $.post('/gm/gamer/search', {name: value}, function (res) {
        data = res
        if (res != null) {
            searchCallback()
        }
    })
}


function doSearch(value) {
    data = null
    if (value == undefined || value == null){
        value = gid
    }
    $.post('/gm/gamer/search', {id: value}, function (res) {
        data = res
        if (res != null) {
            searchCallback()
        }
    })
}


function addplayer() {
    $('#form_player').form('clear')
    $('#form_player_gid').val(gid)
    $('#dialog_player').dialog('open');
}


function save(){
    $.post(URL+ '/save', $('#form_gamer').serialize(), function(res){
        if (res.status){
            vac.alert(res.info);
        }
    });
}