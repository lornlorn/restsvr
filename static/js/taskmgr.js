var $module = $('#module');
var $submit = $('#submit');

$(function () {

    // 添加指令
    $("#joblist").on("click", ".stepdtl-add", function () {
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
    $("#joblist").on("click", ".stepdtl-del", function () {
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

    // $("button.stepdtl-add").click(function () {
    //     var seq = $(this).parent().parent().attr("data-seq"); // $(this)表示获取当前被点击元素的name值
    //     // alert(value);
    //     var html =
    //         "<div class=\"stepdtl\">" +
    //         "    <input data-type=\"check\" type=\"checkbox\">" +
    //         "    <label>IP:</label>" +
    //         "    <input data-type=\"data\" data-name=\"ip\" value=\"192.168.100.101\">" +
    //         "    <label>端口:</label>" +
    //         "    <input data-type=\"data\" data-name=\"port\" value=\"22\">" +
    //         "    <label>用户名:</label>" +
    //         "    <input data-type=\"data\" data-name=\"username\" value=\"test\">" +
    //         "    <label>密码:</label>" +
    //         "    <input data-type=\"data\" data-name=\"password\" value=\"test\">" +
    //         "    <label>命令:</label>" +
    //         "    <input data-type=\"data\" data-name=\"command\" value=\"ls -lrt\">" +
    //         "</div>";

    //     $("div.step[data-seq=" + seq + "]").append(html);
    // });

    // $("button.stepdtl-del").click(function () {
    //     var chkcount = $("input[data-type='check']").length;
    //     var delcount = $("input[data-type='check']:checkbox:checked").length;
    //     if (delcount >= chkcount) {
    //         alert("至少保留一条指令");
    //     } else {
    //         $("input[data-type='check']:checkbox:checked").each(function (index) {
    //             // alert(index);
    //             $(this).parent().remove();
    //             // $(this).attr("data-number", index);
    //         });
    //     };
    // });

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

        $("div#joblist").append(html);
    });

    // 删除步骤 置步骤为失效
    $("button.step-del").click(function () {
        var selects = $("div#joblist input[data-name=\"step-check\"]:checkbox:checked");
        selects.each(function () {
            var step = $(this).parent().parent();
            var title = step.children("p[data-name=\"step-title\"]");
            var title_content = title.text();
            title.text(title_content + "已删除");
            step.attr("data-valid", "false");
        });

    });

    // "立即执行"选中后置"执行时间表达式"为不可用
    $("#ifnow").click(function () {
        if ($("#ifnow").is(":checked")) {
            $("#runtime").attr("disabled", true);
        } else {
            $("#runtime").attr("disabled", false);
        };
    });

    // 提交
    $submit.click(function () {
        var params = {};
        params['module'] = $module.val();
        params['data'] = {};
        $("#joblist .step").each(function (index) {
            var step = $(this);
            if (step.attr("data-valid") == "false") {
                // alert(index);
            } else if (step.attr("data-valid") == "true") {
                var valid_step = $(this);
                var seq = valid_step.attr("data-seq");
                // console.log(valid_step.children(".stepdtl"));
                valid_step.children(".stepdtl").each(function () {
                    console.log($(this));
                });
                // params['data'][seq] = v;
            };
            // var k = $(this).attr('name');
            // var v = $(this).val();
            // params['data'][k] = v;
        });
        console.log('REQUEST : ' + JSON.stringify(params));

        // $.ajax({
        //     url: '/ajax',
        //     type: 'POST',
        //     contentType: "application/json; charset=utf-8",
        //     data: JSON.stringify(params),
        //     async: 'true',
        //     dataType: 'json',
        //     success: function (result) {
        //         console.log('RESPONSE : ' + JSON.stringify(result));
        //         $('#status').text('请求成功');
        //         $('#result').text(result['retcode'] + '|' + result['retmsg']);
        //     },
        //     error: function (result) {
        //         $('#status').text('请求失败');
        //     },
        //     complete: function () {
        //         console.log("Ajax finish");
        //     },
        // });
    });

});
