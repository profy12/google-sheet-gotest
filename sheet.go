package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// Remplacez par le chemin vers votre fichier de clé JSON
	credentialsFile := "service-account-key.json"

	// Charger les informations d'identification du service account
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read service account credentials file: %v", err)
	}

	// Configurer l'authentification avec les informations d'identification du service account
	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse service account key file to config: %v", err)
	}

	client := config.Client(context.Background())

	// Créer le service Google Sheets
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Spécifiez l'ID de la feuille de calcul et la plage de données
	spreadsheetId := "1HUHn9ndFpI3Q4NHg5Rh-i9QLm6fZ1eX4TdajiwW1s0Y"
	readRange := "Promotions!A3:C" // Modifiez selon la plage que vous souhaitez lire

	// Lire les données de Google Sheet
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Afficher les données lues
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data from the spreadsheet:")
		for _, row := range resp.Values {
			fmt.Println(row)
		}
	}
}
