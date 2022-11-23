### 114平台 北京医院挂号 科室余号监控 自动挂号

#### 环境安装
  1. Node
  2. 其他的应该不需要安装

#### appointed-server 为后端项目
##### 开发项目 
  1. 创建mysql 数据库，并导入根目录auto-register.sql 数据
  数据库相关信息
  2. 修改数据库配置appointed-server/src/database/orm-data-source.ts
  3. 本地启动项目
    安装依赖 npm install
    启动 npm run start

#### appointed-web 为前端项目

##### 开发项目 
  1. 本地启动项目
    安装依赖 npm install
    启动 npm run dev
    浏览器打开 http://localhost:9000
  2. 平台默认的账号密码 auto-register@163.com
      admin
      
      
 #### 可拓展的地方
   1. 可配合安卓上的脚本配合使用，每次需要重新登录，发送验证码，通过脚本自动提交到服务，前提条件，项目运行在服务器上或者让自己手机和电脑在同一个网络下
   2. 脚本参考 https://github.com/snailuncle/autojsDemo/tree/master/%E6%8E%A5%E7%A0%81



#### 联系作者
* 如果有好的想法和建议，请联系作者~~
* 平台的话可能会陆续补充

<img src="./images/WechatIMG2.jpeg" width="200px" /> <img src="./images/WechatIMG4.jpeg" width="200px" /> 




