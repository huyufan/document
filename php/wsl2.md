# window10 安装wsl2
- [网站地址:][https://learn.microsoft.com/zh-cn/windows/wsl/install]

## 找wsl2  打开cmd   执行wsl  
- ip addr | grep inet （inet 172.20.160.1/20 ...） 获取ip
- netsh interface portproxy add v4tov4 listenaddress=0.0.0.0 listenport=8080 connectaddress=<WSL_IP> connectport=8080


## 代理
- ip 地址 通过 ipconfig  得到vEthernet(WSL)  下面的ipv4地址
- vim ~ .bashrc
- export http_proxy=http://172.27.128.1:7890
- export https_proxy=http://172.27.128.1:7890
- export all_proxy=socks5://172.27.128.1:7890
- source .bashrc

## 是否安装systemd
- ps -p 1 -o comm=
- sudo apt install systemd  (如果没有 sudo apt update)



