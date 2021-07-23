## Nada

### 项目介绍
Nada参考了JVM,CLI，ZENDVM，V8,Node等运行时的设计，并支实现了JVM的指令集，支持运行Ts,Js,php,java等高级语言,同时又对javascript进行了增强（类型、接口、泛型）、从根本上解决了JS多年的诟病。

### 目的
- Nada作为一个运行时，核心目标是为创业者和企业提供快速试错，快速验证产品，让产品以最低的技术成本快速占领市场的底层服务框架。
- Nada基于GO语言开发，这就意味着你的产品可以轻松移植到各大硬件平台，迅速抢占市场先机。
- 在这个项目中你能学习到虚拟机是如何为高级语言提供强大功能的原理，以及如何自己实现一门编程语言！同时能GET到的原理包括但不限于


- 在这个过程中我们将实现的内容包括但不限于
类型系统、线程、协程、进程调度、IO管理、JIT、debuger套件、编译器。
该项目轮子超多，欢迎大家一起来参与、研究。详情私聊Q184377367。
  - 急寻以下小伙伴
    
    - 工具组：至少熟练一门编程语言,负责各开发组内的增效工具的实现
      
    - VM内核组：熟练GO、熟悉操作系统原理,负责线程，协程，解释器的实现
      
    - JIT组：熟练X86、arm、risv汇编 ,负责字节缓存和实现本地JIT代码的实现
      
    - 编译器组：熟练GO 熟练编译原理，负责编译器开发，将JS，TS,Python，JAVA,编译为虚拟机字节码
      
    - NDK组：熟练go ,负责原生方法包的实现，为高级语言层面提供原生方法的支持，熟悉任意一样即可（GTK、libuv、libenvent、协程、线程、IO管理）

    - debuger组：(熟悉IDEA插件，VSCODE插件)负责VM,debuger功能的实现，至少支持在vscode中对Nada内核进行调试
    - GO语言专家：熟练协程，多线程、数据库、网络原理、操作系统原理、且至少有3门编程语言实战经验；负责对项目中的设计提出优化建议，必要时亲自优化代码。




 

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




### 参与贡献
>  在代码规范上我们有很严格的要求，请认真阅读规范，以确定您能接受规范后再fork!
-  [规范文档](https://gitee.com/grateful/farvm/wikis/%E5%BC%80%E5%8F%91%E8%A7%84%E8%8C%83?sort_id=3481374)
1.  选择擅长的分组
2. [分组列表](https://gitee.com/grateful/farvm/wikis/%E6%93%85%E9%95%BF%E5%88%86%E7%BB%84?sort_id=3481509)
3.  Fork 本仓库
4.  新建 Feat_xxx 分支
5.  提交代码
6.  新建 Pull Request




