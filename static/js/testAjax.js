var $module = $('#module');
var $username = $('#username');
var $password = $('#password');
var $submit = $('#submit');

$(function () {
    $submit.click(function () {
        $('#status').text('');
        $('#result').text('');

        var params = {};
        params['module'] = $module.text();
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
        filter: false,
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
