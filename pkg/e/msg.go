package e

var MsgFlags = map[int]string{
    DEFAULT: "未知错误",
    SUCCESS:                         "ok",
    ERROR:                           "fail",
    INVALID_PARAMS:                  "请求参数错误",
    ERROR_DB:                        "数据库查询失败",

    DATA_NOT_EXIST: "数据不存在",
    // Meets

    ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
    ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
    ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
    ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
    ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
    ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
    ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
    ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
    ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",
    ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
    ERROR_AUTH_TOKEN:                "Token生成失败",
    ERROR_AUTH:                      "Token错误",
    ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
    ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
    ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func Msg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}
