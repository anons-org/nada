## Nada

### Project introduction
Nada refers to the design of JVM,CLI, ZENDVM, V8,Node and other runtime, and implements the instruction set of JVM, supports the running of Ts,Js,PHP,Java and other high-level languages. At the same time, it enhances javascript (type, interface, generics), and fundamentally solves the longstanding complaints of Js.

# # #  
- As a runtime, the core goal of Nada is to provide entrepreneurs and businesses with a basic service framework for fast trial and error, fast product validation, and fast market capture at the lowest technology cost.

### Features

- javascript enhancements. Nada uses the JSS specification to enhance javascript

- Let and function can be ignored
- Annotation support  
  ` ` `  
  Attribute annotation, class annotation, method annotation, interface annotation, custom annotation  
  ` ` `

- Type enhancement  
  ` ` `  
  int bigInt boolean string double float array any  
  ` ` `
- demo  
  ` ` `  
  test(int a, any b){  
  }  
  ` ` `

- Interface Support  
  ` ` `  
  class Test extends A implements IA,IB,IC{  
  }  
  ` ` `
- Static property and method support  
  ` ` `  
  class Test extends A implements IA,IB,IC{  
  static string name;  
  static getName(){  
  }  
  }  
  ` ` `
- interface  
  ` ` `  
  Interface IFoo{  
  }  
  ` ` `



- hot more
- Native distributed
- Support for compiling plug-ins
- Multiple instruction set support


### Development plan
- Stage 1
- Class resolves 60%
- Data type 70%
- Interpreter 2%
- Toolkit 1%
- Jar and zip parsing 0%
- Stage 2
- NDK (NODE API) 0%
- debug agent 0%
- Vscode syntax plugin 0%
- Idea syntax plugin 0%
- LSP Language server
- Multithreaded 0%
- Stage 3
- Cross-platform desktop UI library
- package management




#### Related support
- LSP https://microsoft.github.io/language-server-protocol/
- vscode debug protocol https://vscode.readthedocs.io/en/latest/extensions/example-debuggers/
- https://www.cnblogs.com/cangqinglang/p/14204957.html
- https://www.zhihu.com/question/49043143


### Contribute
> We have very strict requirements on the code specification, please read the specification carefully to make sure that you can accept the specification before fork!
- [specification] (https://gitee.com/grateful/farvm/wikis/%E5%BC%80%E5%8F%91%E8%A7%84%E8%8C%83?sort_id=3481374)
1. Choose groups you are good at
2. The grouping [list] (https://gitee.com/grateful/farvm/wikis/%E6%93%85%E9%95%BF%E5%88%86%E7%BB%84?sort_id=3481509)
3. Fork the warehouse
4. Add Feat_ XXX branch
5. Submit the code
6. Create a Pull Request  