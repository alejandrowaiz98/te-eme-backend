package main

import (
	_ "te-eme-backend/app/adapter/out/firestore"
	"te-eme-backend/app/shared/archetype"
)

func main() {
	archetype.Setup()
}
