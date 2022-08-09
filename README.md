--------------------------------------------------------
Settings 
=> selected:Code And automation-Branches
=> selected:Branch protection rules - master - Edit
--------------------------------------------------------

# Go-Note(Go-笔记)

## 声明:
凡是所参考的编程代码文章和本人自己所写的代码,都是笔记形式代码,只作参考用,不作实际生产环境使用!

## 代码运行环境说明:
a.go版本最起码要go version go1.13.5的,当时写代码和测试的版本是go1.13.5!  
b.基于 go mod 管理包加载!

## 关于提示错误是exec: "git": executable file not found in %PATH%的处理
这是由于第一次安装GitGUI后,没有将;E:\GitGUI\exe\bin加入到系统环境变量中,导致go get执行时报git不在%PATH%中的提示  
执行步骤(这里以windows系统为例):  
a.找到GitGUI或者Git安装目录,在找到Git所属bin目录,然后将该bin目录路径复制下来.  
b.找到 控制面板\系统和安全\系统->高级系统设置->所处顶部导航栏位"高级",下面的"环境变量"按钮.  
c.点击"环境变量"按钮后,弹出框中找到"系统变量"里的"Path"字段,双击打开后,在"变量值"的输入栏中,鼠标光标移动到最后,先输入英文分号(;)隔开之前的路径,再将刚才复制的Git所属bin目录路径(如E:\GitGUI\exe\bin)追加到最后写入.   
d.最后点击"确定",按原路点击"确定"一直到结束弹框显示;再将所打开的IDE或CMD命令行窗口等,关闭下再重新打开,再执行go get命令即可!

## 参考相关文章地址:
1.有关TCP,UDP编程代码  
http://topgoer.com/%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B/socket%E7%BC%96%E7%A8%8B/TCP%E7%BC%96%E7%A8%8B.html
