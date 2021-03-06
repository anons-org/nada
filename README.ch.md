## Nada

### 项目同步
https://github.com/anons-org/nada
https://gitee.com/grateful/nada

### 项目介绍
Nada参考了JVM,CLI，ZENDVM，V8,Node等运行时的设计，并支实现了JVM的指令集，支持运行Ts,Js,PHP,Java等高级语言,同时又对javascript进行了增强（类型、接口、泛型）、从根本上解决了JS多年的诟病。

### 目的
- 能在GO中让PHP,JAVA,JS,TS作为脚本运行
- Nada作为一个运行时，核心目标是为创业者和企业提供快速试错，快速验证产品，让产品以最低的技术成本快速占领市场的底层服务框架。





### 开发计划
- 第一阶段
  - class解析 70%
  - 数据类型 70%
  - 解释器 2%
  - 工具包 1%
    - jar、zip解析 0%
- 第二阶段
   - ndk(对标NODE API) 0%
   - debug agent 0%
   - vscode语法插件 0%
   - idea 语法插件 0%
   - LSP语言服务器
   - 多线程 0%
- 第三阶段
   - 跨平台桌面UI库
    - 包管理
    - ffi
    - JS语法树
    - PHP语法树
    - typescript语法树
    - python 语法树



#### 相关参考
- LSP https://microsoft.github.io/language-server-protocol/
- vscode debug protocol https://vscode.readthedocs.io/en/latest/extensions/example-debuggers/
- Go协程这样用才安全 https://taoshu.in/go/safe-goroutine.html
- Go 语言 map 的并发安全问题 https://taoshu.in/go/go-map-concurrent-misue.html
- JVM 的类加载机制 https://baijiahao.baidu.com/s?id=1662931890502685014&wfr=spider&for=pc


### 参与贡献
-  Fork 本仓库
-  新建 Feat_xxx 分支
-  提交代码
-  新建 Pull Request




