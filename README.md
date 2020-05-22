# log
自己写的一个简单的日志框架, 包含以下功能:
1. 输出log文件, 可自主创建文件夹
2. 支持日志分割(按月\按天\按小时)
3. 支持级别打印
4. 支持级别过滤

初始化如下:

`InitLog(日志路径, 日志文件名, 日志分割格式)`

日志分割格式参数如下:

 `COMMON_FORMAT 通用格式(不分割)`
 
 `MONTH_FROMAT 按月分割`
 
 `DATE_FORMAT 按天分割`
 
 `HOUR_FORMAT 按小时分割`
 
 
 举例:
 
 `import "github.com/boram-gong/log"`
 
 `log.InitLog("/var/log", "mylog", log.DATE_FORMAT)`
 
 打印日志:
 
 `log.ERROR("data")`
 
 设置级别过滤, 参数为级别有五个:
 
 `Info\ Debug\ Warn\ Error\ Fatal`
 
 `log.LogLevelFilter(log.Warn)`
 

更新功能:

    1. 打印日志时附带明确位置即行数
    
    2. 日志清理和日志大小查询
    
    3. 删除按分钟分割，新增按月分割 MONTH_FROMAT



