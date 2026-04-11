package chat

type Role string

const (
	// RoleSystem 系统角色（最高优先级）
	// 用于定义 AI 的全局行为、规则、风格和约束（类似“隐藏提示词”）
	// 通常在对话开始时设置，对后续所有轮次生效
	RoleSystem Role = "system"

	// RoleDeveloper 开发者角色（新规范）
	// 用于开发者注入的指令，优先级通常低于 system、高于 user
	// 适合做中间层控制（如产品逻辑、策略、调试信息等）
	RoleDeveloper Role = "developer"

	// RoleUser 用户角色
	// 表示真实用户输入的内容，是对话的主要驱动力（提问 / 需求 / 指令）
	RoleUser Role = "user"

	// RoleAssistant 助手角色（AI）
	// 表示模型生成的回复内容，是最终返回给用户的结果
	// 同时也会作为上下文参与后续对话
	RoleAssistant Role = "assistant"

	// RoleTool 工具角色（推荐使用）
	// 表示外部工具/函数执行后的返回结果（如数据库查询、API调用等）
	// 用于实现 AI 调用外部能力（RAG / 插件 / Agent）
	RoleTool Role = "tool"
)
