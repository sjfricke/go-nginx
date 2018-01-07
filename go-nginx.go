package main

import (
	// "io/ioutil"
	"flag"
	"fmt"
	// "os"
)

var (
	path     = flag.String("path", "/var/www", "path to where to save directory")
	domain   = flag.String("domain", "", "domain of site")
	suffix   = flag.String("suffix", "html", "folder to append inside your directory for files")
	input    = flag.String("input", "", "input files with line-by-line list of flags for batch run")
	cDefault = flag.Bool("default", "false", "creates a default file instead")

	available string = "/etc/nginx/sites-available"
	enabled   string = "/etc/nginx/sites-enabled"
	fullPath  string
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Default returns configuration for default site
func Default() string {
	return fmt.Sprintf(`
server {
    listen 80 default_server;
    listen [::]:80 default_server;

    root %s;
    index index.html;

    server_name _;

    location / {
            try_files $uri $uri/ =404;
    }
}`, fullPath)
}

// Config returns configuration for normal site
func Config() string {
	return fmt.Sprintf(`
# For HTTPS Support
#server {
#       listen 80;
#       listen [::]:80;
#       server_name %s www.%s;
#       return 301 https://$server_name$request_uri;
#}

server {

        listen 80;
        listen [::]:80;

#       listen 443 ssl http2;
#       listen [::]:443 ssl http2;

#       ssl_certificate /etc/letsencrypt/live/%s/fullchain.pem;
#       ssl_certificate_key /etc/letsencrypt/live/%s/privkey.pem;
#       include snippets/ssl-%s.conf;
#       include snippets/ssl-params.conf;

        root %s;
        index index.html;

        server_name %s www.%s;

        location / {
                try_files $uri $uri/ =404;
        }

#        location ~ /.well-known {
#                allow all;
#        }
}`, *domain, *domain, *domain, *domain, *domain, fullPath, *domain, *domain)
}

func main() {

	flag.Parse()

	if *domain == "" || *path == "" {
		panic("Domain or path is nil")
	}

	fullPath = fmt.Sprintf("%s/%s/%s", *path, *domain, *suffix)

	if *cDefault {
		fmt.Println(Default())
	} else {
		fmt.Println(Config())
	}
}
