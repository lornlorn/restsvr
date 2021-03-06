var $module = $('#module');
var $submit = $('#submit');

$(function () {
    $submit.click(function () {
        $('#status').text('');
        $('#result').text('');

        var params = {};
        params['module'] = $module.val();
        // params['data'] = [];
        params['data'] = {};
        $('#json').find('input[name]').each(function () {
            var k = $(this).attr('name');
            var v = $(this).val();
            // var param = {};
            // param[k] = v;
            // params['data'].push(param);
            params['data'][k] = v;
        });
        console.log('REQUEST : ' + JSON.stringify(params));

        $.ajax({
            url: '/ajax/test',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (result) {
                console.log('RESPONSE : ' + JSON.stringify(result));
                $('#status').text('请求成功');
                $('#result').text(result['retcode'] + '|' + result['retmsg']);
            },
            error: function (result) {
                $('#status').text('请求失败');
            },
            complete: function () {
                console.log("Ajax finish");
            },
        });
    });

    $('#editable-select').editableSelect({
        filter: true,
        effects: 'fade',
        duration: 200,
    });

    $('#sub-select').editableSelect({
        filter: false,
        effects: 'fade',
        duration: 200,
    });

    $('#editable-select').on('select.editable-select', function (element) {
        // do something...
        console.log($(this).val());
        var value = $(this).val();
        $('#sub-select').editableSelect('clear');
        $('#sub-select').val("");
        $('#sub-select').editableSelect('add', value);
    });

});

$("#autocomplete").autocomplete({
    source: function (request, response) {
        var params = {};
        params['module'] = $module.val();
        params['data'] = {};
        params['data']["keyword"] = $("#autocomplete").val();
        $.ajax({
            url: '/ajax/autocomplete/systemlist',
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(params),
            async: 'true',
            dataType: 'json',
            success: function (data) {
                // console.log(data);
                response($.map(data, function (item) {
                    return {
                        id: item.sysid,
                        label: item.sysenname,
                        value: item.sysenname + "-" + item.syscnname,
                    };
                }));
            },
        });
    },
    minLength: 2,
    select: function (event, ui) {
        // console.log(ui.item.id);
        $("#autocomplete").attr("data-id",ui.item.id);
        // console.log($("#autocomplete").attr("data-id"));
    },
});
