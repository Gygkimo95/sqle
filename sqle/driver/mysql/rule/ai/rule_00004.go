package ai

import (
	"strings"

	rulepkg "github.com/actiontech/sqle/sqle/driver/mysql/rule"
	util "github.com/actiontech/sqle/sqle/driver/mysql/rule/ai/util"
	driverV2 "github.com/actiontech/sqle/sqle/driver/v2"
	"github.com/pingcap/parser/ast"
	parserdriver "github.com/pingcap/tidb/types/parser_driver"

	"github.com/actiontech/sqle/sqle/driver/mysql/plocale"
)

const (
	SQLE00004 = "SQLE00004"
)

func init() {
	rh := rulepkg.SourceHandler{
		Rule: rulepkg.SourceRule{
			Name:       SQLE00004,
			Desc:       plocale.Rule00004Desc,
			Annotation: plocale.Rule00004Annotation,
			Category:   plocale.RuleTypeDMLConvention,
			CategoryTags: map[string][]string{
				plocale.RuleCategoryOperand.ID:              {plocale.RuleTagTable.ID},
				plocale.RuleCategorySQL.ID:                  {plocale.RuleTagDDL.ID, plocale.RuleTagIntegrity.ID},
				plocale.RuleCategoryAuditPurpose.ID:         {plocale.RuleTagMaintenance.ID},
				plocale.RuleCategoryAuditAccuracy.ID:        {plocale.RuleTagOffline.ID},
				plocale.RuleCategoryAuditPerformanceCost.ID: {},
			},
			Level:        driverV2.RuleLevelWarn,
			Params:       []*rulepkg.SourceParam{},
			Knowledge:    driverV2.RuleKnowledge{},
			AllowOffline: true,
			Version:      2,
		},
		Message: plocale.Rule00004Message,
		Func:    RuleSQLE00004,
	}
	sourceRuleHandlers = append(sourceRuleHandlers, &rh)
}

/*
==== Prompt start ====
在 MySQL 中，您应该检查 SQL 是否违反了规则(SQLE00004): "在 MySQL 中，建议表的自增字段起始值为0."
您应遵循以下逻辑：
1. 针对 "CREATE TABLE ..." 语句，执行以下检查，若任一条件满足则判定为违反规则：
   1. 语法树中包含 AUTO_INCREMENT 属性的列，且其初始值不等于 0。
2. 针对 "SET ..." 语句，执行以下检查：
   1. 设置的参数为 auto_increment_offset，且其值大于 0。
   若条件满足，则判定为违反规则。
==== Prompt end ====
*/

// ==== Rule code start ====
func RuleSQLE00004(input *rulepkg.RuleHandlerInput) error {
	switch stmt := input.Node.(type) {
	case *ast.CreateTableStmt:
		// "create table"
		if option := util.GetTableOption(stmt.Options, ast.TableOptionAutoIncrement); nil != option {
			//"create table ... auto_increment=..."
			if option.UintValue != 0 {
				// the table option "auto increment" is other than 0
				rulepkg.AddResult(input.Res, input.Rule, SQLE00004)
				return nil
			}
		}
	case *ast.SetStmt:
		for _, variable := range stmt.Variables {
			// 获取设置的变量名
			varName := variable.Name

			// 确认目标对象为'auto_increment_offset'
			if strings.EqualFold(varName, "auto_increment_offset") {
				if v, ok := variable.Value.(*parserdriver.ValueExpr); ok {
					if v.Datum.GetInt64() > 1 {
						rulepkg.AddResult(input.Res, input.Rule, SQLE00004)
						return nil
					}
				}
			}
		}
	}
	return nil
}

// ==== Rule code end ====
