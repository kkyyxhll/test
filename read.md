gin + gorm后端开发:
config目录:放置配置文件信息
yaml文件作用:存放配置信息

创建gin+gorm新项目过程:
🕐go mod init
    得到一个go.mod
🕑go get -u github.com/gin-gonic/gin
    下载gin,得到一个go.sum
🕒go get github.com/spf13/viper
    