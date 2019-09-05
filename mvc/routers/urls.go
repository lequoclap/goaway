package routers

type (
	urls struct {
		STATIC_PATH string
		LOGIN_PATH  string
		HOME_PATH   string
		MEDIUM_PATH string
	}
)

func ReturnURLS() urls {
	var url_patterns urls
	url_patterns.STATIC_PATH = "/static/"
	url_patterns.LOGIN_PATH = "/"
	url_patterns.MEDIUM_PATH = "/medium/"
	url_patterns.HOME_PATH = url_patterns.MEDIUM_PATH + "home/"
	return url_patterns
}
