# dbdog
dbdog is a table definition automatic generation tool!

![screenshot](https://user-images.githubusercontent.com/7962629/86523622-f527d400-bea9-11ea-8622-ea740aff6d7e.png)

## How to use
### install
go get
```
go get -u github.com/IkezoeMakoto/dbdog
```
or
binary download
```
wget "https://github.com/IkezoeMakoto/dbdog/releases/download/v0.0.2/dbdog_$(uname -s)-$(uname -m).zip"
unzip dbdog_linux_amd64.zip
mv dbdog /usr/local/bin
```

### config
your db setting
```
vim config.toml
```

### use
```
dbdog
##########################################################
   ___  ___     __
  / _ \/ _ )___/ /__  ___ _
 / // / _  / _  / _ \/ _ `/
/____/____/\_,_/\___/\_, /
                    /___/
##########################################################
# dbdog is a table definition automatic generation tool! #
```

## output
[output-example.md](output-example.md)