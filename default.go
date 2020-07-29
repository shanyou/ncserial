package ncserial

import (
	"fmt"
	"path/filepath"
)

var (
	defaultProxyConfig = Options{
		{"proxy_redirect", "off"},
		{"proxy_set_header", "Host $host"},
		{"proxy_set_header", "X-Real-IP $remote_addr"},
		{"proxy_set_header", "X-Forwarded-For $proxy_add_x_forwarded_for"},
		{"client_max_body_size", "10m"},
		{"client_body_buffer_size", "128k"},
		{"proxy_connect_timeout", "90"},
		{"proxy_send_timeout", "90"},
		{"proxy_read_timeout", "90"},
		{"proxy_buffers", "32 4k"},
	}

	defaultFastcgiConfig = Options{
		{"fastcgi_param", "SCRIPT_FILENAME $document_root$fastcgi_script_name"},
		{"fastcgi_param", "QUERY_STRING $query_string"},
		{"fastcgi_param", "REQUEST_METHOD $request_method"},
		{"fastcgi_param", "CONTENT_TYPE $content_type"},
		{"fastcgi_param", "CONTENT_LENGTH $content_length"},
		{"fastcgi_param", "SCRIPT_NAME $fastcgi_script_name"},
		{"fastcgi_param", "REQUEST_URI $request_uri"},
		{"fastcgi_param", "DOCUMENT_URI $document_uri"},
		{"fastcgi_param", "DOCUMENT_ROOT $document_root"},
		{"fastcgi_param", "SERVER_PROTOCOL $server_protocol"},
		{"fastcgi_param", "GATEWAY_INTERFACE CGI/1.1"},
		{"fastcgi_param", "SERVER_SOFTWARE nginx/$nginx_version"},
		{"fastcgi_param", "REMOTE_ADDR $remote_addr"},
		{"fastcgi_param", "REMOTE_PORT $remote_port"},
		{"fastcgi_param", "SERVER_ADDR $server_addr"},
		{"fastcgi_param", "SERVER_PORT $server_port"},
		{"fastcgi_param", "SERVER_NAME $server_name"},
		{"fastcgi_index", "index.php"},
		{"fastcgi_param", "REDIRECT_STATUS 200"},
	}

	defaultMimeTypes = Options{
		{"text/html", "html htm shtml"},
		{"text/css", "css"},
		{"text/xml", "xml rss"},
		{"image/gif", "gif"},
		{"image/jpeg", "jpeg jpg"},
		{"application/x-javascript", "js"},
		{"text/plain", "txt"},
		{"text/x-component", "htc"},
		{"text/mathml", "mml"},
		{"image/png", "png"},
		{"image/x-icon", "ico"},
		{"image/x-jng", "jng"},
		{"image/vnd.wap.wbmp", "wbmp"},
		{"application/java-archive", "jar war ear"},
		{"application/mac-binhex40", "hqx"},
		{"application/pdf", "pdf"},
		{"application/x-cocoa", "cco"},
		{"application/x-java-archive-diff", "jardiff"},
		{"application/x-java-jnlp-file", "jnlp"},
		{"application/x-makeself", "run"},
		{"application/x-perl", "pl pm"},
		{"application/x-pilot", "prc pdb"},
		{"application/x-rar-compressed", "rar"},
		{"application/x-redhat-package-manager", "rpm"},
		{"application/x-sea", "sea"},
		{"application/x-shockwave-flash", "swf"},
		{"application/x-stuffit", "sit"},
		{"application/x-tcl", "tcl tk"},
		{"application/x-x509-ca-cert", "der pem crt"},
		{"application/x-xpinstall", "xpi"},
		{"application/zip", "zip"},
		{"application/octet-stream", "deb"},
		{"application/octet-stream", "bin exe dll"},
		{"application/octet-stream", "dmg"},
		{"application/octet-stream", "eot"},
		{"application/octet-stream", "iso img"},
		{"application/octet-stream", "msi msp msm"},
		{"audio/mpeg", "mp3"},
		{"audio/x-realaudio", "ra"},
		{"video/mpeg", "mpeg mpg"},
		{"video/quicktime", "mov"},
		{"video/x-flv", "flv"},
		{"video/x-msvideo", "avi"},
		{"video/x-ms-wmv", "wmv"},
		{"video/x-ms-asf", "asx asf"},
		{"video/x-mng", "mng"},
	}

	defaultEvents = Events{
		WorkerConnections: 204800,
		Use:               "epoll",
		MultiAccept:       true,
	}
)

// NewDefaultHTTPConfig create default http Config
func NewDefaultHTTPConfig() HTTP {
	logformat := "main '$remote_addr - $remote_user [$time_local]  $status " +
		`"$request" $body_bytes_sent "$http_referer" ` +
		`"$http_user_agent" "$http_x_forwarded_for"'`
	http := HTTP{
		DefalutType: "application/octet-stream",
		LogFormat:   logformat,
		MimeTypes:   defaultMimeTypes,
		SendFile:    true,
		ExtConfig:   MergeOptions(defaultProxyConfig, defaultFastcgiConfig),
	}

	return http
}

//NewDefaultRestyConfig create default resty config
func NewDefaultRestyConfig(prefix, logPath, lualibPath string) *Config {
	config := &Config{}
	config.WorkerProcesses = "auto"
	config.LimitNofile = 204800
	config.Events = defaultEvents
	errorLogPath := filepath.Join(logPath, "error.log")
	config.ErrorLog = errorLogPath + " warn"
	config.PId = filepath.Join(logPath, "nginx.pid")
	config.Extras = Options{
		{"env", "RESTY_VERSION"},
		{"env", "RESTY_PREFIX"},
	}
	config.HTTP = NewDefaultHTTPConfig()
	config.HTTP.AccessLog = filepath.Join(logPath, "access.log")
	config.HTTP.ExtConfig = Options{
		{"keepalive_timeout", "60"},
		{"lua_socket_connect_timeout", "3s"},
		{"lua_socket_send_timeout", "3s"},
		{"lua_socket_read_timeout", "3s"},
	}

	if len(lualibPath) > 0 {
		config.HTTP.ExtConfig = append(config.HTTP.ExtConfig,
			[]string{"lua_package_path", fmt.Sprintf("'%s/?.lua;%s/?/init.lua;;'", lualibPath, lualibPath)})
		config.HTTP.ExtConfig = append(config.HTTP.ExtConfig,
			[]string{"lua_package_cpath", fmt.Sprintf("'%s/?.so;;'", lualibPath)})
	}
	return config
}
