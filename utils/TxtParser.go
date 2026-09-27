package utils

import (
	"bufio"
	"log/slog"
	"mime/multipart"
	"regexp"
	"strings"
)

// 文本解析器，用于将整篇小说分成一个个章节并存入数据库
func TxtParserAndSave(file *multipart.File) {
	scanner := bufio.NewScanner(*file)
	re, err := regexp.Compile(`^(第[零一二三四五六七八九十百千万]+[章回]|楔子|第\d+章)\s*[.、：:：\s]*(.*)`)
	if err != nil {
		slog.Info("正则表达式不合法")
		return
	}
	// var f *os.File
	// defer func() {
	// 	err := f.Close()
	// 	if err != nil {
	// 		slog.Info("最后关闭", "error", err.Error())
	// 		return
	// 	}
	// }()
	// write := true
	for scanner.Scan() {
		slog.Info("开始按行读取")
		str := scanner.Text()
		if strings.TrimSpace(str) == "" {
			continue
		}
		if re.MatchString(str) {
			// if f != nil {
			// 	err := f.Close()
			// 	if err != nil {
			// 		slog.Info("f关闭失败")
			// 		return
			// 	}
			// 	slog.Info("f关闭成功")
			// 	f = nil
			// }
			slog.Info("匹配成功 当前行是标题")
			str = strings.ReplaceAll(str, "?", " ")
			str = strings.ReplaceAll(str, ":", " ")
			// var ferr error
			// f, ferr = os.OpenFile("./txt1/"+str+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			// if ferr != nil {
			// 	slog.Info("文件打开或写入失败!", "error", ferr)
			// 	return
			// }
			// write = false
		}
		// if f != nil && write {
		// 	slog.Info("开始写入")
		// 	f.WriteString(str + "\n")
		// }
		// write = true
	}
	if err := scanner.Err(); err != nil {
		slog.Info(err.Error())
	}
}
