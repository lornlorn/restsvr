/* Simplified Chinese translation for the jQuery Timepicker Addon /
/ Written by Will Lu */
(function($) {
	$.timepicker.regional['zh-CN'] = {
		timeOnlyTitle: '选择时间',
		timeText: '时间',
		hourText: '小时',
		minuteText: '分钟',
		secondText: '秒钟',
		millisecText: '毫秒',
		microsecText: '微秒',
		timezoneText: '时区',
		currentText: '当前时间',
		closeText: '关闭',
		dateFormat: 'yy-mm-dd',
		timeFormat: 'HH:mm',
		timeSuffix: '',
		amNames: ['AM', 'A'],
		pmNames: ['PM', 'P'],
		isRTL: false,
		prevText: "上一月", // Display text for previous month link
		nextText: "下一月", // Display text for next month link
		monthNames: [ "1月","2月","3月","4月","5月","6月","7月","8月","9月","10月","11月","12月" ], // Names of months for drop-down and formatting
		dayNamesMin: [ "日","一","二","三","四","五","六" ], // Column headings for days starting at Sunday
		yearSuffix: "年",
		showMonthAfterYear: true, // True if the year select precedes month, false for month then year
	};
	$.timepicker.setDefaults($.timepicker.regional['zh-CN']);
})(jQuery);
