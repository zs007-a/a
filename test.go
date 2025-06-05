
package main

import (
	"fmt"
	"net/http"
)

// 自定义请求处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你好，世界！🎉\n请求路径: %s", r.URL.Path)
}

func main() {
	// 注册路由处理函数
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status": "OK", "message": "API 工作正常"}`)
	})

	// 配置服务器
	port := ":8080"
	fmt.Printf("🚀 服务器已启动，监听端口 %s\n", port)
	fmt.Println("👉 访问地址: http://localhost" + port)
	fmt.Println("🔍 测试端点:")
	fmt.Println("   - 首页: http://localhost:8080")
	fmt.Println("   - API: http://localhost:8080/api")

	// 启动服务器
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("❌ 服务器启动失败: %v\n", err)
	}
}