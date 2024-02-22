import socket
import threading
import socket

middleman_ip = socket.gethostbyname("middleman")
node1_ip = socket.gethostbyname("node1")
node2_ip = socket.gethostbyname("node2")

class P2PHolepuncher:
    def __init__(self, listen_ip, listen_port):
        self.listen_ip = listen_ip
        self.listen_port = listen_port
        # 建立一個 UDP socket
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        # 設定 socket 的參數
        self.sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        # 綁定 socket 到特定的 IP 和 port
        self.sock.bind((self.listen_ip, self.listen_port))

    def run(self):
        while True:
            # 等待來自第一台電腦的請求
            data, addr = self.sock.recvfrom(1024)
            print(addr)
            self.sock.sendto(node2IP.encode(), (node1IP, 8080)) # 收到請求後回傳, node2IP 的位置

            # 如果請求是有效的，則建立一個線程來處理
            if data == b'Hello, world!':
                print("如果請求是有效的，則建立一個線程來處理")
                thread = threading.Thread(target=self.handle_request, args=[addr])
                thread.start()

    def handle_request(self, addr):
        # 向第二台電腦發送一個探測包
        # self.sock.sendto(b'Hello, world!', (addr[0], 8080))
        print("向第二台電腦發送一個探測包", addr[0])
        self.sock.sendto(b'Hello, world!', (node2IP, 8080))

        # 等待來自第二台電腦的回應
        print("等待來自第二台電腦的回應")
        data, node2Addr = self.sock.recvfrom(1024)
        print(data)
        # 如果回應是有效的，則將兩台電腦的 IP 地址和端口號發送給他們
        if data == b'Hello, world!':
            print("如果回應是有效的，則將兩台電腦的 IP 地址和端口號發送給他們")
            self.send_info(addr[0], 8080, node2Addr, 8080)

    def send_info(self, ip1, port1, ip2, port2):
        # 向第一台電腦發送信息
        print("向第一台電腦發送信息", ip1, port1)
        self.sock.sendto(b'ip: {} port: {}'.format(ip2, port2), (ip1, port1))

        # 向第二台電腦發送信息
        print("向第二台電腦發送信息", ip2, port2)
        self.sock.sendto(b'ip: {} port: {}'.format(ip1, port1), (ip2, port2))

if __name__ == "__main__":
    # 建立一個 P2PHolepuncher 實例
    # holepuncher = P2PHolepuncher('127.0.0.1', 8080)
    holepuncher = P2PHolepuncher('0.0.0.0', 8080)

    # 啟動 P2PHolepuncher
    holepuncher.run()

