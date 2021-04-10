package ulog

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 从日志文本中匹配规则, 并着色
type ColorRule struct {
	Text  string
	Color string

	def *ColorDefine
}

type ColorRuleSet struct {
	Rule []*ColorRule
}

// 在规则中与目标文本匹配, 返回对应的颜色
func (self *ColorRuleSet) MatchText(text string) *ColorDefine {

	for _, rule := range self.Rule {
		if strings.Contains(text, rule.Text) {
			return rule.def
		}
	}

	return nil
}

// 从规则文本中加载规则
func (self *ColorRuleSet) Parse(ruleText string) error {

	err := json.Unmarshal([]byte(ruleText), self)
	if err != nil {
		return err
	}

	for _, rule := range self.Rule {

		rule.def = GetColorDefineByName(rule.Color)

		if rule.def == nil {
			return fmt.Errorf("color name not exists: %s", rule.Text)
		}

	}

	return nil
}

func NewColorRuleSet() *ColorRuleSet {
	return &ColorRuleSet{}
}
