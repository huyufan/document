# 安装java
- yum install java-17-openjdk 
## 可以使用alternatives --config java 切换java版本


# 添加 Jenkins 仓库
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat/jenkins.repo

# 导入 GPG 密钥
sudo rpm --import https://pkg.jenkins.io/redhat/jenkins.io.key

# 安装 Jenkins
sudo yum install -y jenkins


# 查看状态 systemctl status jenkins

# 使用 --nogpgcheck 选项跳过 GPG 检查  yum install --nogpgcheck -y jenkins


# gitlab
## 添加GitLab社区版Package
- curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash

## 安装GitLab社区版
- yum install -y gitlab-ce

## 配置GitLab站点Url
- 默认的站点Url配置项是：external_url 'http://gitlab.example.com'

## 启动gitlab gitlab-ctl reconfigure

## 关闭 gitlab-ctl stop

## 密码在 /etc/gitlab/initial_root_password
