import socket
import time


class P2PHolePunching:
    def __init__(self, host, port, server_host, server_port):
        self.host = host
        self.port = port
        self.server_host = server_host
        self.server_port = server_port

        self.udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        self.udp_socket.bind((self.host, self.port))

        self.tcp_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def start(self):
        # 向服務器發送請求
        self.send_request_to_server()

        # 等待對方回應
        self.wait_for_response()

        # 建立 TCP 連接
        self.connect_to_peer()

    def send_request_to_server(self):
        message = "request"
        self.udp_socket.sendto(message.encode(), (self.server_host, self.server_port))

    def wait_for_response(self):
        while True:
            data, address = self.udp_socket.recvfrom(1024)
            if data.decode() == "response":
                break

        self.peer_host = address[0]
        self.peer_port = address[1]

    def connect_to_peer(self):
        self.tcp_socket.connect((self.peer_host, self.peer_port))

        print("Connected to peer:", self.peer_host, self.peer_port)

        while True:
            # 發送數據
            message = input("Enter message: ")
            self.tcp_socket.send(message.encode())

            # 接收數據
            data = self.tcp_socket.recv(1024)
            print("Received message:", data.decode())


if __name__ == "__main__":
    middleman_ip = socket.gethostbyname("middleman")
    node2_ip = socket.gethostbyname("node2")

    # 創建兩個 P2P 打洞實例
    peer_a = P2PHolePunching(middleman_ip, 8080, node2_ip, 8081)
    peer_b = P2PHolePunching(node2_ip, 8081, middleman_ip, 8080)

    # 啟動兩個 P2P 打洞實例
    peer_a.start()
    peer_b.start()

    # 讓兩個 P2P 打洞實例運行
    while True:
        time.sleep(1)
