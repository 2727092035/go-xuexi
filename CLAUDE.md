# Go 学习项目

## 项目用途

前端开发者学习 Go 后端开发的练习项目，包含每个阶段的学习代码和笔记。

## 目录结构

```
go/
├── CLAUDE.md          # 项目规范（本文件）
├── PROGRESS.md        # 学习进度追踪
├── LEARNING_PLAN_60_DAYS.md # 60 天总学习计划
├── phase1-basics/     # 第一阶段：语法基础强化（Day 1-10）
├── phase2-web/        # 第二阶段：Web API 与数据库（Day 11-25）
├── phase3-engineering/# 第三阶段：工程化与中间件（Day 26-40）
├── phase4-advanced/   # 第四阶段：综合项目实战（Day 41-60）
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

## 每日教学内容要求

从 Day 6 开始，后续每天的学习材料要偏“教学型”，不能只有能运行的 demo。每天至少覆盖：

- `demo/main.go`：用分段示例展示当天语法，每段示例前用中文注释说明“这个语法解决什么问题、核心写法是什么、输出说明什么”。
- `demo/main.go` 里只要出现新语法、新标准库函数、新方法调用或新写法，都必须就近解释清楚，不能默认初学者知道。例如 `strings.TrimSpace(raw)` 要说明 `strings` 是标准库包、`TrimSpace` 会去掉字符串前后空白、`raw` 是传入的原始字符串、返回值是清理后的新字符串。
- `notes.md`：必须包含语法说明、代码逐段详解、常见错误、和 JS/TS 的差异、练习验收标准。
- `notes.md` 对新增语法必须写到“用途、参数、返回值、为什么这里要用、如果不用会怎样”。不能只写一句“用于处理字符串”。
- `exercise/main.go`：练习题要从基础到组合递进，不能只给最终答案；复杂题要拆成小函数，方便复盘。
- 如果当天进入项目主题，也要先解释相关语法和标准库 API，再把它放进项目代码里。

`notes.md` 推荐结构：

1. 今天先记住一句话。
2. 语法核心说明。
3. demo 代码逐段详解。
4. 和 JS/TS 的关键差异。
5. 容易踩坑。
6. 练习任务和验收标准。
7. 完成记录和疑问。

## 约定

- 每次学习在对应阶段目录下创建 `dayXX/` 子目录。
- 每个 day 目录必须包含 `demo/`、`exercise/` 和 `notes.md`。
- 学习完成后更新 `PROGRESS.md` 记录进度。
- 代码文件用英文命名，注释可以用中文。
- 示例和练习都要保持 `gofmt` 格式化。
- 后续新增或重写 demo 时，优先加强语法教学和逐段说明，而不是只追求代码短。
- 新人第一次遇到的语法或 API，必须按“来源/用途/参数/返回值/示例输出”讲清楚，再进入组合代码。

## 工作流

1. 开始学习 → 在 `PROGRESS.md` 中标记当天开始。
2. 看示例 → 运行 `dayXX/demo/main.go`，对照注释理解每个语法点。
3. 读详解 → 阅读 `dayXX/notes.md` 的语法说明和 demo 逐段解析。
4. 写练习 → 完成 `dayXX/exercise/main.go`。
5. 记录复盘 → 更新 `dayXX/notes.md`。
6. 学习结束 → 更新 `PROGRESS.md` 状态为已完成，记录收获和疑问。
