【类似Java中的Maven】
![img.png](img/img.png)
![img_1.png](img/img_1.png)
验证：
![img_2.png](img/img_2.png)
![img_3.png](img/img_3.png)
![img_4.png](img/img_4.png)
![img_5.png](img/img_5.png)
## 使用Goland创建GoModule项目
【1】new-project:
![img_6.png](img/img_6.png)
【2】配置：

![img_7.png](img/img_7.png)
【3】选择窗口打开位置：

![img_8.png](img/img_8.png)
【4】项目结构：
![img_9.png](img/img_9.png)
【5】安装第三方包和之前方式一样：
GoModule小试牛刀：web框架Gin安装
录入安装gin的命令：go get -u github.com/gin-gonic/gin

## Goland配置File Warchers
【1】go fmt 格式化（代码如果错乱，利用ctrl+s即可格式化）
![img_10.png](img/img_10.png)
![img_11.png](img/img_11.png)
【2】goimports自动导入包配置
结果：
![img_12.png](img/img_12.png)
如何配置：

![img_13.png](img/img_13.png)
第一次进入提示需要下载：
termi中录入命令：go get golang.org/x/tools/cmd/goimports 
![img_14.png](img/img_14.png)
![img_15.png](img/img_15.png)


配置成功以后：
![img_16.png](img/img_16.png)

PS：每一个项目需要重新配置一次
