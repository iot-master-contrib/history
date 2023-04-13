
# 整体编译
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com
go env -w GOSUMDB=off

app="history"
version="1.0.0"

read -t 5 -p "please input version(default:$version)" ver
if [ -n "${ver}" ];then
	version=$ver
fi


pkg="github.com/zgwit/iot-master/pkg/build"
gitHash=$(git show -s --format=%H)
buildTime=$(date -d today +"%Y-%m-%d %H:%M:%S")

# -w -s
ldflags="-X '$pkg.Version=$version' \
-X '$pkg.GitHash=$gitHash' \
-X '$pkg.BuildTime=$buildTime'"


export GOARCH=amd64

export GOOS=windows
go build -ldflags "$ldflags" -o history.exe main.go

export GOOS=linux
#CC=gcc
go build -ldflags "$ldflags" -o history main.go