# go-nginx
Personal script to setup nginx for a new domain. It goes and create the folder to hold the website, the nginx config file, and links it together. Also batch support makes migrating servers easy

## Download

- `go get github.com/sjfricke/go-nginx`
	- Make sure `$GOPATH/bin` is in your path

## Usage

- **NEED TO RUN AS ROOT** - `sudo go-nginx <args>`
- **NEED TO RESTART NGINX YOURSELF** - `sudo service nginx restart`
	- Can check if configurations all worked with `sudo nginx -t`

```
  -default
        creates a default file instead
  -domain string
        domain of site
  -input string
        input files with line-by-line list of flags for batch run
  -path string
        path to where to save directory (default "/var/www")
  -suffix string
        folder to append inside your directory for files
```

### Examples

- `sudo go-nginx -path /var/www -domain example.com -suffix html`
- `sudo go-nginx -path /var/www -domain example.com -suffix html -default true`
- `sudo go-nginx -input inputExample.txt`
	- [input file](inputExample.txt) is just lines of arguments to run in batch