import socket
import threading
import time

middlemanIP = socket.gethostbyname("middleman")

def main():
    # 建立一個 UDP socket
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    # 設定 socket 的參數
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    # 綁定 socket 到特定的 IP 和 port
    sock.bind(('0.0.0.0', 8080))
    # 等待來自中間人的請求
    print("等待來自中間人的請求")
    data, addr = sock.recvfrom(1024)

    # 如果請求是有效的，則回應中間人
    if data == b'Hello, world!':
        print("如果請求是有效的，則回應中間人", addr[0], addr[1])
        sock.sendto(b'Hello, world!', (addr[0], addr[1]))
        sock.sendto(b'Hello, world!', (addr[0], addr[1]))
        sock.sendto(b'Hello, world!', (addr[0], addr[1]))

if __name__ == "__main__":
    main()
