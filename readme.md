gocmdb
web + mysql

1. go编译环境
2. mysql (mariadb)
3. GOPROXY=https://goproxy.io

4. 创建数据库
    create database gocmdb default charset utf8mb4;
5. 修改数据库连接配置
    conf/db.conf
    dsn=root:881019@tcp(localhost:3306)/gocmdb?charset=utf8mb4&loc=Local&parseTime=True

    用户名:密码@tcp(数据库服务地址:数据库端口)/数据库名称?
6. go run web.go -init
7. go run web.go
    http://ip:port/auth/login

tengxun:
    YBjOre
    endpoint: cvm.tencentcloudapi.com
    region: ap-beijing
    ak: AKIDA3rU2VeTyVurYD4IYt1v1poW1Gn3rq6Q
    sk: 1WkO3ScV8Ypx5drufy5vkrmUheT1G9uH

ali:
    endpoint:ecs-cn-shenzhen.aliyuncs.com
    region:cn-guangzhou
    ak:LTAI4GGKciyqCZ4UutbKo2wv
    sk:rfZ6JoGIKUUA87k2MNRtRaAgNSStNv

promethues:
    Node:
        select delete
        add:
            agent: api register
        models:
            唯一标识: uuid
                     addr
                     username
                     password
                     create_at
                     update_at
                     deleted_at
    Job:
        select delete update 
        models:
            job
    Target:
        select delete update add


orm:
    LoadRelated 可以反向关联查询
Agent:
    1. 注册
    2. 获取配置，生成prometheus配置


初始化admin成功, 默认密码: ctvLP_