package config

// 报表的筛选字段信息

var FilterFiledLst = []map[string]interface{}{
	{"label": "项目", "value": "project"},
	{"label": "日期", "value": "date"},
	{"label": "上线内容", "value": "online_content"},
	{"label": "TAPD链接", "value": "tapd_link"},
	{"label": "文件类型", "value": "file_type"},
	{"label": "上线发起方", "value": "initiator"},
	{"label": "上线类型", "value": "online_type"},
	{"label": "上线原因", "value": "reason"},
	{"label": "根因分析", "value": "detail"},
	{"label": "上线方式", "value": "way"},
	{"label": "maintain时长", "value": "mint"},
}

// 报表字段的可选筛选条件

var FilterLst = []map[string]interface{}{
	{"label": "等于", "value": "=", "status": 0},
	{"label": "不等于", "value": "=", "status": 1},
	{"label": "包含（Like）", "value": "contains", "status": 0},
	{"label": "不包含（Not Like）", "value": "contains", "status": 1},
	{"label": "属于（In）", "value": "in", "status": 0},
	{"label": "不属于（Not In）", "value": "in", "status": 1},
	{"label": "为空", "value": nil, "status": 0},
	{"label": "不为空", "value": nil, "status": 1},
}
