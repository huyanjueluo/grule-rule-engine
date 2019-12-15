// Code generated by ANTLR 4.7.
// DO NOT EDIT!

package parser // grule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasegruleListener is a complete listener for a parse tree produced by gruleParser.
type BasegruleListener struct{}

var _ gruleListener = &BasegruleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegruleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegruleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegruleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegruleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BasegruleListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasegruleListener) ExitRoot(ctx *RootContext) {}

// EnterRuleEntry is called when production ruleEntry is entered.
func (s *BasegruleListener) EnterRuleEntry(ctx *RuleEntryContext) {}

// ExitRuleEntry is called when production ruleEntry is exited.
func (s *BasegruleListener) ExitRuleEntry(ctx *RuleEntryContext) {}

// EnterSalience is called when production salience is entered.
func (s *BasegruleListener) EnterSalience(ctx *SalienceContext) {}

// ExitSalience is called when production salience is exited.
func (s *BasegruleListener) ExitSalience(ctx *SalienceContext) {}

// EnterRuleName is called when production ruleName is entered.
func (s *BasegruleListener) EnterRuleName(ctx *RuleNameContext) {}

// ExitRuleName is called when production ruleName is exited.
func (s *BasegruleListener) ExitRuleName(ctx *RuleNameContext) {}

// EnterRuleDescription is called when production ruleDescription is entered.
func (s *BasegruleListener) EnterRuleDescription(ctx *RuleDescriptionContext) {}

// ExitRuleDescription is called when production ruleDescription is exited.
func (s *BasegruleListener) ExitRuleDescription(ctx *RuleDescriptionContext) {}

// EnterWhenScope is called when production whenScope is entered.
func (s *BasegruleListener) EnterWhenScope(ctx *WhenScopeContext) {}

// ExitWhenScope is called when production whenScope is exited.
func (s *BasegruleListener) ExitWhenScope(ctx *WhenScopeContext) {}

// EnterThenScope is called when production thenScope is entered.
func (s *BasegruleListener) EnterThenScope(ctx *ThenScopeContext) {}

// ExitThenScope is called when production thenScope is exited.
func (s *BasegruleListener) ExitThenScope(ctx *ThenScopeContext) {}

// EnterAssignExpressions is called when production assignExpressions is entered.
func (s *BasegruleListener) EnterAssignExpressions(ctx *AssignExpressionsContext) {}

// ExitAssignExpressions is called when production assignExpressions is exited.
func (s *BasegruleListener) ExitAssignExpressions(ctx *AssignExpressionsContext) {}

// EnterAssignExpression is called when production assignExpression is entered.
func (s *BasegruleListener) EnterAssignExpression(ctx *AssignExpressionContext) {}

// ExitAssignExpression is called when production assignExpression is exited.
func (s *BasegruleListener) ExitAssignExpression(ctx *AssignExpressionContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasegruleListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasegruleListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasegruleListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasegruleListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasegruleListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasegruleListener) ExitPredicate(ctx *PredicateContext) {}

// EnterExpressionAtom is called when production expressionAtom is entered.
func (s *BasegruleListener) EnterExpressionAtom(ctx *ExpressionAtomContext) {}

// ExitExpressionAtom is called when production expressionAtom is exited.
func (s *BasegruleListener) ExitExpressionAtom(ctx *ExpressionAtomContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BasegruleListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BasegruleListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasegruleListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasegruleListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BasegruleListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BasegruleListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BasegruleListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BasegruleListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasegruleListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasegruleListener) ExitVariable(ctx *VariableContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BasegruleListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BasegruleListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BasegruleListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BasegruleListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasegruleListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasegruleListener) ExitConstant(ctx *ConstantContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BasegruleListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BasegruleListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *BasegruleListener) EnterRealLiteral(ctx *RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *BasegruleListener) ExitRealLiteral(ctx *RealLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasegruleListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasegruleListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BasegruleListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BasegruleListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}