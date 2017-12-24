var $module = $('#module');
var $jobinfo = $('#jobinfo');

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
            "<div class=\"step\" data-seq=\"" + nextnum + "\">" +
            "    <p>----------STEP " + nextnum + "----------</p>" +
            "    <p>" +
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

    // 删除步骤
    $("button.step-del").click(function () {

    });

    // "立即执行"选中后置"执行时间表达式"为不可用
    $("#ifnow").click(function () {
        if ($("#ifnow").is(":checked")) {
            $("#runtime").attr("disabled", true);
        } else {
            $("#runtime").attr("disabled", false);
        };
    });

});
