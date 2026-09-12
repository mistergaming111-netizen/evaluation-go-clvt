package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== GESTIONNAIRE DE NOTES ===")
	var ChoixNbrnotes int
	fmt.Println("Combien de notes voulez-vous saisir ?")
	fmt.Scan(&ChoixNbrnotes)
	var listeNotes []int
	for i := 0; i < ChoixNbrnotes; i++{ 
		var NotesChoisies int
		fmt.Print("Note : ")
		fmt.Scan(&NotesChoisies)
		if NotesChoisies < 0 || NotesChoisies > 20 {
			fmt.Println("La note doit etre comprise entre 0 et 20 !")
			i--	
			continue	
			} 
		listeNotes = append(listeNotes, NotesChoisies )
	}
	fmt.Println("=== RÉSULTATS ===")
	somme := 0
	for _, note := range listeNotes {
		somme += note
	}
	moyenne := calculerMoyenne(somme, len(listeNotes))
	fmt.Println("Moyenne :", moyenne)
	NoteMaximale := trouverMaximum(listeNotes)
	fmt.Println("note maximale :", NoteMaximale)
	NoteMinimale := trouverMinimum(listeNotes)
	fmt.Println("note minimale :", NoteMinimale)
	if moyenne < 10 {
		fmt.Println("étudiant admis")
	}else{
		fmt.Println("étudiant pas admis")
	}

	
}

func calculerMoyenne(somme int, nombre int) float64 {
	return float64(somme) / float64(nombre)
}

func trouverMaximum(notes []int) int {
	max := notes[0]
	for _, note := range notes {
		if note > max {
			max = note
		}
	}
	return max
}

func trouverMinimum(notes []int) int {
	min := notes[0]
	for _, note := range notes {
		if note < min {
			min = note		
		}
	}
	return min

}

func afficherResultat(moyenne float64) {
	if moyenne >= 10 {
	fmt.Print("étudiant admis")
	} else {
	fmt.Print("non admis")
	}
    
}