var $module = $('#module');
var $submit = $('#submit');

$(function () {

    // 添加指令
    $("#steplist").on("click", ".stepdtl-add", function () {
        var seq = $(this).parent().parent().attr("data-seq");
        var html =
            "<div class=\"stepdtl\">" +
            "    <input data-type=\"check\" data-name=\"stepdtl-check\" type=\"checkbox\">" +
            "    <label>IP:</label>" +
            "    <input data-type=\"data\" data-name=\"ip\" value=\"192.168.100.101\">" +
            "    <label>端口:</label>" +
            "    <input data-type=\"data\" data-name=\"port\" value=\"22\">" +
            "    <label>用户名:</label>" +
            "    <input data-type=\"data\" data-name=\"username\" value=\"test\">" +
            "    <label>密码:</label>" +
            "    <input data-type=\"data\" data-name=\"password\" value=\"test\">" +
            "    <label>命令:</label>" +
            "    <input data-type=\"data\" data-name=\"command\" value=\"ls -lrt\">" +
            "</div>";

        $("div.step[data-seq=" + seq + "]").append(html);
    });

    // 删除指令
    $("#steplist").on("click", ".stepdtl-del", function () {
        var seq = $(this).parent().parent().attr("data-seq");
        var chkcount = $("div[data-seq='" + seq + "'] input[data-name='stepdtl-check']").length;
        var delcount = $("div[data-seq='" + seq + "'] input[data-name='stepdtl-check']:checkbox:checked").length;
        if (delcount >= chkcount) {
            alert("至少保留一条指令");
        } else {
            $("div[data-seq='" + seq + "'] input[data-name='stepdtl-check']:checkbox:checked").each(function (index) {
                $(this).parent().remove();
            });
        };
    });

    // 添加步骤
    $("button.step-add").click(function () {
        var num = $("div.step").length;
        var nextnum = num + 1;
        var html =
            "<div class=\"step\" data-seq=\"" + nextnum + "\" data-valid=\"true\">" +
            "    <p data-name=\"step-title\">----------STEP " + nextnum + "----------</p>" +
            "    <p data-name=\"step-operation\">" +
            "        <input data-type=\"check\" data-name=\"step-check\" type=\"checkbox\">" +
            "        <button class=\"stepdtl-add\">添加指令</button>" +
            "        <button class=\"stepdtl-del\">删除指令</button>" +
            "    </p>" +
            "    <div class=\"stepdtl\">" +
            "        <input data-type=\"check\" data-name=\"stepdtl-check\" type=\"checkbox\">" +
            "        <label>IP:</label>" +
            "        <input data-type=\"data\" data-name=\"ip\" value=\"192.168.100.101\">" +
            "        <label>端口:</label>" +
            "        <input data-type=\"data\" data-name=\"port\" value=\"22\">" +
            "        <label>用户名:</label>" +
            "        <input data-type=\"data\" data-name=\"username\" value=\"test\">" +
            "        <label>密码:</label>" +
            "        <input data-type=\"data\" data-name=\"password\" value=\"test\">" +
            "        <label>命令:</label>" +
            "        <input data-type=\"data\" data-name=\"command\" value=\"ls -lrt\">" +
            "    </div>" +
            "</div>";

        $("div#steplist").append(html);
    });

    // 删除步骤 置步骤为失效
    $("button.step-del").click(function () {
        var selects = $("div#steplist input[data-name=\"step-check\"]:checkbox:checked");
        selects.each(function () {
            var step = $(this).parent().parent();
            var title = step.children("p[data-name=\"step-title\"]");
            var title_content = title.text();
            title.text(title_content + "已删除");
            step.attr("data-valid", "false");
        });

    });

    // "立即执行"选中后置"执行时间表达式"为不可用
    // $("#iftool").click(function () {
    //     if ($("#iftool").is(":checked")) {
    //         $("#runtime").attr("disabled", true);
    //     } else {
    //         $("#runtime").attr("disabled", false);
    //     };
    // });


    // 根据不同任务类型改变执行时间输入框
    $('#tasktype').change(function () {
        $("#runtime").val("");
        var tasktype = $(this).val();
        switch (tasktype) {
            case "工具":
                console.log("gongju");
                $("#cron").attr("disabled", true);
                $("#runtime").attr("disabled", true);
                $("#cron").attr("type", "hidden");
                $("#runtime").attr("type", "hidden");
                $("#label_runtime").text("");
                break;
            case "定时任务":
                console.log("dsrw");
                $("#cron").attr("disabled", false);
                $("#runtime").attr("disabled", true);
                $("#cron").attr("type", "text");
                $("#runtime").attr("type", "hidden");
                $("#label_runtime").text("执行时间:");
                break;
            case "计划任务":
                console.log("jhrw");
                $("#cron").attr("disabled", true);
                $("#runtime").attr("disabled", false);
                $("#cron").attr("type", "hidden");
                $("#runtime").attr("type", "text");
                $("#label_runtime").text("执行时间:");
                $('#runtime').datetimepicker({
                    controlType: 'select',
                    oneLine: true,
                });
                break;
            default:
                console.log("default");
                alert("页面错误,请联系管理员解决...");
        };
    });

    // 提交
    $submit.click(function () {
        var params = {};
        params['module'] = $module.val();
        params['user'] = "test";
        params['data'] = {};
        params['data']['jobinfo'] = {};
        params['data']['jobinfo']['system'] = $("#system").val();
        params['data']['jobinfo']['name'] = $("#name").val();
        var iftool = "";
        if ($("#iftool").is(":checked")) {
            iftool = "true";
        } else {
            iftool = "false";
        };
        params['data']['jobinfo']['iftool'] = iftool;
        params['data']['jobinfo']['runtime'] = $("#runtime").val();
        params['data']['steplist'] = {};
        $("#steplist .step").each(function () {
            var step = $(this);
            if (step.attr("data-valid") == "false") {
                // 跳过删除的步骤
            } else if (step.attr("data-valid") == "true") {
                var valid_step = $(this);
                var seq = valid_step.attr("data-seq");
                params['data']['steplist'][seq] = {};
                valid_step.children(".stepdtl").each(function (index) {
                    var stepdtl_seq = index + 1;
                    params['data']['steplist'][seq][stepdtl_seq] = {};
                    var stepdtl = $(this);
                    stepdtl.find("input[data-type=\"data\"]").each(function () {
                        var cmd = $(this);
                        var key = cmd.attr("data-name");
                        var value = cmd.val();
                        params['data']['steplist'][seq][stepdtl_seq][key] = value;
                    });
                });

            };
            // var k = $(this).attr('name');
            // var v = $(this).val();
            // params['data'][k] = v;
        });
        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/taskadd',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                // $('#status').text('请求成功');
                // $('#result').text(result['retcode'] + '|' + result['retmsg']);
                console.log("请求成功");
            },
            error: function (result) {
                // $('#status').text('请求失败');
                console.log("请求失败");
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

});

$(function () {
    $(document).tooltip();
});