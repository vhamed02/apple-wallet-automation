package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Apple Wallet Automation</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    body {
      font-family: -apple-system, BlinkMacSystemFont, "SF Pro Display", "Segoe UI", sans-serif;
      background: #000;
      color: #f5f5f7;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .card {
      text-align: center;
      padding: 64px 48px;
      max-width: 520px;
      width: 100%;
    }

    .icon {
      width: 88px;
      height: 88px;
      background: linear-gradient(145deg, #1c1c1e, #2c2c2e);
      border-radius: 28px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 auto 32px;
      box-shadow: 0 8px 32px rgba(0,0,0,0.6);
    }

    .icon svg {
      width: 44px;
      height: 44px;
    }

    h1 {
      font-size: 2rem;
      font-weight: 700;
      letter-spacing: -0.03em;
      color: #f5f5f7;
      margin-bottom: 12px;
    }

    p {
      font-size: 1rem;
      color: #86868b;
      line-height: 1.6;
      margin-bottom: 40px;
    }

    .pills {
      display: flex;
      flex-wrap: wrap;
      gap: 10px;
      justify-content: center;
      margin-bottom: 48px;
    }

    .pill {
      background: #1c1c1e;
      border: 1px solid #2c2c2e;
      border-radius: 100px;
      padding: 6px 16px;
      font-size: 0.8rem;
      color: #86868b;
      letter-spacing: 0.01em;
    }

    .endpoint {
      background: #0a0a0a;
      border: 1px solid #1c1c1e;
      border-radius: 16px;
      padding: 20px 24px;
      text-align: left;
    }

    .endpoint-label {
      font-size: 0.7rem;
      text-transform: uppercase;
      letter-spacing: 0.1em;
      color: #48484a;
      margin-bottom: 8px;
    }

    .endpoint-row {
      display: flex;
      align-items: center;
      gap: 10px;
    }

    .method {
      background: #1d3a2e;
      color: #30d158;
      font-size: 0.72rem;
      font-weight: 600;
      padding: 3px 8px;
      border-radius: 6px;
      letter-spacing: 0.05em;
      flex-shrink: 0;
    }

    .path {
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.88rem;
      color: #f5f5f7;
    }

    .footer {
      margin-top: 40px;
      font-size: 0.75rem;
      color: #3a3a3c;
    }
  </style>
</head>
<body>
  <div class="card">
    <div class="icon">
      <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="2" y="5" width="20" height="14" rx="3" fill="#1c1c1e" stroke="#48484a" stroke-width="1.5"/>
        <rect x="2" y="9" width="20" height="3" fill="#30d158" opacity="0.15"/>
        <rect x="5" y="14" width="4" height="2" rx="1" fill="#30d158"/>
        <rect x="11" y="14" width="3" height="2" rx="1" fill="#48484a"/>
      </svg>
    </div>

    <h1>Wallet Automation</h1>
    <p>Automatically records and categorizes your Apple Pay transactions sent from iPhone Shortcuts.</p>

    <div class="pills">
      <span class="pill">Groceries</span>
      <span class="pill">Restaurant</span>
      <span class="pill">Transport</span>
      <span class="pill">Shopping</span>
      <span class="pill">Health</span>
      <span class="pill">Travel</span>
      <span class="pill">Entertainment</span>
      <span class="pill">Other</span>
    </div>

    <div class="endpoint">
      <div class="endpoint-label">Endpoint</div>
      <div class="endpoint-row">
        <span class="method">POST</span>
        <span class="path">/record/</span>
      </div>
    </div>

    <div class="footer">Requires X-Api-Key header</div>
  </div>
</body>
</html>`)
}
