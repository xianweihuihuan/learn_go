## HTTP 服务器端常用函数

| 函数/方法 | 作用 | 示例 |
|----------|------|-------|
| http.ListenAndServe() | 启动 HTTP 服务器 | `http.ListenAndServe(":8080", nil)` |
| http.ListenAndServeTLS() | 启动 HTTPS 服务器 | `http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)` |
| http.HandleFunc() | 注册路由处理函数 | `http.HandleFunc("/path", handler)` |
| http.Handle() | 注册路由处理对象 | `http.Handle("/path", handler)` |
| http.FileServer() | 创建静态文件服务器 | `http.FileServer(http.Dir("./static"))` |
| http.ServeFile() | 服务单个文件 | `http.ServeFile(w, r, "file.txt")` |
| http.Error() | 返回错误响应 | `http.Error(w, "错误信息", 500)` |
| http.NotFound() | 返回 404 响应 | `http.NotFound(w, r)` |
| http.Redirect() | 重定向 | `http.Redirect(w, r, "/new", 301)` |
| http.SetCookie() | 设置 Cookie | `http.SetCookie(w, &cookie)` |
---
## HTTP 客户端常用函数

| 函数/方法 | 作用 | 示例 |
|-----------|------|---------|
| http.Get() | 发送 GET 请求 | `resp, err := http.Get(url)` |
| http.Post() | 发送 POST 请求 | `http.Post(url, "application/json", body)` |
| http.PostForm() | 发送表单数据 | `http.PostForm(url, formData)` |
| http.NewRequest() | 创建请求 | `req, err := http.NewRequest("GET", url, nil)` |
| http.NewRequestWithContext() | 创建带 Context 的请求 | `http.NewRequestWithContext(ctx, "GET", url, nil)` |
---
## Request 方法

| 方法 | 作用 | 示例 |
|------|------|---------|
| r.URL.Query().Get() | 获取查询参数 | `name := r.URL.Query().Get("name")` |
| r.FormValue() | 获取表单值 | `username := r.FormValue("username")` |
| r.ParseForm() | 解析表单数据 | `r.ParseForm()` |
| r.Header.Get() | 获取请求头 | `ua := r.Header.Get("User-Agent")` |
| r.Cookie() | 获取 Cookie | `cookie, err := r.Cookie("session")` |
| r.Cookies() | 获取所有 Cookie | `cookies := r.Cookies()` |
| r.BasicAuth() | 获取基本认证 | `user, pass, ok := r.BasicAuth()` |
---
## ResponseWriter 方法

| 方法 | 作用 | 示例 |
|-------|------|---------|
| w.Header().Set() | 设置响应头 | `w.Header().Set("Content-Type", "application/json")` |
| w.WriteHeader() | 发送状态码 | `w.WriteHeader(200)` |
| w.Write() | 写入字节 | `w.Write([]byte("data"))` |
| fmt.Fprintf(w, ...) | 格式化写入 | `fmt.Fprintf(w, "Hello %s", name)` |
| io.WriteString(w, ...) | 写入字符串 | `io.WriteString(w, "text")` |
---
## 响应处理方法

| 方法 | 作用 | 示例 |
|-------|------|---------|
| resp.StatusCode | 获取状态码 | `if resp.StatusCode == 200 { ... }` |
| resp.Header.Get() | 获取响应头 | `ct := resp.Header.Get("Content-Type")` |
| io.ReadAll(resp.Body) | 读取响应内容 | `body, _ := io.ReadAll(resp.Body)` |
| json.NewDecoder().Decode() | JSON 解码 | `json.NewDecoder(resp.Body).Decode(&data)` |
| resp.Body.Close() | 关闭响应体 | `defer resp.Body.Close()` |
