# pass-gener

## 非对称加密-密码生成器

### 原理
1. 应用标识符+基础密码(secret)（记心中）
> 如应用标识符是qq, 基础密码是123456;
> secret 自己记心中（本程序任何地方都不会保存，所以也不存在泄漏），所有应用一样，只是应用标识符不同;
> 比如Gmail账号密码构成就是:gm123456, gm表示gmail， 123456是心中记住的 secret。
> 只要应用标识符和secret不变， 每次生成的密码就是一样的， 不记得密码的时候， 直接再生成一次就可以查到密码了。

2. 非对称加密：MD5(md5 16 全小写方式)
> 对 qq123456进行md5得到：12a6b1e7fbfa2ce5

3. 对非对称加密结果进行base16(这里直接用16进制代替)
> 如12a6b1e7fbfa2ce5进行base16得到：31326136623165376662666132636535

4. 对base16结果去掉其中字母，进行想要长度截取
>  如截取base16结果的前8位：31326136
> 如果需要大小写和特殊符号，则可收尾添加字母和特殊符号，如：Qq31326136.
 

### for help:
`pass-gener -h`
> usage: pass-gener [mypassword] [length], eg: pass-gener abc123456 8

### eg:
```shell
s@m pass-gener % ./pass-gener_darwin aa123456 12
input parameters: input password is aa123456, lenth is 12 

MD5 sum into hex number:
8a6f2805b4515ac12058e79e66539be9

last passpord after truncation:
8a6f2805b451

strip alphabet and add prefix & postfix left:
Aa862805451.
```

