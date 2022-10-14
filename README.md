# encrypt_pass

## 非对称加密-密码生成器

### 原理
1. 应用标识符+基础密码（记心中）
> 如应用标识符是qq, 基础密码是123456;
> 基础密码自己记心中，所有应用一样，只是应用标识符不同;
> 比如Gmail账号密码构成就是:gm123456, gm表示gmail， 123456是心中记住的通用基础密码.

2. 非对称加密：MD5(md5 16 全小写方式)
> 对 qq123456进行md5得到：12a6b1e7fbfa2ce5

3. 对非对称加密结果进行base16(这里直接用16进制代替)
> 如12a6b1e7fbfa2ce5进行base16得到：31326136623165376662666132636535

4. 对base16结果去掉其中字母，进行想要长度截取
>  如截取base16结果的前8位：31326136
> 如果需要大小写和特殊符号，则可收尾添加字母和特殊符号，如：Qq31326136.
 

### for help:
`encrypt_pass -h`
> usage: encrypt_pass [mypassword] [length], eg: encrypt_pass abc123456 8

### eg:
```shell
s@m encrypt_pass % ./encrypt_pass_darwin aa123456 12
input parameters: input password is aa123456, lenth is 12 

MD5 sum into hex number:
8a6f2805b4515ac12058e79e66539be9

last passpord after truncation:
8a6f2805b451

strip alphabet and add prefix & postfix left:
Aa862805451.
```

