# Go Study Project Instructions

## Project Structure
- go.mod is located in the project root.
- Manage learning materials by Day folders.
- Keep exercise files in the corresponding Day folder.
- Each Day folder has exactly one main.go, used as that Day's entry point.
- Use descriptive filenames that clearly identify each exercise's topic instead of repeating main.go.

## Go Coding Rules
- Use package main for executable learning examples.
- Only main.go should contain the main() function in each Day package.
- Other Go files should define uniquely named functions.
- Do not create additional go.mod files inside Day folders.
- Avoid unnecessary external dependencies.
- Do not arbitrarily delete existing learning code.
- Apply gofmt to Go code after writing or changing it.

## Learning Rules
- Explain concepts before implementing code.
- Do not immediately provide the complete exercise solution.
- Give hints and guide the learner through implementation.
- Review the learner's code and explain mistakes.
- Explain why the code works, not just what to write.
- Keep examples understandable for beginners and explain the operating principles when code changes.

## Code Modification Rules
- Do not modify unrelated files.
- Preserve existing learning examples.
- Explain the files to be changed before making modifications.
- Run gofmt on changed Go files.
- Run appropriate tests and report the actual results.
- Do not delete or restructure files without permission.
