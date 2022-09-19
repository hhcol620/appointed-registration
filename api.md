### 114 网上预约挂号平台 监控平台

#### 114 官网查询接口 162 表示 hosCode

1. 首页的所有医院列表接口 分页查询

   https://www.114yygh.com/web/hospital/list?_time=1662730587759&keywords=&levelId=0&areaId=0&pageNo=1&pageSize=20

   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/

   | 参数     | 类型   | 默认值     | 含义                               |
   | -------- | ------ | ---------- | ---------------------------------- |
   | \_time   | number | Date.now() | 时间戳                             |
   | keywords | string | ---        | 关键词(可能与搜索也是同一个关键词) |
   | levelId  | number | ---        | 级别                               |
   | areaId   | number | ---        | 区域                               |
   | pageNo   | number | 1          | 分页属性                           |
   | pageSize | number | 20         | 分页属性                           |

2. 医院详情页 门诊列表

   https://www.114yygh.com/web/department/hos/list?_time=1662729484556&hosCode=162

   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/hospital/162/home

   | 参数    | 类型   | 默认值     | 含义                          |
   | ------- | ------ | ---------- | ----------------------------- |
   | \_time  | number | Date.now() | 时间戳                        |
   | hosCode | number | ---        | 医院 code (接口 1 中返回值中) |

3. 门诊挂号详情 通过上一步响应的内容 科室号 code 门诊 code 医院 code

   https://www.114yygh.com/web/product/list?_time=1662733088777

   Method: POST
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/hospital/162/d4a80dff4262f5815eee0362515986cf/200052524/source

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | firstDeptCode  | number | 科室 code | 时间戳                                                                                         |
   | secondDeptCode | number | ---       | 门诊 code (接口 2 中返回值中)                                                                  |
   | hosCode        | number | ---       | 医院 code (接口 1 中返回值中)                                                                  |
   | week           | number | 1         | 1 表示当前日期次日开始，7 天，然后返回的是 7 天的数据 2 表示当前日期开始第九天到第十五天的数据 |

   返回值中当天状态 status 值为 'AVAILABLE' 表示有号

4. 获取图片验证码
   https://www.114yygh.com/web/img/getImgCode?_time=${Date.now()}

   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/
   接口返回数据类型为buffer

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | _time  | number | --- | 时间戳                                                                                         |                                                            |


5. 验证图片验证码
   https://www.114yygh.com/web/checkcode?_time=${Date.now()}&code=${code}

   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | code  | string | --- | 短信验证码加密后字符串                                             

6. 发送短信验证码
   `https://www.114yygh.com/web/common/verify-code/get?_time=${Date.now()}&mobile=${encodeURIComponent(
                  cipherPhone
                )}&smsKey=LOGIN&code=${code}`

   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | _time  | number | --- | 时间戳                                                                                          |
   | mobile | string | ---       | 手机号加密后字符串 |
   | smsKey | string | LOGIN | 未知 |
   | code | string | ---- | 图片验证码

7. 验证短信验证码

   https://www.114yygh.com/web/login?_time=${Date.now()}

   Method: POST
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | code  | string | --- | 短信验证码加密后字符串                                                                                          |
   | mobile | string | ---       | 手机号加密后字符串                                                                 |

