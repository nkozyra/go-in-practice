package main

import (
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile("structured.log", os.O_RDWR|os.O_CREATE, 0755) // # A
	if err != nil {
		panic("could not open log file for writing")
	}
	logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{ // # B
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)
	slog.Info("this is default logging")
	slog.Warn("keep an eye on this, it might be an issue")
	slog.Error("oh no, an error happened here!")
	slog.Debug("this is good while developing ...")
	slog.Info( // # C
		"this is a more complex message",
		slog.String("accepted_values", "key/value pairs with specific types for marshalling"),
		slog.Int("an int:", 30),
		slog.Group("grouped_info",
			slog.String("you_can", "do this too"),
		),
	)
}
