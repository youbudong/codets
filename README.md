# Apifox生成Typescipt类型

使用时当前目录添加apifox.yml文件
```yaml
APIFOX_PROJCET:
  - id: 123        # 项目ID
    name: user     # 生成的文件名
  - id: 1234 
    name: message
APIFOX_TOKEN: xxx  # Apifox Token
OUTPUT_DIR: app/types # 输出目录
```

生成类型文件 `./codets-windows-amd64.exe`

w