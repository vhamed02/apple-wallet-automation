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
      background: #f2f2f7;
      color: #1c1c1e;
      height: 100vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      overflow: hidden;
    }

    .card {
      background: #fff;
      border-radius: 24px;
      border: 1px solid #e5e5ea;
      padding: 48px 56px;
      text-align: center;
      width: 100%;
      max-width: 440px;
    }

    .icon {
      width: 72px;
      height: 72px;
      background: #e8f5ee;
      border-radius: 22px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 auto 24px;
    }

    .icon svg {
      width: 36px;
      height: 36px;
    }

    h1 {
      font-size: 1.5rem;
      font-weight: 700;
      letter-spacing: -0.03em;
      color: #1c1c1e;
      margin-bottom: 8px;
    }

    .sub {
      font-size: 0.9rem;
      color: #8e8e93;
      line-height: 1.55;
      margin-bottom: 32px;
    }

    .divider {
      height: 1px;
      background: #f2f2f7;
      margin-bottom: 24px;
    }

    .meta {
      display: flex;
      justify-content: center;
      gap: 20px;
      margin-bottom: 28px;
    }

    .meta-item {
      text-align: center;
    }

    .meta-value {
      font-size: 0.95rem;
      font-weight: 600;
      color: #1c1c1e;
    }

    .meta-label {
      font-size: 0.7rem;
      color: #aeaeb2;
      margin-top: 2px;
    }

    .github-link {
      display: inline-flex;
      align-items: center;
      gap: 8px;
      background: #1c1c1e;
      color: #f5f5f7;
      text-decoration: none;
      font-size: 0.85rem;
      font-weight: 500;
      padding: 11px 22px;
      border-radius: 100px;
      transition: opacity 0.15s;
    }

    .github-link:hover { opacity: 0.8; }

    .github-link svg {
      width: 16px;
      height: 16px;
      fill: #f5f5f7;
    }
  </style>
</head>
<body>
  <div class="card">
    <div class="icon">
      <svg viewBox="0 0 36 36" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="3" y="9" width="30" height="20" rx="5" fill="#e8f5ee" stroke="#34c759" stroke-width="1.5"/>
        <rect x="3" y="15" width="30" height="5" fill="#34c759" opacity="0.15"/>
        <rect x="7" y="22" width="8" height="3" rx="1.5" fill="#34c759"/>
        <rect x="18" y="22" width="5" height="3" rx="1.5" fill="#d1d1d6"/>
      </svg>
    </div>

    <h1>Wallet Automation</h1>
    <p class="sub">Records and categorizes Apple Pay transactions sent from your iPhone Shortcut.</p>

    <div class="divider"></div>

    <div class="meta">
      <div class="meta-item">
        <div class="meta-value">POST /record/</div>
        <div class="meta-label">Endpoint</div>
      </div>
      <div class="meta-item">
        <div class="meta-value">X-Api-Key</div>
        <div class="meta-label">Auth header</div>
      </div>
      <div class="meta-item">
        <div class="meta-value">8</div>
        <div class="meta-label">Categories</div>
      </div>
    </div>

    <a class="github-link" href="https://github.com/vhamed02/apple-wallet-automation" target="_blank" rel="noopener">
      <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path d="M12 .297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385.6.113.82-.258.82-.577 0-.285-.01-1.04-.015-2.04-3.338.724-4.042-1.61-4.042-1.61C4.422 18.07 3.633 17.7 3.633 17.7c-1.087-.744.084-.729.084-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22 0 1.606-.015 2.896-.015 3.286 0 .315.21.69.825.57C20.565 22.092 24 17.592 24 12.297c0-6.627-5.373-12-12-12"/>
      </svg>
      View on GitHub
    </a>
  </div>
</body>
</html>`)
}
