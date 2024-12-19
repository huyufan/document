# 查看所有连接到 Nginx 的连接数 (总数)
- netstat -ant | grep ':80' 
- netstat -ant | grep ':80' | wc -l
- netstat -ant | grep ':80' | awk '{print $6}' | sort | uniq -c