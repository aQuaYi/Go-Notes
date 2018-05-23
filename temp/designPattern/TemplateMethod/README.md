# 模板方法模式（Template Method）
实现了算法和步骤具体实现方式的分离。
## 定义
模板方法模式定义了一个算法的步骤，并允许子类别为一个或多个步骤提供其实践方式。让子类别在不改变算法架构的情况下，重新定义算法中的某些步骤.

## 模块讲解

Display()的步骤是，
1. 先open()一下
2. 再print()5遍
3. 最后close()

charDisplay和stringDisplay的Display方法，都遵循这样的步骤。但是，他们两的open，print和close又各不相同。 所以，把步骤放入了display结构体中。