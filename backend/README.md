创建数据库<br>
进入mysql root用户后<br>
首先使用`create database yoyu;`创建数据库<br>
然后使用 
`grant all on yoyu.* to yoyu@'localhost' identified by '123456';` 命令以按照默认设置配置数据库。

用户不能超过255位，只能包括字母和数字
密码不能超过255位，否则会返回错误码4xx


数据模型
Post 添加 message_box_id(integer) poster_id(integer)