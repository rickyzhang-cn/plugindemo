## 简介
+ 演示使用golang的plugin package来实现代码动态更新
+ plugin本身只是使用cgo对dlopen，dlsym进行了简单的封装
+ plugin模块没有提供关闭so的调用

## 使用方法
+ 运行./build.sh生成app
+ 运行./build_so.sh 1生成logic_plugin1.so
+ 运行./app
+ 在浏览器中访问相关url增加，删除和查看group store的数据
+ 修改logic中的代码，运行./build_so.sh 2生成logic_plugin2.so
+ 访问[http://127.0.0.1/load?filename=logic_plugin2.so]让app重新加载新的so
