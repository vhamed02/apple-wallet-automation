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
      border-radius: 28px;
      border: 1px solid #e5e5ea;
      padding: 44px 52px;
      width: 100%;
      max-width: 680px;
    }

    /* TOP ROW */
    .top {
      display: flex;
      align-items: center;
      gap: 20px;
      margin-bottom: 28px;
    }

    .icon {
      width: 64px;
      height: 64px;
      background: #e8f5ee;
      border-radius: 18px;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-shrink: 0;
    }

    .icon svg { width: 34px; height: 34px; }

    .top-text h1 {
      font-size: 1.45rem;
      font-weight: 700;
      letter-spacing: -0.03em;
      color: #1c1c1e;
      margin-bottom: 4px;
    }

    .top-text p {
      font-size: 0.875rem;
      color: #8e8e93;
      line-height: 1.5;
    }

    .badge {
      margin-left: auto;
      display: flex;
      align-items: center;
      gap: 6px;
      background: #e8f5ee;
      color: #1a7f3c;
      font-size: 0.72rem;
      font-weight: 600;
      letter-spacing: 0.05em;
      padding: 5px 12px;
      border-radius: 100px;
      white-space: nowrap;
      align-self: flex-start;
    }

    .badge-dot {
      width: 6px;
      height: 6px;
      background: #34c759;
      border-radius: 50%;
      animation: pulse 2s infinite;
    }

    @keyframes pulse {
      0%, 100% { opacity: 1; }
      50%       { opacity: 0.3; }
    }

    /* DIVIDER */
    .divider { height: 1px; background: #f2f2f7; margin-bottom: 24px; }

    /* STATS ROW */
    .stats {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: 12px;
      margin-bottom: 24px;
    }

    .stat {
      background: #f9f9fb;
      border: 1px solid #e5e5ea;
      border-radius: 14px;
      padding: 14px 16px;
      text-align: center;
    }

    .stat-value {
      font-size: 0.95rem;
      font-weight: 650;
      color: #1c1c1e;
      letter-spacing: -0.01em;
    }

    .stat-label {
      font-size: 0.68rem;
      color: #aeaeb2;
      margin-top: 3px;
    }

    /* ENDPOINT + CATEGORIES ROW */
    .mid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 12px;
      margin-bottom: 24px;
    }

    .mid-block {
      background: #f9f9fb;
      border: 1px solid #e5e5ea;
      border-radius: 14px;
      padding: 16px 18px;
    }

    .block-label {
      font-size: 0.65rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.09em;
      color: #aeaeb2;
      margin-bottom: 10px;
    }

    .endpoint-row {
      display: flex;
      align-items: center;
      gap: 8px;
    }

    .method {
      background: #e8f5ee;
      color: #1a7f3c;
      font-size: 0.7rem;
      font-weight: 700;
      padding: 3px 8px;
      border-radius: 6px;
      letter-spacing: 0.04em;
    }

    .path {
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.82rem;
      color: #1c1c1e;
    }

    .fields {
      display: flex;
      flex-wrap: wrap;
      gap: 6px;
    }

    .field {
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.72rem;
      background: #fff;
      border: 1px solid #e5e5ea;
      color: #3c3c43;
      padding: 3px 9px;
      border-radius: 7px;
    }

    .field.req { border-color: #c6ecd1; color: #1a7f3c; background: #f0fbf3; }

    .cats {
      display: flex;
      flex-wrap: wrap;
      gap: 6px;
    }

    .cat {
      font-size: 0.72rem;
      background: #fff;
      border: 1px solid #e5e5ea;
      color: #3c3c43;
      padding: 3px 10px;
      border-radius: 7px;
    }

    /* BOTTOM ROW */
    .bottom {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 12px;
    }

    .storage-note {
      font-size: 0.78rem;
      color: #aeaeb2;
    }

    .storage-note code {
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.74rem;
      background: #f2f2f7;
      padding: 2px 6px;
      border-radius: 5px;
      color: #3c3c43;
    }

    .github-link {
      display: inline-flex;
      align-items: center;
      gap: 7px;
      background: #1c1c1e;
      color: #f5f5f7;
      text-decoration: none;
      font-size: 0.82rem;
      font-weight: 500;
      padding: 10px 20px;
      border-radius: 100px;
      white-space: nowrap;
      transition: opacity 0.15s;
    }

    .github-link:hover { opacity: 0.75; }
    .github-link svg { width: 15px; height: 15px; fill: #f5f5f7; flex-shrink: 0; }
  </style>
</head>
<body>
  <div class="card">

    <!-- TOP -->
    <div class="top">
      <div class="icon">
        <svg viewBox="0 0 34 34" fill="none" xmlns="http://www.w3.org/2000/svg">
          <rect x="2" y="8" width="30" height="20" rx="5" fill="#e8f5ee" stroke="#34c759" stroke-width="1.5"/>
          <rect x="2" y="14" width="30" height="5" fill="#34c759" opacity="0.15"/>
          <rect x="6" y="21" width="8" height="3" rx="1.5" fill="#34c759"/>
          <rect x="17" y="21" width="5" height="3" rx="1.5" fill="#d1d1d6"/>
        </svg>
      </div>
      <div class="top-text">
        <h1>Apple Wallet Automation</h1>
        <p>Records and categorizes Apple Pay transactions from your iPhone Shortcut.</p>
      </div>
      <div class="badge"><span class="badge-dot"></span>Live</div>
    </div>

    <div class="divider"></div>

    <!-- STATS -->
    <div class="stats">
      <div class="stat">
        <div class="stat-value">POST /record/</div>
        <div class="stat-label">Endpoint</div>
      </div>
      <div class="stat">
        <div class="stat-value">X-Api-Key</div>
        <div class="stat-label">Auth header</div>
      </div>
      <div class="stat">
        <div class="stat-value">8</div>
        <div class="stat-label">Categories</div>
      </div>
      <div class="stat">
        <div class="stat-value">JSON</div>
        <div class="stat-label">Per-user storage</div>
      </div>
    </div>

    <!-- ENDPOINT + CATEGORIES -->
    <div class="mid">
      <div class="mid-block">
        <div class="block-label">Request fields</div>
        <div class="endpoint-row" style="margin-bottom:10px">
          <span class="method">POST</span>
          <span class="path">/record/</span>
        </div>
        <div class="fields">
          <span class="field req">merchant</span>
          <span class="field req">amount</span>
          <span class="field">card</span>
        </div>
      </div>
      <div class="mid-block">
        <div class="block-label">Auto-detected categories</div>
        <div class="cats">
          <span class="cat">🛒 Groceries</span>
          <span class="cat">🍔 Restaurant</span>
          <span class="cat">🚗 Transport</span>
          <span class="cat">🛍 Shopping</span>
          <span class="cat">💊 Health</span>
          <span class="cat">🎬 Entertainment</span>
          <span class="cat">✈️ Travel</span>
          <span class="cat">📦 Other</span>
        </div>
      </div>
    </div>

    <!-- BOTTOM -->
    <div class="bottom">
      <span class="storage-note">Data stored in <code>data/&lt;username&gt;.json</code> · Keywords configurable in <code>config.yml</code></span>
      <a class="github-link" href="https://github.com/vhamed02/apple-wallet-automation" target="_blank" rel="noopener">
        <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path d="M12 .297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385.6.113.82-.258.82-.577 0-.285-.01-1.04-.015-2.04-3.338.724-4.042-1.61-4.042-1.61C4.422 18.07 3.633 17.7 3.633 17.7c-1.087-.744.084-.729.084-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22 0 1.606-.015 2.896-.015 3.286 0 .315.21.69.825.57C20.565 22.092 24 17.592 24 12.297c0-6.627-5.373-12-12-12"/>
        </svg>
        View on GitHub
      </a>
    </div>

  </div>
</body>
</html>`)
}
