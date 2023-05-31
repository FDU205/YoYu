创建数据库<br>
进入mysql root用户后<br>
首先使用`create database yoyu;`创建数据库<br>
然后使用 
`grant all on yoyu.* to yoyu@'localhost' identified by '123456';` 命令以按照默认设置配置数据库。

用户不能超过32位
密码不能超过255位，否则会返回错误码4xx

表白墙内容不能超过200个字
提问的内容不能超过200个字
提问箱标题不能超过50个字


数据模型
Post 添加 message_box_id(integer) poster_id(integer)

Visibility 2是匿名， 1是实名
Type 1是追问， 2是回答