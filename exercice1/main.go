package main

import ("fmt")
	
func main() {
	afficherMenu()
	var choix int
	fmt.Print("Veuillez entrer votre choix : ")
	fmt.Scanln(&choix)
	if choix == 0 {
		fmt.Println("Merci d'avoir utilisé le distributeur. Au revoir !")
		return
	} else if choix < 1 || choix > 4 {
		fmt.Println("Choix invalide. Veuillez réessayer.")
		return
	} else {
	afficherBoisson(choix)
	prix := obtenirPrix(choix)
	fmt.Printf("Le prix de votre boisson est de %d €\n", prix)
	fmt.Print("Veuillez entrer le montant que vous insérez : ")
	var montant int
	fmt.Scanln(&montant)
	for montant < prix {

		for montant < prix {
    fmt.Printf("Montant insuffisant. Il vous manque %d €.\n", prix-montant)
    fmt.Print("Veuillez entrer le montant que vous insérez : ")
    var ajout int
    fmt.Scanln(&ajout)
    montant += ajout
}
		}
	}
}

func afficherMenu() {
	fmt.Println("=== DISTRIBUTEUR ===")

	fmt.Println("1 - Eau       : 1 €")
	fmt.Println("2 - Soda      : 2 €")
	fmt.Println("3 - Café      : 2 €")
	fmt.Println("4 - Chocolat  : 3 €")
	fmt.Println("0 - Quitter")
	fmt.Println()
}

func obtenirPrix(choix int) int {
	if choix == 1 {
		return 1
	} else if choix == 2 {
		return 2
	} else if choix == 3 {
		return 2
	} else if choix == 4 {
		return 3
	} else {
		return 0 
	}
}
	
func afficherBoisson(choix int) {
	if choix == 1 {
		fmt.Println("Vous avez choisi de l'eau.")
	} else if choix == 2 {
		fmt.Println("Vous avez choisi un soda.")
	} else if choix == 3 {
		fmt.Println("Vous avez choisi du café.")
	} else if choix == 4 {
		fmt.Println("Vous avez choisi du chocolat.") 
	}
}


