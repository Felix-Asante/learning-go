package maps

import "fmt"

func RunMaps() {
	websites := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Twitter":  "https://twitter.com",
	}

	websites["Google"] = "https://google.com.br"
	websites["Youtube"] = "https://youtube.com"
	fmt.Println(websites)

	// MAKE FUNCTION
	courseRating := make(map[string]int, 3)

	courseRating["Go"] = 5
	courseRating["Python"] = 8
	courseRating["JavaScript"] = 10
	fmt.Println(courseRating)

}
