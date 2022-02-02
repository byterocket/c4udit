一鍵更新neko relay域名

因爲neko的国内域名全部被DNS污染，现只提供IP，可以使用本項目配合cron檢測域名是否指向最新IP並更新最新的IP

# 使用方法
1. 安裝依賴

`pip3 install requests dnspython`

2. 編輯`neko-ddns.py`:
```
res.nameservers = ['1.1.1.1']
```
指定DNS Resolver, 默認值爲`1.1.1.1`

```
token = "<YOUR NEKO API TOKEN HERE>" 
name  = "<需要檢測的線路名字>如<广港链路A>"
```
前往[https://relay.nekoneko.cloud/user/setting](https://relay.nekoneko.cloud/user/setting)獲取API密鑰

```
cf_auth_email = "" # The email used to login 'https://dash.cloudflare.com'
cf_auth_key = "" # Your API Token or Global API Key 
```
前往 [Cloudfalre - Profile - API Tokens](https://dash.cloudflare.com/profile/api-tokens)使用Edit zone DNS模板並在Zone Resources指定你的域名的zone (出於安全考慮，請不要使用global api)

```
cf_auth_method = "token" # Set to "global" for Global API Key or "token" for Scoped API Token
cf_zone_identifier = "" # Can be found in the "Overview" tab of your domain, 在域名的「Overview」頁面中找到域名的Zone ID
cf_record_name = "" # Which record you want to be synced
```

3. 定時執行

使用`crontab -e`添加新規則即可
`30 * * * * /path/to/neko-ddns.py`


<h1 align=center><code>c4udit</code></h1>

## Introduction

`c4udit` is a static analyzer for solidity contracts based on regular
expressions specifically crafted for [Code4Rena](https://code4rena.com) contests.

It is capable of finding low risk issues and gas optimization documented in the
[c4-common-issues](https://github.com/byterocket/c4-common-issues) repository.

Note that `c4udit` uses [c4-common-issues](https://github.com/byterocket/c4-common-issues)'s issue identifiers.

## Installation

First you need to have the Go toolchain installed. You can find instruction [here](https://go.dev/doc/install).

Then install `c4udit` with:
```
$ go install github.com/byterocket/c4udit@latest
```

To just build the binary:
```
$ git clone https://github.com/byterocket/c4udit
$ cd c4udit/
$ go build .
```
Now you should be able to run `c4udit` with:
```
$ ./c4udit
```

## Usage

```
Usage:
	c4udit [flags] [files...]

Flags:
	-h    Print help text.
	-s    Save report as file.
```

## Example

Running `c4udit` against the `examples` directory:
```
$ ./c4udit examples/
Files analyzed:
- examples/Test.sol

Issues found:
 G001:
  examples/Test.sol::4 => uint256 a = 0;
  examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {

 G002:
  examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {

 G007:
  examples/Test.sol::6 => string b = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

 G008:
  examples/Test.sol::13 => i = i / 2;

 L001:
  examples/Test.sol::16 => token.transferFrom(msg.sender, address(this), 100);

 L003:
  examples/Test.sol::1 => pragma solidity ^0.8.0;
```

Using the `-s` flag, `c4udit` will create a report in markdown format.
For an example check out the report in the `examples` directory [here](./examples/c4udit-report.md).


## License

Note that this tool is licensed as [free software](./LICENSE)!
