var $module = $('#module');
var $jobinfo = $('#jobinfo');

$(function () {

    $("#joblist").on("click", ".stepdtl-add", function () {
        var seq = $(this).parent().parent().attr("data-seq"); // $(this)表示获取当前被点击元素的name值
        // alert(value);
        var html =
            "<div class=\"stepdtl\">" +
            "    <input data-type=\"check\" type=\"checkbox\">" +
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

    $("#joblist").on("click", ".stepdtl-del", function () {
        var seq = $(this).parent().parent().attr("data-seq");
        // alert(selectstep);
        var chkcount = $("div[data-seq='"+seq+"'] input[data-type='check']").length;
        var delcount = $("div[data-seq='"+seq+"'] input[data-type='check']:checkbox:checked").length;
        // var chkcount = $("input[data-type='check']").length;
        // var delcount = $("input[data-type='check']:checkbox:checked").length;
        if (delcount >= chkcount) {
            alert("至少保留一条指令");
        } else {
            $("div[data-seq='"+seq+"'] input[data-type='check']:checkbox:checked").each(function (index) {
                // alert(index);
                $(this).parent().remove();
                // $(this).attr("data-number", index);
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

    $("button.step-add").click(function () {
        var num = $("div.step").length; // $(this)表示获取当前被点击元素的name值
        // alert(num);
        var nextnum = num + 1;
        var html =
            "<div class=\"step\" data-seq=\"" + nextnum + "\">" +
            "    <p>----------STEP " + nextnum + "----------</p>" +
            "    <p>" +
            "        <button class=\"stepdtl-add\">添加指令</button>" +
            "        <button class=\"stepdtl-del\">删除指令</button>" +
            "    </p>" +
            "    <div class=\"stepdtl\">" +
            "        <input data-type=\"check\" type=\"checkbox\">" +
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

});

$("div[data-number]").each(function (index) {
    alert(index);
    $(this).attr("data-number", index);
});