package models

import (
	"log/slog"

	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 雪花算法生成id
var Node *snowflake.Node

func InitDB() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3309)/sw?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("数据库连接失败")
		return
	}
	Node, err = snowflake.NewNode(1)
	if err != nil {
		slog.Error("雪花算法初始化失败")
		return
	}

	g := gen.NewGenerator(gen.Config{
		// 生成文件的输出路径
		OutPath: "./domain/query",
		// 生成的 Model 包名，这个包放在query所在父级目录下
		ModelPkgPath: "po",

		// 生成模式: 这里选择要不要生成查询函数
		Mode: gen.WithoutContext | gen.WithQueryInterface,

		// 以下配置控制生成 Model 的字段属性 (可选)
		FieldNullable:     true,  // 数据库字段可空时，生成指针类型
		FieldCoverable:    false, // 字段有默认值时，生成指针类型以处理零值问题
		FieldSignable:     false, // 检测整数字段是否为无符号类型
		FieldWithIndexTag: true,  // 生成 gorm index 标签
		FieldWithTypeTag:  true,  // 生成 gorm type 标签
	})
	// 3. 使用数据库连接
	g.UseDB(DB)

	// 4. 指定要生成代码的表
	// 方式A: 生成所有表
	g.ApplyBasic(g.GenerateAllTable()...)

	// 方式B: 生成指定的表，并可以自定义模型 (推荐)
	// 后面写哪个表
	// userModel := g.GenerateModel("swuser")

	// 方式C: 生成表时进行更精细的控制
	// g.GenerateModel("products",
	//     gen.FieldType("price", "decimal.Decimal"), // 自定义字段类型
	//     gen.FieldIgnore("deleted_at"),             // 忽略某个字段
	// )

	// 将生成的结构体写入文件
	// g.ApplyBasic(userModel, g.GenerateModel("book"))

	// 执行生成
	g.Execute()
	slog.Info("数据库初始化完成")
}
