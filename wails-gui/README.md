# System Tools GUI - Wails Application

这是一个使用 Wails 框架构建的系统工具集合 GUI 应用，整合了多个实用的脚本工具。

## 🛠️ 功能特性

### 1. DNS 缓存清理 (Clear DNS Cache)
- **功能**: 清除系统 DNS 缓存
- **支持平台**: Windows, macOS, Linux
- **对应原脚本**: `clear-dns.bat`

### 2. Git 批量拉取 (Git Pull All)
- **功能**: 批量拉取指定目录下所有 Git 仓库的最新代码
- **输入**: 磁盘盘符（Windows）和工作目录路径
- **对应原脚本**: `git-pull-all.bat`

### 3. 强制结束进程 (Kill Process)
- **功能**: 根据 PID 强制结束进程
- **输入**: 进程 ID (PID)
- **对应原脚本**: `kill-task-force.bat`

### 4. 查看进程信息 (View PID)
- **功能**: 根据 PID 查看进程详细信息
- **输入**: 进程 ID (PID)
- **对应原脚本**: `view-pid-occupied.bat`

### 5. 查看端口占用 (View Port)
- **功能**: 查看指定端口的占用情况
- **输入**: 端口号
- **对应原脚本**: `view-port-occupied.bat`

### 6. Laravel 队列 (Laravel Queue)
- **功能**: 运行 Laravel 定时任务
- **输入**: 磁盘盘符（Windows）和 Laravel 项目路径
- **对应原脚本**: `laravel-queue.bat`

### 7. Docker Skeleton
- **功能**: 获取 Hyperf Skeleton Docker 容器启动命令
- **对应原脚本**: `skeleton.bat`

### 8. Supervisor 管理命令
- **功能**: 显示 Supervisor 进程管理常用命令
- **对应原脚本**: `supervisor.txt`

### 9. Fiddler 图片下载脚本
- **功能**: 显示并复制 Fiddler 自动下载图片的 JScript 脚本
- **对应原脚本**: `fiddler-donw-img.js`

### 10. 系统信息 (System Info)
- **功能**: 显示当前系统信息（操作系统、架构、Go 版本、CPU 数量）

## 📁 项目结构

```
wails-gui/
├── app.go              # Go 后端逻辑（包含所有工具函数）
├── main.go             # Wails 应用入口
├── go.mod              # Go 模块依赖
├── go.sum              # Go 依赖校验
├── wails-gui           # 编译后的可执行文件
├── frontend/
│   └── dist/
│       └── index.html  # 前端界面（标签页式 UI）
└── README.md           # 本说明文档
```

## 🚀 快速开始

### 前置要求
- Go 1.21+
- Node.js 18+ (用于前端构建)
- Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

### 开发模式运行

```bash
cd /workspace/wails-gui
wails dev
```

### 生产环境编译

```bash
cd /workspace/wails-gui
wails build
```

编译后的可执行文件位于 `build/bin/` 目录。

### 直接运行（仅后端测试）

```bash
cd /workspace/wails-gui
go run .
```

## 💡 使用说明

1. **启动应用**后，你会看到一个标签页式的界面
2. **点击不同标签**切换到对应的工具面板
3. **填写必要参数**（如需要）
4. **点击操作按钮**执行相应功能
5. **查看结果**在下方结果显示区域

## 🔧 后端 API 方法

以下是 Go 后端提供的主要方法：

| 方法名 | 参数 | 返回值 | 说明 |
|--------|------|--------|------|
| `ClearDNS()` | 无 | `ScriptToolResult` | 清除 DNS 缓存 |
| `GitPullAll(disk, path)` | disk: 磁盘，path: 路径 | `ScriptToolResult` | Git 批量拉取 |
| `KillTaskForce(pid)` | pid: 进程 ID | `ScriptToolResult` | 强制结束进程 |
| `ViewPIDOccupied(pid)` | pid: 进程 ID | `ScriptToolResult` | 查看进程信息 |
| `ViewPortOccupied(port)` | port: 端口号 | `ScriptToolResult` | 查看端口占用 |
| `RunLaravelQueue(disk, path)` | disk: 磁盘，path: 路径 | `ScriptToolResult` | 运行 Laravel 队列 |
| `RunSkeletonDocker()` | 无 | `ScriptToolResult` | 获取 Docker 命令 |
| `GetSupervisorCommands()` | 无 | `ScriptToolResult` | 获取 Supervisor 命令 |
| `GetFiddlerScript()` | 无 | `ScriptToolResult` | 获取 Fiddler 脚本 |
| `GetSystemInfo()` | 无 | `ScriptToolResult` | 获取系统信息 |

### ScriptToolResult 结构

```go
type ScriptToolResult struct {
    Success bool   `json:"success"`  // 是否成功
    Message string `json:"message"`  // 消息
    Output  string `json:"output"`   // 输出内容
}
```

## ⚙️ 跨平台支持

应用会自动检测操作系统并使用相应的命令：

- **Windows**: 使用 `ipconfig`, `taskkill`, `tasklist`, `netstat` 等
- **macOS**: 使用 `dscacheutil`, `kill`, `ps`, `lsof` 等
- **Linux**: 使用 `systemctl`, `kill`, `ps`, `lsof` 等

## 📝 注意事项

1. **权限要求**: 某些操作（如清除 DNS、结束进程）可能需要管理员/root 权限
2. **Windows 路径**: 在 Windows 上使用 Git Pull 和 Laravel 功能时，需要指定盘符（如 D）
3. **Docker 命令**: Docker Skeleton 功能仅提供命令文本，需要手动在终端执行
4. **Fiddler 脚本**: 需要在 Fiddler 的 Rules 菜单中打开 Script Editor 并粘贴脚本

## 🎨 界面预览

应用采用现代化暗色主题设计，包含：
- 响应式布局
- 标签页导航
- 实时结果反馈
- 代码高亮显示
- 一键复制功能

## 📄 原始脚本来源

本项目整合了以下原始脚本的功能：
- `clear-dns.bat` - DNS 刷新
- `git-pull-all.bat` - Git 批量拉取
- `kill-task-force.bat` - 强制结束进程
- `view-pid-occupied.bat` - 查看 PID 占用
- `view-port-occupied.bat` - 查看端口占用
- `laravel-queue.bat` - Laravel 队列
- `skeleton.bat` - Docker Skeleton
- `supervisor.txt` - Supervisor 命令
- `fiddler-donw-img.js` - Fiddler 图片下载

---

**Built with ❤️ using Wails**
