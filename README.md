# 可易

校园二手交易应用的后端实现

## Features
- 二手交易

## Usage

### Config
You have to export these environment variables.

| Name           | Description     | Value                                               |
|----------------|-----------------|-----------------------------------------------------|
| MODE           | 运行环境            | dev, test, production                               |
| SITE_NAME      | 站点名称            | 可易                                                  |
| DB_URL         | 数据库 URL (mysql) | user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true |
| REDIS_URL      | 如果为空，使用内存缓存     | redis:6379                                          |
| EMAIL_HOST     | EMAIL_HOST      |                                                     |
| EMAIL_PORT     | EMAIL_PORT      |                                                     |
| EMAIL_USER     | 也是默认的发件人        |                                                     |
| EMAIL_PASSWORD | EMAIL_PASSWORD  |                                                     |

### Build
```shell
git clone https://github.com/hasbai/keyi.git
cd keyi
# install swag and generate docs
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseInternal --parseDependency --parseDepth 1 # to generate the latest docs, this should be run before compiling
# build and run
go build -o keyi.exe
./keyi.exe
```

### Test
Please export `MODE=test`

### API Docs
Please visit http://localhost:8000/docs after running app

## Badge

[//]: # ([![build]&#40;https://github.com/hasbai/keyi/actions/workflows/master.yaml/badge.svg&#41;]&#40;https://github.com/hasbai/keyi/actions/workflows/master.yaml&#41;)
[//]: # ([![dev build]&#40;https://github.com/hasbai/keyi/actions/workflows/dev.yaml/badge.svg&#41;]&#40;https://github.com/hasbai/keyi/actions/workflows/dev.yaml&#41;)

[![stars](https://img.shields.io/github/stars/hasbai/keyi)](https://github.com/hasbai/keyi/stargazers)
[![issues](https://img.shields.io/github/issues/hasbai/keyi)](https://github.com/hasbai/keyi/issues)
[![pull requests](https://img.shields.io/github/issues-pr/hasbai/keyi)](https://github.com/hasbai/keyi/pulls)

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

### Powered by

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/hasbai/keyi/issues/new) or [Submit PRs](https://github.com/hasbai/keyi/compare).

We are now in rapid development, any contribution would be of great help. 
For the developing roadmap, please visit [this issue](https://github.com/hasbai/keyi/issues/1).

### Contributors

This project exists thanks to all the people who contribute.

<a href="https://github.com/hasbai/keyi/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=hasbai/keyi"  alt="contributors"/>
</a>

## Licence

[![license](https://img.shields.io/github/license/hasbai/keyi)](https://github.com/hasbai/keyi/blob/master/LICENSE)
© hasbai
