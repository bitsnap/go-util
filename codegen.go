package goutil

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

func GenerateInto(
	dir string,
	headerFunc func(string) string,
	templates map[string][]func() string,
) {
	// Count total jobs
	totalJobs := 0
	for _, funcs := range templates {
		totalJobs += len(funcs)
	}

	type GeneratedContent struct {
		Filename string
		Priority int
		Content  string
	}

	outputChan := make(chan GeneratedContent, totalJobs)
	var wg sync.WaitGroup

	// Process each template generator in parallel
	for filename, funcs := range templates {
		for priority, generateFunc := range funcs {
			if generateFunc == nil {
				continue
			}
			wg.Add(1)
			go func(filename string, priority int, generate func() string) {
				defer wg.Done()
				outputChan <- GeneratedContent{
					Filename: filename,
					Priority: priority,
					Content:  generate(),
				}
			}(filename, priority, generateFunc)
		}
	}

	// Close the channel after all workers are done
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	// Collect and organize rendered content by filename and priority
	renderedContent := make(map[string][]string)
	for filename := range templates {
		// Preinitialize slices with capacity equal to the number of functions
		renderedContent[filename] = make([]string, len(templates[filename]))
	}

	for output := range outputChan {
		// Ensure no out-of-bounds issues
		if len(renderedContent[output.Filename]) <= output.Priority {
			newSize := output.Priority + 1
			temp := make([]string, newSize)
			copy(temp, renderedContent[output.Filename])
			renderedContent[output.Filename] = temp
		}
		renderedContent[output.Filename][output.Priority] = output.Content
	}

	// Write rendered content to files
	for filename, contentParts := range renderedContent {
		// Filter out empty entries
		nonEmptyParts := slices.DeleteFunc(contentParts, func(part string) bool {
			return part == ""
		})

		if len(nonEmptyParts) == 0 {
			continue
		}

		// Generate final file content
		concatContent := strings.Join(nonEmptyParts, "\n")
		finalContent := strings.TrimSpace(headerFunc(concatContent) + concatContent)

		// Write to file
		outputPath := filepath.Join(dir, filename)
		if err := os.WriteFile(outputPath, []byte(finalContent), 0o644); err != nil {
			Logger().Errorw("Failed to write file", "error", err, "filename", filename)
		}
	}
}
