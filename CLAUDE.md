# Go 学习项目

## 项目用途

前端开发者学习 Go 后端开发的练习项目，包含每个阶段的学习代码和笔记。

## 目录结构

```
go/
├── CLAUDE.md          # 项目规范（本文件）
├── PROGRESS.md        # 学习进度追踪
├── phase1-basics/     # 第一阶段：语法基础（Day 1-14）
├── phase2-web/        # 第二阶段：Web 开发（Day 15-35）
├── phase3-engineering/# 第三阶段：工程化（Day 36-56）
├── phase4-advanced/   # 第四阶段：进阶实战（Day 57-90）
└── projects/          # 完整项目（博客 API、短链接等）
```

## Day 目录约定

每个 day 目录按下面结构组织，避免同一个 package 下出现多个 `func main()`：

```
dayXX/
├── demo/main.go       # 当天示例代码
├── exercise/main.go   # 当天练习代码
└── notes.md           # 当天笔记和复盘
```

运行方式：

```bash
go run phase1-basics/dayXX/demo/main.go
go run phase1-basics/dayXX/exercise/main.go
go test ./...
```

## 约定

- 每次学习在对应阶段目录下创建 `dayXX/` 子目录。
- 每个 day 目录必须包含 `demo/`、`exercise/` 和 `notes.md`。
- 学习完成后更新 `PROGRESS.md` 记录进度。
- 代码文件用英文命名，注释可以用中文。
- 示例和练习都要保持 `gofmt` 格式化。

## 工作流

1. 开始学习 → 在 `PROGRESS.md` 中标记当天开始。
2. 看示例 → 运行 `dayXX/demo/main.go`。
3. 写练习 → 完成 `dayXX/exercise/main.go`。
4. 记录复盘 → 更新 `dayXX/notes.md`。
5. 学习结束 → 更新 `PROGRESS.md` 状态为已完成，记录收获和疑问。
