### 思路

~~未解决，golang里面不知道如何创建动态二维数组~~

1. 首先，生成一个二维数组
2. 依次插入数据

总结：生成二维数组
- 生成一个res:=make([][]int,n)
- 再循环res[i]=make([]int,m)来插入到res

> 数组(Array)<br><br>
> Time：60ms<br><br>
> Beats：97.85%