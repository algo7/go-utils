package files

import "os"

// FileExists checks if a specific file exists at the given path and confirms it is not a directory.
// It returns true if the file exists and false if the file does not exist or if there are other errors like permission issues.
// Example usage:
//
//	if FileExists("file.txt") {
//		fmt.Println("File exists")
//	} else {
//		fmt.Println("File does not exist or cannot be accessed")
//	}
func FileExists(filename string) bool {

	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist
			return false
		}
		// Handle other possible errors like permission issues
		return false
	}

	return !info.IsDir() // Ensure the path is not a directory
}
