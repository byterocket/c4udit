package analyzer

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Run an analysis of Solidity contracts in `path`.
// Argument `issues` encodes the Issues to search for.
func Run(issues []Issue, paths []string) (*Report, error) {
	report := &Report{
		Issues:           issues,
		FilesAnalyzed:    []string{},
		FindingsPerIssue: make(map[string][]Finding),
	}

	for _, path := range paths {
		err := run(report, path)
		if err != nil {
			return &Report{}, nil
		}
	}

	return report, nil
}

func run(report *Report, path string) error {
	pathInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if pathInfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}

		for _, file := range files {
			err = run(report, filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
		}
	} else {
		file := path

		// Only analyze Solidity files
		if !strings.HasSuffix(file, ".sol") {
			return nil
		}

		findingsPerIssue, err := analyzeFile(report.Issues, file)
		if err != nil {
			return err
		}

		// Add file and findings to report.
		report.FilesAnalyzed = append(report.FilesAnalyzed, file)
		for _, issue := range report.Issues {
			report.FindingsPerIssue[issue.Identifier] = append(report.FindingsPerIssue[issue.Identifier],
				findingsPerIssue[issue.Identifier]...,
			)
		}
	}

	return nil
}

func analyzeFile(issues []Issue, file string) (map[string][]Finding, error) {
	findings := make(map[string][]Finding)

	readFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		for _, issue := range issues {
			matched, _ := regexp.MatchString(issue.Pattern, line)
			if matched {
				findings[issue.Identifier] = append(findings[issue.Identifier], Finding{
					IssueIdentifier: issue.Identifier,
					File:            file,
					LineNumber:      lineNumber,
					LineContent:     strings.TrimSpace(line),
				})
			}
		}
	}

	return findings, nil
}
