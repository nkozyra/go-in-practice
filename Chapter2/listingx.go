package main

import "log"

var foo = `
201.153.135.148 - - [23/Jan/2024:21:25:42 +0000] "GET /viewforum.php?f=1&sid=bc617614dd55d3f02d96ec8c4be8cc9c HTTP/1.1" 200 15986 "https://forums.hipinion.com/viewtopic.php?f=1&t=175770&sid=f26b85e882a9c9a8553e03ba28987ec2" "Mozilla/5.0 (Android 13; Mobile; rv:121.0) Gecko/121.0 Firefox/121.0"
68.173.115.246 - - [23/Jan/2024:21:25:42 +0000] "GET /download/file.php?avatar=1557_1350242254.gif HTTP/1.1" 200 5951 "https://forums.hipinion.com/viewtopic.php?f=1&t=141575&start=5790" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
3.224.220.101 - - [23/Jan/2024:21:25:42 +0000] "GET /search.php?search_id=active_topics&sid=1063a3644c6148fa0e600685845bd206 HTTP/1.1" 200 3327 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Amazonbot/0.1; +https://developer.amazon.com/support/amazonbot)"
166.199.242.95 - - [23/Jan/2024:21:25:42 +0000] "GET /search.php?search_id=egosearch HTTP/1.1" 200 14244 "-" "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1"
172.59.64.75 - - [23/Jan/2024:21:25:42 +0000] "GET /viewforum.php?f=1 HTTP/1.1" 200 15928 "https://forums.hipinion.com/viewtopic.php?f=1&t=175864&start=90" "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36"
108.56.78.52 - - [23/Jan/2024:21:25:42 +0000] "GET /download/file.php?avatar=33388_1686532875.jpg HTTP/1.1" 200 72565 "https://forums.hipinion.com/viewtopic.php?f=1&t=175856" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
34.220.88.70 - - [23/Jan/2024:21:25:43 +0000] "GET /viewtopic.php?f=1&t=3808&sid=746d3165dfbb40e99b8cbd413e3fe37d&start=8760 HTTP/1.1" 301 193 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/605.1.15 (KHTML, like Gecko; compatible; FriendlyCrawler/1.0) Chrome/120.0.6099.216 Safari/605.1.15"
`

func main() {
	log.Println(foo)
}
