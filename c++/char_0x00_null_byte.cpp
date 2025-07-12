#include <iomanip>
#include <sstream>
#include <iostream>

using namespace std;

/**
 * @file null_check.cpp
 * @brief 測試在字元陣列中遇到 0x00 (null byte) 時的行為。
 *
 * 此程式碼建立一個包含 0x00 的 char 陣列，並嘗試將其內容以十六進位格式輸出。
 * 迴圈在遇到 0x00 時會停止，藉此觀察 null byte 對於字元處理的影響。
 *
 * 用於測試和說明 C/C++ 中 null byte (0x00) 在字串處理時的特殊意義。
 */
int main() {
    // char buf[]= "Hello world";
    char buf[5]= {0x01, 0x02, 0x00, 0x03, 0x04};
    size_t size = sizeof(buf) - 1; // exclude null terminator

    auto ptr = buf;

    std::ostringstream oss;
    for (const char* p = ptr; *p && p < buf + size; ++p) {
        oss << std::hex << std::setw(2) << std::setfill('0') << (static_cast<unsigned int>(static_cast<unsigned char>(*p))) << " ";
        std::cout << *p << std::endl;
    }
    std::cout << oss.str() << std::endl;
    std::cout << size << std::endl;

    return 0;
}
