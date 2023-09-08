#!/usr/bin/env python3
"""
Very simple TCP socket in python
Usage::
    ./tcp-server.py [<port>]
"""

import socket

def main(port=80):
    # 匯入 socket 模組
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    # 設定 socket 的參數
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

    # 綁定 socket 到特定的 IP 和 port
    sock.bind(('0.0.0.0', port))
    print(f"bind the sock: 0.0.0.0 {port}")

    # 開始監聽連線
    sock.listen(1)

    # 接受連線
    conn, addr = sock.accept()

    # 印出 client IP
    print(addr[0])

    # 關閉連線
    conn.close()

if __name__ == "__main__":
    from sys import argv
    if len(argv) == 2:
        main(port=int(argv[1]))
    else:
        main()
