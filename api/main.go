package main

import (
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"deej.io/mml-htmx-demo/api/components"
	"deej.io/mml-htmx-demo/api/mml"
)

func main() {
	connections := make(map[string]struct{})
	rolls := 0

	mmlServerWSURI := os.Getenv("MML_SERVER_WS_URI")
	if mmlServerWSURI == "" {
		mmlServerWSURI = "ws://localhost:8081"
	}

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("GET /")
		components.Client(mmlServerWSURI).Render(r.Context(), w)
	})

	http.HandleFunc("/mml", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("GET /mml")
		light := mml.Light{
			ID:        "light",
			Type:      mml.LightTypePoint,
			Intensity: 300,
			X: 5,
			Y: 10,
			Z: 5,
		}
		components.Init(light).Render(r.Context(), w)
	})

	http.HandleFunc("/light", func(w http.ResponseWriter, r *http.Request) {
		colour := fmt.Sprintf("rgb(%d,%d,%d)", rand.Intn(256), rand.Intn(256), rand.Intn(256))
		light := mml.Light{
			ID:        "light",
			Type:      mml.LightTypePoint,
			Intensity: 300,
			Color:     colour,
			X: 5,
			Y: 10,
			Z: 5,
		}

		components.Light(light).Render(r.Context(), w)
	})

	http.HandleFunc("/uptime", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		ts, _ := strconv.Atoi(r.Form.Get("ts"))
		slog.Info("uptime", slog.Any("url", r.URL), slog.Any("form", r.Form))
		uptime := time.Millisecond * time.Duration(ts)
		uptimeMinutes := int(uptime.Minutes())
		uptimeSeconds := int(uptime.Seconds()) - uptimeMinutes * 60
		var uptimeText string
		if uptimeMinutes > 0 {
			uptimeText = fmt.Sprintf("Uptime: %dm %2ds", uptimeMinutes, uptimeSeconds)
		} else {
			uptimeText = fmt.Sprintf("Uptime: %ds", uptimeSeconds)
		}
		label := fmt.Sprintf(`
			<m-label
				content="%s"
				id="uptime-label"
				y="3"
				width="5"
				height="0.5"
				color="#bfdbfe"
				font-color="#172554"
				alignment="center"
				hx-trigger="every 1s"
				hx-get="/uptime"
				hx-vals="js:{ts: document.timeline.currentTime}"
				hx-swap="outerHTML"
			></m-label>
		`, uptimeText)
		io.WriteString(w, label)
	})


	http.HandleFunc("/connected", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.Form.Get("connectionId")
		connections[id] = struct{}{}
		slog.Info("/connected", slog.String("connectionId", id), slog.Int("active", len(connections)))
		components.ConnectedClients(len(connections)).Render(r.Context(), w)
	})

	http.HandleFunc("/disconnected", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.Form.Get("connectionId")
		delete(connections, id)
		slog.Info("/disconnected", slog.String("connectionId", id), slog.Int("active", len(connections)))
		components.ConnectedClients(len(connections)).Render(r.Context(), w)
	})

	rollMap := [6][3]int{
		{0, 0, 0},
		{0, 0, 270},
		{270, 0, 0},
		{90, 0, 0},
		{0, 0, 90},
		{180, 0, 0},
	}

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		slog.Info("roll", slog.Any("url", r.URL), slog.Any("form", r.Form))
		from, _ := strconv.Atoi(r.Form.Get("from"))
		ts, _ := strconv.Atoi(r.Form.Get("ts"))
		startTime := int64(ts)
		to := rand.Intn(6)
		current := rollMap[from]
		target := rollMap[to]
		anims := []components.Animation {
			{ID: "y-up-anim", Easing: components.EaseOutSine,     Attr: "y",  StartTimeMs: startTime, DurationMs: 300, Start: 1, End: 3, Loop: false},
			{ID: "y-down-anim", Easing: components.EaseOutBounce, Attr: "y",  StartTimeMs: startTime + 300, DurationMs: 500, Start: 3, End: 1, Loop: false},
			{ID: "rx-anim", Easing: components.EaseOutCubic,      Attr: "rx", StartTimeMs: startTime, DurationMs: 500, Start: current[0], End: target[0], Loop: false},
			{ID: "ry-anim", Easing: components.EaseOutCubic,      Attr: "ry", StartTimeMs: startTime, DurationMs: 500, Start: current[1], End: target[1], Loop: false},
			{ID: "rz-anim", Easing: components.EaseOutCubic,      Attr: "rz", StartTimeMs: startTime, DurationMs: 500, Start: current[2], End: target[2], Loop: false},
		}
		rolls++
		components.All(
			components.Dice(to, anims),
			components.DiceClickLabel(rolls, true),
		).Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", nil)
}
