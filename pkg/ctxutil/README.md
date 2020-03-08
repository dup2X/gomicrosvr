### 使用

1. 提供Setter/Getter方法设置或者获取Context的值。 

例如：

``` go
    import "ctxutil"
    //设置traceID
    ctxutil.SetTraceID("heheda~")
    //获取traceID
    traceID, _ := ctxutil.GetTraceID()
```

2. 如果需要把Context中的内容打印出来。

``` go
    import "ctxutil"
    //将Context的内容转换成字符串，ctx为你的context.Context
    str := ctxutil.TraceString(ctx)
```

3. 如果需要添加新的Setter/Getter方法。

例如：要往Context中添加City的内容，步骤如下：

``` go
    //Step 1: 添加City 的Key
    const (
    	_ keytype = iota
    	traceIDKey
    	spanIDKey
    	hintCodeKey
    	hintContentKey
    	langKey
    
    	passengerIDKey
    	driverIDKey
    	orderIDKey
    	districtKey
    	redisTimeKey
    	dbTimeKey
    	tokenKey
    
    	callerKey
    	degradeKey
    
    	httpRequestKey
    	//这是你添加的City的Key
    	cityKey 
    )
    
    //Step 2: 添加Setter方法
    func SetCity(ctx context.Context, city string) context.Context {
        return context.WithValue(ctx, cityKey, city)
    }

    //Step 3: 添加Getter方法
    func GetCity(ctx context.Context) (string, error) {
        	return getString(ctx, cityKey)
    }
    
    //Step 4: 为了让city这个字段可以打印出来，添加到Context字段的映射表
    //这样当你调用ctxutil.TraceString的时候，就可以把city打印出来了。
    var builtInKeys = []keytype{
    	traceIDKey,
    	spanIDKey,
    	redisTimeKey,
    	dbTimeKey,
    	passengerIDKey,
    	driverIDKey,
    	orderIDKey,
    	districtKey,
    	hintCodeKey,
    	hintContentKey,
    	langKey,
    	tokenKey,
    	//这是你添加的内容
    	cityKey,
    }
    
    var keyName = map[keytype]string{
    	traceIDKey:     "trace_id",
    	spanIDKey:      "span_id",
    	passengerIDKey: "passenger_id",
    	driverIDKey:    "driver_id",
    	orderIDKey:     "order_id",
    	districtKey:    "district",
    	redisTimeKey:   "redisTime",
    	dbTimeKey:      "dbTime",
    	hintCodeKey:    "hintCode",
    	hintContentKey: "hintContent",
    	langKey:        "lang",
    	tokenKey:       "token",
    	//这是你添加的内容
    	cityKey:        "city",
    }
```

### 为什么要context util

官方的context.Context本质上是一个链表，通过它暴露的Value方法，可以把它看成是php中的array。

它方便但同时也有php array的不足，即当多人协作开发时，你不知道context里面有什么。

类型系统是一个思路，如果有了类型系统，我们通过IDE的提示就能知道context中有哪些field，就可以去取对应的值。

但这并不是一个好的方式，context官方文档中明确反对把context内嵌到结构体中。

同时自定义结构体的另一个坏处是不符合官方函数签名，很多标准库的方法要求第一个参数必须是context.Context。

虽然可以让自定义的类型实现context.Context接口规定的方法，但是到了函数中，也就只能调用context.Context暴露的几个方法了。

想要恢复原状只能类型推断，但是这样又引入了很多麻烦。

另一种思路是开发一个lib，这个lib通过暴露一系列方法，从而让context看起来像是有类型一样。

这和golang本身实现对象方法的方式比较类似，在go中给一个对象添加方法：

``` go
type some struct{}

func (s some) Do() {

}

//使用时
var a some
a.Do()
```

而lib的方式和这非常类似：

``` go
func Do(ctx context.Context) {

}

//使用时
Do(ctx)
```

本质上就是把对象方法的receiver作为函数的第一个参数传进去，逻辑上间接地给context加上了一些方法。

这是一种非侵入式的方法，lib提供了一些setter getter操作context，对于使用方来说，有需要就用，没需要就不用，完全不需要修改自己现有的代码。

提供setter和getter本质上还是想避免人为地拼写key，容易出错。比如set时写的key是hello，get时误写成了hallo。

如果定义为同一个const，那么还得考虑这些定义应该放在什么地方，会不会出现循环引用。gettter和setter屏蔽了这些细节，用户不需要管key，key是内置的。

当然，这里提供的key只是业务中必须或者常用的，各个业务自有需求的key，目前还是需要业务方自己调用context.WithValue去set。我们后续将提供方法供业务去set自定义key。

### 输出context

context的另一个问题就是，用户无法得知它内部到底存了些什么，用户必须记得自己set过哪些key，才能把它们一一取出来打日志。

ctxutil提供了Format方法，会自动检测内置的key是否被设置过，然后把这些key的值格式化成一个字符串。ctxutil对于每一个内置key都暴露一个setter和一个getter，用户无法得知具体的key。

但是这种方式进行格式化也是有局限性的，这也是为什么我们后续将提供方法让业务set自定义key。

因为如果业务方自己通过context.WithValue设置key，ctxutil无法得知具体的key有哪些。

而如果通过ctxutil暴露的方法进行设置，ctxutil能够把设置过的key进行缓存，从而在格式化的时候可以输出context内部的值。

当由于context需要保证线程安全，需要考虑的方面很多，具体的实现目前还在讨论中，之后会支持用户自定义key。

文档：
`godoc -http=':端口号' go.intra.xiaojukeji.com/gobiz/ctxutil`

### 联系
有问题请找联系： dengyi@didichuxing.com 或者 liudeen@didichuxing.com