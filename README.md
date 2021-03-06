



# 数据结构

[TOC]



## 基本概念

### 数据元素

数据的基本单位，在计算机程序中通常作为一个整体进行考虑和处理；一个数据元素可以由多个数据项组成，数据项是数据不可分割的最小单位。

### 数据对象

性质相同的数据元素的集合，是数据的子集。

### 数据结构

是相互之间存在一种或者多种特定关系的数据元素的集合

1. 集合：结构中的数据元素之间除了“属于同一个集合”的关系外，别无其他的关系
2. 线性结构：结构中的数据元素之间存在一个对一个的关系
3. 树形结构：结构中的数据元素之间存在一个对多个的关系
4. 图状结构和网状结构：结构中的数据元素之间存在多个对多个的关系。

数据的逻辑结构一般是数学模型即数学定义上的逻辑关系，数据结构是数据在计算机内部存储的物理结构也叫存储结构。

### 抽象数据类型（ADT）

是指一个数据模型以及定义在该模型上的一组操作。

``` ADT抽象数据类型定义格式
ADT抽象数据类型名{
	数据对象:<>
	数据关系:<>
	基本操作:<对数据的增删改查操作及其他处理操作>
}
```

## 算法

定义：对特定问题求解步骤的一种描述，它是指令的有限序列，其中每一条指令表示一个或多个操作。

### 特性：

1. 有穷性：一个算法必须总是在执行有穷步之后结束，且每一步都可在有穷时间内完成。
2. 确定性：算法每一条指令都必须有确切的含义，读者理解时不会产生二义性，且在任何条件下，算法只有唯一的一条执行路径，相同输入只能有相同输出。
3. 可行性：一个算法是能行的，即算法中描述的操作都是可以通过已经实现的基本运算执行有限次来实现。
4. 输入：一个算法有零个或多个输入，这些输入取自于某个特定的对象的集合。
5. 输出：一个算法有一个或多个的输出，这些输出是同输入有着特殊关系的量。（不一定有输入，但是一定有输出，输出一定和输入有关系）

### 算法设计的目标

1. 正确性：算法应当满足具体问题的需求。大体分为四个层次
   1. 程序不含语法错误
   2. 程序对于几组输入数据能够得出满足规格说明要求的结果
   3. 程序对于精心选择的典型、苛刻而带有刁难性的几组输入数据能够产生满足规格说明要求的结果
   4. 程序对于一切合法的输入数据都能产生满足规格要求的结果
2. 可读性：算法是给人看的，得让人能看得懂，机器是肯定懂的。
3. 健壮性：当输入数据非法时，算法也能适当地做出反应或进行处理，而不会产生莫名其妙地输出结果。
4. 效率和低存储需求：通俗易懂地讲又快又轻巧，存储占用少，耗时短。

### 算法效率的度量

一般情况下，算法中基本操作的重复执行的次数是问题规模`n`的某个函数`f(n)`，算法的时间度量记作：
$$
T(n)=O(f(n))
$$
语句中的`频度`指的是该语句重复执行的次数。

因为在程序里面，基本语句的执行对于机器来说可以忽略不计，计算复杂度时一般采用执行次数最多的程序块，不管实际执行情况如何，皆取理论上的最大次数，程序块的频度基本是1，n，n^2等，所以时间复杂度分别为
$$
常量阶：O(1),
线性阶O(n),
平方阶O(n^)等
$$
由此可见，我们应该尽量选用多项式阶的算法，避免使用复杂度高如指数阶的算法。

空间复杂度
$$
S(n)=O(f(n))
$$
略

## 线性表



## 栈和队列

