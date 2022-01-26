package analyzer

import (
	"fmt"
	"strings"
)

// Report is the end result of an analysis containing the files analyzed,
// the issues searched for and a map of findings per issue.
type Report struct {
	Issues        []Issue
	FilesAnalyzed []string
	// Key is Issue Identifier
	FindingsPerIssue map[string][]Finding
}

// Issue represents an Issue to search for in the codebase.
// The pattern field is a RegEx string which must compile.
type Issue struct {
	Identifier string
	Severity   Severity
	Title      string
	Link       string
	Pattern    string
}

// Finding represents a possible Issue found in the codebase.
type Finding struct {
	IssueIdentifier string
	File            string
	LineNumber      int
	LineContent     string
}

// Severity type defining the severity level for an Issue.
type Severity int

// The Severity Enum.
const (
	GASOP Severity = iota
	NC
	LOW
)

// Markdown returns the report as string in markdown style.
func (r Report) Markdown() string {
	const c4uditRepoLink = "https://github.com/byterocket/c4udit"
	// Issue output in Code4Rena format:
	// ### {{ issue.Title }}
	//
	// #### Impact
	// Issue information: [{{ issue.Identifier }}]({{ issue.Link }})
	//
	// #### Findings
	// {{ _, finding := range findings: finding.String() }}
	//
	// #### Tools used
	// [c4udit]({{ c4uditRepoLink }})
	//
	buf := strings.Builder{}

	buf.WriteString("# c4udit Report\n")
	buf.WriteString("\n")

	buf.WriteString("## Files analyzed\n")
	for _, f := range r.FilesAnalyzed {
		buf.WriteString("- " + f + "\n")
	}
	buf.WriteString("\n")

	buf.WriteString("## Issues found\n")
	buf.WriteString("\n")
	for _, issue := range r.Issues {
		findings := r.FindingsPerIssue[issue.Identifier]
		if len(findings) == 0 {
			continue
		}

		buf.WriteString("### " + issue.Title + "\n")
		buf.WriteString("\n")

		// Impact
		buf.WriteString("#### Impact\n")
		buf.WriteString("Issue Information: [" + issue.Identifier + "]" + "(" + issue.Link + ")" + "\n")
		buf.WriteString("\n")

		// Findings
		buf.WriteString("#### Findings:\n")
		buf.WriteString("```\n")
		for _, finding := range findings {
			buf.WriteString(finding.String())
		}
		buf.WriteString("```\n")

		// Tools used
		buf.WriteString("#### Tools used\n")
		buf.WriteString("[c4udit](" + c4uditRepoLink + ")\n")

		buf.WriteString("\n")
	}

	return buf.String()
}

func (r Report) String() string {
	// Build files string.
	files := "Files analyzed:\n"
	for _, f := range r.FilesAnalyzed {
		files += fmt.Sprintf("- %s\n", f)
	}
	files += "\n"

	// Build issues string.
	issues := "Issues found:\n"
	for i, issue := range r.Issues {
		// Get findings for issue
		findings := r.FindingsPerIssue[issue.Identifier]

		// Skip if no findings
		if len(findings) == 0 {
			continue
		}

		// Add findings per issue
		issues += " " + issue.Identifier + ":\n"
		for _, finding := range findings {
			issues += "  " + finding.String()
		}

		// Add newline if not last issue
		if i+1 != len(r.Issues) {
			issues += "\n"
		}
	}

	return files + issues
}

func (i Issue) String() string {
	return i.Identifier
}

func (f Finding) String() string {
	return fmt.Sprintf("%s::%d => %s\n", f.File, f.LineNumber, f.LineContent)
}

func (s Severity) String() string {
	return []string{
		"Gas Optimization",
		"Non-Critical",
		"Low Risk",
	}[s]
}
