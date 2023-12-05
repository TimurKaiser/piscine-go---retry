package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// check vérifie si une erreur est survenue et panique si c'est le cas.

func check(err error) {
	if err != nil {
		fmt.Println("Une erreur est survenue :")
		panic(err)
	}
}

func formatPunctuation(text string) string {
	// Remplace les espaces avant les ponctuations par des espaces après.
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " ,", ", ")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", "; ")

	// Gère les groupes de ponctuations comme ... ou !?.
	text = strings.ReplaceAll(text, " ...", "...")
	text = strings.ReplaceAll(text, " !?", "!?")

	// Gère les apostrophes ' correctement.
	text = regexp.MustCompile(`\s+'`).ReplaceAllString(text, " '")
	text = regexp.MustCompile(`'\s+`).ReplaceAllString(text, "' ")

	// Gère le cas où il y a plus d'un mot entre les apostrophes.
	text = regexp.MustCompile(`' (.+?) '`).ReplaceAllString(text, "'$1'")

	return text
}

// filterEmptyStrings supprime les chaînes vides d'une slice.

func filterEmptyStrings(words []string) []string {
	var filtered []string
	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}
	return filtered
}

// IsVowel vérifie si un rune est une voyelle.

func IsVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y', 'A', 'E', 'I', 'O', 'U', 'Y':
		return true
	default:
		return false
	}
}

func main() {
	arg := os.Args[1:]
	// Vérifie si le nombre correct d'arguments a été fourni.
	if len(arg) != 2 {
		fmt.Println("Veuillez fournir les input et output")
		return
	}

	data, err := os.ReadFile(arg[0])
	// Utilise la fonction check pour gérer toute erreur de lecture.
	check(err)
	input := string(data)

	// Vérifie si l'entrée est vide.
	if input == "" {
		fmt.Println("Aucune entrée fournie.")
		return
	}

	// Sépare l'entrée en mots.
	words := strings.Split(input, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]

		// Effectue des actions en fonction des balises trouvées.
		switch word {
		case "(hex)", "(bin)":
			if i == 0 {
				continue
			}
			prevWord := words[i-1]
			var value int64
			var err error

			if word == "(hex)" {
				value, err = strconv.ParseInt(prevWord, 16, 64)
			} else if word == "(bin)" {
				value, err = strconv.ParseInt(prevWord, 2, 64)
			}
			check(err) // Gère les erreurs de conversion immédiatement après la conversion.
			words[i-1] = fmt.Sprint(value)
			words[i] = "" // Efface la directive de transformation après utilisation.

		case "(low)", "(up)", "(cap)":
			if i == 0 {
				continue
			}
			prevWord := words[i-1]
			var transformed string

			switch word {
			case "(low)":
				transformed = strings.ToLower(prevWord)
			case "(up)":
				transformed = strings.ToUpper(prevWord)
			case "(cap)":
				transformed = strings.Title(prevWord)
			}
			words[i-1] = transformed
			words[i] = ""

		case "(low,", "(up,", "(cap,":
			if i < 2 {
				continue
			}
			// Gère les transformations avec un décalage.
			offsetStr := strings.TrimSuffix(words[i+1], ")")
			offset, err := strconv.Atoi(offsetStr)
			// Utilise la fonction check pour gérer les erreurs de conversion.
			check(err)

			// Applique les transformations sur les mots précédents en fonction de l'offset.
			for j := 1; j <= offset && i-j >= 0; j++ {
				switch word {
				case "(low,":
					// Met en minuscules les mots précédents selon l'offset.
					words[i-j] = strings.ToLower(words[i-j])
				case "(up,":
					// Met en majuscules les mots précédents selon l'offset.
					words[i-j] = strings.ToUpper(words[i-j])
				case "(cap,":
					// Met en majuscule la première lettre des mots précédents selon l'offset.
					words[i-j] = strings.Title(words[i-j])
				}
			}
			words[i] = ""
			words[i+1] = ""

		case "a":
			// Gère la transformation de 'a' en 'an' devant un mot commençant par 'h' ou une voyelle.
			if i+1 < len(words) && (strings.HasPrefix(words[i+1], "h") || strings.IndexFunc(words[i+1], IsVowel) == 0) {
				words[i] = "an"
			}
		}
	}

	filteredWords := filterEmptyStrings(words)
	joinedWords := strings.Join(filteredWords, " ")

	// Formate la ponctuation.
	formattedOutput := formatPunctuation(joinedWords)
	// Écrit le résultat dans le fichier de destination.
	err = os.WriteFile(arg[1], []byte(formattedOutput), 0666)
	check(err)
}
