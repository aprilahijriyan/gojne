package lib

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ParseCsrfToken(e *colly.HTMLElement) string {
	return e.DOM.Find("input[name=_token]").AttrOr("value", "")
}

func ParseDeliveryDetail(h *colly.HTMLElement) ResultDelivery {
	delivery := ResultDelivery{}
	h.DOM.Find(".tile_stats_count").Each(func(i int, e *goquery.Selection) {
		var value = e.Find("h4").Text()
		if value == "" {
			value = e.Find("h3").Text()
		}
		if value == "" {
			value = e.Find("h2").Text()
		}
		detail := DetailDelivery{}
		detail.Title = strings.TrimSpace(e.Find(".count_top").Text())
		detail.Value = strings.TrimSpace(value)
		delivery = append(delivery, detail)
	})
	return delivery
}

func ParseDetailHistory(h *colly.HTMLElement) DetailHistory {
	history := DetailHistory{}
	receiverElem := h.DOM.Find("html body div.container.body div.right_col div div.row div.col-md-12 div.x_panel div.x_content div.dashboard-widget-content div.col-md-3.col-sm-3.col-xs-12").First()
	receiverName := strings.TrimSpace(receiverElem.Find("h2").Text())
	receiverDetail := ReceiverInfo{
		Name:         receiverName,
		Relationship: strings.TrimSpace(receiverElem.Find("h4").Text()),
	}
	history.Receiver = receiverDetail
	h.DOM.Find("ul > li").Each(func(i int, e *goquery.Selection) {
		history.Data = append(history.Data, DataHistory{
			Title: strings.TrimSpace(e.Find(".title").Text()),
			Date:  strings.TrimSpace(e.Find(".byline").Text()),
		})
	})
	return history
}

func ParseDetailShipment(h *colly.HTMLElement) ResultShipment {
	var shipment = ResultShipment{}
	h.DOM.Find("html body div.container.body div.right_col div div.row div.col-md-12 div.x_panel .tile").Each(func(i int, e *goquery.Selection) {
		shipment = append(shipment, DetailShipment{
			Title: strings.TrimSpace(e.Find("span").Text()),
			Value: strings.TrimSpace(e.Find("h4").Text()),
		})
	})
	return shipment
}
