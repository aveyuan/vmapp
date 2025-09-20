package middleware

// 静态文件不需要中间件的
var ExtViewData []string = []string{"/api/", "/static/", "/uploads/", "/sitemap"}
var ExtFront []string = []string{"/static/", "/uploads/"}
var Igf = "igf"