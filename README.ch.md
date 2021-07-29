## Nada

### 项目同步
https://github.com/anons-org/nada
https://gitee.com/grateful/nada

### 项目介绍
Nada参考了JVM,CLI，ZENDVM，V8,Node等运行时的设计，并支实现了JVM的指令集，支持运行Ts,Js,PHP,Java等高级语言,同时又对javascript进行了增强（类型、接口、泛型）、从根本上解决了JS多年的诟病。

### 目的
- Nada作为一个运行时，核心目标是为创业者和企业提供快速试错，快速验证产品，让产品以最低的技术成本快速占领市场的底层服务框架。

### 特有功能

- javascript增强,Nada使用jss规范对javascript进行了增强
    
    - let、function可忽略
    - 注解支持
      ```
      属性注解、类注解、方法注解、接口注解、自定义注解
      ```

    - 类型增强
      ```
      int  bigInt boolean string  double float array any
      ```
    - demo
      ```
      test(int a, any b){
      }
      ```

    - 接口支持
      ```
      class Test extends A implements IA,IB,IC{
      }
      ```
    - 静态属性和方法支持
      ```
      class Test extends A implements IA,IB,IC{
          static string name;
          static getName(){
          }
      }
      ```
    - 接口
      ```
      Interface IFoo{
      }
      ```



- 热更
- 原生分布式
- 支持编译插件
- 多指令集支持


### 开发计划
- 第一阶段
  - class解析 60%
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




#### 相关支持
- LSP https://microsoft.github.io/language-server-protocol/
- vscode debug protocol https://vscode.readthedocs.io/en/latest/extensions/example-debuggers/


### 参与贡献
>  在代码规范上我们有很严格的要求，请认真阅读规范，以确定您能接受规范后再fork!
-  [规范文档](https://gitee.com/grateful/farvm/wikis/%E5%BC%80%E5%8F%91%E8%A7%84%E8%8C%83?sort_id=3481374)
1.  选择擅长的分组
2. [分组列表](https://gitee.com/grateful/farvm/wikis/%E6%93%85%E9%95%BF%E5%88%86%E7%BB%84?sort_id=3481509)
3.  Fork 本仓库
4.  新建 Feat_xxx 分支
5.  提交代码
6.  新建 Pull Request