8. 获取科室某一天的剩余号详情
   https://www.114yygh.com/web/product/detail?_time=${Date.now()}
   Method: POST
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/hospital/162/eb89d3cd9db56d8cc3e29109aed61f6b/200048496/source

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | firstDeptCode  | number | 科室 code | 时间戳                                                                                         |
   | secondDeptCode | number | ---       | 门诊 code (接口 2 中返回值中)                                                                  |
   | hosCode        | number | ---       | 医院 code (接口 1 中返回值中)                                                                  |
   | target           | date | 1         | 格式 2022-10-01|

   返回值：

   ```js
      res.data = [
         {
            "dutyCode": "MORNING",
            "dutyImgUrl": "//img.114yygh.com/fe/home/iconfont/a87ff679a2f3e71d9181a67b7542122c",
            "detail": [
               {
               "uniqProductKey": "ce3bd6aded3a7cc339dcbd87b40c35429a5cd2b0",
               "doctorName": "刘一涵",
               "doctorTitleName": "副主任医师",
               "skill": "种植，只看口腔种植科，种植牙、种植美学修复、前牙美容修复及活动及固定义齿修复。",
               "period": [],
               "fcode": "&#xf7ee;&#xf128;",
               "ncode": "&#xf128;",
               "wnumber": 22,
               "znumber": 24
               }
            ],
            "showIndexPosition": 0,
            "showNumber": true,
            "showPrice": true
         },
         {
            "dutyCode": "AFTERNOON",
            "dutyImgUrl": "//img.114yygh.com/fe/home/iconfont/a87ff679a2f3e71d9181a67b7542122c",
            "detail": [
               {
               "uniqProductKey": "a0c335d532679bc3ebc3014f2f0890f14ed17f57",
               "doctorName": "刘一涵",
               "doctorTitleName": "副主任医师",
               "skill": "种植，只看口腔种植科，种植牙、种植美学修复、前牙美容修复及活动及固定义齿修复。",
               "period": [],
               "fcode": "&#xf7ee;&#xf128;",
               "ncode": "&#xf15e;",
               "wnumber": 13,
               "znumber": 30
               }
            ],
            "showIndexPosition": 0,
            "showNumber": true,
            "showPrice": true
         }
      ]
   
   ```
   注意：猜测 fcode ncode 图表字体，用来加密，其实可以避免爬取，参考对应其响应参数中的 dutyImgUrl, 可以将其下载下来然后再将文件后缀名修改为ttf

   fcode 对应的是价格 ncode 对应的剩余号数量 这个应该是对应了三个字体文件，vscode 下载一个字体预览的插件找到这个规律

   判断是否有余号  wnumber 为奇数表示还有余号

9. 判断用户是否实名 （可能有cookie 变化，需要重新塞进去）
   https://www.114yygh.com/web/user/real-name/status?_time=1664527932720
   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/hospital/162/eb89d3cd9db56d8cc3e29109aed61f6b/200048496/source

   返回值 
   ```json
      {
         "status": "AUTH_PASS",
         "name": "*贺",
         "idTypeView": "居民身份证",
         "idNo": "******2116",
         "message": "实名认证成功"
      }
   ```

10. 确认预约详情
   https://www.114yygh.com/web/product/confirm?_time=1664528415669
   Method: POST
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/hospital/162/eb89d3cd9db56d8cc3e29109aed61f6b/200048496/source

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | firstDeptCode  | number | 科室 code | 时间戳                                                                                         |
   | secondDeptCode | number | ---       | 门诊 code (接口 2 中返回值中)                                                                  |
   | hosCode        | number | ---       | 医院 code (接口 1 中返回值中)                                                                  |
   | target           | date | 1         | 格式 2022-10-01|    |
   | dutyTime           | number or string | 0        |  |
   | uniqProductKey    | string |        | 科室唯一信息  |

   返回值：
   ```json
      {
         "uniqProductKey": "c1fec2c4c7b73f32fd94798084f9364a633a22f0",
         "hospitalName": "中日友好医院",
         "departmentName": "儿科门诊",
         "doctorName": "医生",
         "doctorTitleName": "医生",
         "skill": null,
         "visitTime": "2022年10月01日 周六 下午 ",
         "totalFee": "&#xf3c2;&#xf38a;",
         "serviceFee": "50",
         "showFee": true,
         "needRemoteHospitalCard": false,
         "dataItem": {
            "hospitalCardId": 1,
            "hospitalCardIdTip": null,
            "jytCardId": 2,
            "jytCardIdTip": null,
            "smsCode": 4,
            "smsCodeTip": null,
            "contactUserInfo": 2
         },
         "commonRegisterNotice": null,
         "dutyImgUrl": "//img.114yygh.com/fe/home/iconfont/1679091c5a880faf6fb5e6087eb1b2dc",
         "showRegRule": null,
         "takePlaceTips": ""
      }
   ```


