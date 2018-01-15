# sort.Search()的使用方式

## 二分搜索算法
基维百科上关于[二分搜索算法](https://zh.wikipedia.org/zh-cn/%E4%BA%8C%E5%88%86%E6%90%9C%E7%B4%A2%E7%AE%97%E6%B3%95)的定义是

> 在计算机科学中，二分搜索（英语：binary search），也称折半搜索（英语：half-interval search）[1]、对数搜索（英语：logarithmic search）[2]，是一种在有序数组中查找某一特定元素的搜索算法。搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，则搜索过程结束；如果某一特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，则代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半。

## sort.Search的用法

    func Search(n int, f func(int) bool) int

官方文档这样描述该方法：
>Search()方法回使用“二分查找”算法来搜索某指定切片[0:n]，并返回能够使f(i)=true的最小的i（0&lt;=i&lt;n）值，并且会假定，如果f(i)=true，则f(i+1)=true，即对于切片[0:n]，  
>i之前的切片元素会使f()函数返回false，i及i之后的元素会使f()函数返回true。但是，当
>在切片中无法找到时f(i)=true的i时（此时切片元素都不能使f()函数返回true），Search()
>方法会返回`n`。

Search()函数一个常用的使用方式是搜索元素x是否在已经`排好顺序`的切片s中：

```golang
    x := 11
    s := []int{3, 6, 8, 11, 45} //注意已经升序排序
    pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
    if pos < len(s) && s[pos] == x {
        fmt.Println(x, "在s中的位置为：", pos)
    } else {
        fmt.Println("s不包含元素", x)
    }
```

官方文档还给出了一个猜数字的小程序：

```golang
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
```   

## 总结
1. 被查找的切片或者数组a，需要是`排好顺序`的。
1. `排好顺序`的意思是，存在一个i，使得对于所有的x>=i,f(x)返回true