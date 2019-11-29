package e

var MsgFlags = map[int]string {
    SUCCESS:                    "success",
    ERROR:                      "error",
    INVALID_PARAMS:             "请求参数错误",
    ERROR_EXIST_SERVICES:       "服务已存在",
    ERROR_EXIST_SERVICES_FAIL:  "获取服务信息失败",
    ERROR_NOT_EXIST_SERVICES:   "服务不存在",
    ERROR_GET_SERVICES_FAIL:    "获取所有服务失败",
    ERROR_COUNT_SERVICES_FAIL:  "统计服务失败",
    ERROR_ADD_SERVICES_FAIL:    "新增服务失败",
    ERROR_EDIT_SERVICES_FAIL:   "修改服务失败",
    ERROR_DELETE_SERVICES_FAIL: "删除服务失败",

}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}