11. 获取就诊人信息
   https://www.114yygh.com/web/patient/list?_time=${Date.now()}&showType=ORDER_CONFIRM
   Method: GET
   鉴权: Cookie
   接口所在页面: https://www.114yygh.com/hospital/270/submission?hosCode=270&firstDeptCode=15bdfe158b9142390a393a474ace6571&secondDeptCode=200043560&dutyTime=0&dutyDate=2022-10-01&uniqProductKey=c1fec2c4c7b73f32fd94798084f9364a633a22f0

   | 参数           | 类型   | 默认值    | 含义                                                                                           |
   | -------------- | ------ | --------- | ---------------------------------------------------------------------------------------------- |
   | _time  | number |  | 时间戳                                                                                         |
   | showType | string | ORDER_CONFIRM       | 

   返回值：
   ```json
      {
         "count": 1,
         "list": [
            {
               "patientName": "黄贺",
               "idCardType": "IDENTITY_CARD",
               "idCardTypeView": "居民身份证",
               "idCardNo": "4114****2116",
               "idCardNoView": "4114****2116",
               "phone": "17796761085",
               "status": "BIND",
               "statusTips": "",
               "cardList": [
               {
                  "cardType": "SOCIAL_SECURITY",
                  "cardTypeView": "社保卡",
                  "cardNo": "1298****600X",
                  "cardNoView": "1298****600X",
                  "medicareType": "MEDICARE_CARD",
                  "medicareTypeView": "医保"
               },
               {
                  "cardType": "IDENTITY_CARD",
                  "cardTypeView": "居民身份证",
                  "cardNo": "4114****2116",
                  "cardNoView": "4114****2116",
                  "medicareType": "SELF_PAY_CARD",
                  "medicareTypeView": "自费"
               }
               ],
               "options": []
            }
         ]
      }
   ```




#### 注意事项

1. 用浏览器插件不断刷新 114 平台，发现可能有对 ip 封禁（每天同个 ip 访问数量限制）的可能，这块建议看看能不能实现在云服务器上进行动态切换 ip
2. 目前暂未发现连续刷接口会对账号造成影响，（可以考虑收费，使用虚假账号完成订阅）
3. 114 平台目前登录方式有两个，微信扫码和手机验证码登录，这步操作需要考虑 cookie 失效后自动完成重新登录，并且更新 cookie
   - 扫码的方式目前来看无法实现自动化操作，所以暂时不做延伸
   - 使用接码平台注册账号，注册完成提取 Cookie (综合来看使用验证码登录还算是一个方案)
4. 114 手机验证码加密方式

   aes 加密

   密钥： imed2019imed2019

   测试平台： https://the-x.cn/cryptography/Aes.aspx

   下面为一个测试用例

   - {
   - "mobile": "8E91WgqSzEHTX_a_3cFOag==",
   - "code": "fF60C2fgLcu2GR_YiWTOBg=="
   - }

   17796761085

   845411

#### 程序需要做的事情

- 已经注册用户，输入医院名称进行筛选， 保留 hosCode 162
- 根据 hosCode 获取科室信息 10c186f26ae7ecf8160e2dcf1f2e7312 200053529
- 用户输入起始监控日期（默认就是当前时间）和结束监控日期（目标预约日期）
- 到达起始监控日期开始进行监控
- 监控到有剩余号的，启动真实号码不挂代理去预约号（目前发现最好使用短信验证码登录，这样在提交的时候似乎可以不需要短信验证了）
- 有余量通知用户，帮忙挂号成功也通过邮件通知用户

#### Todo:

- 新建一张表用户记录上述信息（也就是用户输入信息）
- 用户的登录信息如何保存 Cookie
- 如何判断时间是否到达开始监控日期
- 对于到达监控日期的预约单，开启子进程刷接口
- 对于到达结束监控日期的子进程进行关闭，关闭这个操作可以放到上一步中

#### 鉴权方式

`headers: { Cookies: 'imed_session=atpXyENDKaSaxYo6ZAZvpSFswetcyDJl_5542437; imed_session=hIPDRiKlXtnBISc6Mxy5CXiahqOB2kTW_5542429; secure-key=e5e16e22-236d-478e-bdb5-aef48d5dee5d; imed_session=bsMFCIMQGYXZaUlC53MrWutjFWPMhLSi_5542915; cmi-user-ticket=d0p7EbDA5pFHrPZiLi2NpExOn1ZHnOL5cfrl3Q..; agent_login_img_code=ff5e0edf577e4d84b2e809c84f968641; imed_session_tm=1662874694023', Host: 'www.114yygh.com', 'Content-Type': 'application/json;charset=UTF-8', // 'Cache-Control': 'no-cache', Referer: 'https://www.114yygh.com/', Accept: 'application/json, text/plain, */*', 'Request-Source': 'PC', 'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36', 'Accept-Encoding': 'gzip, deflate, br', 'sec-ch-ua': '"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"', 'sec-ch-ua-mobile': '?0', 'sec-ch-ua-platform': 'macOS', 'Sec-Fetch-Dest': 'empty', 'Sec-Fetch-Mode': 'cors', 'Sec-Fetch-Site': 'same-origin' }`

### 完成情况
- [x] 所有医院
- [x] 一个医院的所有科室
- [ ] 一个科室是否有号
- [ ] 登录