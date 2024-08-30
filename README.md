# secureuuid
安全 UUID 工具，J secureuuid 的 Go 实现

# 参数
* `-v` `-version` - 打印软件版本
* `-l` `-license` - 打印软件许可
* `-r` `-random` - 指定随机 UUIDv4
* `-u <uuid>` `-uuid <uuid>` - 指定固定 UUID
* `-n <uuid>` `-namespace <uuid>` - 指定命名空间
* `-md5 <uuid>` - 指定 UUIDv3 名称
* `-sha1 <uuid>` - 指定 UUIDv5 名称
* `-c` `-compact` - 打印紧凑UUID（xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx）
* `-s` `-standard` - 打印标准UUID（xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx）
* `-no-newline` - 不打印回车符号

# 用处
## 生成

<details>

<summary>随机生成安全 UUIDv4（密码级别随机）</summary>

```shell
secureuuid -r -s
```

```
11e416c4-6962-496a-887f-02be4e4bfd79

```

</details>

## 验证

<details>

<summary>验证 UUID 正确性，如果 UUID 合法，程序返回 0，否则返回 125</summary>

可以尝试不搭配 `-compact` 与 `-standard`，这样只会在错误流输出错误信息

```shell
secureuuid -u abcdefgh-ijkl-mnop-qrst-uvwxyz012345 -no-newline
echo $?
```

```
125
```

</details>

## 转换

<details>

<summary>将有无连接符号的 UUID 互相转换</summary>

```shell
secureuuid -u 11e416c4-6962-496a-887f-02be4e4bfd79 -c
```

```
11e416c46962496a887f02be4e4bfd79

```

<hr />

```shell
secureuuid -u 11e416c46962496a887f02be4e4bfd79 -s
```

```
11e416c4-6962-496a-887f-02be4e4bfd79

```

</details>

## 哈希

<details>

<summary>生成 UUIDv3（MD5）或 UUIDv5（SHA1）</summary>

```shell
secureuuid -n 6ba7b810-9dad-11d1-80b4-00c04fd430c8 -sha1 flowerinsnow.online -s
```

```
bd2b3986-1d4c-54f6-91e3-c02d797310a5

```

<details>

<summary>可以使用预设命名空间</summary>

```shell
secureuuid -n dns -sha1 flowerinsnow.online -s
```

```
bd2b3986-1d4c-54f6-91e3-c02d797310a5

```

</details>

</details>