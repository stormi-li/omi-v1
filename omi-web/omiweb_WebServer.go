package omiweb

import (
	"embed"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	omiclient "github.com/stormi-li/omi-v1/omi-manager"
)

type WebServer struct {
	router          *router
	redisClient     *redis.Client
	omiWebClient    *omiclient.Client
	omiServerClient *omiclient.Client
	serverName      string
	weight          int
	embeddedSource  embed.FS
	embedModel      bool
	cache           *fileCache
}

func newWebServer(redisClient *redis.Client, omiWebClient, omiServerClient *omiclient.Client, serverName string, weight int) *WebServer {
	return &WebServer{
		router:          newRouter(omiServerClient.NewSearcher()),
		redisClient:     redisClient,
		omiWebClient:    omiWebClient,
		omiServerClient: omiServerClient,
		serverName:      serverName,
		weight:          weight,
	}
}

func (webServer *WebServer) EmbedSource(embeddedSource embed.FS) {
	webServer.embeddedSource = embeddedSource
	webServer.embedModel = true
}

func (webServer *WebServer) SetCache(cacheDir string, maxSize int) {
	var err error
	webServer.cache, err = newFileCache(cacheDir, maxSize)
	if err != nil {
		panic(err)
	}
}

func (webServer *WebServer) handleFunc(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) > 0 && webServer.router.Has(parts[1]) {
		pathRequestResolution(r, webServer.router)
		httpProxy(w, r, webServer.cache)
		websocketProxy(w, r)
		return
	}

	filePath := r.URL.Path
	if r.URL.Path == "/" {
		filePath = index_path
	}
	filePath = target_path + filePath
	var data []byte
	if webServer.embedModel {
		data, _ = webServer.embeddedSource.ReadFile(filePath)
	} else {
		data, _ = os.ReadFile(filePath)
	}
	w.Write(data)
}

func (webServer *WebServer) Listen(address string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		webServer.handleFunc(w, r)
	})
	webServer.omiWebClient.NewRegister(webServer.serverName, webServer.weight).Register(address)
	log.Println("omi web server: " + webServer.serverName + " is running on http://" + address)
	err := http.ListenAndServe(":"+strings.Split(address, ":")[1], nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
