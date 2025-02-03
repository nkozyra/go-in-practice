package main

validLanguages = []string{
	"en",
	"sp",
	"fr",
	"de",
}

func init() {
	flag.StringVar(&userLanguage, "lang", "en", "language to use (en, sp, fr, de).")
	flag.Parse()
	if slices.Contains(validLanguages, userLanguage) == false {
		log.Fatalf("Invalid language %s. Please use one of %v", userLanguage, validLanguages)
	}
}