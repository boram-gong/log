# log
一个简单的日志框架, 包含一下功能:
1. 输出log文件, 可自主创建文件夹
2. 支持日志分割(按天\按小时\按分钟)
3. 支持级别打印
4. 支持级别过滤

初始化如下:

`InitLog(日志路径, 日志文件名, 日志分割格式)`

日志分割格式参数如下:

 `	COMMON_FORMAT 通用格式(不分割); DATE_FORMAT 按天分割; HOUR_FORMAT 按小时分割; MINUTE_FROMAT 按分钟分割`
 
 举例:
 
 `import "github.com/boram-gong/log"`
 
 `log.InitLog("/var/log", "mylog", log.DATE_FORMAT)`
 
 打印日志:
 
 `log.ERROR("data")`



