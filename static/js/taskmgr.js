var $module = $('#module');
var $jobinfo = $('#jobinfo');

$(function () {
    $("button.stepdtl-add").click(function () {
        var value = $(this).attr("id"); // $(this)表示获取当前被点击元素的name值
        alert(value);
    });
});

$("div[data-number]").each(function(index) {
    alert(index);
    $(this).attr("data-number", index);
});