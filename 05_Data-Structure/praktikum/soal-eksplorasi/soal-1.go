package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CourseData struct {
	CourseName               string
	LearningSatisfaction    int
	MentorSatisfaction      int
}

func main() {
	file, err := os.Open("survey.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var courses []CourseData

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		learningSatisfaction, _ := strconv.Atoi(record[3])
		mentorSatisfaction, _ := strconv.Atoi(record[4])

		courses = append(courses, CourseData{
			CourseName:             record[2],
			LearningSatisfaction:  learningSatisfaction,
			MentorSatisfaction:    mentorSatisfaction,
		})
	}

	mentorAvg := calculateAverageMentorSatisfaction(courses)
	fmt.Println("Rata-rata kepuasan mentor secara keseluruhan:", mentorAvg)

	learningAvg := calculateAverageLearningSatisfaction(courses)
	fmt.Println("Rata-rata kepuasan belajar secara keseluruhan:", learningAvg)

	highestMentorCourse := findCourseWithHighestMentorSatisfaction(courses)
	fmt.Println("Nama kursus dengan rata-rata kepuasan mentor tertinggi:", highestMentorCourse)
}

func calculateAverageMentorSatisfaction(courses []CourseData) float64 {
	total := 0
	for _, course := range courses {
		total += course.MentorSatisfaction
	}
	return float64(total) / float64(len(courses))
}

func calculateAverageLearningSatisfaction(courses []CourseData) float64 {
	total := 0
	for _, course := range courses {
		total += course.LearningSatisfaction
	}
	return float64(total) / float64(len(courses))
}

func findCourseWithHighestMentorSatisfaction(courses []CourseData) string {
	highest := 0.0 // Menggunakan float64
	var highestCourse string
	for _, course := range courses {
		avg := float64(course.MentorSatisfaction) / float64(len(courses))
		if avg > highest {
			highest = avg
			highestCourse = course.CourseName
		}
	}
	return highestCourse
}

