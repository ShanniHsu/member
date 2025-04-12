# 使用官方Golang映像
FROM golang:1.24.2 AS builder

# 設定工作目錄為/app,接下來的指令(COPY、RUN)都會以/app為根目錄
WORKDIR /app

# 複製go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

COPY config.json .

# 複製程式碼,複製到container的/app目錄
COPY . .

# 編譯 使用go build生成main的執行檔
RUN go build -o main

# 使用更小的映像來部署 (第二階段,不含go compiler,只放執行用的東西) (multi-stage build)
FROM debian:bookworm-slim

WORKDIR /root/

# 從第一階段builder把main可執行檔複製進來
COPY --from=builder /app/main .

COPY --from=builder /app/config.json .

# 開放container的8080port (不是實際開放,只是聲明)
EXPOSE 8080

# 啟動container時執行的命令
CMD ["./main"]