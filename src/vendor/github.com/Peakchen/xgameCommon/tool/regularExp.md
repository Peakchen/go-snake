# golang 正则用法

    text := `Hello 世界！123 Go.`
    reg := regexp.MustCompile(`[a-z]+`)  // 查找连续的小写字母
    fmt.Printf("%q\n", reg.FindAllString(text, -1)    // 输出结果["ello" "o"]

    reg = regexp.MustCompile(`[^a-z]+`)     // 查找连续的非小写字母
    fmt.Printf("%q\n", reg.FindAllString(text, -1))     // ["H" " 世界！123 G" "."]

    reg = regexp.MustCompile(`[\w]+`)   // 查找连续的单词字母
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello" "123" "Go"]

    reg = regexp.MustCompile(`[^\w\s]+`)  // 查找连续的非单词字母、非空白字符
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["世界！" "."]

    reg = regexp.MustCompile(`[[:upper:]]+`)   // 查找连续的大写字母
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["H" "G"]

    reg = regexp.MustCompile(`[[:^ascii:]]+`)   // 查找连续的非 ASCII 字符
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["世界！"]

    reg = regexp.MustCompile(`[\pP]+`)   // 查找连续的标点符号
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["！" "."]

    reg = regexp.MustCompile(`[\PP]+`)   // 查找连续的非标点符号字符
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello 世界" "123 Go"]

    reg = regexp.MustCompile(`[\p{Han}]+`)   // 查找连续的汉字
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["世界"]

    reg = regexp.MustCompile(`[\P{Han}]+`)   // 查找连续的非汉字字符
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello " "！123 Go."]

    reg = regexp.MustCompile(`Hello|Go`)   // 查找 Hello 或 Go
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello" "Go"]

    reg = regexp.MustCompile(`^H.*\s`)   // 查找行首以 H 开头，以空格结尾的字符串
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello 世界！123 "]

    reg = regexp.MustCompile(`(?U)^H.*\s`)   // 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello "]

    reg = regexp.MustCompile(`(?i:^hello).*Go`)   // 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello 世界！123 Go"]

    reg = regexp.MustCompile(`\QGo.\E`)   // 查找 Go.
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Go."]

    reg = regexp.MustCompile(`(?U)^.* `)   // 查找从行首开始，以空格结尾的字符串（非贪婪模式）
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello "]

    reg = regexp.MustCompile(` [^ ]*$`)   // 查找以空格开头，到行尾结束，中间不包含空格字符串
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // [" Go."]

    reg = regexp.MustCompile(`(?U)\b.+\b`)   // 查找“单词边界”之间的字符串
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello" " 世界！" "123" " " "Go"]

    reg = regexp.MustCompile(`[^ ]{1,4}o`)   // 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello" "Go"]

    reg = regexp.MustCompile(`(?:Hell|G)o`)   // 查找 Hello 或 Go
    fmt.Printf("%q\n", reg.FindAllString(text, -1))   // ["Hello" "Go"]

    reg = regexp.MustCompile(`(?PHell|G)o`)   // 查找 Hello 或 Go，替换为 Hellooo、Gooo
    fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))   // "Hellooo 世界！123 Gooo."

    reg = regexp.MustCompile(`(Hello)(.*)(Go)`)   // 交换 Hello 和 Go
    fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))   // "Go 世界！123 Hello."

    reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|]`)
    fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))