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
      min-height: 100vh;
    }

    a { color: inherit; text-decoration: none; }

    /* ── HERO ── */
    .hero {
      background: #fff;
      border-bottom: 1px solid #e5e5ea;
      padding: 72px 24px 64px;
      text-align: center;
    }

    .hero-icon {
      width: 96px;
      height: 96px;
      background: linear-gradient(145deg, #e8f5ee, #d1f0dc);
      border-radius: 30px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 auto 28px;
      box-shadow: 0 4px 24px rgba(52, 199, 89, 0.15);
    }

    .hero-icon svg {
      width: 48px;
      height: 48px;
    }

    .badge {
      display: inline-flex;
      align-items: center;
      gap: 6px;
      background: #e8f5ee;
      color: #1a7f3c;
      font-size: 0.75rem;
      font-weight: 600;
      letter-spacing: 0.05em;
      text-transform: uppercase;
      padding: 5px 12px;
      border-radius: 100px;
      margin-bottom: 20px;
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
      50% { opacity: 0.4; }
    }

    .hero h1 {
      font-size: 2.6rem;
      font-weight: 700;
      letter-spacing: -0.04em;
      color: #1c1c1e;
      margin-bottom: 14px;
      line-height: 1.1;
    }

    .hero h1 span {
      color: #34c759;
    }

    .hero p {
      font-size: 1.05rem;
      color: #6e6e73;
      max-width: 480px;
      margin: 0 auto;
      line-height: 1.65;
    }

    /* ── STATS ── */
    .stats {
      display: flex;
      justify-content: center;
      gap: 0;
      background: #fff;
      border-bottom: 1px solid #e5e5ea;
      border-top: 1px solid #e5e5ea;
      margin-top: 0;
    }

    .stat {
      flex: 1;
      max-width: 180px;
      text-align: center;
      padding: 28px 16px;
      border-right: 1px solid #e5e5ea;
    }

    .stat:last-child { border-right: none; }

    .stat-value {
      font-size: 1.8rem;
      font-weight: 700;
      letter-spacing: -0.03em;
      color: #1c1c1e;
    }

    .stat-label {
      font-size: 0.75rem;
      color: #8e8e93;
      margin-top: 4px;
      letter-spacing: 0.01em;
    }

    /* ── LAYOUT ── */
    .container {
      max-width: 900px;
      margin: 0 auto;
      padding: 48px 24px 80px;
    }

    .section-title {
      font-size: 0.7rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.1em;
      color: #8e8e93;
      margin-bottom: 16px;
    }

    /* ── HOW IT WORKS ── */
    .steps {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 16px;
      margin-bottom: 40px;
    }

    .step {
      background: #fff;
      border-radius: 18px;
      padding: 28px 24px;
      border: 1px solid #e5e5ea;
    }

    .step-num {
      width: 32px;
      height: 32px;
      background: #f2f2f7;
      border-radius: 10px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 0.85rem;
      font-weight: 700;
      color: #34c759;
      margin-bottom: 16px;
    }

    .step h3 {
      font-size: 0.95rem;
      font-weight: 600;
      color: #1c1c1e;
      margin-bottom: 6px;
    }

    .step p {
      font-size: 0.82rem;
      color: #8e8e93;
      line-height: 1.55;
    }

    /* ── API REFERENCE ── */
    .api-card {
      background: #fff;
      border-radius: 18px;
      border: 1px solid #e5e5ea;
      overflow: hidden;
      margin-bottom: 40px;
    }

    .api-header {
      padding: 20px 24px;
      border-bottom: 1px solid #f2f2f7;
      display: flex;
      align-items: center;
      gap: 12px;
    }

    .method {
      background: #e8f5ee;
      color: #1a7f3c;
      font-size: 0.75rem;
      font-weight: 700;
      padding: 4px 10px;
      border-radius: 8px;
      letter-spacing: 0.05em;
    }

    .api-path {
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.95rem;
      color: #1c1c1e;
      font-weight: 500;
    }

    .api-desc {
      font-size: 0.82rem;
      color: #8e8e93;
      margin-left: auto;
    }

    .api-body {
      padding: 24px;
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 24px;
    }

    .api-section-label {
      font-size: 0.7rem;
      font-weight: 600;
      text-transform: uppercase;
      letter-spacing: 0.08em;
      color: #8e8e93;
      margin-bottom: 12px;
    }

    .param-list {
      display: flex;
      flex-direction: column;
      gap: 10px;
    }

    .param {
      display: flex;
      align-items: flex-start;
      gap: 10px;
    }

    .param-name {
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.8rem;
      color: #1c1c1e;
      background: #f2f2f7;
      padding: 3px 8px;
      border-radius: 6px;
      flex-shrink: 0;
    }

    .param-req {
      font-size: 0.68rem;
      color: #ff3b30;
      font-weight: 600;
      margin-top: 2px;
      flex-shrink: 0;
    }

    .param-opt {
      font-size: 0.68rem;
      color: #8e8e93;
      font-weight: 500;
      margin-top: 2px;
      flex-shrink: 0;
    }

    .param-desc {
      font-size: 0.8rem;
      color: #6e6e73;
      line-height: 1.4;
    }

    .code-block {
      background: #f2f2f7;
      border-radius: 12px;
      padding: 16px;
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.78rem;
      color: #1c1c1e;
      line-height: 1.7;
      overflow-x: auto;
    }

    .code-block .key { color: #5e5ce6; }
    .code-block .str { color: #34c759; }
    .code-block .punc { color: #8e8e93; }

    /* ── CATEGORIES ── */
    .categories-grid {
      display: grid;
      grid-template-columns: repeat(4, 1fr);
      gap: 12px;
      margin-bottom: 40px;
    }

    .cat-card {
      background: #fff;
      border: 1px solid #e5e5ea;
      border-radius: 16px;
      padding: 20px 18px;
      text-align: center;
    }

    .cat-emoji {
      font-size: 1.8rem;
      margin-bottom: 8px;
      display: block;
    }

    .cat-name {
      font-size: 0.85rem;
      font-weight: 600;
      color: #1c1c1e;
      margin-bottom: 4px;
    }

    .cat-examples {
      font-size: 0.72rem;
      color: #8e8e93;
      line-height: 1.5;
    }

    /* ── RESPONSE ── */
    .response-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 16px;
      margin-bottom: 40px;
    }

    .response-card {
      background: #fff;
      border-radius: 18px;
      border: 1px solid #e5e5ea;
      overflow: hidden;
    }

    .response-header {
      padding: 14px 18px;
      border-bottom: 1px solid #f2f2f7;
      display: flex;
      align-items: center;
      gap: 8px;
    }

    .status-dot {
      width: 8px;
      height: 8px;
      border-radius: 50%;
    }

    .status-201 { background: #34c759; }
    .status-401 { background: #ff9f0a; }

    .response-status {
      font-size: 0.8rem;
      font-weight: 600;
      color: #1c1c1e;
    }

    .response-label {
      font-size: 0.75rem;
      color: #8e8e93;
      margin-left: auto;
    }

    .response-body {
      padding: 16px 18px;
    }

    /* ── CURL ── */
    .curl-card {
      background: #1c1c1e;
      border-radius: 18px;
      overflow: hidden;
      margin-bottom: 40px;
    }

    .curl-header {
      padding: 14px 20px;
      border-bottom: 1px solid #2c2c2e;
      display: flex;
      align-items: center;
      gap: 8px;
    }

    .curl-dot {
      width: 10px;
      height: 10px;
      border-radius: 50%;
    }

    .curl-title {
      font-size: 0.75rem;
      color: #48484a;
      margin-left: 4px;
      font-family: "SF Mono", "Fira Code", monospace;
    }

    .curl-body {
      padding: 20px;
      font-family: "SF Mono", "Fira Code", monospace;
      font-size: 0.8rem;
      line-height: 1.8;
      color: #e5e5ea;
      overflow-x: auto;
      white-space: pre;
    }

    .curl-body .cmd { color: #64d2ff; }
    .curl-body .flag { color: #ff9f0a; }
    .curl-body .url { color: #30d158; }
    .curl-body .hkey { color: #bf5af2; }
    .curl-body .hval { color: #ffd60a; }
    .curl-body .jkey { color: #64d2ff; }
    .curl-body .jval { color: #30d158; }

    /* ── FOOTER ── */
    .footer {
      text-align: center;
      font-size: 0.78rem;
      color: #aeaeb2;
      padding: 32px 24px;
      border-top: 1px solid #e5e5ea;
      background: #fff;
    }

    /* ── RESPONSIVE ── */
    @media (max-width: 680px) {
      .hero h1 { font-size: 1.8rem; }
      .steps { grid-template-columns: 1fr; }
      .categories-grid { grid-template-columns: repeat(2, 1fr); }
      .api-body { grid-template-columns: 1fr; }
      .response-grid { grid-template-columns: 1fr; }
      .stats { flex-wrap: wrap; }
      .stat { border-right: none; border-bottom: 1px solid #e5e5ea; max-width: 100%; }
      .api-desc { display: none; }
    }
  </style>
</head>
<body>

  <!-- HERO -->
  <div class="hero">
    <div class="hero-icon">
      <svg viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
        <rect x="4" y="12" width="40" height="26" rx="6" fill="#e8f5ee" stroke="#34c759" stroke-width="1.5"/>
        <rect x="4" y="19" width="40" height="7" fill="#34c759" opacity="0.12"/>
        <rect x="10" y="30" width="10" height="4" rx="2" fill="#34c759"/>
        <rect x="24" y="30" width="7" height="4" rx="2" fill="#d1d1d6"/>
        <circle cx="36" cy="10" r="6" fill="#34c759"/>
        <path d="M33 10l2 2 4-4" stroke="#fff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
    </div>
    <div class="badge"><span class="badge-dot"></span>Live</div>
    <h1>Apple Wallet <span>Automation</span></h1>
    <p>Automatically captures and categorizes every Apple Pay transaction sent from your iPhone Shortcut — no manual logging, ever.</p>
  </div>

  <!-- STATS -->
  <div class="stats">
    <div class="stat">
      <div class="stat-value">8</div>
      <div class="stat-label">Categories</div>
    </div>
    <div class="stat">
      <div class="stat-value">POST</div>
      <div class="stat-label">Method</div>
    </div>
    <div class="stat">
      <div class="stat-value">JSON</div>
      <div class="stat-label">Storage</div>
    </div>
    <div class="stat">
      <div class="stat-value">API Key</div>
      <div class="stat-label">Auth</div>
    </div>
  </div>

  <div class="container">

    <!-- HOW IT WORKS -->
    <div class="section-title">How it works</div>
    <div class="steps">
      <div class="step">
        <div class="step-num">1</div>
        <h3>Pay with Apple Pay</h3>
        <p>Make any purchase on your iPhone using Apple Wallet as you normally would.</p>
      </div>
      <div class="step">
        <div class="step-num">2</div>
        <h3>Shortcut fires</h3>
        <p>An iOS Shortcut detects the transaction and instantly POSTs the details to this server.</p>
      </div>
      <div class="step">
        <div class="step-num">3</div>
        <h3>Recorded &amp; categorized</h3>
        <p>The server validates your API key, detects the category from the merchant name, and saves it to your personal JSON log.</p>
      </div>
    </div>

    <!-- API REFERENCE -->
    <div class="section-title">API Reference</div>
    <div class="api-card">
      <div class="api-header">
        <span class="method">POST</span>
        <span class="api-path">/record/</span>
        <span class="api-desc">Record a transaction</span>
      </div>
      <div class="api-body">
        <div>
          <div class="api-section-label">Request fields</div>
          <div class="param-list">
            <div class="param">
              <span class="param-name">merchant</span>
              <span class="param-req">required</span>
              <span class="param-desc">Merchant or store name used for category detection</span>
            </div>
            <div class="param">
              <span class="param-name">amount</span>
              <span class="param-req">required</span>
              <span class="param-desc">Transaction amount, any currency format</span>
            </div>
            <div class="param">
              <span class="param-name">card</span>
              <span class="param-opt">optional</span>
              <span class="param-desc">Card label, e.g. Visa Classic</span>
            </div>
          </div>
          <br/>
          <div class="api-section-label">Required header</div>
          <div class="param-list">
            <div class="param">
              <span class="param-name">X-Api-Key</span>
              <span class="param-req">required</span>
              <span class="param-desc">Your personal API key defined in credentials.yml</span>
            </div>
          </div>
        </div>
        <div>
          <div class="api-section-label">Request body</div>
          <div class="code-block"><span class="punc">{</span>
  <span class="key">"merchant"</span><span class="punc">:</span> <span class="str">"Yerevan City Komitas"</span><span class="punc">,</span>
  <span class="key">"amount"</span><span class="punc">:</span>   <span class="str">"֏26 307,00"</span><span class="punc">,</span>
  <span class="key">"card"</span><span class="punc">:</span>     <span class="str">"Visa Classic"</span>
<span class="punc">}</span></div>
        </div>
      </div>
    </div>

    <!-- EXAMPLE CURL -->
    <div class="section-title">Example request</div>
    <div class="curl-card">
      <div class="curl-header">
        <div class="curl-dot" style="background:#ff5f57"></div>
        <div class="curl-dot" style="background:#febc2e"></div>
        <div class="curl-dot" style="background:#28c840"></div>
        <span class="curl-title">terminal</span>
      </div>
      <div class="curl-body"><span class="cmd">curl</span> <span class="flag">-X POST</span> <span class="url">https://applepay.vendorex.shop/record/</span> \
  <span class="flag">-H</span> <span class="hkey">'X-Api-Key'</span><span class="hval">: 'YOUR_API_KEY'</span> \
  <span class="flag">-H</span> <span class="hval">'Content-Type: application/json'</span> \
  <span class="flag">-d</span> '<span class="punc">{</span>
    <span class="jkey">"merchant"</span>: <span class="jval">"Yerevan City Komitas"</span>,
    <span class="jkey">"amount"</span>:   <span class="jval">"֏26 307,00"</span>,
    <span class="jkey">"card"</span>:     <span class="jval">"Visa Classic"</span>
  <span class="punc">}</span>'</div>
    </div>

    <!-- RESPONSES -->
    <div class="section-title">Responses</div>
    <div class="response-grid">
      <div class="response-card">
        <div class="response-header">
          <div class="status-dot status-201"></div>
          <span class="response-status">201 Created</span>
          <span class="response-label">Transaction saved</span>
        </div>
        <div class="response-body">
          <div class="code-block"><span class="punc">{</span>
  <span class="key">"id"</span><span class="punc">:</span>          <span class="str">"uuid-v4"</span><span class="punc">,</span>
  <span class="key">"category"</span><span class="punc">:</span>    <span class="str">"Groceries"</span><span class="punc">,</span>
  <span class="key">"recorded_at"</span><span class="punc">:</span> <span class="str">"2026-08-16T..."</span>
<span class="punc">}</span></div>
        </div>
      </div>
      <div class="response-card">
        <div class="response-header">
          <div class="status-dot status-401"></div>
          <span class="response-status">401 Unauthorized</span>
          <span class="response-label">Not recorded</span>
        </div>
        <div class="response-body">
          <div class="code-block"><span class="punc">{</span>
  <span class="key">"error"</span><span class="punc">:</span> <span class="str">"missing api key"</span>
<span class="punc">}</span>

<span class="punc">{</span>
  <span class="key">"error"</span><span class="punc">:</span> <span class="str">"invalid api key"</span>
<span class="punc">}</span></div>
        </div>
      </div>
    </div>

    <!-- CATEGORIES -->
    <div class="section-title">Auto-detected categories</div>
    <div class="categories-grid">
      <div class="cat-card">
        <span class="cat-emoji">🛒</span>
        <div class="cat-name">Groceries</div>
        <div class="cat-examples">Yerevan City, Carrefour, Walmart, Tesco…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">🍔</span>
        <div class="cat-name">Restaurant</div>
        <div class="cat-examples">KFC, McDonald's, Starbucks, Subway…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">🚗</span>
        <div class="cat-name">Transport</div>
        <div class="cat-examples">Uber, Bolt, Taxi, Parking, Fuel…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">🛍</span>
        <div class="cat-name">Shopping</div>
        <div class="cat-examples">Amazon, Zara, Nike, IKEA, Mall…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">💊</span>
        <div class="cat-name">Health</div>
        <div class="cat-examples">Pharmacy, Hospital, Gym, Clinic…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">🎬</span>
        <div class="cat-name">Entertainment</div>
        <div class="cat-examples">Cinema, Netflix, Steam, Concert…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">✈️</span>
        <div class="cat-name">Travel</div>
        <div class="cat-examples">Hotel, Airbnb, Booking, Resort…</div>
      </div>
      <div class="cat-card">
        <span class="cat-emoji">📦</span>
        <div class="cat-name">Other</div>
        <div class="cat-examples">Anything that doesn't match a keyword</div>
      </div>
    </div>

  </div>

  <!-- FOOTER -->
  <div class="footer">
    Apple Wallet Automation &nbsp;·&nbsp; Keywords configurable in <code>config.yml</code> &nbsp;·&nbsp; Data stored per-user as JSON
  </div>

</body>
</html>`)
}
