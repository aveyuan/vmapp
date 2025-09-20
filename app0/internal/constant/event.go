package constant

const (
	EventLogCtx = "vevent_log_ctx"
)

type ArticleStatus int8

const (
	ArticleStatusAll   ArticleStatus = -100
	ArticleStatusCheck ArticleStatus = -1
	ArticleStatusDarft ArticleStatus = 0
	ArticleStatusOk    ArticleStatus = 1
	ArticleStatusFail  ArticleStatus = 2
)
