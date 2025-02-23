package main

import (
	"fmt"
)

// 1.	Facts and Rules:
// •	Fact: Represents a simple statement like parent(alice, bob).
// •	Rule: Represents a conditional logic like grandparent(X, Z) :- parent(X, Y), parent(Y, Z).
// 2.	Knowledge Base:
// •	Stores facts and rules.
// •	Provides a Query method to check if a given fact can be derived.
// 3.	Query Evaluation:
// •	Checks if the fact matches directly with existing facts.
// •	Evaluates rules by checking if their conditions (body) are satisfied.
// 4.	Unification:
// •	Matches facts based on predicates and arguments.
// •	Does not handle variable substitution in this simplified example.
// 5.	Recursive Evaluation:
// •	Enables chaining of facts and rules, simulating logic inference.

// Advantages and Limitations

// •	Advantages:
// •	This system allows declarative rule-based reasoning.
// •	Demonstrates the essence of logic programming in a procedural language.
// •	Limitations:
// •	Lacks advanced unification with variable binding.
// •	Performance can degrade with complex rules and queries.

// This implementation illustrates the basic principles of logic programming in Go while keeping the structure simple.
// Fact represents a simple fact in the logic system.
type Fact struct {
	Predicate string
	Arguments []string
}

// Rule represents a logical rule with conditions.
type Rule struct {
	Head      Fact
	Body      []Fact
}

// KnowledgeBase holds facts and rules for the logic system.
type KnowledgeBase struct {
	Facts []Fact
	Rules []Rule
}

// AddFact adds a new fact to the knowledge base.
func (kb *KnowledgeBase) AddFact(fact Fact) {
	kb.Facts = append(kb.Facts, fact)
}

// AddRule adds a new rule to the knowledge base.
func (kb *KnowledgeBase) AddRule(rule Rule) {
	kb.Rules = append(kb.Rules, rule)
}

// Query searches for matching facts or rules in the knowledge base.
func (kb *KnowledgeBase) Query(query Fact) bool {
	// Check for direct matching facts.
	for _, fact := range kb.Facts {
		if matchFacts(query, fact) {
			return true
		}
	}

	// Check rules to derive new facts.
	for _, rule := range kb.Rules {
		if matchFacts(query, rule.Head) {
			// Evaluate the body of the rule.
			if kb.evaluateBody(rule.Body) {
				return true
			}
		}
	}

	return false
}

// Evaluate the body of a rule (all conditions must be true).
func (kb *KnowledgeBase) evaluateBody(body []Fact) bool {
	for _, condition := range body {
		if !kb.Query(condition) {
			return false
		}
	}
	return true
}

// Match facts (basic unification based on predicates and arguments).
func matchFacts(f1, f2 Fact) bool {
	if f1.Predicate != f2.Predicate || len(f1.Arguments) != len(f2.Arguments) {
		return false
	}
	for i, arg := range f1.Arguments {
		if arg != f2.Arguments[i] {
			return false
		}
	}
	return true
}

// Main function demonstrates the logic programming system.
func mainL() {
	kb := &KnowledgeBase{}

	// Add some facts.
	kb.AddFact(Fact{"parent", []string{"alice", "bob"}})
	kb.AddFact(Fact{"parent", []string{"bob", "charlie"}})

	// Add a rule: grandparent(X, Z) :- parent(X, Y), parent(Y, Z).
	kb.AddRule(Rule{
		Head: Fact{"grandparent", []string{"X", "Z"}},
		Body: []Fact{
			{"parent", []string{"X", "Y"}},
			{"parent", []string{"Y", "Z"}},
		},
	})

	// Query for a grandparent relationship.
	query := Fact{"grandparent", []string{"alice", "charlie"}}
	if kb.Query(query) {
		fmt.Println("Query is true:", query)
	} else {
		fmt.Println("Query is false:", query)
	}
}