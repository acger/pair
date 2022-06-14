Acger Pair 是基于微服务框架go-zero而搭建的一个练手项目，供大家学习、参考

Acger Pair is a project built to learn go-zero, enjoy your self

项目基于TiDB、Elasticsearch、Filebeat、Kafka、Redis等技术栈开发，也包含了gorm、validator、copier、json-iterator、go-queue、go-elasticsearch、websocket、qiniu等等常用库，有兴趣的小伙伴翻一下代码就可以看到相关实践。服务端总共只有三个模块，用户模块 - user， 匹配模块 - pair， 聊天通讯模块 - chat，每个模块都有自己得api与rpc服务，手写的代码不过两三百行。自从使用了go-zero做开发，~~摸鱼时间翻倍~~生产效率直接提高100% ，go-zero yyds！

- 这里顺手安利一下**tidb**，在群里看到经常有小伙伴问分布式事务怎么搞，其实把这个问题交给分布式数据库就好了。当你使用tidb的时候，跟使用单体的mysql数据库几乎是一模一样的，零代码入侵。分布式事务中的并发处理、数据的拆合、节点间的调度全部都在tidb内部完成，应用端可以无感知的使用，非常舒适。

- 还有提一下的是，项目使用Github Actions + K8S 实现自动化CI/CD部署。部署流程也相当简单，需要用到的文件都有保存在项目当中，详细可以查看以下目录：
```
/.github/workflows  --- actions工作流的配置文件
/deploy/actions/dockerfile --- 部署时要用的dockerfile文件（去除了goproxy设置）
/deploy/k8s --- 初次部署到k8s需要用到的manifest文件
/deploy/k8s/temp  --- temp文件夹下也有dockerfile，方便手动制作镜像
```
