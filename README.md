# HET EJDID Backend - Hazem Jdey - ISAMM 2022

C'est la partie serveur de mon projet.


## Dépendedances

- Golang
- Mysql

## Installation
- Créer une base de données
- Changer le nom du fichier ".env-example" en ".env" et modifier les informations de la base de donnée.
- Modifier selon votre choix le port dans "main.go" (par défaut 8081).
- Lancer :
```
go mod download

go run main.go
```