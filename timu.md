## golang 面试题目

### 目录
1. [= 和 := 的区别](#jump1)
2. [指针的作用](#jump2)
3. [golang 允许多个返回值吗](#jump3)
4. [拼接字符串](#jump4)





### 题目
1. = 和 := 的区别<a id="jump1"></a> 
=是赋值语句，:=是定义变量。
2. 指针的作用<a id="jump2"></a>  
指针可以指向数据的地址。
3. golang 允许多个返回值吗<a id="jump3"></a> 
允许，例如 err 
4. 拼接字符串<a id="jump4"></a>
- ***+***  
会遍历两个字符串，开辟一块新的空间来存储拼接字符串。
- fmt.Sprintf  
由于采用接口参数，要用反射获取，有性能损耗。
- strings.Builder  
内部实现指针+切片，避免变量拷贝    
writeString()拼接  
String()返回,直接将[]byte转换为String。
- bytes.Buffer  
缓冲byte类型的缓冲器，存放的都是byte  
底层也是一个[]byte切片
- strings.join
基于strings.Builder实现，提前进行容量分配  
减少内存分配，很高效。

***理解***:如何实现高效的拼接字符串?内存损耗小,
速度快。为什么"+" 和 "strings.Builder"性能差别
这么大?底层分配内存方式不同,"+"会开辟一块新空间来
拼接，而string.Builder底层是指针+切片,容量申请
不同。当反复拼接字符串时，"+"每次都需要申请新空间，
而底层切片的话只有在当前切片容量不够时才会继续申请。
那么为什么"strings.Builder"又比"bytes.Buffer"
快呢，那是因为"bytes.Buffer"在返回时重新申请了
一块空间。

***性能比较***  
string.Builder > bytes.Buffer > + > fmt.Sprintf

