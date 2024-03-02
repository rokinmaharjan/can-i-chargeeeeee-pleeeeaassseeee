package charger

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func FindCharger() bool {
	log.Println("Finding chargers")

	const available = "Available"
	const chargerStatusSelector = "span.MuiChip-label.MuiChip-labelSmall.css-1pjtbja"

	ctx, cancel := chromedp.NewContext(context.Background())
	// Ensure browser resources are released when no longer needed
	defer cancel()

	var firstChargerStatus string
	var secondChargerStatus string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://charge.id/ao-61"),
		// Wait for the page to load and render JavaScript
		chromedp.Sleep(2*time.Second),
		chromedp.Text(chargerStatusSelector, &firstChargerStatus, chromedp.ByQueryAll),

		chromedp.Navigate("https://charge.id/ao-62"),
		// Wait for the page to load and render JavaScript
		chromedp.Sleep(2*time.Second),
		chromedp.Text(chargerStatusSelector, &secondChargerStatus, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Fatal("Error while fetching charger status:", err)
	}

	if firstChargerStatus == available || secondChargerStatus == available {
		return true
	}

	return false
}
