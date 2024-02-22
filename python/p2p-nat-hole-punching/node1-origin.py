import socket
import threading
import time

def main():
    # 建立一個 UDP socket
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    # 設定 socket 的參數
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

    # 綁定 socket 到特定的 IP 和 port
    sock.bind(('0.0.0.0', 8080))

    # 生成一個隨機的 port
    my_port = random.randint(10000, 20000)

    # 建立一個線程來發送探測包
    thread = threading.Thread(target=send_probe, args=[sock, my_port])
    thread.start()

    # 等待另一台電腦的回應
    while True:
        data, addr = sock.recvfrom(1024)
        if data == b'ip: {} port: {}'.format(ip2, port2):
            print('收到回應')
            print(addr[0])
            print(data)
            break

def send_probe(sock, my_port):
    # 向中間人發送一個探測包
    sock.sendto(b'Hello, world!', ('127.0.0.1', 8080))

    # 等待 1 秒
    time.sleep(1)

if __name__ == "__main__":
    main()
