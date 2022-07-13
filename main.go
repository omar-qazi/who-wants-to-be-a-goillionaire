package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type QuizQuestions []struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answers  struct {
		AnswerA string `json:"answer_a"`
		AnswerB string `json:"answer_b"`
		AnswerC string `json:"answer_c"`
		AnswerD string `json:"answer_d"`
		AnswerE string `json:"answer_e"`
		AnswerF string `json:"answer_f"`
	} `json:"answers"`
	CorrectAnswers struct {
		AnswerACorrect string `json:"answer_a_correct"`
		AnswerBCorrect string `json:"answer_b_correct"`
		AnswerCCorrect string `json:"answer_c_correct"`
		AnswerDCorrect string `json:"answer_d_correct"`
		AnswerECorrect string `json:"answer_e_correct"`
		AnswerFCorrect string `json:"answer_f_correct"`
	} `json:"correct_answers"`
	Category   string `json:"category"`
	Difficulty string `json:"difficulty"`
}

func main() {
	questions := GetQuizQuestions()
	var playername string

	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome to... Who Wants To Be A Goillionaire!")

	fmt.Println("\nHow to play:")
	fmt.Println(" - You will be given 10 questions.")
	fmt.Println(" - For each one that you answer correctly, you will be awarded with $100,000")
	fmt.Println(" - Give one wrong answer, and you lose all the money.")
	fmt.Println(" - Thank you for playing and best of luck!")

	fmt.Println("\n\nPress enter your name to start playing...")
	fmt.Scanln(&playername)

	score := 0
	for i := 0; i < len(questions); i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("%s, you have won $%d00,000 !\n\n", playername, score)

		fmt.Print("\nQuestion:\n", questions[i].Question, "\n\n")
		if questions[i].Answers.AnswerA != "" {
			fmt.Println("a. ", questions[i].Answers.AnswerA)
		}
		if questions[i].Answers.AnswerB != "" {
			fmt.Println("b. ", questions[i].Answers.AnswerB)
		}
		if questions[i].Answers.AnswerC != "" {
			fmt.Println("c. ", questions[i].Answers.AnswerC)
		}
		if questions[i].Answers.AnswerD != "" {
			fmt.Println("d. ", questions[i].Answers.AnswerD)
		}
		if questions[i].Answers.AnswerE != "" {
			fmt.Println("e. ", questions[i].Answers.AnswerE)
		}
		if questions[i].Answers.AnswerF != "" {
			fmt.Println("f. ", questions[i].Answers.AnswerF)
		}

		//fmt.Println(questions[i].CorrectAnswers)

		var userAns string
		fmt.Print("\nEnter response (a,b,c,d,e,f): ")
		fmt.Scanln(&userAns)

		if userAns == "a" && questions[i].CorrectAnswers.AnswerACorrect == "true" ||
			userAns == "b" && questions[i].CorrectAnswers.AnswerBCorrect == "true" ||
			userAns == "c" && questions[i].CorrectAnswers.AnswerCCorrect == "true" ||
			userAns == "d" && questions[i].CorrectAnswers.AnswerDCorrect == "true" ||
			userAns == "e" && questions[i].CorrectAnswers.AnswerECorrect == "true" ||
			userAns == "f" && questions[i].CorrectAnswers.AnswerFCorrect == "true" {
			score += 1
			fmt.Printf("\nYou have won $%d00,000 !\n", score)
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Sorry, you lost. Better luck next time.")
			fmt.Printf("\n\n%s, your high score was $%d00,000 !\n\n", playername, score)
			os.Exit(0)
		}
	}
}

func GetQuizQuestions() QuizQuestions {
	const baseUrl = `https://quizapi.io/api/v1/questions`
	const apiKey = `la9Eq8Rle8k5iw06vJKFc9dcXbYFkL11oxqqKqMe`
	const limit = `10`
	const category = `Linux`
	const difficulty = `easy`
	const tags = `terminal`
	const requestUrl = baseUrl + `?apiKey=` + apiKey + `&limit=` + limit + `&category=` + category + `&difficulty=` + difficulty

	response, err := http.Get(requestUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	questions := QuizQuestions{}
	jsonErr := json.Unmarshal(content, &questions)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return questions
}
