# 配置说明：

## 配置数据，在于方便读取，数据安全，可便捷于开发和非开发人员操作
### 数据加载过程中，需要数据正确性保障，以及对于非正常数据的警告提示，正常加载数据到内存

<!--

resource load interface:

type ICommonConfig interface {
	ComfireAct(data interface{}) (errlist []string)
	DataRWAct(data interface{}) (errlist []string)
}

use for example with templateconfig.go

you can use tool explore exls for config struct.
tool link: 
    xExcel2x: https://github.com/Peakchen/xExcel2x
    xExport4Go: https://github.com/Peakchen/xExport4Go
    
-->

<!-- 
    serverConfig:  服务器相关配置
    LogicConfig:   逻辑相关配置
-->