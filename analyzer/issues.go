package analyzer

// AllIssues returns the list of all issues.
func AllIssues() []Issue {
	return append(GasOpIssues(), LowRiskIssues()...)
}

// GasOpIssues returns the list of all gas optimization issues.
func GasOpIssues() []Issue {
	return []Issue{
		// G001 - Don't Initialize Variables with Default Value
		{
			"G001",
			GASOP,
			"Don't Initialize Variables with Default Value",
			"https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g001---dont-initialize-variables-with-default-value",
			"(= 0|= false)",
		},
		// G002 - Cache Array Length Outside of Loop
		{
			"G002",
			GASOP,
			"Cache Array Length Outside of Loop",
			"https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g002---cache-array-length-outside-of-loop",
			".length",
		},
		// G003 - Use != 0 instead of > 0 for Unsigned Integer Comparison
		{
			"G003",
			GASOP,
			"Use != 0 instead of > 0 for Unsigned Integer Comparison",
			"https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g003---use--0-instead-of--0-for-unsigned-integer-comparison",
			"(>0|> 0)",
		},
		// G006 - Use immutable for OpenZeppelin AccessControl's Roles Declarations
		{
			"G006",
			GASOP,
			"Use immutable for OpenZeppelin AccessControl's Roles Declarations",
			"https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g006---use-immutable-for-openzeppelin-accesscontrols-roles-declarations",
			"keccak",
		},
		// G007 - Long Revert Strings
		{
			"G007",
			GASOP,
			"Long Revert Strings",
			"https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g007---long-revert-strings",
			"\".{33,}\"", // Anything between "'s with at least 33 characters
		},
		// G008 - Shift Right instead of Dividing by 2
		{
			"G008",
			GASOP,
			"Shift Right instead of Dividing by 2",
			"https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g008---shift-right-instead-of-dividing-by-2",
			"(/2|/ 2)",
		},
	}
}

// LowRiskIssues returns the list of all low risk issues.
func LowRiskIssues() []Issue {
	return []Issue{
		// L001 - Unsafe ERC20 Operation(s)
		{
			"L001",
			LOW,
			"Unsafe ERC20 Operation(s)",
			"https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l001---unsafe-erc20-operations",
			"(\\.transfer\\(|\\.transferFrom\\(|\\.approve\\()", // ".tranfer(", ".transferFrom(" or ".approve("
		},
		// L003 - Unspecific Compiler Version Pragma
		{
			"L003",
			LOW,
			"Unspecific Compiler Version Pragma",
			"https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l003---unspecific-compiler-version-pragma",
			"pragma solidity (\\^|>)", // "pragma solidity ^" or "pragma solidity >"
		},
	}
}
