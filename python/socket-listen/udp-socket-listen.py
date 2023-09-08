#!/usr/bin/env python3
"""
Very simple UDP socket in python
Usage::
    ./udp-server.py [<port>]
"""

import socket

def main(port=80):
    # 建立 UDP 套接字
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    # 綁定套接字到指定的端口
    sock.bind(("0.0.0.0", port))
    print(f"bind the sock: 0.0.0.0 {port}")

    # 進入循環，等待接收 UDP 數據包
    while True:
        # 接收 UDP 數據包
        data, addr = sock.recvfrom(1024)

        # 解析數據包，並印出 client IP
        client_ip = addr[0]
        print(f"Received UDP data from {client_ip}")

if __name__ == "__main__":
    from sys import argv
    if len(argv) == 2:
        main(port=int(argv[1]))
    else:
        main()
