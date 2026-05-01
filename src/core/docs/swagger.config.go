package docs

import (
	"github.com/gofiber/fiber/v3"
)

const swaggerUIHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>DMS Service Documentation</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
  <style>
    body { margin: 0; padding: 0; }
    #swagger-ui { max-width: 1200px; margin: 0 auto; padding: 20px; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
  <script>
    window.onload = () => {
      SwaggerUIBundle({
        url: "/docs/swagger.yaml",
        dom_id: "#swagger-ui",
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIBundle.SwaggerUIStandalonePreset],
        layout: "BaseLayout",
        deepLinking: true,
      });
    };
  </script>
</body>
</html>`

func SetUp(app *fiber.App) {
	app.Get("/docs", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString(swaggerUIHTML)
	})

	app.Get("/docs/swagger.yaml", func(c fiber.Ctx) error {
		return c.SendFile("./src/core/docs/swagger.yaml")
	})
}
