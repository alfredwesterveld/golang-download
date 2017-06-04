# golang-download
P.S: Learning expirements but could be useful anyway

Downloads list of urls from stdin line by line to disc

```bash
$ cat urls.txt
https://d325nkq7ptgsx2.cloudfront.net/5c5a-5095-2018-2800-1469641555844.jpg
https://d325nkq7ptgsx2.cloudfront.net/8ed7-5095-2800-1996-1469641634984.jpg

$ cat urls.txt | go run main.go; # pipe Stdin to program
```
