$(document).ready(function(){

    showAllSimulationMeter();



    //在meter_cell下面添加模拟表记列表
    function showAllSimulationMeter(){
        console.log("运行显示所有表记函数")
        //清空所有节点，只留下最后一个用于添加的
        $("#add_meter").prevAll().remove();
        //获取所有表记信息
        getAllSimulationMeter();

        console.log("显示所有表记函数成功")
    }

    //添加一个格子
    function addOneSimulationMeterCell(id,location,address,ipv4,cycle){
        var addressuse = new Array();
        addressuse[0] = address[0].toString(16);
        addressuse[1] = address[1].toString(16);
        addressuse[2] = address[2].toString(16);
        addressuse[3] = address[3].toString(16);
        addressuse[4] = address[4].toString(16);
        addressuse[5] = address[5].toString(16);
        var addressShow = addressuse[0]+'-'+addressuse[1]+'-'+addressuse[2]+'-'+
            addressuse[3]+'-'+addressuse[4]+'-'+addressuse[5]
        var ipv4use = new Array();
        ipv4use[0] = ipv4[0];
        ipv4use[1] = ipv4[1];
        ipv4use[2] = ipv4[2];
        ipv4use[3] = ipv4[3];
        ipv4use[4] = ipv4[4]*Math.pow(2,8)+ ipv4[5];
        var ipv4Show = ipv4use[0]+"."+ipv4use[1]+"."+ipv4use[2]+"."+
            ipv4use[3]+":"+ipv4use[4];
        var html = ""
        html += '<div class="col-xs-3 ">'+
            '<div class="panel panel-default panel panel-info" >'+
            '<div class="panel-heading" align="center">'+
            '模拟表记'+
            '</div>'+
            '<div class="panel-body meter-body" align="center">'+
            '<form>'+
            '<div class="input-group input-group">'+
            '<input type="hidden" name="id" value="' +
            id+
            '"/>'+
            // '<span class="input-group-addon" id="basic-addon1">地理位置</span>'+
            // '<input type="text" class="form-control" placeholder="地理位置" value="'+
            // location+
            // '" name="location" aria-describedby="basic-addon1">'+
            '</div>'+
            '<div class="input-group input-group">'+
            '<span class="input-group-addon" id="basic-addon2">表记地址</span>'+
            '<input type="text" class="form-control" placeholder="表记地址" value="'+
            addressShow+
            '" name="address" aria-describedby="basic-addon2">'+
            '</div>'+
            '<div class="input-group input-group">'+
            '<span class="input-group-addon" id="basic-addon3">目标IPv4</span>'+
            '<input type="text" class="form-control" placeholder="目标IPv4" value="'+
            ipv4Show+
            '" name="ipv4" aria-describedby="basic-addon3">'+
            '</div>'+
            '<div class="input-group input-group">'+
            '<span class="input-group-addon" id="basic-addon4">时间周期</span>'+
            '<input type="number" class="form-control" value="'+
            cycle+
            '" name="cycle" max="24" min="1" step="1">'+
            '</div>'+
            '</form>'+
            '</div>'+
            '<div class="panel-footer" align="right">'+
            '<button  class="btn btn-primary">修改</button> '+
            '<button  class="btn btn-danger">删除</button>'+
            '</div>'+
            '</div>'+
            '</div>';
        $("#add_meter").before(html);


    }

    //从服务器获取模拟表记列表
    function getAllSimulationMeter(){
        $.ajax({
            type:"get",
            url:"/meter",
            async:true,
            beforeSend:function(xhr){
                console.log("发送前")
            },
            success:function(data,textStatus,jqXHR){
                console.log(data)
                //把其中的表记信息数组提取出来
                var obj = JSON.parse(data);
                var dataList = obj.Data
                //遍历数组
                for (var i = 0; i < dataList.length; i++) {
                    var id = dataList[i].id
                    var location = dataList[i].location;
                    var address = dataList[i].address;
                    var ipv4 = dataList[i].ipv4;
                    var cycle = dataList[i].cycle;
                    addOneSimulationMeterCell(id,location,address,ipv4,cycle);
                    // addOneCell()
                }
                var meterList = $("#add_meter").prevAll();
                for (var i=0;i<meterList.length;i++){
                    $(meterList[i]).find("button:even").click(function () {
                        var newId = $(this).parents(".col-xs-3").find('input[name="id"]').val();
                        var newLocation = $(this).parents(".col-xs-3").find('input[name="location"]').val();
                        var newAddress = $(this).parents(".col-xs-3").find('input[name="address"]').val();
                        var newIpv4 = $(this).parents(".col-xs-3").find('input[name="ipv4"]').val();
                        var newCycle = $(this).parents(".col-xs-3").find('input[name="cycle"]').val();
                        // alert(newId+newLocation+newAddress+newIpv4+newCycle);
                        click_update(newId,newLocation,newAddress,newIpv4,newCycle);
                    });
                    $(meterList[i]).find("button:odd").click(function () {
                        var newId = $(this).parents(".col-xs-3").find('input[name="id"]').val();
                        click_delete(newId);
                    });
                }



                // var meterList = $("#add_meter").prevAll();      //所有“表记格子”
                // $("#add_meter").prevAll().find("button:even").click(function () {
                //
                //     click_update(11);
                // });
                // $("#add_meter").prevAll().find("button:odd").click(function () {
                //     click_delete(10);
                // });

                        // meterList
                // meterList.find("button").click(function () {
                //     alert("点击时间")
                // });
                // console.log(textStatus)
                // console.log(jqXHR)
            },
            error:function(xhr,textStatus){
                console.log('错误')
                alert("服务器获取数据失败")
                // console.log(xhr)
                // console.log(textStatus)
            },
            complete:function(){
                console.log('结束')
            }
        });
    }

    $("#btn_create").click(function () {
        $("#form_create").attr("method","post");
        $("#form_create").attr("action","/meter");
        // 创建Input
        var my_input = $('<input type="text" name="type" />');
        my_input.attr('value', "create");
        // 附加到Form
        $("#form_create").append(my_input);
        $("#form_create").submit();
    });

//id,location,address,ipv4,cycle
    function click_update(id,location,address,ipv4,cycle) {
        $.ajax({
            url:'/meter',
            type:'POST', //GET
            async:true,    //或false,是否异步
            data:{
                id:id,
                type:'update',
                location:location,
                address:address,
                ipv4:ipv4,
                cycle:cycle
            },
            timeout:5000,    //超时时间
            dataType:'html',    //返回的数据格式：json/xml/html/script/jsonp/text
            beforeSend:function(xhr){
                console.log(xhr)
                console.log('发送前')
            },
            success:function(data,textStatus,jqXHR){
                console.log(data)
                // console.log(textStatus)
                // console.log(jqXHR)
            },
            error:function(xhr,textStatus){
                console.log('错误')
                console.log(xhr)
                console.log(textStatus)
            },
            complete:function(){
                console.log('结束')
            }
        })
    }
//id
    function click_delete(id) {
        $.ajax({
            url:'/meter',
            type:'POST', //GET
            async:true,    //或false,是否异步
            data:{
                id:id,
                type:'delete'
            },
            timeout:5000,    //超时时间
            dataType:'html',    //返回的数据格式：json/xml/html/script/jsonp/text
            beforeSend:function(xhr){
                console.log(xhr)
                console.log('发送前')
            },
            success:function(data,textStatus,jqXHR){
                console.log(data)
                showAllSimulationMeter();
            },
            error:function(xhr,textStatus){
                console.log('错误')
                console.log(xhr)
                console.log(textStatus)
            },
            complete:function(){
                console.log('结束')
            }
        })
    }


});

