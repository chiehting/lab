package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	webview "github.com/webview/webview_go"
	"howett.net/plist"
)

// InfoPlist read Info.plist structure
type InfoPlist struct {
	DefaultHomePage string `plist:"DefaultHomePage"`
}

func main() {
	var homeURL string

	// 檢查是否從 URL Scheme 啟動
	if len(os.Args) > 1 {
		// 解析 URL scheme (例如: browsertool://open?url=https://example.com)
		if strings.HasPrefix(os.Args[1], "browsertool://") {
			u, err := url.Parse(os.Args[1])
			if err == nil {
				params := u.Query()
				if targetURL := params.Get("url"); targetURL != "" {
					homeURL = targetURL
				}
			}
		} else {
			// 直接傳入 URL
			homeURL = os.Args[1]
		}
	}

	// 如果沒有從參數獲得 URL，嘗試從 Info.plist 讀取
	if homeURL == "" {
		homeURL = getDefaultHomePageFromPlist()
	}

	// 如果還是沒有，詢問使用者
	if homeURL == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("請輸入要瀏覽的網址: ")
		homeURL, _ = reader.ReadString('\n')
		homeURL = strings.TrimSpace(homeURL)
	}

	// 確保網址有正確的協議
	if !strings.HasPrefix(homeURL, "http://") && !strings.HasPrefix(homeURL, "https://") {
		homeURL = "https://" + homeURL
	}

	// 建立 webview
	w := webview.New(true)
	defer w.Destroy()

	// 設定視窗
	w.SetTitle("Go 瀏覽器工具")
	w.SetSize(1280, 800, webview.HintNone)

	// 注入工具列
	w.Init(`
		// 建立工具列
		const toolbar = document.createElement('div');
		toolbar.id = 'custom-toolbar';
		toolbar.style.cssText = 'position: fixed; top: 0; left: 0; width: 100%; height: 50px; background: linear-gradient(to bottom, #f8f9fa, #e9ecef); border-bottom: 1px solid #dee2e6; z-index: 9999; display: flex; align-items: center; padding: 0 15px; gap: 10px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);';
		
		// 建立按鈕樣式
		const buttonStyle = 'padding: 8px 16px; background-color: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; font-size: 14px; font-weight: 500; transition: all 0.2s;';
		
		// 建立返回按鈕
		const backButton = document.createElement('button');
		backButton.innerHTML = '← 返回';
		backButton.style.cssText = buttonStyle;
		backButton.onclick = function() { window.history.back(); };
		
		// 建立前進按鈕
		const forwardButton = document.createElement('button');
		forwardButton.innerHTML = '前進 →';
		forwardButton.style.cssText = buttonStyle;
		forwardButton.onclick = function() { window.history.forward(); };
		
		// 建立 Home 按鈕
		const homeButton = document.createElement('button');
		homeButton.innerHTML = '🏠 首頁';
		homeButton.style.cssText = buttonStyle;
		homeButton.onclick = function() { window.goHome(); };
		
		// 建立重新整理按鈕
		const refreshButton = document.createElement('button');
		refreshButton.innerHTML = '🔄 重新整理';
		refreshButton.style.cssText = buttonStyle;
		refreshButton.onclick = function() { window.location.reload(); };
		
		// 建立網址列
		const urlBar = document.createElement('input');
		urlBar.type = 'text';
		urlBar.style.cssText = 'flex: 1; padding: 8px 12px; border: 1px solid #ced4da; border-radius: 4px; font-size: 14px; margin: 0 10px;';
		urlBar.value = window.location.href;
		urlBar.onkeypress = function(e) {
			if (e.key === 'Enter') {
				let url = this.value;
				if (!url.startsWith('http://') && !url.startsWith('https://')) {
					url = 'https://' + url;
				}
				window.location.href = url;
			}
		};
		
		// 監聽頁面變化更新網址列
		setInterval(() => {
			if (urlBar.value !== window.location.href) {
				urlBar.value = window.location.href;
			}
		}, 100);
		
		// 將元素加入工具列
		toolbar.appendChild(backButton);
		toolbar.appendChild(forwardButton);
		toolbar.appendChild(homeButton);
		toolbar.appendChild(refreshButton);
		toolbar.appendChild(urlBar);
		
		// 調整頁面內容的位置
		document.body.style.marginTop = '50px';
		
		// 將工具列加入頁面
		document.body.appendChild(toolbar);
		
		// 按鈕懸停效果
		[backButton, forwardButton, homeButton, refreshButton].forEach(btn => {
			btn.onmouseover = function() { this.style.opacity = '0.8'; };
			btn.onmouseout = function() { this.style.opacity = '1'; };
		});
	`)

	// 綁定函數
	w.Bind("goHome", func() {
		w.Navigate(homeURL)
	})

	// 導航到首頁
	w.Navigate(homeURL)

	// 執行
	w.Run()
}

// 從 Info.plist 讀取預設首頁
func getDefaultHomePageFromPlist() string {
	// 獲取執行檔路徑
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}

	// 構建 Info.plist 路徑
	// 如果是 .app bundle，Info.plist 在 ../Info.plist
	plistPath := filepath.Join(filepath.Dir(execPath), "..", "Info.plist")

	// 讀取 plist 檔案
	file, err := os.Open(plistPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	var info InfoPlist
	decoder := plist.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		return ""
	}

	return info.DefaultHomePage
}
