/*
 * Socket 的技術定義：
 * 
 * Socket 是操作系統提供的一種抽象，用於網絡通信的編程接口（API）
 * 它是應用程序與網絡協議棧之間的橋樑
 */

#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

// === Socket 的核心概念 ===

// 1. Socket 描述符（Socket Descriptor）
// 就像文件描述符一樣，是一個整數，代表一個網絡連接
int sockfd;

// 2. Socket 地址結構
struct sockaddr_in server_addr = {
    .sin_family = AF_INET,           // 地址族：IPv4
    .sin_port = htons(8080),         // 端口號：8080
    .sin_addr.s_addr = INADDR_ANY    // IP 地址：任意地址
};

// === Socket 的基本操作 ===

// 創建 Socket
sockfd = socket(AF_INET,      // 地址族（IPv4）
                SOCK_STREAM,  // Socket 類型（TCP）
                0);           // 協議（默認）

if (sockfd < 0) {
    perror("socket creation failed");
    exit(1);
}

// Socket 就是這個整數 sockfd！
printf("Socket 文件描述符: %d\n", sockfd);

/*
 * 重要理解：
 * - Socket 本身就是一個整數（文件描述符）
 * - 這個整數指向內核中的一個數據結構
 * - 該數據結構包含了網絡連接的所有信息
 */