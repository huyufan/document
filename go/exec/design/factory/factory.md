# 简单工厂

## 用于创建单一产品，将所有子类的创建过程集中在一个工厂中， 如要修改，只需修改一个工厂即可。简单工厂经常和单例模式一起使用， 例如用简单工厂创建缓存对象（文件缓存），某天需要改用redis缓存，修改工厂即可。

# 抽象工厂

## 常用于创建一整个产品族，而不是单一产品。 通过选择不同的工厂来达到目的，其优势在于可以通过替换工厂而快速替换整个产品族。 例如上面的例子美国工厂生产美国girl，中国工厂生产中国girl。